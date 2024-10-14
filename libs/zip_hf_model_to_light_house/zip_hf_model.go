package zip_hf_model_to_light_house

import (
	"bufio"
	"encoding/json"
	"eternal-infer-worker/libs/lighthouse"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var ZipChunkSize = 200 //MB
var BASH_EXEC = "/usr/bin/bash"
var THREADS = runtime.NumCPU() - 1

type HFModelZipFile struct {
	File string `json:"file"`
	Hash string `json:"hash"`
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
		script = fmt.Sprintf("sudo cat %v.zip.part* | pigz -d | docker load", modelFolder)
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

func DownloadHFModelFromLightHouse(hash string, hfDir string, isZkSync bool) error {
	info, err := getHFModelResultFromLightHouse(hash)
	if err != nil {
		return fmt.Errorf("error when get model info from light house hash : %v err :%v ", hash, err)
	}
	err = downloadZipFileFromLightHouse(info, hfDir)
	if err != nil {
		return fmt.Errorf("error when download zip chunk file:%v ", err)
	}
	scriptFile, err := getScriptUnZipFile(info.Model, hfDir, isZkSync)
	if err != nil {
		return fmt.Errorf("error when get unzip script file:%v ", err)
	}
	log.Println("Start unzip list files")
	output, err := ExecuteCommand(scriptFile)
	if err != nil {
		return fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}
	log.Println("Success unzip list files")
	return nil
}
