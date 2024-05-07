package lighthouse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
	"github.com/pkg/errors"
)

type Response struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

const (
	IPFSGateway = "https://gateway.lighthouse.storage/ipfs/"
)

func DownloadDataSimple(hash string) ([]byte, string, error) {
	// https://gateway.lighthouse.storage/ipfs/QmXPGcEHCi1ZmbHFwScuP4ZJ2iv9YjTJMUroUTJUnFXxxj
	urlLink := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash)

	resp, err := http.Get(urlLink)
	if err != nil {
		return nil, "", errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", errors.WithStack(err)
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

	return &respBody, nil
}

func UploadData(apikey, fileName string, data []byte) (string, error) {
	urlLink := "https://node.lighthouse.storage/api/v0/add"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}
	if _, err = fw.Write(data); err != nil {
		return "", err
	}

	w.Close()

	req, err := http.NewRequest(
		"POST",
		urlLink,
		&b,
	)

	if err != nil {
		return "", errors.WithStack(err)
	}

	client := &http.Client{}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", w.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apikey))

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	log.Println("body", string(body))

	var respBody Response

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return respBody.Hash, nil
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
