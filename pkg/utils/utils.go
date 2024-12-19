package utils

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

type (
	Environment string
	Mode        string
)

const (
	Develop    Environment = "develop"
	Staging    Environment = "staging"
	Production Environment = "production"

	Api    Mode = "api"
	Worker Mode = "worker"
)

// HandleSigterm -- Handles Ctrl+C or most other means of "controlled" shutdown gracefully.
// Invokes the supplied func before exiting.
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}

func IsEnvProduction() bool {
	return viper.GetString("ENV") == string(Production)
}

func IsWorker() bool {
	return viper.GetString("MODE") == string(Worker)
}

func IsApi() bool {
	return viper.GetString("MODE") == string(Api)
}

func ToPointer[E any](s E) *E {
	return &s
}
