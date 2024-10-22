package manager

import (
	"encoding/json"
	"errors"
	"eternal-infer-worker/libs/dockercmd"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/file"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/libs/zip_hf_model_to_light_house"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/client"
)

func checkModelFileExist(filePath string) bool {
	return file.CheckFileExist(filePath)
}

func isMultiPartsModelURL(url string) bool {
	if strings.Contains(url, ",") {
		return true
	}
	return false
}

func downloadFile(url, dest string) (string, error) {
	log.Info("[downloadFile] url: ", url, " ,dest: ", dest)
	return file.DownloadChunkedDataDest(url, dest)
}

func downloadSingleFile(url, dest string) (string, error) {
	log.Info("[downloadSingleFile] url: ", url, " ,dest: ", dest)
	return file.DownloadFile(url, dest)
}

func downloadMultiPartsModelDest(url, path, filename string) (string, error) {
	urls := strings.Split(url, ",")
	var filePaths []string
	for idx, u := range urls {
		partPath := path + "/" + filename + "_" + fmt.Sprintf("%d", idx)
		filePath, err := downloadFile(lighthouse.IPFSGateway+u, partPath)
		if err != nil {
			return "", err
		}
		filePaths = append(filePaths, filePath)
	}

	err := file.MergeFiles(filePaths, path+"/"+filename)
	if err != nil {
		return "", err
	}

	return path + "/" + filename, nil
}

func (m *ModelInstance) SetupDocker() error {
	/*
		t := time.Now()
		defer func() {
			log.Info(fmt.Sprintf("[SetupDocker] - %v took %v \n", m.ModelInfo.ModelAddr, time.Since(t)))
		}()*/

	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	if !m.TrainingRequest.ZKSync {
		filePath := ""
		var err error
		targetImageName := m.ModelInfo.ModelAddr
		fileExistValid := false
		log.Info("[SetupDocker] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " ,targetImageName: ", targetImageName)
		exist := checkModelFileExist(m.ModelPath + "/model.zip")
		if exist {
			log.Info("[SetupDocker] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " ,targetImageName: ", targetImageName, " ,exist: ", exist)

			match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.ModelFileHash, m.ModelPath+"/model.zip")
			if err != nil {
				log.Error("[SetupDocker] Check model file hash got error", err)
				return err
			}

			log.Info("[SetupDocker] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " ,targetImageName: ", targetImageName, " ,exist: ", exist, " ,match: ", match)
			if match {
				fileExistValid = true
				filePath = m.ModelPath + "/model.zip"
			} else {
				err = file.RemoveFile(m.ModelPath + "/model.zip")
				if err != nil {
					log.Error("[SetupDocker][Err] Remove model file got error", err)
					return err
				}

				err1 := dockercmd.RemoveImage(targetImageName)
				if err1 != nil {
					log.Error("[SetupDocker][Err] RemoveImage err: ", err)
				}
			}
		}

		if !fileExistValid {
			isMultiParts := isMultiPartsModelURL(m.ModelInfo.Metadata.ModelURL)
			urls := m.ModelInfo.Metadata.ModelURL
			if strings.Contains(urls, "ipfs://") {
				urls = strings.ReplaceAll(m.ModelInfo.Metadata.ModelURL, "ipfs://", lighthouse.IPFSGateway)
			}

			if isMultiParts {
				filePath, err = downloadMultiPartsModelDest(urls, m.ModelPath, "model.zip")
				if err != nil {
					log.Error("[SetupDocker][Err] Download multi parts model got error", err)
					return err
				}
			} else {
				log.Info("[SetupDocker][Err] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " ,targetImageName: ", targetImageName, " ,exist: ", exist, " , Downloading")

				//filePath, err = downloadFile(urls, m.ModelPath+"/model.zip") //old
				filePath, err = downloadSingleFile(urls, m.ModelPath+"/model.zip") //new
				if err != nil {
					log.Error("[SetupDocker][Err] Download file with IPFS gateway got error", err)
					return err
				}
			}

			start := time.Now().UTC()
			log.Info("[SetupDocker][DEBUG] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " , Starting...")
			match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.ModelFileHash, filePath)
			log.Info("[SetupDocker][DEBUG] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " , End: ", time.Now().UTC().Sub(start).Seconds(), " seconds")

			if err != nil {
				log.Error("[SetupDocker][Err] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " ,targetImageName: ", targetImageName, " ,exist: ", exist, " , error ", err)
				return err
			}

			if !match {
				err = errors.New("Model file hash not match")
				log.Error("[SetupDocker][Err] - checkModelFileExist: ", m.ModelPath+"/model.zip", " ,ModelFileHash: ", m.ModelInfo.Metadata.ModelFileHash, " ,targetImageName: ", targetImageName, " ,exist: ", exist, " , error ", err)
				return err
			}
		} else {
			log.Warning("[SetupDocker] Model file already exist")
		}

		err = dockercmd.LoadLocalImageWithCustomName(filePath, targetImageName)
		if err != nil {
			log.Error("[SetupDocker][Err]  Load local image got error", err)
			return err
		}

		m.Loaded = true
		log.Info("[SetupDocker] - loaded - filePath", filePath)
	} else {
		//check image existed or not
		img, _ := dockercmd.GetImageInfo(m.ModelInfo.ModelAddr)

		//log.Info("[SetupDocker]  GetImageInfo: ", m.ModelInfo.ModelAddr, " ,img: ", img.ID)
		if img.ID == "" {
			temp := strings.Split(m.ModelInfo.Metadata.ModelURL, "/")
			hash := temp[len(temp)-1]
			out, err := zip_hf_model_to_light_house.DownloadHFModelFromLightHouse(hash, m.ModelPath, m.TrainingRequest.ZKSync)
			if err != nil {
				log.Error("[SetupDocker][Err]  Download model zkchain got error", err)
				return err
			}

			//rename docker images from real name to model-address, for convenient with our flow. EX:  nikolasigmoid/flux-black-forest ->0x9874732a8699fca824a9a7d948f6bcd30a141238
			msg := string(out)
			if !strings.Contains(msg, "Loaded image:") {
				str := fmt.Sprintf("[SetupDocker][Err] cannot get image name from: \"%s\"", msg)
				log.Warning(str)
				err = errors.New(str)
				return err
			}

			imageName := strings.TrimSpace(strings.ReplaceAll(msg, "Loaded image:", ""))
			if imageName == "" {
				str := fmt.Sprintf("[SetupDocker][Err] cannot get image name from string: %s", imageName)
				log.Warning(str)
				err = errors.New(str)
				return err
			}

			//rename image to: model_address
			imageNameTag := strings.Split(imageName, ":")
			if len(imageNameTag) < 2 {
				str := fmt.Sprintf("[SetupDocker][Err] cannot get image name + tag: %s", imageName)
				log.Warning(str)
				err = errors.New(str)
				return err
			}

			//oldName := imageNameTag[0]
			//tag := imageNameTag[1]

			nameName := fmt.Sprintf("%s:%s", m.ModelInfo.ModelAddr, "latest")
			err = dockercmd.RenameImage(imageName, nameName)
			if err != nil {
				log.Error("[SetupDocker][Err] cannot update image name", err)
				return err
			}

			log.Info("[SetupDocker][Success]  loaded and changed image name from: ", imageName, " ,to: ", nameName)
		}

	}

	return nil
}

func (m *ModelInstance) SetupDockerVerifier() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()
	filePath := ""
	var err error

	fileExistValid := false
	exist := checkModelFileExist(m.ModelPath + "/verifier.zip")
	if exist {
		match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.VerifierFileHash, m.ModelPath+"/verifier.zip")
		if err != nil {
			log.Error("Check model file hash got error", err)
			return err
		}

		if match {
			fileExistValid = true
			filePath = m.ModelPath + "/verifier.zip"
		}
	}

	if !fileExistValid {
		isMultiParts := isMultiPartsModelURL(m.ModelInfo.Metadata.VerifierURL)
		urls := m.ModelInfo.Metadata.VerifierURL
		if strings.Contains(urls, "ipfs://") {
			urls = strings.ReplaceAll(m.ModelInfo.Metadata.VerifierURL, "ipfs://", lighthouse.IPFSGateway)
		}

		if isMultiParts {
			filePath, err = downloadMultiPartsModelDest(urls, m.ModelPath, "verifier.zip")
			if err != nil {
				log.Error("Download multi parts model got error", err)
				return err
			}
		} else {
			filePath, err = downloadFile(urls, m.ModelPath+"/verifier.zip")
			if err != nil {
				log.Error("Download file with IPFS gateway got error", err)
				return err
			}
		}

		match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.VerifierFileHash, filePath)
		if err != nil {
			log.Error("Check model file hash got error", err)
			return err
		}

		if !match {
			log.Error("Model file hash not match")
			return err
		}
	} else {
		log.Warning("Model file already exist")
	}

	err = dockercmd.LoadLocalImageWithCustomName(filePath, m.ModelInfo.ModelAddr+"_validator")
	if err != nil {
		log.Error("Load local image got error", err)
		return err
	}

	m.VerifierLoaded = true

	return nil
}

func (m *ModelInstance) StartDocker() error {
	//if m.TrainingRequest["ZKSync"] == true {
	//	// TODO
	//} else {}
	var err error
	t := time.Now()
	defer func() {

		if err != nil {
			log.Error(fmt.Sprintf("[StartDocker] ModelAddress: %v, port: %s, DisableGPU: %v,  containerID: %s, took %v \n", m.ModelInfo.ModelAddr, m.Port, m.DisableGPU, m.containerID, time.Since(t)))
		} else {
			log.Info(fmt.Sprintf("[StartDocker] ModelAddress: %v, port: %s, DisableGPU: %v,  containerID: %s, took %v  \n", m.ModelInfo.ModelAddr, m.Port, m.DisableGPU, m.containerID, time.Since(t)))
		}
	}()

	existedContainer, err := dockercmd.GetContainerByName(m.ModelInfo.ModelAddr)
	if err != nil {
		log.Error(fmt.Sprintf("[StartDocker][ERR][GetContainerByName] ModelAddress: %v, DisableGPU: %v,  err: %v  \n", m.ModelInfo.ModelAddr, m.DisableGPU, err))
		return err
	}

	if existedContainer != nil {
		log.Infof("[StartDocker][DEBUG][GetContainerByName] container: %v, status: %s  \n", m.ModelInfo.ModelAddr, existedContainer.Status)
		if strings.Contains(existedContainer.Status, "up") || strings.Contains(existedContainer.Status, "running") {
			m.Ready = true
			return nil
		}
	}

	m.actionLock.Lock()
	defer m.actionLock.Unlock()
	resultMountDir := filepath.Join("/home/eternal_ai/infer-results/" + m.ModelInfo.ModelAddr)
	err = os.MkdirAll(resultMountDir, os.ModePerm)
	if err != nil {
		log.Printf("[StartDocker][ERR] ModelAddress: %v, resultMountDir: %s,  err: %v  \n", m.ModelInfo.ModelAddr, resultMountDir, err)
		return err
	}

	//log.Printf("[StartDocker][DEBUG][CreateAndStartContainer] ModelAddress: %v, resultMountDir: %s  \n", m.ModelInfo.ModelAddr, resultMountDir)
	//TODO - get image's name
	ctnInfo, err := dockercmd.CreateAndStartContainer(m.ModelInfo.ModelAddr, m.ModelInfo.ModelAddr, m.Port, resultMountDir, m.DisableGPU)
	if err != nil {
		log.Errorf("[StartDocker][ERR][CreateAndStartContainer] ModelAddress: %v, resultMountDir: %s, DisableGPU: %v,  err: %v  \n", m.ModelInfo.ModelAddr, resultMountDir, m.DisableGPU, err)
		return err
	}

	//get container again, because it would be nil if it was not started.
	if existedContainer == nil {
		existedContainer, err = dockercmd.GetContainerByName(m.ModelInfo.ModelAddr)
		if err != nil {
			log.Error(fmt.Sprintf("[StartDocker][ERR][GetContainerByName] ModelAddress: %v, DisableGPU: %v,  err: %v  \n", m.ModelInfo.ModelAddr, m.DisableGPU, err))
			return err
		}
	}

	m.Port = fmt.Sprintf("%v", existedContainer.Ports[0].PublicPort)
	m.containerID = ctnInfo.ID
	m.ResultDir = resultMountDir

	//log.Infof("[StartDocker][DEBUG][WaitForContainerToReady] ModelAddress: %v, resultMountDir: %s, port: %s, DisableGPU: %v,  containerID: %s  \n", m.ModelInfo.ModelAddr, resultMountDir, m.Port, m.DisableGPU, m.containerID)
	err = dockercmd.WaitForContainerToReady(m.containerID)
	if err != nil {
		log.Error(fmt.Sprintf("[StartDocker][ERR][WaitForContainerToReady] ModelAddress: %v, resultMountDir: %s, port: %s, DisableGPU: %v, containerID: %s, err: %v  \n", m.ModelInfo.ModelAddr, resultMountDir, m.Port, m.DisableGPU, m.containerID, err))
		return err
	}

	m.Ready = true
	return nil
}

func (m *ModelInstance) StartDockerVerifier() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	resultMountDir := filepath.Join(getCurrentDir(), "verify-results/"+m.ModelInfo.ModelAddr)

	err := os.MkdirAll(resultMountDir, os.ModePerm)
	if err != nil {
		log.Error("Create model mount dir got error", err)
		return err
	}
	containerName := m.ModelInfo.ModelAddr + "_validator"

	ctnInfo, err := dockercmd.CreateAndStartContainer(containerName, containerName, m.Port, resultMountDir, m.DisableGPU)
	if err != nil {
		log.Error("Create and start container got error", err)
		return err
	}

	log.Info("Container ID:", ctnInfo.ID)
	m.verifierContainerID = ctnInfo.ID
	m.VerifyDir = resultMountDir

	m.VerifierReady = true

	return nil
}

func (m *ModelInstance) PauseDocker() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	ctnInfo, err := dockercmd.GetContainerByName(m.ModelInfo.ModelAddr)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}
	if ctnInfo == nil {
		return nil
	}
	m.containerID = ctnInfo.ID

	err = dockercmd.PauseContainer(m.containerID)
	if err != nil {
		log.Error("Pause container got error", err)
		return err
	}

	m.Ready = false
	return nil
}

func (m *ModelInstance) PauseVerifierDocker() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	ctnInfo, err := dockercmd.GetContainerByName(m.ModelInfo.ModelAddr + "_validator")
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}
	if ctnInfo == nil {
		return nil
	}
	m.verifierContainerID = ctnInfo.ID

	err = dockercmd.PauseContainer(m.verifierContainerID)
	if err != nil {
		log.Error("Pause container got error", err)
		return err
	}
	m.VerifierReady = false
	return nil
}

func (m *ModelInstance) UnpauseDocker() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	containerStatus, err := dockercmd.GetContainerInfo(m.containerID)
	if err != nil {
		log.Error("Get container info got error", err)
		if !client.IsErrNotFound(err) {
			return err
		}
	}
	if containerStatus != nil {

		switch containerStatus.State.Status {
		case "running":
			return nil
		case "exited", "dead":
			err = dockercmd.RemoveContainer(m.containerID)
			if err != nil {
				log.Error("Remove container got error", err)
				return err
			}
		case "paused":
			err = dockercmd.UnpauseContainer(m.containerID)
			if err != nil {
				log.Error("Unpause container got error", err)
				return err
			}
			m.Ready = true
		}

	} else {

	}

	return nil
}

func (m *ModelInstance) UnpauseVerifierDocker() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	containerStatus, err := dockercmd.GetContainerInfo(m.verifierContainerID)
	if err != nil {
		log.Error("Get container info got error", err)
		if !client.IsErrNotFound(err) {
			return err
		}
	}
	if containerStatus != nil {

		switch containerStatus.State.Status {
		case "running":
			return nil
		case "exited", "dead":
			err = dockercmd.RemoveContainer(m.verifierContainerID)
			if err != nil {
				log.Error("Remove container got error", err)
				return err
			}
		case "paused":
			err = dockercmd.UnpauseContainer(m.verifierContainerID)
			if err != nil {
				log.Error("Unpause container got error", err)
				return err
			}
			m.VerifierReady = true
		}

	}

	return nil
}

func (m *ModelInstance) RemoveDocker() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	err := dockercmd.RemoveContainer(m.containerID)
	if err != nil {
		log.Error("Remove container got error", err)
		return err
	}

	m.Ready = false
	return nil
}

func (m *ModelInstance) RemoveVerifierDocker() error {
	m.actionLock.Lock()
	defer m.actionLock.Unlock()

	err := dockercmd.RemoveContainer(m.verifierContainerID)
	if err != nil {
		log.Error("Remove container got error", err)
		return err
	}

	m.VerifierReady = false
	return nil
}

func (m *ModelInstance) IsRunning() (bool, error) {
	containerStatus, err := dockercmd.GetContainerInfo(m.containerID)
	if err != nil {
		log.Error("Get container info got error", err)
		if !client.IsErrNotFound(err) {
			return false, err
		}
	}
	if containerStatus != nil {
		if containerStatus.State.Status == "running" {
			return true, nil
		}
	}
	return false, nil
}

func (m *ModelInstance) IsVerifierRunning() (bool, error) {
	containerStatus, err := dockercmd.GetContainerInfo(m.verifierContainerID)
	if err != nil {
		log.Error("Get container info got error", err)
		if !client.IsErrNotFound(err) {
			return false, err
		}
	}
	if containerStatus != nil {
		if containerStatus.State.Status == "running" {
			return true, nil
		}
	}
	return false, nil
}

func (m *ModelInstance) GetTrainingRequest() error {
	url := fmt.Sprintf("%s/api/dojo/model-info-by-model-address/%s", "https://api-dojo2.eternalai.org", m.ModelInfo.ModelAddr)
	log.Info("[GetTrainingRequest] - url: ", url)

	response, err := http.Get(url)
	if err != nil {
		log.Error("[GetTrainingRequest][Error] - url: ", url, " ,err: ", err)
		return err
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("[GetTrainingRequest][Error] - url: ", url, " ,err: ", err)
		return err
	}

	data := &APIResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Error("[GetTrainingRequest][Error] - url: ", url, " ,err: ", err)
		return err
	}

	m.TrainingRequest = &data.Data
	//log.Info("[GetTrainingRequest][Success] - url: ", url, " ,TrainingRequest: ", m.TrainingRequest)
	return nil
}
