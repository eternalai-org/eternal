package watcher

import (
	"testing"

	"eternal-infer-worker/config"
	"eternal-infer-worker/coordinator"
	"eternal-infer-worker/manager"
)

func TestStakeForWorker(t *testing.T) {
	// tskw := &TaskWatcher{
	// 	networkCfg: NetworkConfig{
	// 		RPC: "https://rpc.fluxchain.eternalai.org",
	// 	},
	// 	taskContract: "0x430583bdff80c5be1536ed82f9c8090eef68e2f6",
	// 	account:      "0d3cd1703a70aac5c20f81d1227ee59ec1ae3d70a5782d83bad5543565a6abf3",
	// }
	// tskw.getPendingTaskFromContractZk()
}

func TestTaskWatcher_CheckAssignmentCompleted(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		networkCfg       NetworkConfig
		version          string
		taskContract     string
		account          string
		modelsDir        string
		lighthouseAPI    string
		mode             string
		id               int
		numOfWorker      int
		modelManager     *manager.ModelManager
		coordinator      *coordinator.Coordinator
		zkSycn           bool
		paymasterAddr    string
		paymasterToken   string
		paymasterZeroFee bool
		daoToken         string
		daoTokenName     string
		chaincfg         *config.ChainConfig
		// Named input parameters for target function.
		assignmentID string
		want         bool
		wantErr      bool
	}{
		{
			assignmentID: "6901",
			networkCfg: NetworkConfig{
				RPC: "https://rpc.fluxchain.eternalai.org",
			},
			taskContract: "0x430583bdff80c5be1536ed82f9c8090eef68e2f6",
			account:      "0d3cd1703a70aac5c20f81d1227ee59ec1ae3d70a5782d83bad5543565a6abf3",
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tskw, err := NewTaskWatcher(tt.networkCfg, tt.version, tt.taskContract, tt.account, tt.modelsDir, tt.lighthouseAPI, tt.mode, tt.id, tt.numOfWorker, tt.modelManager, tt.coordinator, tt.zkSycn, tt.paymasterAddr, tt.paymasterToken, tt.paymasterZeroFee, tt.daoToken, tt.daoTokenName, tt.chaincfg)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got, gotErr := tskw.CheckAssignmentCompleted(tt.assignmentID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CheckAssignmentCompleted() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CheckAssignmentCompleted() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("CheckAssignmentCompleted() = %v, want %v", got, tt.want)
			}
		})
	}
}
