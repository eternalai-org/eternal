package manager

import (
	"eternal-infer-worker/libs/eaimodel"
	"math/big"
	"sync"
	"testing"
)

func TestModelInstance_SetupDocker2(t *testing.T) {

}

func TestModelInstance_SetupDocker(t *testing.T) {
	type fields struct {
		ModeID        string
		ModelAddress  string
		ModelType     eaimodel.ModelType
		ModelPath     string
		ModelUrl      string
		ModelFileHash string
		Port          string
		inferLock     sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test ModelInstance_SetupDocker",
			fields: fields{
				ModeID:        "1",
				ModelAddress:  "photon",
				ModelType:     eaimodel.ModelTypeImage,
				ModelPath:     "photon",
				ModelUrl:      "https://gateway.lighthouse.storage/ipfs/QmdkKEjx2fauzbPh1j5bUiQXrUG5Ft36pJGHS8awrN89Dc",
				ModelFileHash: "photon",
				Port:          "5001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelInstance{
				ModelInfo: eaimodel.ModelInfoContract{
					ModelID:   big.NewInt(1),
					ModelAddr: tt.fields.ModelAddress,
					Metadata: eaimodel.ModelMetadata{
						ModelType:     tt.fields.ModelType,
						ModelURL:      tt.fields.ModelUrl,
						ModelFileHash: tt.fields.ModelFileHash,
						Version:       1,
						ModelName:     "test",
					},
				},
				ModelPath: tt.fields.ModelPath,
				Port:      tt.fields.Port,
			}
			if err := m.SetupDocker(); (err != nil) != tt.wantErr {
				t.Errorf("ModelInstance.SetupDocker() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := m.StartDocker(false); (err != nil) != tt.wantErr {
				t.Errorf("ModelInstance.StartDocker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
