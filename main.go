package main

import (
	"eternal-infer-worker/chains/base"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"
	_ "net/http/pprof"
	"sync"
	"time"
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

	taskWatcher := libs.NewTasksWatcher(b)

goto_here:
	verifed := taskWatcher.Verify()
	if !verifed {

		err := taskWatcher.MakeVerify()
		if err != nil {
			//only log
		}

		goto goto_here
	}

	//get and process tasks
	var wg sync.WaitGroup

	for {
		wg.Add(2)

		go taskWatcher.GetPendingTasks(&wg)

		go taskWatcher.ExecueteTasks(&wg)

		wg.Wait()
		time.Sleep(5 * time.Second)
	}

}
