package manager

import (
	"errors"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/libs/eaimodel"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ModelManager struct {
	modelsDir     string
	lck           sync.RWMutex
	currentModels map[string]*ModelInstance
	rpc           string
	nodeMode      string
	workerHub     string
	status        string
	disableGPU    bool
	zksync        bool
	llm           bool
}

func NewModelManager(modelsDir, rpcEndpoint, nodeMode, workerHub string, disableGPU bool, zkSync bool, llm bool) *ModelManager {
	return &ModelManager{
		modelsDir:     modelsDir,
		rpc:           rpcEndpoint,
		currentModels: make(map[string]*ModelInstance),
		nodeMode:      nodeMode,
		workerHub:     workerHub,
		status:        "initializing",
		disableGPU:    disableGPU,
		zksync:        zkSync,
		llm:           llm,
	}
}

func (m *ModelManager) GetStatus() string {
	return m.status
}

func (m *ModelManager) setStatus(text string) {
	m.status = text
}

func (m *ModelManager) GetLoadeModels() map[string]*ModelInstance {
	m.lck.RLock()
	defer m.lck.RUnlock()
	res := make(map[string]*ModelInstance)
	for k, v := range m.currentModels {
		res[k] = v
	}
	return res
}

// func (m *ModelManager) WatchAndPreloadModels() {
// 	for {
// 		time.Sleep(5 * time.Second)
// 		modelsList, err := m.getHubModels()
// 		if err != nil {
// 			log.Println("Get models error: ", err)
// 			continue
// 		}

// 		loadedModels := m.GetLoadeModels()
// 		loadedModelsMap := make(map[string]bool)

// 		for _, modelAddr := range loadedModels {
// 			loadedModelsMap[modelAddr] = true
// 		}

// 		modelToPreload := []string{}
// 		for _, modelAddr := range modelsList {
// 			if _, ok := loadedModelsMap[strings.ToLower(modelAddr.String())]; !ok {
// 				modelToPreload = append(modelToPreload, modelAddr.String())
// 			}
// 		}

// 		if len(modelToPreload) > 0 {
// 			m.setStatus("preloading new models")
// 			log.Println("Preload models: ", modelToPreload)
// 			err := m.PreloadModels(modelToPreload)
// 			if err != nil {
// 				log.Println("Preload models error: ", err)
// 			}
// 		} else {
// 			log.Println("No new models to preload")
// 			m.setStatus("all model preloaded")
// 		}
// 	}
// }

// func (m *ModelManager) getHubModels() ([]common.Address, error) {
// 	client, err := ethclient.Dial(m.rpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	hubAddress := common.HexToAddress(m.workerHub)

// 	workerHub, err := abi.NewWorkerHub(hubAddress, client)
// 	if err != nil {
// 		return nil, err
// 	}

// 	models, err := workerHub.WorkerHubCaller.Models()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return models, nil
// }

func (m *ModelManager) PreloadModels(list []string) error {
	m.lck.Lock()
	defer m.lck.Unlock()
	//log.Println("Preloading models")

	for _, modelAddress := range list {
		if modelAddress == "" {
			continue
		}

		// tui.UI.UpdateUI(tui.UIMessageData{
		// 	Action: tui.UIActionShowSpinner,
		// 	Text:   "Preloading model " + modelAddress,
		// })

		log.Println("[ModelManager].PreloadModels - modelAddress: ", modelAddress)
		err := m.loadModel(modelAddress)
		if err != nil {
			return err
		}
	}

	//log.Println("[ModelManager]. Preloading models done")
	return nil
}

func (m *ModelManager) LoadModelTest(modelAddress string) error {
	return m.loadModel(modelAddress)
}

func (m *ModelManager) loadModel(modelAddress string) error {
	var loadErr error
	defer func() {
		if loadErr != nil {
			m.currentModels[strings.ToLower(modelAddress)] = nil
			m.status = "failed"
		}
	}()
	m.status = "loading"
	client, err := ethclient.Dial(m.rpc)
	if err != nil {
		loadErr = err
		return err
	}

	modelInfo, err := eaimodel.GetModelInfoFromContract(modelAddress, client)
	if err != nil {
		loadErr = err
		return err
	}

	inst := &ModelInstance{
		ModelInfo:  *modelInfo,
		ZKSync:     m.zksync,
		ModelPath:  filepath.Join(m.modelsDir, modelInfo.ModelID.String()),
		Port:       fmt.Sprintf("%v", randPort()),
		DisableGPU: m.disableGPU,
		LLM:        m.llm,
	}
	if modelInfo.Metadata.ModelType == eaimodel.ModelTypeText {
		inst.LLM = true
	}

	m.currentModels[strings.ToLower(modelAddress)] = inst

	err = inst.SetupDocker()
	if err != nil {
		loadErr = err
		return err
	}

	err = inst.StartDocker()
	if err != nil {
		loadErr = err
		return err
	}

	/*if m.nodeMode == watcher.MODE_VALIDATOR {
		err = inst.SetupDockerVerifier()
		if err != nil {
			loadErr = err
			return err
		}

		err = inst.StartDockerVerifier()
		if err != nil {
			loadErr = err
			return err
		}
	}*/

	m.status = "loaded"
	return nil
}

func (m *ModelManager) GetModelInstance(modelContract string) (*ModelInstance, error) {
	m.lck.RLock()
	defer m.lck.RUnlock()
	inst, ok := m.currentModels[strings.ToLower(modelContract)]
	if !ok {
		return nil, errors.New("Model not found")
	}
	return inst, nil
}

func (m *ModelManager) IsModelReady(modelAddress string) (bool, bool) {
	m.lck.RLock()
	defer m.lck.RUnlock()
	modelAddress = strings.ToLower(modelAddress)
	if modelInst, ok := m.currentModels[modelAddress]; ok {
		return modelInst.Ready, modelInst.VerifierReady
	}

	return false, false
}

func (m *ModelManager) pauseAllInstances(exception string) error {
	for _, modelInst := range m.currentModels {
		if strings.EqualFold(modelInst.ModelInfo.ModelAddr, exception) {
			continue
		}
		err := modelInst.PauseDocker()
		if err != nil {
			log.Println("Pause model error: ", err)
			return err
		}
		if m.nodeMode == libs.MODE_VALIDATOR {
			err := modelInst.PauseVerifierDocker()
			if err != nil {
				log.Println("Pause model verifier error: ", err)
				return err
			}
		}

	}
	return nil
}

func (m *ModelManager) PauseInstance(modelAddress string) error {
	m.lck.Lock()
	defer m.lck.Unlock()

	modelAddress = strings.ToLower(modelAddress)
	modelInst, ok := m.currentModels[modelAddress]
	if !ok {
		return errors.New("Model not found")
	}

	err := modelInst.PauseDocker()
	if err != nil {
		log.Println("Pause model error: ", err)
		return err
	}
	if m.nodeMode == libs.MODE_VALIDATOR {
		err := modelInst.PauseVerifierDocker()
		if err != nil {
			log.Println("Pause model verifier error: ", err)
			return err
		}
	}

	return nil
}

func (m *ModelManager) RemoveAllInstanceDocker() error {
	m.lck.Lock()
	defer m.lck.Unlock()
	for idx, modelInst := range m.currentModels {
		log.Println("Removing model docker: ", modelInst.ModelInfo.ModelAddr, idx)
		err := modelInst.RemoveDocker()
		if err != nil {
			return err
		}
		if m.nodeMode == libs.MODE_VALIDATOR {
			err := modelInst.RemoveVerifierDocker()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *ModelManager) MakeReady(modelAddress string) error {
	m.lck.Lock()
	defer m.lck.Unlock()

	err := m.pauseAllInstances(strings.ToLower(modelAddress))
	if err != nil {
		return err
	}

	modelAddress = strings.ToLower(modelAddress)
	modelInst, ok := m.currentModels[modelAddress]
	if !ok {
		err := m.loadModel(modelAddress)
		if err != nil {
			log.Println("Load model error: ", err)
			return err
		}
	}

	if !modelInst.Ready || (!modelInst.VerifierReady && m.nodeMode == libs.MODE_VALIDATOR) {
		err = m.startModelInst(modelAddress)
		if err != nil {
			log.Println("Start model error: ", err)
			return err
		}
	}

	return nil
}

func (m *ModelManager) startModelInst(modelAddress string) error {
	defer func() {
		time.Sleep(2 * time.Second)
	}()
	modelAddress = strings.ToLower(modelAddress)
	modelInst, ok := m.currentModels[modelAddress]
	if !ok {
		return errors.New("Model not found")
	}
	err := modelInst.StartDocker()
	if err != nil {
		log.Println("Start model error: ", err)
		return err
	}

	if m.nodeMode == libs.MODE_VALIDATOR {
		err := modelInst.StartDockerVerifier()
		if err != nil {
			log.Println("Start model verifier error: ", err)
			return err
		}
	}
	return nil
}

func (m *ModelManager) GetFiles(dirPath string) []fs.FileInfo {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Error(fmt.Sprintf("[getFiles] Failed to read directory: %v", dirPath), err)
		return []fs.FileInfo{}
	}

	resp := []fs.FileInfo{}

	subDirpaths := []string{}
	for _, file := range files {
		if !file.IsDir() {
			resp = append(resp, file)
		} else {
			subDirPath := filepath.Join(dirPath + file.Name())
			subDirpaths = append(subDirpaths, subDirPath)
		}
	}

	for _, sf := range subDirpaths {
		f := m.GetFiles(sf)
		resp = append(resp, f...)
	}

	return resp
}

func (m *ModelManager) RemoveTheGeneratedFile(modelAddress string) {
	for {
		now := time.Now().UTC()

		// Specify the directory you want to list files from
		dirPath := filepath.Join(MountDir + modelAddress)

		// Read the directory
		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Error(fmt.Sprintf("[RemoveTheGeneratedFile] Failed to read directory: %v", dirPath), err)
			return
		}

		// Loop through the files and print their names
		log.Infof("[RemoveTheGeneratedFile] %d files", len(files))
		limitDate := now.Add(time.Hour * 24 * -1)
		log.Infof("[RemoveTheGeneratedFile] limit date %s", limitDate.String())
		for _, file := range files {
			// Check if it's not a directory
			if !file.IsDir() {
				creationTime := file.ModTime()
				log.Infof("[RemoveTheGeneratedFile] file %s creationTime %s", file.Name(), creationTime.String())
				filePath := dirPath + "/" + file.Name()
				if creationTime.Before(limitDate) {
					if strings.Contains(filePath, ".png") { // only remove png files
						err = os.Remove(filePath)
						if err != nil {
							log.Errorf("[RemoveTheGeneratedFile] error while removing file: %v err %v", filePath, err)
							continue
						}
					}
				}
			}
		}

		time.Sleep(time.Hour * 15)
	}
}
