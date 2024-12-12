package watcher

import (
	"testing"
)

func TestStakeForWorker(t *testing.T) {
	tskw := &TaskWatcher{
		networkCfg: NetworkConfig{
			RPC: "https://rpc.fluxchain.eternalai.org",
		},
		taskContract: "0x430583bdff80c5be1536ed82f9c8090eef68e2f6",
		account:      "0d3cd1703a70aac5c20f81d1227ee59ec1ae3d70a5782d83bad5543565a6abf3",
	}
	tskw.getPendingTaskFromContractZk()
}
