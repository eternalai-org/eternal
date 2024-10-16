package watcher

import (
	"encoding/json"
	"errors"
	"eternal-infer-worker/libs/dockercmd"
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
	if modelInst.TrainingRequest == nil || modelInst.TrainingRequest.ZKSync == false {
		return tskw.executeWorkerTaskDefault(modelInst, task, ext, runnerInst)
	} else {
		return tskw.executeWorkerTaskDefaultZk(modelInst, task, ext, runnerInst)
	}
}

func (tskw *TaskWatcher) executeWorkerTaskDefault(modelInst *manager.ModelInstance, task *types.TaskInfo, ext string, newRunner *runner.RunnerInstance) error {
	output := fmt.Sprintf("%s/%v.%v", dockercmd.OUTPUT_RESULT_DIR, task.TaskID, ext)
	err := newRunner.Run(output)
	if err != nil {
		log.Error("run task error: ", err)
		return err
	}
	resultData, err := readResultFile(fmt.Sprintf("%v/%v.%v", modelInst.ResultDir, task.TaskID, ext))
	if err != nil {
		log.Error("read result file error: ", err)
		return err
	}
	log.Info("uploading result: ", fmt.Sprintf("%v_result.%v", task.TaskID, ext))
	cid, err := lighthouse.UploadData(tskw.lighthouseAPI, fmt.Sprintf("%v_result.%v", task.TaskID, ext), resultData)
	if err != nil {
		log.Error("upload data error: ", err)
		return err
	}
	resultLink := fmt.Sprintf("ipfs://%v", cid)

	taskResult := eaimodel.TaskResult{
		ResultURI: resultLink,
		Data:      nil,
		Storage:   eaimodel.LightHouseStorageType,
	}

	resultData, err = json.Marshal(taskResult)
	if err != nil {
		log.Error("marshal result error: ", err)
		return err
	}

	log.Info("\nsubmitting result for task %v size %v\n", task.TaskID, len(resultData))
	err = tskw.SubmitResult(task.AssignmentID, resultData)
	if err != nil {
		log.Error("submit result error: ", err)
		return err
	}
	return nil
}
