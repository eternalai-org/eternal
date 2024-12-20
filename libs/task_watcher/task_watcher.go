package task_watcher

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"sync"

	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/pkg/logger"

	"go.uber.org/zap"
)

var TaskChecker = make(map[string]bool)

type TaskWatcher struct {
	runnerLock   sync.RWMutex
	chain        interfaces.IChain
	cnf          *config.Config
	IsStaked     *bool
	currentBlock uint64
}

func NewTasksWatcher(base interfaces.IChain, cnf *config.Config) *TaskWatcher {
	return &TaskWatcher{
		chain: base,
		cnf:   cnf,
	}
}

func (t *TaskWatcher) GetPendingTasks(ctx context.Context, wg *sync.WaitGroup) chan *interfaces.Task {
	taskQueue := make(chan *interfaces.Task, 10)

	go func() {
		logger.AtLog.Info("Waiting task...")
		defer wg.Done()
		defer close(taskQueue)
		fBlock := t.chain.FromBlock(t.currentBlock)
		tBlock := t.chain.ToBlock()
		// fBlock = 23945180 - 1
		// tBlock = fBlock + 2

		tasks, err := t.chain.GetPendingTasks(ctx, fBlock, tBlock)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks",
				zap.Uint64("from_block", fBlock),
				zap.Uint64("to_block", tBlock),
				zap.Error(err),
			)
		}
		logger.GetLoggerInstanceFromContext(ctx).Info("GetPendingTasks",
			zap.Uint64("from_block", fBlock),
			zap.Uint64("to_block", tBlock),
			zap.Int("tasks", len(tasks)),
		)

		for _, v := range tasks {
			logger.GetLoggerInstanceFromContext(ctx).Info("GetPendingTasks.item",
				zap.Uint64("from_block", fBlock),
				zap.Uint64("to_block", tBlock),
				zap.String("task_id", v.TaskID),
				zap.String("inference_id", v.InferenceID),
				zap.String("assigment_id", v.AssignmentID),
				zap.String("assignment_role", v.AssignmentRole),
				zap.Bool("is_batch", v.IsBatch),
				zap.Int("batch_len", len(v.BatchInfers)),
			)
			if v, ok := TaskChecker[v.TaskID]; ok && v {
				continue
			}

			TaskChecker[v.TaskID] = true
			taskQueue <- v
		}
		t.currentBlock = tBlock
		return
	}()
	return taskQueue
}

func (t *TaskWatcher) ExecueteTasks(ctx context.Context, wg *sync.WaitGroup, taskQueue chan *interfaces.Task) {
	go func() {
		logger.AtLog.Info("Execuete task...")
		defer wg.Done()
		for task := range taskQueue {
			assigmentID, ok := big.NewInt(0).SetString(task.AssignmentID, 10)
			if !ok {
				continue
			}

			taskResult, err := t.executeTasks(ctx, task)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("ExecueteTasks",
					zap.Any("assigment_id", task.AssignmentID),
					zap.String("inference_id", task.AssignmentID),
					zap.Error(err),
				)
				continue
			}

			resultData, err := json.Marshal(taskResult)
			if err != nil {
				continue
			}

			tx, err := t.chain.SubmitTask(ctx, assigmentID, resultData)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("ExecueteTasks",
					zap.Any("assigment_id", task.AssignmentID),
					zap.Any("resultData", resultData),
					zap.String("inference_id", task.AssignmentID),
					zap.Error(err),
				)
				continue
			}

			logger.GetLoggerInstanceFromContext(ctx).Info("ExecueteTasks",
				zap.Any("assigment_id", task.AssignmentID),
				zap.String("inference_id", task.AssignmentID),
				zap.String("tx", tx.Hash().Hex()),
			)
		}
	}()
}

func (t *TaskWatcher) executeTasks(ctx context.Context, task *interfaces.Task) (*interfaces.TaskResult, error) {
	res := &interfaces.TaskResult{}
	result := []byte{}
	if len(task.BatchInfers) == 0 {
		for _, b := range task.BatchInfers {
			seed := libs.CreateSeed(b.PromptInput, task.TaskID)
			obj, err := t.InferChatCompletions(ctx, b.PromptInput, "", seed)
			if err != nil {
				return nil, err
			}
			_b, err := json.Marshal(obj)
			if err != nil {
				return nil, err
			}
			b.PromptOutput = string(_b)
		}

		objJson, err := json.Marshal(task.BatchInfers)
		if err != nil {
			return nil, err
		}

		result = objJson

	} else {
		seed := libs.CreateSeed(task.Params, task.TaskID)
		obj, err := t.InferChatCompletions(ctx, task.Params, "", seed)
		if err != nil {
			return nil, err
		}

		objJson, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}

		result = objJson

	}

	res.Storage = interfaces.LightHouseStorageType
	res.Data = result
	ext := "txt"
	url, err := lighthouse.UploadData(t.cnf.LighthouseKey, fmt.Sprintf("%v_result.%v", task.TaskID, ext), res.Data)
	if err != nil {
		return nil, err
	}
	res.ResultURI = "ipfs://" + url
	logger.GetLoggerInstanceFromContext(ctx).Info("executeTasks", zap.Any("res", res))

	return res, nil
}

func (t *TaskWatcher) InferChatCompletions(ctx context.Context, prompt string, model string, seed uint64) (*interfaces.LLMInferResponse, error) {
	res := &interfaces.LLMInferResponse{}
	infer := interfaces.LLMInferRequest{
		Model: model,
		Seed:  seed,
	}
	infer.Messages = []interfaces.LLMInferMessage{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	url := t.cnf.ApiUrl
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", t.cnf.ApiKey)

	inferJSON, err := json.Marshal(infer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(inferJSON))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("call api err", zap.Error(err))
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (t *TaskWatcher) Verify() bool {
	if t.IsStaked != nil && *t.IsStaked {
		return true
	}

	isStake, err := t.chain.IsStaked()
	if err != nil {
		isStake = false
		t.IsStaked = &isStake
		logger.AtLog.Error(err)
	}
	t.IsStaked = &isStake

	return *t.IsStaked
}

func (t *TaskWatcher) MakeVerify() error {
	err := t.chain.StakeForWorker()
	if err != nil {
		return err
	}

	err = t.chain.JoinForMinting()
	if err != nil {
		return err
	}
	return nil
}
