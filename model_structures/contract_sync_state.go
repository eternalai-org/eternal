package model_structures

type ContractSyncState struct {
	ContractAddress  string `json:"contract_address"`
	Job              string `json:"job"`
	LastSyncedBlock  uint64 `json:"last_synced_block"`
	ResyncToBlock    uint64 `json:"resync_to_block"`
	ResyncFromBlock  uint64 `json:"resync_from_block"`
	ClearDataAndSync bool   `json:"clear_data_and_sync"`
}

func (ContractSyncState) CollectionName() string {
	return "contract_sync_state"
}
