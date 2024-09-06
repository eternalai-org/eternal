package watcher

import "testing"

func TestStakeForWorker(t *testing.T) {
	tskw := &TaskWatcher{
		networkCfg: NetworkConfig{
			RPC: "https://eternal-ai3.tc.l2aas.com/rpc",
		},
		taskContract: "0xb0F6c20163170958f9935121378a3ed3E8d6263d",
		account:      "d726701ba59f743a4df4f35dd2c1c3ca29fd29de1d73522a4eb5cb38d8fd3ef9",
	}
	tskw.stakeForWorker()
}
