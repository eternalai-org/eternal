package types

import (
	"encoding/json"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/model_structures"
)

type TaskSubmitRequest struct {
	ModelContract string `json:"model_contract"`
	Params        string `json:"params"`
	Value         string `json:"value"`
}

type TaskInfo struct {
	TaskID string `json:"task_id"`
	// ModelAddress     string `json:"model_name"`
	// ModelID       string `json:"model_id"`
	AssignmentID   string `json:"assignment_id"`
	ModelContract  string `json:"model_contract"`
	Params         string `json:"params"`
	Value          string `json:"value"`
	AssignmentRole string `json:"assignment_role"`

	ZKSync       bool                 `json:"zk_sync"`
	Requestor    string               `json:"requestor"`
	InferenceID  string               `json:"inference_id"`
	TaskResult   *eaimodel.TaskResult `json:"task_result"`
	Status       uint8                `json:"status"`
	Retry        int
	BatchInfers  []*model_structures.BatchInferHistory
	ExternalData *model_structures.AgentInferExternalData
	IsBatch      bool
}

func (a TaskInfo) String() string {
	jsonStr, _ := json.Marshal(a)
	return string(jsonStr)
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
