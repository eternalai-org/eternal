package model_structures

import "strings"

type StateDB interface {
	Query(contractAddress, jobName string, v interface{}) interface{}
}

type ContractSyncState struct {
	ID               int    `json:"id"`
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

func (o *ContractSyncState) Query(contractAddress, jobName string, v interface{}) interface{} {
	input := v.(*[]ContractSyncState)
	for _, i := range *input {
		if strings.EqualFold(i.ContractAddress, contractAddress) && strings.EqualFold(i.Job, jobName) {
			return i
		}
	}
	return nil
}
