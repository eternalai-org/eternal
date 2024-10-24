package file

import (
	"archive/zip"
	"eternal-infer-worker/libs"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
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
	log.Println("[DownloadChunkedDataDest] - link: ", link, " ,dest: ", dest)

	// check folder exist first, if not create
	destDir := filepath.Dir(dest)
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		log.Println("[DownloadChunkedDataDest][Err][MkdirAll] ", link, " ,dest: ", dest, " ,err: ", err)
		return "", err
	}

	fileSize, err := getFileSizeFromLink(link)
	if err != nil {
		log.Println("[DownloadChunkedDataDest][Err][getFileSizeFromLink] ", link, " ,dest: ", dest, " ,err: ", err)
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
	//log.Println("parts: ", parts)
	part := int64(0)

	log.Println("[DownloadChunkedDataDest] - link: ", link, ", fileSize: ", fileSize, " ,chunkSize: ", chunkSize, " ,start: ", start, " ,end: ", end, " ,parts: ", parts, " ,part: ", part)
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

		log.Println("[DownloadChunkedDataDest] - link: ", link, "maxConcurrentChunk: ", maxConcurrentChunk)
		if maxConcurrentChunk == 0 {
			break
		}
		startPart := part

		//log.Println("maxConcurrentChunk: ", maxConcurrentChunk)
		wg.Add(int(maxConcurrentChunk))
		for i := int64(0); i < maxConcurrentChunk; i++ {
			go func(fpart int64) {
				defer wg.Done()

				start := fpart * chunkSize
				end := start + chunkSize
				if end > fileSize {
					end = fileSize
				}

				//log.Println("part: ", fpart, "/", parts, end-start, "bytes", start, end)
				log.Println("[DownloadChunkedDataDest] - download  - link: ", link, " ,fpart: ", fpart, "bytes: ", start, "-", end)
				data, err := downloadChunkedData(link, int(start), int(end)-1)
				if err != nil {
					//log.Println("Download chunked data got error", err)
					log.Println("[DownloadChunkedDataDest][downloadChunkedData][Err] - download  - link: ", link, " ,fpart: ", fpart, "bytes: ", start, "-", end, " ,error: ", err)
					return
				}
				if len(data) == 0 {
					log.Println("[DownloadChunkedDataDest][downloadChunkedData][Err] - download  - link: ", link, " ,fpart: ", fpart, "bytes: ", start, "-", end, " ,error: empty part")
					return
				}

				dataBufferLck.Lock()
				dataBuffer[fpart] = data
				dataBufferLck.Unlock()
				log.Println("[DownloadChunkedDataDest] - download  - link: ", link, "- done!!")

			}(part)

			part++
		}

		wg.Wait()

		var data []byte

		if len(dataBuffer) == 0 {
			err = errors.New("No data downloaded")

			log.Println("[DownloadChunkedDataDest][downloadChunkedData][Err] - download  - link: ", link, " ,error: ", err)
			return "", err
		}

		if len(dataBuffer) != int(maxConcurrentChunk) {
			err = errors.New("Missing data")
			log.Println("[DownloadChunkedDataDest][downloadChunkedData][Err] - download  - link: ", link, " ,error: ", err)
			return "", err
		}

		for i := int64(startPart); i < part; i++ {
			if len(dataBuffer[i]) == 0 {
				return "", errors.New("Empty data")
			}
			data = append(data, dataBuffer[i]...)
		}

		err = WriteFile(filePath, data, os.ModePerm)
		if err != nil {
			log.Println("[DownloadChunkedDataDest][downloadChunkedData][Err] - download  - link: ", link, " ,error: ", err)
			return "", err
		}

		downloadedSize += int64(len(data))
		log.Println("[DownloadChunkedDataDest] - link: ", link, " ,fileSize/ downloadedSize", fileSize, "/", downloadedSize)

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

func DownloadFile(link string, dest string) (string, error) {
	log.Println("[DownloadFile] - link: ", link, " ,dest: ", dest)

	// check folder exist first, if not create
	destDir := filepath.Dir(dest)
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		log.Println("[DownloadFile][Err][MkdirAll] ", link, " ,dest: ", dest, " ,err: ", err)
		return "", err
	}

	filePath := path.Join(dest)
	if _, err = os.Stat(filePath); err == nil {
		return filePath, nil
	}

	//download here
	resp, err := http.Get(link)
	if err != nil {
		log.Println("[DownloadFile][Err][http.Get] ", link, " ,dest: ", dest, " ,err: ", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("download get http code: %d", resp.StatusCode))
		log.Println("[DownloadFile][Err][http.Get] ", link, " ,dest: ", dest, " ,err: ", err)
		return "", err
	}

	//end download here
	out, err := os.Create(filePath)
	if err != nil {
		log.Println("[DownloadFile][Err][os.Create] ", link, " ,dest: ", dest, " ,filePath: ", filePath, " ,err: ", err)
		return "", err
	}

	defer out.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Println("[DownloadFile][Err][io.Copy] ", link, " ,dest: ", dest, " ,filePath: ", filePath, " ,err: ", err)
		return "", err
	}

	log.Println("[DownloadFile] - link: ", link, " ,dest: ", dest, " ,filePath: ", filePath, " ,n: ", n)
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

func ReadFile(fn string) ([]byte, error) {
	_b, err := os.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return _b, nil
}

func WillUpdateVersion(releaseVersion string) (bool, error) {
	willUpdate := false
	_b, err := ReadFile(libs.VERSION_FILENAME)
	if err != nil {
		//code will be updated if log is not created.
		return true, err
	}

	str := string(_b)
	if !strings.EqualFold(releaseVersion, str) {
		willUpdate = true
	}

	return willUpdate, nil
}

func UpdateVersionLog(releaseVersion string) error {
	f, err := os.Create(libs.VERSION_FILENAME)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(releaseVersion)
	if err != nil {
		return err
	}

	return nil
}
