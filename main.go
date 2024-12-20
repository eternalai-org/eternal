package main

import (
	"context"
	"flag"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"

	"eternal-infer-worker/chains/base"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"

	"eternal-infer-worker/libs/googlecloud"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/libs/task_watcher"
	"eternal-infer-worker/pkg/logger"
)

var configFile string

func main() {
	// init flag
	flag.StringVar(&configFile, "config-file", "env/development.yml", "Config file path")
	flag.Parse()

	cnf, err := config.ReadConfig(configFile)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	gcsCfg := googlecloud.GCS{
		ProjectId: os.Getenv("GCS_PROJECT_ID"),
		Bucket:    os.Getenv("GCS_BUCKET"),
		Auth:      os.Getenv("GCS_AUTH"),
		Endpoint:  os.Getenv("GCS_ENDPOINT"),
		Region:    os.Getenv("GCS_REGION"),
		AccessKey: os.Getenv("GCS_ACCESS_KEY"),
		SecretKey: os.Getenv("GCS_SECRET_KEY"),
	}

	_, err = lighthouse.NewDataGCStorage(lighthouse.GCS{
		ProjectId: gcsCfg.ProjectId,
		Bucket:    gcsCfg.Bucket,
		Auth:      gcsCfg.Auth,
		Endpoint:  gcsCfg.Endpoint,
		Region:    gcsCfg.Region,
		AccessKey: gcsCfg.AccessKey,
		SecretKey: gcsCfg.SecretKey,
	})
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	b, err := base.NewBaseChain(cnf)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	taskWatcher := task_watcher.NewTasksWatcher(b, cnf)

goto_here:
	verifed := taskWatcher.Verify()
	if !verifed {
		err := taskWatcher.MakeVerify()
		if err != nil {
			logger.AtLog.Error(err)
		}

		time.Sleep(time.Second * libs.TimeToWating)
		goto goto_here
	}

	// get and process tasks
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	for {
		wg.Add(1)
		taskWatcher.GetPendingTasks(ctx, wg)

		wg.Add(1)
		go taskWatcher.ExecueteTasks(ctx, wg)

		wg.Wait()
		time.Sleep(time.Second * libs.TimeToWating)
	}
}
