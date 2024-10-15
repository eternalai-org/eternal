package watcher

import (
	"bytes"
	"errors"
	"eternal-infer-worker/libs/dockercmd"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/file"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/types"
	"fmt"
	"log"
	"strings"
)

// TODO: @liam
func (tskw *TaskWatcher) retrieveRequestResult(requestID string) (*eaimodel.TaskResult, error) {
	return nil, nil
	// ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	// if err != nil {
	// 	return nil, err
	// }

	// hubAddress := common.HexToAddress(tskw.taskContract)

	// workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	// if err != nil {
	// 	return nil, err
	// }

	// requestIDInt := new(big.Int)
	// requestIDInt.SetString(requestID, 10)

	// requestInfo, err := workerHub.WorkerHubCaller.GetInference(nil, requestIDInt)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(requestInfo.Result) == 0 {
	// 	return nil, errors.New("no result")
	// }

	// taskResult := &eaimodel.TaskResult{}

	// err = json.Unmarshal([]byte(requestInfo.Result), taskResult)
	// if err != nil {
	// 	return nil, err
	// }

	// return taskResult, nil
}

func downloadInferResult(resultURI string, filePath string) (string, error) {
	resultURI = strings.ReplaceAll(resultURI, "ipfs://", lighthouse.IPFSGateway)
	log.Println("[downloadInferResult] resultURI: ", resultURI, " ,filePath: ", filePath)
	return file.DownloadChunkedDataDest(resultURI, filePath)
}

func (tskw *TaskWatcher) executeVerifierTask(task *types.TaskInfo) error {
	// retrieveInferResult and check if the result is correct

	// downloadInferResult to verifier docker volume (orgResult)
	// retry infer locally and write result to verifier docker volume (verfResult)

	requestResult, err := tskw.retrieveRequestResult(task.TaskID)
	if err != nil {
		log.Println("retrieve request result error: ", err)
		return err
	}

	newRunner := tskw.GetRunner(task.TaskID)
	if newRunner == nil {
		log.Println("runner not found", task.TaskID)
		return errors.New("runner not found")
	}

	modelInst, err := tskw.modelManager.GetModelInstance(task.ModelContract)
	if err != nil {
		log.Println("get model instance error: ", err)
		return err
	}

	finalResult := &bytes.Buffer{}
	ext := modelInst.GetExt()
	output := fmt.Sprintf("%s/%v.%v", dockercmd.OUTPUT_RESULT_DIR, task.TaskID, ext)

	err = newRunner.Run(output)
	if err != nil {
		log.Println("run task error: ", err)
		return err
	}

	result := newRunner.GetResult()
	resultData, err := readResultFile(result)
	if err != nil {
		log.Println("read result file error: ", err)
		return err
	}
	finalResult = bytes.NewBuffer(resultData)

	orgResultPath := modelInst.VerifyDir + fmt.Sprintf("/%v/org_result.%v", task.TaskID, ext)
	verfResultPath := modelInst.VerifyDir + fmt.Sprintf("/%v/verf_result.%v", task.TaskID, ext)

	_ = finalResult

	_, err = downloadInferResult(requestResult.ResultURI, orgResultPath)
	if err != nil {
		log.Println("download infer result error: ", err)
		return err
	}

	err = writeResultFile(verfResultPath, resultData)
	if err != nil {
		log.Println("write result file error: ", err)
		return err
	}

	// run verifier
	// TODO
	return nil
}
