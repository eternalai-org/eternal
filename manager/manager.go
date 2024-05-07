package manager

import (
	"errors"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/tui"
	"fmt"
	"log"
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
}

func NewModelManager(modelsDir, rpcEndpoint, nodeMode string) *ModelManager {
	return &ModelManager{
		modelsDir:     modelsDir,
		rpc:           rpcEndpoint,
		currentModels: make(map[string]*ModelInstance),
		nodeMode:      nodeMode,
	}
}

func (m *ModelManager) GetLoadeModels() []string {
	m.lck.RLock()
	defer m.lck.RUnlock()
	var res []string
	for k := range m.currentModels {
		res = append(res, k)
	}
	return res
}

func (m *ModelManager) WatchAndPreloadModels() {
	// TODO: @liam watch for new model and load
}

func (m *ModelManager) PreloadModels(list []string) error {
	m.lck.Lock()
	defer m.lck.Unlock()
	log.Println("Preloading models")

	tui.UI.UpdateSectionText(tui.UIMessageData{
		Section: tui.UISectionStatusText,
		Color:   "waiting",
		Text:    "Preloading models",
	})

	for _, modelAddress := range list {
		if modelAddress == "" {
			continue
		}

		// tui.UI.UpdateUI(tui.UIMessageData{
		// 	Action: tui.UIActionShowSpinner,
		// 	Text:   "Preloading model " + modelAddress,
		// })

		err := m.loadModel(modelAddress)
		if err != nil {
			return err
		}
	}

	log.Println("Preloading models done")
	return nil
}

func (m *ModelManager) loadModel(modelAddress string) error {

	client, err := ethclient.Dial(m.rpc)
	if err != nil {
		return err
	}

	modelInfo, err := eaimodel.GetModelInfoFromContract(modelAddress, client)
	if err != nil {
		return err
	}

	inst := &ModelInstance{
		ModelInfo: *modelInfo,

		ModelPath: filepath.Join(m.modelsDir, modelInfo.ModelID.String()),
		Port:      fmt.Sprintf("%v", randPort()),
	}

	m.currentModels[strings.ToLower(modelAddress)] = inst

	err = inst.SetupDocker()
	if err != nil {
		return err
	}

	// err = inst.StartDocker()
	// if err != nil {
	// 	return err
	// }

	if m.nodeMode == "verifier" {
		err = inst.SetupDockerVerifier()
		if err != nil {
			return err
		}

		// err = inst.StartDockerVerifier()
		// if err != nil {
		// 	return err
		// }
	}
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

func (m *ModelManager) pauseAllInstances() error {
	for _, modelInst := range m.currentModels {
		err := modelInst.PauseDocker()
		if err != nil {
			log.Println("Pause model error: ", err)
			return err
		}
		if m.nodeMode == "verifier" {
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
	if m.nodeMode == "verifier" {
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
		if m.nodeMode == "verifier" {
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

	err := m.pauseAllInstances()
	if err != nil {
		return err
	}

	modelAddress = strings.ToLower(modelAddress)
	_, ok := m.currentModels[modelAddress]
	if !ok {
		err := m.loadModel(modelAddress)
		if err != nil {
			log.Println("Load model error: ", err)
			return err
		}
	}

	err = m.startModelInst(modelAddress)
	if err != nil {
		log.Println("Start model error: ", err)
		return err
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

	if m.nodeMode == "verifier" {
		err := modelInst.StartDockerVerifier()
		if err != nil {
			log.Println("Start model verifier error: ", err)
			return err
		}
	}
	return nil
}
