package manager

import "testing"

func TestManager_LoadModel(t *testing.T) {
	modelManager := NewModelManager("./models", "https://rpc.eternalai.bvm.network", "miner", "0x01cfa8f0d8467cd0b03c93d6232d355b0c588f74", false)
	err := modelManager.LoadModelTest("0x355c5f5ffa7e8a7dce1dc21f7453756a128c9ba4")
	if err != nil {
		return
	}
}
