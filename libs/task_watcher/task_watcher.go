package task_watcher

import (
	"sync"

	"eternal-infer-worker/chains/interfaces"

	"github.com/davecgh/go-spew/spew"
)

type TasksWatcher struct {
	taskQueue  chan *interfaces.Tasks
	runnerLock sync.RWMutex
	base       interfaces.IChain
}

func NewTasksWatcher(base interfaces.IChain) *TasksWatcher {
	return &TasksWatcher{
		taskQueue: make(chan *interfaces.Tasks, 1),
	}
}

func (t *TasksWatcher) GetPendingTasks(wg *sync.WaitGroup) {
	tasks, err := t.base.GetPendingTasks(23909302, 23909467)
	if err != nil {
		return
	}

	for _, v := range tasks {
		t.taskQueue <- v
	}
}

func (t *TasksWatcher) ExecueteTasks(wg *sync.WaitGroup) {
	task := <-t.taskQueue
	spew.Dump(task)
}

func (t *TasksWatcher) Verify() bool {
	return true
}
