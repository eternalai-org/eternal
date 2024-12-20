package config

import (
	"os"

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
	ApiUrl            string
	ApiKey            string
	LighthouseKey     string
}

func ReadConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg.Rpc = os.Getenv("CHAIN_RPC")
	cfg.Account = os.Getenv("ACCOUNT_PRIV")
	cfg.StakingHubAddress = os.Getenv("STAKING_HUB_ADDRESS")
	cfg.WorkerHubAddress = os.Getenv("WORKER_HUB_ADDRESS")
	cfg.ApiUrl = os.Getenv("API_URL")
	cfg.ApiKey = os.Getenv("API_KEY")
	cfg.LighthouseKey = os.Getenv("LIGHT_HOUSE_API_KEY")
	return cfg, nil
}
