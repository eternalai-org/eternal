package main

import (
	"context"
	"flag"
	_ "net/http/pprof"
	"sync"
	"time"

	"eternal-infer-worker/chains/base"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs/task_watcher"
	"eternal-infer-worker/pkg/logger"
)

var configFile string

const TimeToWating time.Duration = 5

func main() {
	// init flag
	flag.StringVar(&configFile, "config-file", "env/development.yml", "Config file path")
	flag.Parse()

	cnf, err := config.ReadConfig(configFile)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	b, err := base.NewBaseChain(cnf)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	taskWatcher := task_watcher.NewTasksWatcher(b)

goto_here:
	verifed := taskWatcher.Verify()
	if !verifed {
		err := taskWatcher.MakeVerify()
		if err != nil {
			logger.AtLog.Error(err)
		}

		time.Sleep(time.Second * TimeToWating)
		goto goto_here
	}

	// get and process tasks
	wg := &sync.WaitGroup{}

	for {
		logger.AtLog.Info("Waiting new task...")
		ctx := context.Background()

		wg.Add(1)
		go taskWatcher.GetPendingTasks(ctx, wg)

		wg.Add(1)
		go taskWatcher.ExecueteTasks(ctx, wg)

		wg.Wait()
		time.Sleep(TimeToWating * time.Second)
	}
}
