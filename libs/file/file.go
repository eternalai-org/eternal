package file

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

func UnzipFile(zipFilePath string, dst string) error {
	archive, err := zip.OpenReader(zipFilePath)
	if err != nil {
		log.Println("Error opening zip file...", err)
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		log.Println("Unzipping file... ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return errors.New("invalid file path")
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			log.Println("Error creating directory...", err)
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			log.Println("Error opening file...", err)
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			log.Println("Error opening file in archive", err)
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			log.Println("Error copying file to destination...", err)
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	log.Println("Done unzip file!")
	return nil
}

func getFileSizeFromLink(link string) (int64, error) {
	resp, err := http.Head(link)
	if err != nil {
		return 0, err
	}
	if resp.Header.Get("Content-Length") != "" {
		size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			return 0, err
		}
		return size, nil
	}

	contentRange := resp.Header.Get("Content-Range")

	if contentRange == "" {
		return 0, errors.New("Content-Range not found")
	}

	parts := strings.Split(contentRange, "/")
	if len(parts) < 2 {
		return 0, errors.New("Content-Range not found")
	}

	size, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func DownloadChunkedDataDest(link string, dest string) (string, error) {
	// check folder exist first, if not create
	destDir := filepath.Dir(dest)
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileSize, err := getFileSizeFromLink(link)
	if err != nil {
		return "", err
	}

	filePath := path.Join(dest)
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	var chunkSize int64 = 1024 * 1024 * 10 // 10MB
	var start int64 = 0
	var end int64 = chunkSize

	parts := fileSize/chunkSize + 1
	log.Println("parts: ", parts)
	part := int64(0)

	maxConcurrentChunk := int64(20)

	_ = maxConcurrentChunk
	_ = part
	_ = end
	_ = start

	downloadedSize := int64(0)

	for downloadedSize < fileSize {
		var wg sync.WaitGroup

		dataBuffer := make(map[int64][]byte)
		dataBufferLck := sync.Mutex{}

		chunkLeft := parts - part
		if chunkLeft < maxConcurrentChunk {
			maxConcurrentChunk = chunkLeft
		}
		if maxConcurrentChunk == 0 {
			break
		}
		startPart := part

		log.Println("maxConcurrentChunk: ", maxConcurrentChunk)
		wg.Add(int(maxConcurrentChunk))
		for i := int64(0); i < maxConcurrentChunk; i++ {
			go func(fpart int64) {
				defer wg.Done()

				start := fpart * chunkSize
				end := start + chunkSize
				if end > fileSize {
					end = fileSize
				}

				log.Println("part: ", fpart, "/", parts, end-start, "bytes", start, end)
				data, err := downloadChunkedData(link, int(start), int(end)-1)
				if err != nil {
					log.Println("Download chunked data got error", err)
					return
				}
				if len(data) == 0 {
					log.Println("Empty data", fpart)
					return
				}

				dataBufferLck.Lock()
				dataBuffer[fpart] = data
				dataBufferLck.Unlock()

			}(part)

			part++
		}

		wg.Wait()

		var data []byte

		if len(dataBuffer) == 0 {
			return "", errors.New("No data downloaded")
		}

		if len(dataBuffer) != int(maxConcurrentChunk) {
			return "", errors.New("Missing data")
		}

		for i := int64(startPart); i < part; i++ {
			if len(dataBuffer[i]) == 0 {
				return "", errors.New("Empty data")
			}
			data = append(data, dataBuffer[i]...)
		}

		err = WriteFile(filePath, data, os.ModePerm)
		if err != nil {
			return "", err
		}

		downloadedSize += int64(len(data))

		log.Println("downloadedSize: ", downloadedSize, "fileSize: ", fileSize)

		// ----

		// if end > fileSize {
		// 	end = fileSize
		// }

		// log.Println("part: ", part, "/", parts)
		// data, err := downloadChunkedData(link, int(start), int(end))
		// if err != nil {
		// 	return "", err
		// }

		// err = WriteFile(filePath, data, os.ModePerm)
		// if err != nil {
		// 	return "", err
		// }

		// start = end + 1
		// end += chunkSize

		// part++

	}

	return filePath, nil
}

func DownloadChunkedData(link string, dir string) (string, error) {
	// check folder exist first, if not create
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileSize, err := getFileSizeFromLink(link)
	if err != nil {
		return "", err
	}

	filePath := path.Join(dir, fmt.Sprintf("model.zip"))
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	var chunkSize int64 = 1024 * 1024 * 10 // 10MB
	var start int64 = 0
	var end int64 = chunkSize

	// parts := fileSize / chunkSize
	// fmt.Println("parts: ", parts)
	part := 0
	for start < fileSize {
		if end > fileSize {
			end = fileSize
		}

		// fmt.Println("part: ", part, "start: ", start, "end: ", end)
		data, err := downloadChunkedData(link, int(start), int(end))
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

func downloadChunkedData(link string, start, end int) ([]byte, error) {
	urlLink := fmt.Sprintf("%s", link)

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

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
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

func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func MergeFiles(files []string, output string) error {
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, file := range files {
		in, err := os.Open(file)
		if err != nil {
			return err
		}

		_, err = io.Copy(out, in)
		in.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
