package manager

import (
	"time"
)

const INIT_PRICE_AIRC20 = "100000000000000" // 1 / 21e6 * 1e18

type TrainingRequestType int

const (
	TrainingRequestTypeUserDefine TrainingRequestType = iota + 1
	TrainingRequestTypeHuggingFace
)

type TrainingRequestExportStatus string

const (
	TrainingRequestExportStatusWaiting   TrainingRequestExportStatus = "waiting"
	TrainingRequestExportStatusRunning   TrainingRequestExportStatus = "running"
	TrainingRequestExportStatusCompleted TrainingRequestExportStatus = "completed"
	TrainingRequestExportStatusFailed    TrainingRequestExportStatus = "failed"
)

type TrainingRequestStatus string

type TrainingRequestERC20Info struct {
	ERC20Address         string    `json:"erc20_address" bson:"erc20_address"`
	ERC20IssueTime       time.Time `json:"erc20_issue_time" bson:"erc20_issue_time"`
	ERC20Description     string    `json:"erc20_description" bson:"erc20_description"`
	ERC20Name            string    `json:"erc20_name" bson:"erc20_name"`
	ERC20Symbol          string    `json:"erc20_symbol" bson:"erc20_symbol"`
	ERC20Supply          string    `json:"erc20_supply" bson:"erc20_supply"`
	TokenPrice           string    `json:"token_price" bson:"token_price"`
	TokenPriceNumber     float64   `json:"token_price_number" bson:"token_price_number"`
	TokenMarketCap       string    `json:"token_market_cap" bson:"token_market_cap"`
	TokenMarketCapNumber float64   `json:"token_market_cap_number" bson:"token_market_cap_number"`

	CurrentDonate   string `json:"current_donate" bson:"current_donate"`
	TotalInvestment string `json:"total_investment" bson:"-"`
	TotalBackers    int    `json:"total_backers" bson:"-"`

	IsERC20Deployed bool   `json:"is_erc20_deployed" bson:"is_erc20_deployed"`
	ERC20Owner      string `json:"erc20_owner" bson:"erc20_owner"`
}

type APIResponse struct {
	Status int             `json:"status"`
	Data   TrainingRequest `json:"data"`
}

type TrainingRequest struct {
	ModelID          string                `json:"model_id" bson:"model_id"`
	Params           string                `json:"params" bson:"params"`
	Status           TrainingRequestStatus `json:"status" bson:"status"`
	ModelCheckingLog string                `json:"model_checking_log" bson:"model_checking_log"`
	Result           string                `json:"result" bson:"result"`
	IsOnchain        bool                  `json:"is_onchain" bson:"is_onchain"`
	Description      string                `json:"description" bson:"description"`

	OutputUUID  string   `bson:"output_uuid" json:"output_uuid"`
	OutputLink  string   `bson:"output_link" json:"output_link"`
	Progress    int      `bson:"progress" json:"progress"`
	ExecutedAt  int64    `bson:"executed_at" json:"executed_at"`
	CompletedAt int64    `bson:"completed_at" json:"completed_at"`
	Error       string   `bson:"error" json:"error"`
	Logs        string   `bson:"logs" json:"logs"`
	ErrLogs     string   `bson:"err_logs" json:"err_logs"`
	Datasets    []string `bson:"datasets" json:"datasets"`

	PredictNumber int `json:"predict_number" bson:"predict_number"`
	//LastPredict         *ModelPredictHistory `json:"last_predict" bson:"-"`
	Category string `json:"category" bson:"category"`
	//CurrentListing      *ModelMarket         `json:"current_listing" bson:"-"`

	LastUpdatedBlock    uint64 `json:"last_updated_block" bson:"last_updated_block"`
	Thumbnail           string `json:"thumbnail" bson:"thumbnail"`
	ClaimedDeployReward bool   `json:"claimed_deploy_reward" bson:"claimed_deploy_reward"`
	TrainingType        string `json:"training_type" bson:"training_type"`

	TrainingRequestERC20Info `bson:",inline" json:",inline"`

	Type         TrainingRequestType         `json:"type" bson:"type"`
	ExportStatus TrainingRequestExportStatus `json:"export_status" bson:"export_status"`
	ExportResult string                      `json:"export_result" bson:"export_result"`
	ExportLog    string                      `json:"export_log" bson:"export_log"`

	ModelMetadata          string `json:"model_metadata" bson:"model_metadata"`
	ModelAddressUserDefine string `json:"model_address_user_define" bson:"model_address_user_define"`

	FileSize             int64   `json:"file_size" bson:"file_size"`
	ChargeFee            string  `json:"charge_fee" bson:"charge_fee"`
	ChargeFeeNumber      float64 `json:"charge_fee_number" bson:"charge_fee_number"`
	ChargeFeeTokenSymbol string  `json:"charge_fee_token_symbol" bson:"charge_fee_token_symbol"`
	ChargeReceiveAddress string  `json:"charge_receive_address" bson:"charge_receive_address"`
	ChargePaidTxHash     string  `json:"charge_paid_tx_hash" bson:"charge_paid_tx_hash"`

	SortOrder        int    `json:"sort_order" bson:"sort_order"`
	HuggingFaceId    string `bson:"hugging_face_id" json:"hugging_face_id"`
	DataTypeExporter string `json:"data_type_exporter" bson:"data_type_exporter"`

	ZKSync          bool           `json:"zk_sync" bson:"zk_sync"`
	ZKSyncNetwork   *ZKSyncNetwork `json:"zk_sync_network" bson:"zk_sync_network"`
	ModelStorageUrl string         `json:"model_storage_url" bson:"model_storage_url"`
}

type ZKSyncNetwork struct {
	RPC              string `json:"rpc" bson:"rpc"`
	ChainId          string `json:"chain_id" bson:"chain_id"`
	WorkerHubAddress string `json:"worker_hub_address" bson:"worker_hub_address"`
	PaymasterToken   string `json:"paymaster_token" bson:"paymaster_token"`
	PaymasterAddress string `json:"paymaster_address" bson:"paymaster_address"`
	NFTAddress       string `json:"nft_address" bson:"nft_address"`
	PaymasterFeeZero bool   `json:"paymaster_fee_zero" bson:"paymaster_fee_zero"`

	EAIERC20 string `json:"eaierc_20" bson:"eaierc_20"`
	Name     string `json:"name" bson:"name"`
	Explorer string `json:"explorer" bson:"explorer"`
}
