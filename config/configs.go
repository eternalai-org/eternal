package config

import (
	"encoding/json"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/libs/eth"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type ChainConfig struct {
	ChainId          string `json:"chain_id"`
	Rpc              string `json:"rpc"`
	Explorer         string `json:"explorer"`
	EaiErc20         string `json:"eai_erc20"`
	Name             string `json:"name"`
	NftAddress       string `json:"nft_address"`
	PaymasterAddress string `json:"paymaster_address"`
	PaymasterFeeZero bool   `json:"paymaster_fee_zero"`
	PaymasterToken   string `json:"paymaster_token"`
	WorkerhubAddress string `json:"workerhub_address"`
	ZkSync           bool   `json:"zk_sync"`
	EaiNative        bool   `json:"eai_native"`
	DAOToken         string `json:"dao_token"`
	DAOTokenName     string `json:"dao_token_name"`
	APIUrl           string `json:"api_url"` // if this field is filled out, API will be used instead of docker
	APIKey           string `json:"api_key"` // if the `APIUrl` is filled out, API will be used instead of docker
	ModelID          string `json:"model_id"`
	ModelName        string `json:"model_name"`
}

const (
	ETERNAL_CHAIN = "43338"
	HERMES_CHAIN  = "45762"
	SYMBIOSIS     = "45762"
	FLUX_CHAIN    = "222673"
	BASE_CHAIN    = "8453"
)

var ChainConfigs = map[string]ChainConfig{
	ETERNAL_CHAIN: {
		ChainId:          "43338",
		Rpc:              "https://node.eternalai.org/",
		Explorer:         "https://explorer.eternalai.org",
		EaiErc20:         "EAI Subnet",
		Name:             "",
		NftAddress:       "",
		PaymasterAddress: "",
		PaymasterFeeZero: false,
		PaymasterToken:   "",
		WorkerhubAddress: "",
		ZkSync:           false,
		EaiNative:        false,
		APIUrl:           "",
		APIKey:           "",
	},
	/*"222672": {
		ChainId:          "222672",
		Rpc:              "https://rpc.eternalai.bvm.network",
		Explorer:         "https://explorer.eternalai.bvm.network",
		EaiErc20:         "",
		Name:             "Llama 405B Subnet",
		NftAddress:       "0x78a77967b521a3e92e34ccd52438c0cb2815e898",
		PaymasterAddress: "0xf40a14473f649d15cd63d38f3ca68c4cbc301f3c",
		PaymasterFeeZero: true,
		PaymasterToken:   "0xcdbe9d69d5d9a98d85384c05b462d16a588b53fa",
		WorkerhubAddress: "0x01cfa8f0d8467cd0b03c93d6232d355b0c588f74",
		ZkSync:           true,
		EaiNative:        true,
	},*/
	FLUX_CHAIN: {
		ChainId:          "222673",
		Rpc:              "https://rpc.fluxchain.eternalai.org",
		Explorer:         "https://explorer.fluxchain.eternalai.org",
		EaiErc20:         "",
		Name:             "Flux Subnet",
		NftAddress:       "0x747d9e388d71c11f1e69b556e10f50dbbe4c2d3b",
		PaymasterAddress: "0xf40a14473f649d15cd63d38f3ca68c4cbc301f3c",
		PaymasterFeeZero: true,
		PaymasterToken:   "0xcdbe9d69d5d9a98d85384c05b462d16a588b53fa",
		WorkerhubAddress: "0x430583bdff80c5be1536ed82f9c8090eef68e2f6",
		ZkSync:           true,
		EaiNative:        true,
		DAOToken:         "0x2fDF1e58F61EDE27A1BEa5E329A68dcfB081968b",
		DAOTokenName:     "IMAGINE",
		APIUrl:           "https://api.eternalai.org/v1/chat/completions",
		APIKey:           "",
		ModelID:          "500007",
		ModelName:        "hf.co/lmstudio-community/INTELLECT-1-Instruct-GGUF:Q8_0",
	},
	/*
		HERMES_CHAIN: {
			ChainId:          "45762",
			Rpc:              "https://rpc.hermeschain.eternalai.org",
			Explorer:         "https://explorer.hermeschain.eternalai.org",
			EaiErc20:         "",
			Name:             "Uncensored",
			NftAddress:       "0x97c671381dabae0ae24554120dce2e9b0baeb3cd",
			PaymasterAddress: "0xf40a14473f649d15cd63d38f3ca68c4cbc301f3c",
			PaymasterFeeZero: true,
			PaymasterToken:   "0xcdbe9d69d5d9a98d85384c05b462d16a588b53fa",
			WorkerhubAddress: "0x87e9b8630c1e20dd86451ae15af7663d006f089c",
			ZkSync:           true,
			EaiNative:        true,
			DAOToken:         "0x5211b000cce15fd7ac100e75a157a876dd30bef0",
			DAOTokenName:     "UNCENSORED",

			APIUrl:    "https://api.eternalai.org/v1/chat/completions",
			APIKey:    "",
			ModelID:   "500007",
			ModelName: "PrimeIntellect/INTELLECT-1-Instruct",
		},*/

	SYMBIOSIS: {
		ChainId:          "45762",
		Rpc:              "https://rpc.symbiosis.eternalai.org",
		Explorer:         "https://explorer.symbiosis.eternalai.org",
		EaiErc20:         "",
		Name:             "Uncensored",
		NftAddress:       "0x97c671381dabae0ae24554120dce2e9b0baeb3cd", ///HERMES_MAINNET_COLLECTION_ADDRESS
		PaymasterAddress: "0xf40a14473f649d15cd63d38f3ca68c4cbc301f3c",
		PaymasterFeeZero: true,
		PaymasterToken:   "0xcdbe9d69d5d9a98d85384c05b462d16a588b53fa",
		WorkerhubAddress: "0x87e9B8630c1E20dd86451AE15aF7663D006f089c", //HERMES_MAINNET_WORKER_HUB_ADDRESS
		ZkSync:           true,
		EaiNative:        true,
		DAOToken:         "0x5211b000CCe15fd7aC100E75a157a876dd30bef0", //HERMES_MAINNET_DAO_TOKEN_ADDRESS
		DAOTokenName:     "UNCENSORED",

		APIUrl:    "",
		APIKey:    "",
		ModelID:   "500007",
		ModelName: "hf.com/lmstudio-community/INTELLECT-1-Instruct-GGUF:Q8_0",
	},

	BASE_CHAIN: {
		ChainId:          "8453",
		Rpc:              "https://base.llamarpc.com",
		Explorer:         "https://basescan.org",
		EaiErc20:         "0x4b6bf1d365ea1a8d916da37fafd4ae8c86d061d7",
		Name:             "Base network",
		NftAddress:       "0x1e4c5b7a0568a2cb96bebe5c70472c140fc847a0", ///HERMES_MAINNET_COLLECTION_ADDRESS
		PaymasterAddress: "",
		PaymasterFeeZero: true,
		PaymasterToken:   "",
		WorkerhubAddress: "0xa1d2f74c345ff1d97b8fc72e061903cd84c66f69", //HERMES_MAINNET_WORKER_HUB_ADDRESS
		ZkSync:           true,
		EaiNative:        true,
		DAOToken:         "0x2fb0108f90724f63da4360d39c588124eaeb3f7d", //HERMES_MAINNET_DAO_TOKEN_ADDRESS
		DAOTokenName:     "UNCENSORED",

		APIUrl:    "",
		APIKey:    "",
		ModelID:   "",
		ModelName: "hf.com/lmstudio-community/INTELLECT-1-Instruct-GGUF:Q8_0",
	},
}

var ModelsLLM = map[string]string{
	"500005":       "NousResearch/Hermes-3-Llama-3.1-8B",
	"999999999999": "NousResearch/Hermes-3-Llama-3.1-8B",
}

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

	ZKSync           bool         `json:"zk_sync"`
	PaymasterFeeZero bool         `json:"paymaster_fee_zero"`
	PaymasterAddress string       `json:"paymaster_address"`
	PaymasterToken   string       `json:"paymaster_token"`
	PaymasterZeroFee bool         `json:"paymaster_zero_fee"`
	DAOToken         string       `json:"dao_token"`
	DAOTokenName     string       `json:"dao_token_name"`
	ChainCfg         *ChainConfig `json:"chain_cfg"`
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
	cfg := new(Config)
	err := godotenv.Load()
	rpc := flag.String("rpc", "", "(optional) rpc url of the network")
	// ws := flag.String("ws", "", "ws url of the network")
	workerHub := flag.String("worker-hub", "", "(optional) task contract address")
	account := flag.String("account_priv", "", "(optional) account private key")
	if *account == "" {
		*account = os.Getenv("ACCOUNT_PRIV")
	}

	modelsDir := flag.String("models-dir", "", "(optional) models dir")
	port := flag.Int("port", 0, "(optional) port of the server")
	modeValidator := flag.Bool("validator", false, "(optional) run as validator")
	noGPU := flag.Bool("no-gpu", false, "(optional) disable gpu")
	noUpdateOnStart := flag.Bool("no-update-on-start", false, "(optional) disable update on start")
	silentMode := flag.Bool("silent", false, "(optional) silent mode")
	chain := flag.String("chain", "", "(optional) chain")
	if *chain == "" {
		*chain = os.Getenv("CHAIN_ID")
	}

	wallet := flag.Bool("wallet", false, "wallet cmd ('-wallet help' for more info)")

	help := flag.Bool("help", false, "show help")

	lighthouseAPI := flag.String("lighthouse", "", "(*REQUIRE) lighthouse api")
	if *lighthouseAPI == "" {
		*lighthouseAPI = os.Getenv("LIGHT_HOUSE_API_KEY")
	}

	apiKey := flag.String("api-key", "", "(*REQUIRE) lighthouse api")
	if *apiKey == "" {
		*apiKey = os.Getenv("API_KEY")
	}

	apiURL := flag.String("api-url", "", "(optional) call api instead of docker")
	if *apiURL == "" {
		*apiURL = os.Getenv("API_URL")
	}

	flag.Parse()
	mode := libs.MODE_MINER
	if *modeValidator {
		mode = libs.MODE_VALIDATOR
	}

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	/*
		cfg, err := readCfgFile()
		if err != nil {
			return nil, nil, err
		}*/

	chainConfig := ChainConfigs[*chain]
	cfg.ZKSync = chainConfig.ZkSync
	if cfg.ZKSync {
		cfg.PaymasterToken = chainConfig.PaymasterToken
		cfg.PaymasterAddress = chainConfig.PaymasterAddress
		cfg.PaymasterZeroFee = chainConfig.PaymasterFeeZero
		cfg.DAOToken = chainConfig.DAOToken
		cfg.DAOTokenName = chainConfig.DAOTokenName
	}

	if *port != 0 {
		cfg.Port = *port
	}

	if *modelsDir != "" {
		cfg.ModelsDir = *modelsDir
	}

	if *rpc != "" {
		cfg.RPC = *rpc
		if cfg.ZKSync {
			cfg.RPC = chainConfig.Rpc
		}
	} else {
		cfg.RPC = chainConfig.Rpc
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
		if cfg.ZKSync {
			cfg.WorkerHub = strings.ToLower(chainConfig.WorkerhubAddress)
		}
	} else {
		cfg.WorkerHub = strings.ToLower(chainConfig.WorkerhubAddress)
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

	//cfg.setDefaultValue() >????

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

	cfg.ChainCfg = &chainConfig
	cfg.ChainCfg.APIKey = *apiKey

	if apiURL != nil && *apiURL != "" {
		cfg.ChainCfg.APIUrl = *apiURL
	}
	_b, _ := json.Marshal(cfg)
	fmt.Println("cfg: ", string(_b))
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
		c.NodeMode = libs.MODE_MINER
	}

	if c.Account == "" {
		genPriv, _, _, _ := eth.GenerateAddress()
		c.Account = genPriv
	}
}

func IsUsedAPI(cfg *ChainConfig) (string, bool) {
	if cfg != nil && cfg.APIUrl != "" {
		return cfg.APIUrl, true
	}

	return "", false
}
