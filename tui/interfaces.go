package tui

import "eternal-infer-worker/types"

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
}

type ModelManagerI interface {
	GetStatus() string
}
