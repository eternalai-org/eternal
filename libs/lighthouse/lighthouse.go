package lighthouse

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"eternal-infer-worker/libs/eth"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"github.com/gabriel-vasile/mimetype"
	"github.com/ipfs/boxo/blockservice"
	blockstore "github.com/ipfs/boxo/blockstore"
	chunker "github.com/ipfs/boxo/chunker"
	offline "github.com/ipfs/boxo/exchange/offline"
	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs/importer/balanced"
	uih "github.com/ipfs/boxo/ipld/unixfs/importer/helpers"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/pkg/errors"
)

type Response struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

type gcstorage struct {
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle
	projectId  string
	ctx        context.Context
	formatType string
}

type GCS struct {
	ProjectId string
	Bucket    string
	Auth      string
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
}

var GCSStorage *gcstorage

const (
	IPFSGateway = "https://gateway.lighthouse.storage/ipfs/"
)

func NewDataGCStorage(config GCS) (*gcstorage, error) {
	// Creates a Google Cloud client from config GC Auth key
	jsonKey, _ := base64.StdEncoding.DecodeString(config.Auth)
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(jsonKey)))
	if err != nil {
		return nil, err
	}

	// Creates a Bucket instance.
	bucket := client.Bucket(config.Bucket)

	// Init our GCStorage object
	gcStorage := gcstorage{
		bucketName: config.Bucket,    // get bucket name from config
		bucket:     bucket,           // assign bucket object
		client:     client,           // assign client object
		ctx:        ctx,              // assign context object
		projectId:  config.ProjectId, // assign project id, not required
	}

	GCSStorage = &gcStorage
	return &gcStorage, nil
}

func DownloadDataSimple(hash string) ([]byte, string, error) {
	if strings.Contains(hash, "ipfs://") {
		hash = strings.Replace(hash, "ipfs://", "", 1)
	}
	// https://gateway.lighthouse.storage/ipfs/QmXPGcEHCi1ZmbHFwScuP4ZJ2iv9YjTJMUroUTJUnFXxxj
	// urlLink := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash)

	urlLink := fmt.Sprintf("https://cdn.eternalai.org/upload/%s", hash)

	resp, err := http.Get(urlLink)
	if err != nil {
		return nil, "", fmt.Errorf("error when try get reponse :%v", err)
	}
	if resp == nil {
		return nil, "", fmt.Errorf("error when try get reponse ==nil")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return DownloadDataSimpleFromLighthouse(hash)
		}
		return nil, "", fmt.Errorf("error when try get data url :%v => reponse code :%v", urlLink, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error when read body from  reponse :%v", err)
	}

	mtype := mimetype.Detect(body)

	return body, mtype.String(), nil
}

func DownloadDataSimpleFromLighthouse(hash string) ([]byte, string, error) {
	if strings.Contains(hash, "ipfs://") {
		hash = strings.Replace(hash, "ipfs://", "", 1)
	}
	// https://gateway.lighthouse.storage/ipfs/QmXPGcEHCi1ZmbHFwScuP4ZJ2iv9YjTJMUroUTJUnFXxxj
	urlLink := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash)

	resp, err := http.Get(urlLink)
	if err != nil {
		return nil, "", fmt.Errorf("error when try get reponse :%v", err)
	}
	if resp == nil {
		return nil, "", fmt.Errorf("error when try get reponse ==nil")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("error when try get data url :%v => reponse code :%v", urlLink, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error when read body from  reponse :%v", err)
	}

	mtype := mimetype.Detect(body)

	return body, mtype.String(), nil
}

func downloadChunkedData(hash string, start, end int) ([]byte, error) {
	urlLink := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash)

	req, err := http.NewRequest("GET", urlLink, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}

func DownloadChunkedData(hash string, modelDir string) (string, error) {
	// check folder exist first, if not create
	err := os.MkdirAll(modelDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileInfo, err := GetFileInfo(hash)
	if err != nil {
		return "", err
	}

	fileSize, err := strconv.ParseInt(fileInfo.FileSizeInBytes, 10, 64)
	if err != nil {
		return "", err
	}

	filePath := path.Join(modelDir, fmt.Sprintf("model.zip"))
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	var chunkSize int64 = 1024 * 1024 * 10 // 10MB
	var start int64 = 0
	var end int64 = chunkSize

	parts := fileSize / chunkSize
	log.Println("parts: ", parts)
	part := 0
	for start < fileSize {
		if end > fileSize {
			end = fileSize
		}

		log.Println("part: ", part, "start: ", start, "end: ", end)
		data, err := downloadChunkedData(hash, int(start), int(end))
		if err != nil {
			return "", err
		}

		err = WriteFile(filePath, data, os.ModePerm)
		if err != nil {
			return "", err
		}

		start = end + 1
		end += chunkSize

		part++

	}

	return filePath, nil
}

func WriteFile(name string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

type FileInfo struct {
	FileSizeInBytes string `json:"fileSizeInBytes"`
	Cid             string `json:"cid"`
	Encryption      bool   `json:"encryption"`
	FileName        string `json:"fileName"`
	MimeType        string `json:"mimeType"`
	TxHash          string `json:"txHash"`
	Error           struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func GetFileInfo(hash string) (*FileInfo, error) {
	urlLink := fmt.Sprintf("https://api.lighthouse.storage/api/lighthouse/file_info?cid=%s", hash)

	resp, err := http.Get(urlLink)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Println("body", string(body))

	var respBody FileInfo

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if respBody.Error.Code != 0 {
		return &respBody, errors.New(respBody.Error.Message)
	}

	return &respBody, nil
}

func UploadData(apikey, fileName string, data []byte) (string, error) {
	cidHash := Cid(data)

	bucket := "eternal-ai"
	object := "upload/" + cidHash
	ctx := context.Background()

	buf := bytes.NewBuffer(data)

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := GCSStorage.client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err := io.Copy(wc, buf); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %w", err)
	}

	return cidHash, nil
}

func GenAPIkey(privkey string) (string, error) {
	priv, address, err := eth.GetAccountInfo(privkey)
	if err != nil {
		return "", err
	}

	msg, err := getMsgToSign(address.String())
	if err != nil {
		return "", err
	}

	sig, err := eth.SignMessage(msg, priv)
	if err != nil {
		return "", err
	}

	apikey, err := getAPIKey(address.String(), sig)
	if err != nil {
		return "", err
	}
	return apikey, nil
}

func getMsgToSign(address string) (string, error) {
	urlLink := fmt.Sprintf("https://api.lighthouse.storage/api/auth/get_message?publicKey=%s", address)

	req, err := http.NewRequest(
		"GET",
		urlLink,
		nil,
	)
	if err != nil {
		return "", errors.WithStack(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	log.Println("body", string(body))

	return string(body), nil
}

func getAPIKey(address, sig string) (string, error) {
	urlLink := "https://api.lighthouse.storage/api/auth/create_api_key"

	reqBody := map[string]string{
		"publicKey":     address,
		"signedMessage": sig,
		"keyName":       "eternal-node",
	}

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		urlLink,
		bytes.NewBuffer(reqBodyJSON),
	)
	if err != nil {
		return "", errors.WithStack(err)
	}

	client := &http.Client{}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	log.Println("body", string(body))

	return string(body), nil
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// func Cid(data []byte) string {
// 	// Create an IPLD UnixFS chunker with size 1 MiB
// 	chunks := chunker.NewSizeSplitter(bytes.NewReader(data), 1024*1024)

// 	// Concatenate the chunks to build the DAG
// 	var buf bytes.Buffer
// 	for {
// 		chunk, err := chunks.NextBytes()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			panic(err)
// 		}

// 		buf.Write(chunk)
// 	}

// 	// Calculate the CID for the DAG
// 	hash, err := mh.Sum(buf.Bytes(), mh.SHA2_256, -1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create a CID version 1 (with multibase encoding base58btc)
// 	c := cid.NewCidV0(hash)

// 	// Print the CID as a string
// 	fmt.Println("IPFS CID in Golang:", c.String())

// 	return c.String()
// }

func Cid(data []byte) string {
	ds := dsync.MutexWrap(datastore.NewNullDatastore())
	bs := blockstore.NewBlockstore(ds)
	bs = blockstore.NewIdStore(bs)
	bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)
	ufsImportParams := uih.DagBuilderParams{
		Maxlinks:   uih.BlockSizeLimit, // Default max of 174 links per block
		RawLeaves:  false,
		CidBuilder: cid.V0Builder{},
		Dagserv:    dsrv,
		NoCopy:     false,
	}
	reader := bytes.NewReader(data)
	ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(reader, chunker.DefaultBlockSize)) // 256KiB chunks
	if err != nil {
		return cid.Undef.String()
	}
	nd, err := balanced.Layout(ufsBuilder)
	if err != nil {
		return cid.Undef.String()
	}
	return nd.Cid().String()
}

func fileExistOnNetwork(data []byte) (string, bool, error) {
	cid := Cid(data)
	fileInfo, err := GetFileInfo(cid)
	if err != nil {
		log.Error("fileExistOnNetwork - fileInfo: ", fileInfo, " ,err: ", err)
		if fileInfo != nil && fileInfo.Error.Code == 404 {
			return "", false, nil
		}
		return "", false, nil
	}

	return cid, true, nil
}

func DownloadToFile(hash string, filePath string) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := os.Stat(filePath); err == nil {
		return nil
	}
	data, _, err := DownloadDataSimple(hash)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		f.Close()
		os.Remove(filePath)
		return err
	}
	err = f.Close()
	if err != nil {
		os.Remove(filePath)
	}
	return nil
}

func DownloadDataSimpleWithRetry(hash string) ([]byte, string, error) {
	var err error
	var byteResp []byte
	var mimeType string
	for i := 0; i < 3; i++ {
		byteResp, mimeType, err = DownloadDataSimple(hash)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		return byteResp, mimeType, nil
	}

	return byteResp, mimeType, nil
}
