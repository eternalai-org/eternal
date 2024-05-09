package config

import (
	"encoding/json"
	"eternal-infer-worker/libs/eth"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Port          int    `json:"port"`
	ModelsDir     string `json:"models_dir"`
	RPC           string `json:"rpc"`
	NodeMode      string `json:"node_mode"`
	WorkerHub     string `json:"worker_hub"`
	Account       string `json:"account"`
	LighthouseAPI string `json:"lighthouse_api"`
}

const (
	defaultRPC       = "https://eternal-ai3.tc.l2aas.com/rpc"
	defaultWorkerHub = "0xb0F6c20163170958f9935121378a3ed3E8d6263d"
	defaultModelsDir = "./models"
	defaultPort      = 5656
)

func ReadConfig() (*Config, error) {
	var cfg *Config
	mode := "miner"

	rpc := flag.String("rpc", "", "(optional) rpc url of the network")
	// ws := flag.String("ws", "", "ws url of the network")
	workerHub := flag.String("worker-hub", "", "(optional) task contract address")
	account := flag.String("account", "", "(optional) account private key")
	modelsDir := flag.String("models-dir", "", "(optional) models dir")
	port := flag.Int("port", 0, "(optional) port of the server")
	modeValidator := flag.Bool("validator", false, "(optional) run as validator")

	lighthouseAPI := flag.String("lighthouse-api", "", "(*REQUIRE) lighthouse api")

	flag.Parse()

	if *modeValidator {
		mode = "validator"
	}

	cfg, err := readCfgFile()
	if err != nil {
		return nil, err
	}

	if *port != 0 {
		cfg.Port = *port
	}

	if *modelsDir != "" {
		cfg.ModelsDir = *modelsDir
	}

	if *rpc != "" {
		cfg.RPC = *rpc
	}

	if mode != "" {
		cfg.NodeMode = mode
	}

	if *workerHub != "" {
		cfg.WorkerHub = *workerHub
	}

	if *account != "" {
		cfg.Account = *account
	}

	if *lighthouseAPI != "" {
		cfg.LighthouseAPI = *lighthouseAPI
	}

	if cfg.LighthouseAPI == "" {
		fmt.Println("*lighthouseAPI", *lighthouseAPI)

		flag.PrintDefaults()
		return nil, fmt.Errorf("lighthouse api is required")
	}

	cfg.setDefaultValue()

	err = cfg.save()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func readCfgFile() (*Config, error) {
	var cfg Config
	cfgPath := "./cfg.json"

	cfgFile, err := os.OpenFile(cfgPath, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			//return empty config
			return &cfg, nil
		}
		return nil, err
	}

	defer cfgFile.Close()

	jsonParser := json.NewDecoder(cfgFile)

	if err = jsonParser.Decode(&cfg); err != nil {
		if err.Error() == "EOF" {
			return &cfg, nil
		}
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) save() error {
	cfgPath := "./cfg.json"

	cfgFile, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer cfgFile.Close()

	jsonParser := json.NewEncoder(cfgFile)

	jsonParser.SetIndent("", "  ")
	if err = jsonParser.Encode(c); err != nil {
		return err
	}

	return nil
}

func (c *Config) setDefaultValue() {
	if c.RPC == "" {
		c.RPC = defaultRPC
	}

	if c.WorkerHub == "" {
		c.WorkerHub = defaultWorkerHub
	}

	if c.ModelsDir == "" {
		c.ModelsDir = defaultModelsDir
	}

	if c.Port == 0 {
		c.Port = defaultPort
	}

	if c.NodeMode == "" {
		c.NodeMode = "miner"
	}

	if c.Account == "" {
		genPriv, _, _, _ := eth.GenerateAddress()
		c.Account = genPriv
	}
}
