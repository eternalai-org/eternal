package zip_hf_model_to_light_house

import (
	"bufio"
	"encoding/json"
	"eternal-infer-worker/libs/file"
	"eternal-infer-worker/libs/lighthouse"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var ZipChunkSize = 200 //MB
var BASH_EXEC = "/usr/bin/bash"
var THREADS = runtime.NumCPU() - 1

type HFModelZipFile struct {
	File   string `json:"file"`
	Hash   string `json:"hash"`
	Index  int
	Worker string
}
type HFModelInLightHouse struct {
	Model     string           `json:"model"`
	NumOfFile int              `json:"num_of_file"`
	Files     []HFModelZipFile `json:"files"`
}

func ExecuteCommand(fileCmd string) ([]byte, error) {
	commandId := strconv.FormatInt(time.Now().UnixMicro(), 10)
	fileLog := fmt.Sprintf("/tmp/log_%v.txt", commandId)
	execCmd := fmt.Sprintf("%v %v  2>&1 | /usr/bin/tee %v", BASH_EXEC, fileCmd, fileLog)
	fileExec := fmt.Sprintf("/tmp/bash_%v.sh", commandId)
	err := os.WriteFile(fileExec, []byte(execCmd), 0644)
	if err != nil {
		return nil, err
	}
	command := exec.Command(BASH_EXEC, fileExec)
	out, err := command.Output()
	if err != nil {
		return out, err
	}
	out, err = os.ReadFile(fileLog)
	log.Println("[ExecuteCommand][Debug] out: ", string(out))
	return out, err
}

func getScriptZipFile(modelFolder string, hfDir string) (string, error) {
	filePath := fmt.Sprintf("/tmp/hf-zip-model-%v.sh", modelFolder)
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			return "", fmt.Errorf("error removing file:%v", err)
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file:%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("cd %v \n ", hfDir))
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}
	_, err = file.WriteString(fmt.Sprintf("sudo rm -Rf %v.zip.part-* \n ", modelFolder))
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}
	_, err = file.WriteString(fmt.Sprintf("sudo tar -cf - %v | sudo pigz --best -p %v | sudo split -b %vM - %v.zip.part-", modelFolder, THREADS, ZipChunkSize, modelFolder))
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}
	return filePath, nil
}

func getScriptUnZipFile(modelFolder string, hfDir string, isZkSync bool) (string, error) {
	filePath := fmt.Sprintf("/tmp/hf-unzip-model-%v.sh", modelFolder)
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			return "", fmt.Errorf("error removing file:%v", err)
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file:%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("cd %v \n ", hfDir))
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}

	script := fmt.Sprintf("sudo cat %v.zip.part-* | sudo pigz -p %v -d | sudo tar -xf -", modelFolder, 2)
	if isZkSync {
		//use the new cmd
		script = fmt.Sprintf("sudo cat %v.zip.part* | pigz -d | docker load ", modelFolder)
	}

	log.Println("[getScriptUnZipFile]", "modelFolder: ", modelFolder, " ,hfDir: ", hfDir, " ,script: ", script)

	_, err = file.WriteString(script)
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}

	return filePath, nil
}

func getListZipFile(modelFolder string, hfDir string) ([]string, error) {
	filePath := fmt.Sprintf("/tmp/list-zip-model-%v.sh", modelFolder)
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			return nil, fmt.Errorf("error removing file:%v", err)
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error creating file:%v", err)
	}
	_, err = file.WriteString(fmt.Sprintf("rm /tmp/list_file_%v.txt \n", modelFolder))
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error write file:%v", err)
	}
	_, err = file.WriteString(fmt.Sprintf("cd %v \n", hfDir))
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error write file:%v", err)
	}
	_, err = file.WriteString(fmt.Sprintf("sudo ls %v.zip.part-* > /tmp/list_file_%v.txt \n", modelFolder, modelFolder))
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error write file:%v", err)
	}
	file.Close()

	output, err := ExecuteCommand(fmt.Sprintf("/tmp/list-zip-model-%v.sh ", modelFolder))
	if err != nil {
		return nil, fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	file, err = os.Open(fmt.Sprintf("/tmp/list_file_%v.txt", modelFolder))
	if err != nil {
		return nil, fmt.Errorf("error opening file:%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	listFile := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		listFile = append(listFile, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file:,%v", err)
	}
	return listFile, nil
}

func uploadHFModelResultToLightHouse(info *HFModelInLightHouse, apiKey string) (string, error) {
	data, _ := json.Marshal(info)
	cid, err := lighthouse.UploadData(apiKey, info.Model, data)
	if err != nil {
		return "", err
	}
	return cid, nil
}

func getHFModelResultFromLightHouse(hash string) (*HFModelInLightHouse, error) {
	data, _, err := lighthouse.DownloadDataSimple(hash)
	if err != nil {
		return nil, err
	}
	var result HFModelInLightHouse
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func downloadZipFileFromLightHouse(info *HFModelInLightHouse, hfDir string) error {
	input := make(chan HFModelZipFile)
	errors := make(chan error)

	log.Println("[downloadZipFileFromLightHouse]", "THREADS", THREADS)
	for i := 0; i < THREADS; i++ {
		go func() {
			for file := range input {
				log.Println("Start download ", "file", file.File, "hash", file.Hash, "hfDir", hfDir)
				var err error
				for i := 0; i < 10; i++ {
					err = lighthouse.DownloadToFile(file.Hash, fmt.Sprintf("%v/%v", hfDir, file.File))
					if err != nil {
						time.Sleep(2 * time.Minute)
						continue
					}
					break
				}
				if err != nil {
					log.Println("Error when try down file from light house", "file", file.File, "hash", file.Hash, "err", err.Error())
				} else {
					log.Println("Success download ", "file", file.File, "hash", file.Hash, "hfDir", hfDir)
				}
				errors <- err
			}
		}()
	}

	go func() {
		for _, file := range info.Files {
			input <- file
		}
	}()
	errCount := 0
	done := 0
	for err := range errors {
		if err != nil {
			errCount++
		}
		done++
		if done == len(info.Files) {
			close(input)
			close(errors)
			break
		}
	}
	if errCount > 0 {
		return fmt.Errorf("error when downloading zip file:%v, errCount/total : (%v/%v)", info.Model, errCount, done)
	}
	log.Println("success download all zip file:", "model", info.Model, "errCount", errCount, "total", done)
	return nil
}

type DownloadFileChan struct {
	Data HFModelZipFile
	Msg  *string
	Err  error
}

func DownloadHFModelFromLightHouse(hash string, hfDir string, isZkSync bool) ([]byte, error) {
	info, err := getHFModelResultFromLightHouse(hash)
	if err != nil {
		return nil, fmt.Errorf("error when get model info from light house hash : %v err :%v ", hash, err)
	}
	//err = downloadZipFileFromLightHouseNew(info, hfDir)
	err = downloadZipFileFromLightHouseConcurrentNew(info, hfDir) // use this concurrent instead
	if err != nil {
		return nil, fmt.Errorf("error when download zip chunk file:%v ", err)
	}
	scriptFile, err := getScriptUnZipFile(info.Model, hfDir, isZkSync)
	if err != nil {
		return nil, fmt.Errorf("error when get unzip script file:%v ", err)
	}

	log.Println("Start unzip list files: ", scriptFile)
	output, err := ExecuteCommand(scriptFile)
	if err != nil {
		return nil, fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	log.Println("Success unzip list files")
	return output, nil
}

func DownloadHFFile(wg *sync.WaitGroup, hfFile HFModelZipFile, modelPath string, dlChan chan DownloadFileChan) {
	defer wg.Done()
	downloadedFilePath := new(string)
	var err error
	_dl := ""

	defer func() {
		dlChan <- DownloadFileChan{
			Err:  err,
			Data: hfFile,
			Msg:  downloadedFilePath,
		}
	}()

	url := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hfFile.Hash)
	dest := fmt.Sprintf("%s/%s", modelPath, hfFile.File)

	f, err := os.Open(dest)
	if err == nil && f != nil {
		//skip if existed
		downloadedFilePath = &dest
		return
	}

	_dl, err = file.DownloadFile(url, dest)
	downloadedFilePath = &_dl

}

func downloadZipFileFromLightHouseNew(info *HFModelInLightHouse, hfDir string) error {
	//index downloaded files for sorting.
	if info.NumOfFile == 0 {
		return nil
	}

	errorNum := 0
	var wg sync.WaitGroup
	total := info.NumOfFile
	limit := 8 //8 process at once
	lastPageItems := total % limit

	retryDownload := []HFModelZipFile{}

	_t := float64(total) / float64(limit)
	totalPage := int(math.Ceil(_t))

	for page := 1; page <= totalPage; page++ {
		dlChan := make(chan DownloadFileChan)

		offset := (page - 1) * limit
		start := offset
		end := offset + limit
		if page != totalPage {
			wg.Add(limit)
		} else {
			end = offset + lastPageItems
			wg.Add(lastPageItems)
		}

		for i := start; i < end; i++ {
			go DownloadHFFile(&wg, info.Files[i], hfDir, dlChan)
		}

		for i := start; i < end; i++ {
			dFChan := <-dlChan
			if dFChan.Err != nil {
				errorNum++
				log.Println("[DownloadFile][Error] File: ", dFChan.Data.File, " ,error: ", dFChan.Err)
				retryDownload = append(retryDownload, dFChan.Data)

				//remove file if error
				os.Remove(fmt.Sprintf("%s/%s", hfDir, dFChan.Data.File))

			} else {
				log.Println("[DownloadFile][Success] File: ", dFChan.Data.File, " ,filePath: ", *dFChan.Msg)
			}
		}

		wg.Wait()
		time.Sleep(time.Second * 2)
	}

	if errorNum != 0 {
		retryInfo := HFModelInLightHouse{
			Model:     info.Model,
			NumOfFile: len(retryDownload),
			Files:     retryDownload,
		}

		return downloadZipFileFromLightHouseNew(&retryInfo, hfDir)
		//return errors.New("error while download file")
	}

	return nil
}

func _downloadHFFile(wg *sync.WaitGroup, input <-chan HFModelZipFile, hfDir string, output chan DownloadFileChan, wkName string) {
	defer wg.Done()

	var wg1 sync.WaitGroup
	for data := range input {
		//sleep 5 seconds before download
		time.Sleep(time.Second * 5)
		wg1.Add(1)
		data.Worker = wkName
		DownloadHFFile(&wg1, data, hfDir, output)
	}

	wg1.Wait()
}

func downloadZipFileFromLightHouseConcurrentNew(info *HFModelInLightHouse, hfDir string) error {
	workers := 8
	queueLimit := 20

	//index downloaded files for sorting.
	if info.NumOfFile == 0 {
		return nil
	}

	queue := make(chan HFModelZipFile, queueLimit)
	retryDownload := []HFModelZipFile{}
	errorNum := 0
	var wg sync.WaitGroup
	dlChan := make(chan DownloadFileChan, info.NumOfFile)

	//init workers
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		wkName := fmt.Sprintf("worker-%d", i)
		go _downloadHFFile(&wg, queue, hfDir, dlChan, wkName)
	}

	//push data to queue
	go func() {
		for i, data := range info.Files {
			data.Index = i
			queue <- data
		}
		close(queue)
	}()

	//listen data from outChan
	for range info.Files {
		dFChan := <-dlChan
		if dFChan.Err != nil {
			errorNum++
			log.Println("[DownloadFile][Error][", dFChan.Data.Worker, "] File: ", dFChan.Data.File, " ,error: ", dFChan.Err)
			retryDownload = append(retryDownload, dFChan.Data)

			//remove file if error
			os.Remove(fmt.Sprintf("%s/%s", hfDir, dFChan.Data.File))

		} else {
			log.Println("[DownloadFile][Success][", dFChan.Data.Worker, "] File: ", dFChan.Data.File, " ,filePath: ", *dFChan.Msg)
		}
	}

	wg.Wait()
	if errorNum != 0 {
		retryInfo := HFModelInLightHouse{
			Model:     info.Model,
			NumOfFile: len(retryDownload),
			Files:     retryDownload,
		}

		return downloadZipFileFromLightHouseNew(&retryInfo, hfDir)
		//return errors.New("error while download file")
	}
	return nil
}
