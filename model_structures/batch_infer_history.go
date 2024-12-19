package model_structures

import "time"

const (
	BatchInferHistoryStatusPending       string = "pending"
	BatchInferHistoryStatusAgentInferred string = "agent-inferred"
	BatchInferHistoryStatusQueueHandled  string = "queue-handled"
	BatchInferHistoryStatusCompleted     string = "completed"

	BatchInferHistoryStatusFailed string = "failed"
)

const (
	ToolsetTypeDefault          string = "default"
	ToolsetTypeReplyMentions    string = "reply_mentions"
	ToolsetTypeReplyNonMentions string = "reply_non_mentions"
	ToolsetTypeFollow           string = "follow"
	ToolsetTypePost             string = "post"
	ToolsetTypeCreateToken      string = "create_token"
)

type BatchInferHistory struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id" bson:"user_id"`
	UserAddress string `json:"user_address" bson:"user_address"`
	Toolset     string `json:"toolset" bson:"toolset"`

	AgentContractAddress string `json:"agent_contract_address" bson:"agent_contract_address"`
	ContractAgentID      string `json:"contract_agent_id" bson:"contract_agent_id"`
	ChainID              string `json:"chain_id" bson:"chain_id"`

	AssistantID     string `json:"assistant_id" bson:"assistant_id"`
	PromptInput     string `json:"prompt_input" bson:"prompt_input"`
	SystemPrompt    string `bson:"system_prompt" json:"system_prompt"`
	PromptInputHash string `bson:"prompt_input_hash" json:"prompt_input_hash"`

	AgentType        int    `bson:"agent_type" json:"agent_type"`
	TwitterSnapshot  string `bson:"twitter_snapshot" json:"twitter_snapshot"` // file coin hash
	UserInfoSnapshot string `json:"user_info_snapshot" bson:"user_info_snapshot"`

	OutputMaxCharacter uint   `json:"output_max_character" bson:"output_max_character"`
	PromptOutput       string `json:"prompt_output" bson:"prompt_output"`
	PromptOutputHash   string `bson:"prompt_output_hash" json:"prompt_output_hash"`

	Status string `json:"status" bson:"status"`
	Log    string `json:"log" bson:"log"`

	InferID string `json:"infer_id" bson:"infer_id"` // when call to agent contract to create infer
	ModelID string `json:"model_id" bson:"model_id"`

	InscribeTxHash               string `json:"inscribe_tx_hash" bson:"inscribe_tx_hash"`
	SubmitSolutionInscribeTxHash string `json:"submit_solution_tx_hash" bson:"submit_solution_tx_hash"`

	BtcInscribeTxHash               string    `json:"btc_inscribe_tx_hash" bson:"btc_inscribe_tx_hash"`
	BtcSubmitSolutionInscribeTxHash string    `bson:"btc_submit_solution_inscribe_tx_hash" json:"btc_submit_solution_inscribe_tx_hash"`
	InferWalletAddress              string    `json:"infer_wallet_address" bson:"infer_wallet_address"`
	SubmitInferAt                   time.Time `json:"submit_infer_at" bson:"submit_infer_at"`

	AssignmentAddresses   []string `json:"assignment_addresses" bson:"assignment_addresses"`
	SubmitSolutionAddress string   `json:"submit_solution_address" bson:"submit_solution_address"`

	CommitTxHash []string `json:"commit_tx_hash" bson:"commit_tx_hash"`
	RevealTxHash []string `json:"reveal_tx_hash" bson:"reveal_tx_hash"`

	BtcCommitInscribeTxHash []string `bson:"btc_commit_inscribe_tx_hash" json:"btc_commit_inscribe_tx_hash"`
	BtcRevealInscribeTxHash []string `bson:"btc_reveal_inscribe_tx_hash" json:"btc_reveal_inscribe_tx_hash"`
}
