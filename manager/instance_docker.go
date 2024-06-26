package manager

import (
	"eternal-infer-worker/libs/dockercmd"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/file"
	"eternal-infer-worker/libs/lighthouse"
	"fmt"
	"log"
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
	return file.DownloadChunkedDataDest(url, dest)
}

func downloadMultiPartsModelDest(url, path, filename string) (string, error) {
	urls := strings.Split(url, ",")
	var filePaths []string
	for idx, u := range urls {
		partPath := path + "/" + filename + "_" + string(idx)
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
	t := time.Now()
	defer func() {
		log.Printf("SetupDocker %v took %v \n", m.ModelInfo.ModelAddr, time.Since(t))
	}()
	m.actionLock.Lock()
	defer m.actionLock.Unlock()
	filePath := ""
	var err error
	targetImageName := m.ModelInfo.ModelAddr
	fileExistValid := false
	exist := checkModelFileExist(m.ModelPath + "/model.zip")
	if exist {
		match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.ModelFileHash, m.ModelPath+"/model.zip")
		if err != nil {
			log.Println("Check model file hash got error", err)
			return err
		}

		if match {
			fileExistValid = true
			filePath = m.ModelPath + "/model.zip"
		} else {
			err = file.RemoveFile(m.ModelPath + "/model.zip")
			if err != nil {
				log.Println("Remove model file got error", err)
				return err
			}
			dockercmd.RemoveImage(targetImageName)
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
				log.Println("Download multi parts model got error", err)
				return err
			}
		} else {
			filePath, err = downloadFile(urls, m.ModelPath+"/model.zip")
			if err != nil {
				log.Println("Download file with IPFS gateway got error", err)
				return err
			}
		}

		match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.ModelFileHash, filePath)
		if err != nil {
			log.Println("Check model file hash got error", err)
			return err
		}

		if !match {
			log.Println("Model file hash not match")
			return err
		}
	} else {
		log.Println("Model file already exist")
	}

	err = dockercmd.LoadLocalImageWithCustomName(filePath, targetImageName)
	if err != nil {
		log.Println("Load local image got error", err)
		return err
	}

	m.Loaded = true

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
			log.Println("Check model file hash got error", err)
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
				log.Println("Download multi parts model got error", err)
				return err
			}
		} else {
			filePath, err = downloadFile(urls, m.ModelPath+"/verifier.zip")
			if err != nil {
				log.Println("Download file with IPFS gateway got error", err)
				return err
			}
		}

		match, err := eaimodel.CheckModelFileHash(m.ModelInfo.Metadata.VerifierFileHash, filePath)
		if err != nil {
			log.Println("Check model file hash got error", err)
			return err
		}

		if !match {
			log.Println("Model file hash not match")
			return err
		}
	} else {
		log.Println("Model file already exist")
	}

	err = dockercmd.LoadLocalImageWithCustomName(filePath, m.ModelInfo.ModelAddr+"_validator")
	if err != nil {
		log.Println("Load local image got error", err)
		return err
	}

	m.VerifierLoaded = true

	return nil
}

func (m *ModelInstance) StartDocker() error {
	t := time.Now()
	defer func() {
		log.Printf("StartDocker %v took %v \n", m.ModelInfo.ModelAddr, time.Since(t))
	}()
	m.actionLock.Lock()
	defer m.actionLock.Unlock()
	resultMountDir := filepath.Join(getCurrentDir(), "infer-results/"+m.ModelInfo.ModelAddr)

	err := os.MkdirAll(resultMountDir, os.ModePerm)
	if err != nil {
		log.Println("Create model mount dir got error", err)
		return err
	}

	ctnInfo, err := dockercmd.CreateAndStartContainer(m.ModelInfo.ModelAddr, m.ModelInfo.ModelAddr, m.Port, resultMountDir, m.DisableGPU)
	if err != nil {
		log.Println("Create and start container got error", err)
		return err
	}

	existedContainer, err := dockercmd.GetContainerByName(m.ModelInfo.ModelAddr)
	if err != nil {
		return err
	}
	m.Port = fmt.Sprintf("%v", existedContainer.Ports[0].PublicPort)

	log.Println("Container ID:", ctnInfo.ID)
	m.containerID = ctnInfo.ID
	m.ResultDir = resultMountDir

	err = dockercmd.WaitForContainerToReady(m.containerID)
	if err != nil {
		log.Println("Wait for container to ready got error", err)
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
		log.Println("Create model mount dir got error", err)
		return err
	}
	containerName := m.ModelInfo.ModelAddr + "_validator"

	ctnInfo, err := dockercmd.CreateAndStartContainer(containerName, containerName, m.Port, resultMountDir, m.DisableGPU)
	if err != nil {
		log.Println("Create and start container got error", err)
		return err
	}

	log.Println("Container ID:", ctnInfo.ID)
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
		log.Println("Pause container got error", err)
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
		log.Println("Pause container got error", err)
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
		log.Println("Get container info got error", err)
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
				log.Println("Remove container got error", err)
				return err
			}
		case "paused":
			err = dockercmd.UnpauseContainer(m.containerID)
			if err != nil {
				log.Println("Unpause container got error", err)
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
		log.Println("Get container info got error", err)
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
				log.Println("Remove container got error", err)
				return err
			}
		case "paused":
			err = dockercmd.UnpauseContainer(m.verifierContainerID)
			if err != nil {
				log.Println("Unpause container got error", err)
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
		log.Println("Remove container got error", err)
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
		log.Println("Remove container got error", err)
		return err
	}

	m.VerifierReady = false
	return nil
}

func (m *ModelInstance) IsRunning() (bool, error) {
	containerStatus, err := dockercmd.GetContainerInfo(m.containerID)
	if err != nil {
		log.Println("Get container info got error", err)
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
		log.Println("Get container info got error", err)
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
