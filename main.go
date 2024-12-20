package main

import (
	_ "net/http/pprof"
	"sync"
	"time"

	"eternal-infer-worker/chains/base"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs/task_watcher"
)

func main() {
	cnf, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	b, err := base.NewBaseChain(cnf)
	if err != nil {
		panic(err)
	}

	taskWatcher := task_watcher.NewTasksWatcher(b)

goto_here:
	verifed := taskWatcher.Verify()
	if !verifed {

		// err := taskWatcher.MakeVerify()
		// if err != nil {
		// 	// only log
		// }

		time.Sleep(time.Second * 5)
		goto goto_here
	}

	// get and process tasks
	wg := &sync.WaitGroup{}

	for {
		wg.Add(1)
		go taskWatcher.GetPendingTasks(wg)

		wg.Add(1)
		go taskWatcher.ExecueteTasks(wg)

		wg.Wait()
		time.Sleep(5 * time.Second)
	}
}
