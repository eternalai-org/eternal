package manager

import "testing"

func TestManager_LoadModel(t *testing.T) {
	modelManager := NewModelManager("/home/models", "https://rpc.fluxchain.eternalai.org", "miner", "0x430583bdff80c5be1536ed82f9c8090eef68e2f6", false)
	err := modelManager.LoadModelTest("0x9874732a8699fca824a9a7d948f6bcd30a141238")
	if err != nil {
		return
	}
}

func TestManager_LoadModelZkyncFalse(t *testing.T) {
	modelManager := NewModelManager("/home/models", "https://node.eternalai.org/", "miner", "0x05726BF187938c06d6C832dc493E3Df70fe735c8", false)
	err := modelManager.LoadModelTest("0x8901edfde509d2ddc462cde2cea65f6e8ef27216")
	if err != nil {
		return
	}
}
