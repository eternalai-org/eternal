package types

type TaskSubmitRequest struct {
	ModelContract string `json:"model_contract"`
	Params        string `json:"params"`
	Value         string `json:"value"`
}

type TaskInfo struct {
	TaskID string `json:"task_id"`
	// ModelAddress     string `json:"model_name"`
	// ModelID       string `json:"model_id"`
	ModelContract string `json:"model_contract"`
	Params        string `json:"params"`
	Requestor     string `json:"requestor"`
	Value         string `json:"value"`
}

type TaskRunnerInfo struct {
	TaskID    string `json:"task_id"`
	ModelName string `json:"model_name"`
	Params    string `json:"params"`
	Result    string `json:"result"`
	Error     string `json:"error"`
	IsDone    bool   `json:"is_done"`
	Txhash    string `json:"txhash"`
}
