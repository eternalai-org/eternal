package tui

import (
	"eternal-infer-worker/types"
	"time"
)

type TaskManagerI interface {
	GetCurrentRunningTasks() []types.TaskRunnerInfo
	GetWorkerAddress() string
	GetWorkerBalance() string
	Unregister() error
	GetProcessedTasks() uint64
	GetSessionEarning() string
	GetAssignedModel() string
	StakeStatus() string
	GetStakedAmount() string
	GetUnstakeInfo() (string, time.Time, int64)
	ReclaimStake() error
	ClaimMiningReward() error
	Restake() error
	GetMiningReward() string
	GetHubGlobalInfo() (*types.HubGlobalInfo, error)
	HasNewVersion() (bool, string)
}

type ModelManagerI interface {
	GetStatus() string
}
