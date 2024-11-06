package types

type WorkerInfo struct {
	Address                     string `json:"address"`
	Balance                     string `json:"balance"`
	StakeStatus                 string `json:"stake_status"`
	StakedAmount                string `json:"staked_amount"`
	MiningReward                string `json:"mining_reward"`
	PendingUnstakeAmount        string `json:"pending_unstake_amount"`
	PendingUnstakeUnlockAt      string `json:"pending_unstake_unlock_at"`
	PendingUnstakeUnlockAtBlock string `json:"pending_unstake_unlock_at_block"`
	AssignModel                 string `json:"assign_model"`
}

type HubGlobalInfo struct {
	TotalValidators uint64
	TotalMiners     uint64
	TotalModels     uint64
	FeePercent      uint16
	UnstakeDelay    string
}
