package watcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/pkg/logger"
	"eternal-infer-worker/runner"
	"eternal-infer-worker/types"

	"go.uber.org/zap"

	log "github.com/sirupsen/logrus"
)

func (tskw *TaskWatcher) executeWorkerTask(task *types.TaskInfo) error {
	runnerInst := tskw.GetRunner(task.TaskID)
	logger.GetLoggerInstanceFromContext(context.Background()).Info("executeWorkerTask", zap.Any("runnerInst", runnerInst))
	if runnerInst == nil {
		log.Error("runner not found", task.TaskID)
		return errors.New("runner not found")
	}

	modelInst, err := tskw.modelManager.GetModelInstance(task.ModelContract)
	if err != nil {
		log.Error("get model instance error: ", err)
		return err
	}
	// execute to get result from docker container
	ext := modelInst.GetExt()
	if modelInst.ZKSync == false {
		log.Info(fmt.Sprintf("executeWorkerTaskDefault process task %s", task.TaskID))
		return tskw.executeWorkerTaskDefault(modelInst, task, ext, runnerInst)
	} else {
		log.Info(fmt.Sprintf("executeWorkerTaskDefaultZk process task %s", task.TaskID))
		return tskw.executeWorkerTaskDefaultZk(modelInst, task, ext, runnerInst)
	}
}

func (tskw *TaskWatcher) runDockerToGetValue(modelInst *manager.ModelInstance, task *types.TaskInfo, ext string, newRunner *runner.RunnerInstance, setDone bool) (*eaimodel.TaskResult, error) {
	var err error
	logs := new([]zap.Field)
	*logs = []zap.Field{
		zap.String("inferenceId", task.InferenceID),
		zap.String("task", task.TaskID),
		zap.String("role", task.AssignmentRole),
		zap.Bool("modelInst.LLM", modelInst.LLM),
		zap.Bool("r.task.IsBatch", task.IsBatch),
	}

	defer func() {
		if err != nil {
			*logs = append(*logs, zap.Error(err))
			logger.GetLoggerInstanceFromContext(context.Background()).Error("runDockerToGetValue", *logs...)
		} else {
			logger.GetLoggerInstanceFromContext(context.Background()).Info("runDockerToGetValue", *logs...)
		}
	}()

	outputFile := fmt.Sprintf("%v.%v", task.TaskID, ext)
	err = newRunner.Run(outputFile, setDone)
	if err != nil {
		return nil, err
	}

	resultData, err := readResultFile(fmt.Sprintf("%v/%v.%v", modelInst.ResultDir, task.TaskID, ext))
	if err != nil {
		return nil, err
	}

	// TODO - HERE
	cid, err := lighthouse.UploadData(tskw.lighthouseAPI, fmt.Sprintf("%v_result.%v", task.TaskID, ext), resultData)
	if err != nil {
		return nil, err
	}
	resultLink := fmt.Sprintf("ipfs://%v", cid)

	taskResult := eaimodel.TaskResult{
		ResultURI: resultLink,
		Storage:   eaimodel.LightHouseStorageType,
		Data:      nil,
	}

	*logs = append(*logs, zap.Any("taskResult", taskResult))
	return &taskResult, nil
}

func (tskw *TaskWatcher) executeWorkerTaskDefault(modelInst *manager.ModelInstance, task *types.TaskInfo, ext string, newRunner *runner.RunnerInstance) error {
	var err error
	taskResult := &eaimodel.TaskResult{}
	log.Info(fmt.Sprintf("executeWorkerTaskDefault runDockerToGetValue for task %s", task.TaskID))
	taskResult, err = tskw.runDockerToGetValue(modelInst, task, ext, newRunner, true)
	if err != nil {
		log.Error(fmt.Sprintf("executeWorkerTaskDefault runDockerToGetValue task %s Result error: %v", task.TaskID, err))
		return err
	}

	resultData, err := json.Marshal(taskResult)
	if err != nil {
		log.Error(fmt.Sprintf("executeWorkerTaskDefault marshal result task %s error: %v", task.TaskID, err))
		return err
	}
	log.Info(fmt.Sprintf("executeWorkerTaskDefault task %s result %s", task.TaskID, string(resultData)))

	err = tskw.SubmitResult(task.AssignmentID, resultData)
	if err != nil {
		log.Error(fmt.Sprintf("executeWorkerTaskDefault submit result data task %s error: %v", task.TaskID, err))
		return err
	}
	log.Info(fmt.Sprintf("executeWorkerTaskDefault submit result data task %s success", task))
	return nil
}
