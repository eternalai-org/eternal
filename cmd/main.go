package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"eternal/internal/setting"

	"eternal/pkg/logger"
	"eternal/pkg/utils"

	_ "eternal/db/migrate"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	migrate "github.com/xakep666/mongo-migrate"

	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/zap"
)

var (
	envFlag    string
	configFile string
	err        error
	cr         *cron.Cron
)

func main() {
	// init flag
	flag.StringVar(&envFlag, "env", "development", "Config env: development, production")
	flag.StringVar(&configFile, "config-file", "env/development.yml", "Config file path")
	flag.Parse()

	if utils.Environment(envFlag) == utils.Production {
		viper.AutomaticEnv()
	} else {
		viper.SetConfigFile(configFile)
		if err = viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	// init logger
	if viper.GetString("LOGGER_FORMAT") == "json" {
		logger.InitLoggerDefault(viper.GetBool("LOGGER_ENABLE_DEBUG"))
	} else {
		logger.InitLoggerDefaultDev()
	}

	// first step: need to define service
	s := setting.NewSetting()
	if utils.IsWorker() {
		cr = cron.New(
			cron.WithLogger(cron.VerbosePrintfLogger(logger.AtLog)),
		)
		// worker
		s.InitWorkerProcesses(cr)
		cr.Start()
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.AtLog.Info("Shutting down server...")

	defer func() {
		if utils.IsWorker() {
			cr.Stop()
		}
	}()

	logger.AtLog.Info("Server exiting")
}

func migrateData(mongoDb *mongo.Database) {
	migrate.SetDatabase(mongoDb)
	if migrateErr := migrate.Up(context.Background(), -1); migrateErr != nil {
		logger.AtLog.Logger.Error(
			"Migrate error",
			zap.Error(migrateErr))
	}
}
