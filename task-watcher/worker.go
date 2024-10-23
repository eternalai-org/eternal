package watcher

import (
	"encoding/json"
	"errors"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/runner"
	"eternal-infer-worker/types"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (tskw *TaskWatcher) executeWorkerTask(task *types.TaskInfo) error {
	runnerInst := tskw.GetRunner(task.TaskID)
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
	outputFile := fmt.Sprintf("%v.%v", task.TaskID, ext)
	err := newRunner.Run(outputFile, setDone)
	if err != nil {
		log.Error("run task error: ", err)
		return nil, err
	}
	resultData, err := readResultFile(fmt.Sprintf("%v/%v.%v", modelInst.ResultDir, task.TaskID, ext))
	if err != nil {
		log.Error("read result file error: ", err)
		return nil, err
	}
	log.Info("uploading result: ", fmt.Sprintf("%v_result.%v", task.TaskID, ext))
	cid, err := lighthouse.UploadData(tskw.lighthouseAPI, fmt.Sprintf("%v_result.%v", task.TaskID, ext), resultData)
	if err != nil {
		log.Error("upload data error: ", err)
		return nil, err
	}
	resultLink := fmt.Sprintf("ipfs://%v", cid)

	taskResult := eaimodel.TaskResult{
		ResultURI: resultLink,
		Storage:   eaimodel.LightHouseStorageType,
		Data:      nil,
	}
	return &taskResult, nil
}

func (tskw *TaskWatcher) executeWorkerTaskDefault(modelInst *manager.ModelInstance, task *types.TaskInfo, ext string, newRunner *runner.RunnerInstance) error {
	log.Info(fmt.Sprintf("executeWorkerTaskDefault runDockerToGetValue for task %s", task.TaskID))
	taskResult, err := tskw.runDockerToGetValue(modelInst, task, ext, newRunner, true)
	if err != nil {
		log.Error("executeWorkerTaskDefault runDockerToGetValue taskResult error: ", err)
		return err
	}
	resultData, err := json.Marshal(taskResult)
	if err != nil {
		log.Error("executeWorkerTaskDefault marshal result error: ", err)
		return err
	}
	log.Info(fmt.Sprintf("executeWorkerTaskDefault result %s", string(resultData)))

	err = tskw.SubmitResult(task.AssignmentID, resultData)
	if err != nil {
		log.Error("executeWorkerTaskDefault submit result data error: ", err)
		return err
	}
	log.Info(fmt.Sprintf("executeWorkerTaskDefault submit result data success"))
	return nil
}
