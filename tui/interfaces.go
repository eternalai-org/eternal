package tui

import "eternal-infer-worker/types"

type TaskManagerI interface {
	GetCurrentRunningTasks() []types.TaskRunnerInfo
	GetWorkerAddress() string
	GetWorkerBalance() string
	UnstakeAndQuit() error
	GetProcessedTasks() uint64
	GetAssignedModel() string
}

type ModelManagerI interface {
	GetStatus() string
}
