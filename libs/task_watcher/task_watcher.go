package task_watcher

import (
	"context"
	"sync"

	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/pkg/logger"

	"go.uber.org/zap"
)

type TasksWatcher struct {
	taskQueue  chan *interfaces.Tasks
	runnerLock sync.RWMutex
	chain      interfaces.IChain
}

func NewTasksWatcher(base interfaces.IChain) *TasksWatcher {
	return &TasksWatcher{
		taskQueue: make(chan *interfaces.Tasks, 1),
		chain:     base,
	}
}

func (t *TasksWatcher) GetPendingTasks(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	tasks, err := t.chain.GetPendingTasks(ctx, 23909302, 23909467)
	if err != nil {
		return
	}
	for _, v := range tasks {
		t.taskQueue <- v
	}
}

func (t *TasksWatcher) ExecueteTasks(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	task := <-t.taskQueue
	logger.GetLoggerInstanceFromContext(ctx).Info("ExecueteTasks", zap.Any("task", task))
}

func (t *TasksWatcher) MakeVerify() error {
	return nil
}

func (t *TasksWatcher) Verify() bool {
	return true
}
