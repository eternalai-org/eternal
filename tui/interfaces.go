package tui

import "eternal-infer-worker/types"

type TaskManagerI interface {
	GetCurrentRunningTasks() []types.TaskRunnerInfo
	GetWorkerAddress() string
	GetWorkerBalance() string
}
