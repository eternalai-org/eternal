package tui

import (
	"eternal-infer-worker/types"
	"time"
)

type TaskManagerI interface {
	GetCurrentRunningTasks() []types.TaskRunnerInfo
	GetWorkerAddress() string
	GetWorkerBalance() string
	UnregisterAndQuit() error
	GetProcessedTasks() uint64
	GetSessionEarning() string
	GetAssignedModel() string
	StakeStatus() string
	GetStakedAmount() string
	GetUnstakeInfo() (string, time.Time)
	ReclaimStake() error
}

type ModelManagerI interface {
	GetStatus() string
}
