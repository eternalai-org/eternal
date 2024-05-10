package eaimodel

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TaskResult struct {
	ResultURI string `json:"result_uri"`
}

type ModelInfoContract struct {
	ModelID   *big.Int
	ModelAddr string
	OwnerAddr string

	Metadata ModelMetadata
}

type ModelMetadata struct {
	Version         uint64       `json:"version"`
	ModelName       string       `json:"model_name"`
	ModelType       ModelType    `json:"model_type"`
	ModelURL        string       `json:"model_url"`
	ModelFileHash   string       `json:"model_file_hash"`
	MinHardwareTier HardwareTier `json:"min_hardware"`

	VerifierURL      string `json:"verifier_url"`
	VerifierFileHash string `json:"verifier_file_hash"`
}

type ModelType string

const (
	ModelTypeText  ModelType = "text"
	ModelTypeImage ModelType = "image"
)

type HardwareTier int64

const (
	HardwareTier_1 HardwareTier = 1
	HardwareTier_2 HardwareTier = 2
	HardwareTier_3 HardwareTier = 3
)

func GetModelInfoFromContract(modelAddr string, ethClient *ethclient.Client) (*ModelInfoContract, error) {

	// fakeModel := &ModelInfoContract{
	// 	ModelID:   new(big.Int).SetInt64(1),
	// 	ModelAddr: strings.ToLower(modelAddr),
	// 	Metadata: ModelMetadata{
	// 		Version:   1,
	// 		ModelName: "Fake Model",
	// 		ModelType: "text",
	// 		// ModelURL:      "https://gateway.lighthouse.storage/ipfs/QmdkKEjx2fauzbPh1j5bUiQXrUG5Ft36pJGHS8awrN89Dc",
	// 		// ModelFileHash: "492b2b3dea3003d359fe1b2cee77a22777d8a9faf942ab6dee81e6bfadaadec4",
	// 		ModelURL:         "https://gateway.lighthouse.storage/ipfs/QmcFYMYpVodkpT6t1fVmWNjPnUnnQbXvwpqyheXvPGKUr8",
	// 		ModelFileHash:    "7f1f29cb884c5b2f4d072b99afcb87f32cbe4adc88cffedab15ffc9fd30887ae",
	// 		MinHardwareTier:  HardwareTier_1,
	// 		VerifierURL:      "https://gateway.lighthouse.storage/ipfs/QmdkKEjx2fauzbPh1j5bUiQXrUG5Ft36pJGHS8awrN89Dc",
	// 		VerifierFileHash: "492b2b3dea3003d359fe1b2cee77a22777d8a9faf942ab6dee81e6bfadaadec4",
	// 	},
	// }

	// return fakeModel, nil

	addr := common.HexToAddress(modelAddr)

	modelContract, err := NewHybridModel(addr, ethClient)
	if err != nil {
		return nil, err
	}
	modelID, err := modelContract.HybridModelCaller.Identifier(nil)
	if err != nil {
		return nil, err
	}
	if modelID.Cmp(new(big.Int).SetInt64(0)) <= 0 {
		log.Println("model id is invalid", modelID)
		return nil, errors.New("model id is invalid")
	}

	modelName, err := modelContract.HybridModelCaller.Name(nil)
	if err != nil {
		return nil, err
	}

	newModelInfo := &ModelInfoContract{
		ModelID:   modelID,
		ModelAddr: strings.ToLower(modelAddr),
	}

	metadata, err := modelContract.HybridModelCaller.Metadata(nil)
	if err != nil {
		return nil, err
	}

	modelMetadata := &ModelMetadata{}

	err = json.Unmarshal([]byte(metadata), modelMetadata)
	if err != nil {
		return nil, err
	}

	modelMetadata.ModelName = modelName

	newModelInfo.Metadata = *modelMetadata
	owner, err := modelContract.HybridModelCaller.Owner(nil)
	if err != nil {
		return nil, err
	}

	newModelInfo.OwnerAddr = strings.ToLower(owner.String())

	return newModelInfo, nil

}

func CheckModelFileHash(modelHash string, filePath string) (bool, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer f.Close()

	buf := make([]byte, 1024*1024)
	h := sha256.New()

	for {
		bytesRead, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				return false, err
			}
			break
		}
		h.Write(buf[:bytesRead])
	}

	log.Printf("checksum: %s\n", hex.EncodeToString(h.Sum(nil)))

	if hex.EncodeToString(h.Sum(nil)) != modelHash {
		return false, nil
	}

	return true, nil
}
