package config

import (
	"encoding/json"
	"eternal-infer-worker/libs/eth"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CmdType struct {
	Cmd  string
	Args []string
}

type Config struct {
	Port                 int    `json:"port"`
	ModelsDir            string `json:"models_dir"`
	RPC                  string `json:"rpc"`
	NodeMode             string `json:"node_mode"`
	WorkerHub            string `json:"worker_hub"`
	Account              string `json:"account"`
	LighthouseAPI        string `json:"lighthouse_api"`
	DisableGPU           bool   `json:"disable_gpu"`
	DisableUpdateOnStart bool   `json:"disable_update_on_start"`
	SilentMode           bool   `json:"silent_mode"`

	ZKSync           bool   `json:"zk_sync"`
	PaymasterFeeZero bool   `json:"paymaster_fee_zero"`
	PaymasterAddress string `json:"paymaster_address"`
	PaymasterToken   string `json:"paymaster_token"`
}

const (
	defaultRPC       = "https://node.eternalai.org"
	defaultWorkerHub = "0x05726BF187938c06d6C832dc493E3Df70fe735c8"
	// defaultRPC    = "https://eternal-ai3.tc.l2aas.com/rpc" //Testnet
	// defaultWorkerHub = "0xb0F6c20163170958f9935121378a3ed3E8d6263d" //Testnet
	defaultModelsDir = "./models"
	defaultPort      = 5656
)

func ReadConfig() (*Config, *CmdType, error) {
	var cfg *Config

	rpc := flag.String("rpc", "", "(optional) rpc url of the network")
	// ws := flag.String("ws", "", "ws url of the network")
	workerHub := flag.String("worker-hub", "", "(optional) task contract address")
	account := flag.String("account", "", "(optional) account private key")
	modelsDir := flag.String("models-dir", "", "(optional) models dir")
	port := flag.Int("port", 0, "(optional) port of the server")
	modeValidator := flag.Bool("validator", false, "(optional) run as validator")
	noGPU := flag.Bool("no-gpu", false, "(optional) disable gpu")
	noUpdateOnStart := flag.Bool("no-update-on-start", false, "(optional) disable update on start")
	silentMode := flag.Bool("silent", false, "(optional) silent mode")

	wallet := flag.Bool("wallet", false, "wallet cmd ('-wallet help' for more info)")

	help := flag.Bool("help", false, "show help")

	lighthouseAPI := flag.String("lighthouse", "", "(*REQUIRE) lighthouse api")

	flag.Parse()
	mode := "miner"
	if *modeValidator {
		mode = "validator"
	}

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	cfg, err := readCfgFile()
	if err != nil {
		return nil, nil, err
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

	cfg.SilentMode = *silentMode

	if *noGPU {
		cfg.DisableGPU = true
	}
	if *noUpdateOnStart {
		cfg.DisableUpdateOnStart = true
	}

	if *workerHub != "" {
		//check if invalid contract
		if len(*workerHub) != 42 {
			return nil, nil, fmt.Errorf("invalid worker hub address")
		}
		cfg.WorkerHub = *workerHub
	}

	if *account != "" {
		//check if invalid account
		_, _, err := eth.GetAccountInfo(*account)
		if err != nil {
			return nil, nil, err
		}
		cfg.Account = *account
	}

	if *lighthouseAPI != "" {
		cfg.LighthouseAPI = *lighthouseAPI
	}

	if cfg.LighthouseAPI == "" {
		fmt.Println("*lighthouseAPI", *lighthouseAPI)

		flag.PrintDefaults()
		return nil, nil, fmt.Errorf("lighthouse api is required")
	}

	cfg.setDefaultValue()

	if *wallet {
		a := os.Args[1:]
		ab := strings.Join(a, " ")
		cmdArgs := strings.Split(ab, "-wallet ")
		if len(cmdArgs) <= 1 {
			flag.PrintDefaults()
			os.Exit(0)
		}
		args := strings.Split(cmdArgs[1], " ")
		return cfg, &CmdType{Cmd: "wallet", Args: args}, nil
	}

	err = cfg.save()
	if err != nil {
		return nil, nil, err
	}

	return cfg, nil, nil
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
