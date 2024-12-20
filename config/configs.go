package config

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

const (
	IPFSPrefix = "ipfs://"
)

type Config struct {
	Rpc               string
	Account           string
	StakingHubAddress string
	WorkerHubAddress  string
}

func ReadConfig() (*Config, error) {
	cfg := new(Config)

	err := godotenv.Load("/Users/macbook_autonomous/projects/new-bitcoin-city/eternal/config.env")
	if err != nil {
		return nil, err
	}

	cfg.Rpc = os.Getenv("CHAIN_RPC")
	cfg.Account = os.Getenv("ACCOUNT_PRIV")
	cfg.StakingHubAddress = os.Getenv("STAKING_HUB_ADDRESS")
	cfg.WorkerHubAddress = os.Getenv("WORKER_HUB_ADDRESS")
	spew.Dump(cfg)
	return cfg, nil
}
