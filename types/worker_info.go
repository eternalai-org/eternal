package types

type WorkerInfo struct {
	Address                string `json:"address"`
	StakeStatus            string `json:"stake_status"`
	StakedAmount           string `json:"staked_amount"`
	MiningReward           string `json:"mining_reward"`
	PendingUnstakeAmount   string `json:"pending_unstake_amount"`
	PendingUnstakeUnlockAt string `json:"pending_unstake_unlock_at"`
	AssignModel            string `json:"assign_model"`
}
