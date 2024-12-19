package lighthouse

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	chunker "github.com/ipfs/boxo/chunker"
	uih "github.com/ipfs/boxo/ipld/unixfs/importer/helpers"

	"github.com/go-resty/resty/v2"
	"github.com/ipfs/boxo/blockservice"
	"github.com/ipfs/boxo/blockstore"
	"github.com/ipfs/boxo/exchange/offline"
	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs/importer/balanced"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
)

type ILighthouse interface {
	FileInfo(ctx context.Context, cid string) (*FileInfo, error)
	DownloadFile(ctx context.Context, link string, rangeHeader string) ([]byte, error)
	DownloadFileByCid(ctx context.Context, cid string, rangeHeader string) ([]byte, error)
	UploadDataFile(ctx context.Context, filename string, byte []byte) (*Response, error)
	UploadDataFileByUrl(ctx context.Context, url string) (*Response, error)
	FilesUploaded(ctx context.Context) ([]*FileInfo, error)
}

type lighthouse struct {
	apiKey string
}

// Response represents the response structure for file operations
type Response struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"` // Note: Hash in response is CID.
	Size string `json:"Size"`
}

func (r *Response) FileCoinLink() string {
	return fmt.Sprintf("%s/%s", IPFSGatewayUrl, r.Hash)
}

// FileInfo represents the metadata information of a file
type FileInfo struct {
	Id         string `json:"id"`
	CreatedAt  int64  `json:"createdAt"`
	LastUpdate int64  `json:"lastUpdate"`

	FileSizeInBytes string `json:"fileSizeInBytes"`
	Cid             string `json:"cid"`
	Encryption      bool   `json:"encryption"`
	FileName        string `json:"fileName"`
	MimeType        string `json:"mimeType"`
	TxHash          string `json:"txHash"`
	PublicKey       string `json:"publicKey"`
	Status          string `json:"status"`

	Error *Error `json:"error"`
}
type ListFileResp struct {
	FileList   []*FileInfo `json:"fileList"`
	TotalFiles int64       `json:"totalFiles"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	FilesUploadedUrl string = "https://api.lighthouse.storage/api/user/files_uploaded"
	IPFSGatewayUrl   string = "https://gateway.lighthouse.storage/ipfs"
	UploadDataUrl    string = "https://node.lighthouse.storage/api/v0/add"
	FileInfoUrl      string = "https://api.lighthouse.storage/api/lighthouse/file_info"

	ChunkSize = 10 * 1024 * 1024 // 10MB
)

func NewLighthouse(apiKey string) ILighthouse {
	return &lighthouse{
		apiKey: apiKey,
	}
}

func (l *lighthouse) FilesUploaded(ctx context.Context) ([]*FileInfo, error) {
	resp := &ListFileResp{}
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, err := client.R().SetContext(ctx).
		SetAuthScheme("Bearer").
		SetAuthToken(l.apiKey).
		SetResult(resp).
		Get(FilesUploadedUrl)
	if err != nil {
		return nil, err
	}

	return resp.FileList, nil
}

func (l *lighthouse) UploadDataFileByUrl(ctx context.Context, rawUrl string) (*Response, error) {
	parsedUrl, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return nil, err
	}

	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	res, err := client.R().SetContext(ctx).Get(rawUrl)
	if err != nil {
		return nil, err
	}

	fileName := strings.Replace(res.Header().Get("Content-Disposition"), "attachment; filename=", "", -1)
	fileName = strings.Replace(fileName, "attachment;filename=", "", -1)
	if fileName == "" {
		filePath := parsedUrl.Path
		fileName = path.Base(filePath)
	}

	return l.UploadDataFile(ctx, fileName, res.Body())
}

func (l *lighthouse) UploadDataFile(ctx context.Context, filename string, byteData []byte) (*Response, error) {
	var err error
	b := *bytes.NewBuffer(byteData)

	// Create a multipart writer to upload the file
	multipartBuffer := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(multipartBuffer)
	defer multipartWriter.Close()

	// Create the form file
	fw, err := multipartWriter.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	// Write the file data to the form file
	if _, err = fw.Write(b.Bytes()); err != nil {
		return nil, err
	}

	// Set the content type
	contentType := multipartWriter.FormDataContentType()

	// Initialize the HTTP client and make the request
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	r, err := client.R().SetContext(ctx).
		SetHeader("accept", "application/json").
		SetHeader("content-type", contentType).
		SetAuthScheme("Bearer").
		SetAuthToken(l.apiKey).
		SetBody(multipartBuffer).
		Post(UploadDataUrl)
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	if err := json.Unmarshal(r.Body(), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *lighthouse) FileInfo(ctx context.Context, cid string) (*FileInfo, error) {
	resp := &FileInfo{}
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	r, err := client.R().SetContext(ctx).
		SetQueryParam("cid", cid).SetResult(resp).
		Get(FileInfoUrl)
	if err != nil {
		return nil, err
	}

	if r.StatusCode() == http.StatusNotFound || resp.Cid == "" {
		return nil, errors.New("Not Found")
	}

	if resp.Error != nil && resp.Error.Code != 0 {
		return nil, errors.New(resp.Error.Message)
	}

	return resp, nil
}

func (l *lighthouse) DownloadFileByCid(ctx context.Context, cid string, rangeHeader string) ([]byte, error) {
	link := fmt.Sprintf("%s/%s", IPFSGatewayUrl, cid)
	return l.DownloadFile(ctx, link, "")
}

// downloadData fetches data from a URL with optional range header
func (l *lighthouse) DownloadFile(ctx context.Context, link string, rangeHeader string) ([]byte, error) {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if rangeHeader != "" {
		client.SetHeader("Range", rangeHeader)
	}
	resp, err := client.R().SetContext(ctx).Get(link)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("downloadData failed with unexpected status code %d for URL %s", resp.StatusCode(), link)
	}

	return resp.Body(), nil
}

func (l *lighthouse) extractCidFromByte(_ context.Context, byte []byte) (string, error) {
	bs := blockstore.NewIdStore(
		blockstore.NewBlockstore(dsync.MutexWrap(datastore.NewNullDatastore())),
	)

	bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)

	ufsImportParams := uih.DagBuilderParams{
		Maxlinks:   uih.BlockSizeLimit, // Default max of 174 links per block
		CidBuilder: cid.V0Builder{},
		Dagserv:    dsrv,
	}

	reader := bytes.NewReader(byte)
	ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(reader, chunker.DefaultBlockSize)) // 256KiB chunks
	if err != nil {
		return "", err
	}

	nd, err := balanced.Layout(ufsBuilder)
	if err != nil {
		return "", err
	}

	return nd.Cid().String(), nil
}

// fileExistOnNetwork checks if a file already exists on the network
func (l *lighthouse) fileExistOnNetwork(ctx context.Context, data []byte) (string, bool, error) {
	cid, err := l.extractCidFromByte(ctx, data)
	if err != nil {
		return "", false, err
	}

	fileInfo, err := l.FileInfo(ctx, cid)
	if err != nil {
		return "", false, err
	}

	if fileInfo != nil && fileInfo.Error.Code == http.StatusNotFound {
		return "", false, nil
	}

	return cid, true, nil
}
