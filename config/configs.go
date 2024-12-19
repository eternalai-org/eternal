package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Rpc               string
	Account           string
	StakingHubAddress string
}

func ReadConfig() (*Config, error) {
	cfg := new(Config)

	err := godotenv.Load("config.env")
	if err != nil {
		return nil, err
	}

	cfg.Rpc = os.Getenv("CHAIN_RPC")
	cfg.Account = os.Getenv("ACCOUNT_PRIV")
	cfg.StakingHubAddress = os.Getenv("STAKING_HUB_ADDRESS")
	return cfg, nil
}
