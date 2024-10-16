package runner

import (
	"crypto/sha256"
	"encoding/hex"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/types"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
)

type RunnerInstance struct {
	modelManager     *manager.ModelManager
	isDone           bool
	task             *types.TaskInfo
	result           string
	expectResultType eaimodel.ModelType
	err              string
}

func NewRunnerInstance(modelManager *manager.ModelManager, task *types.TaskInfo) (*RunnerInstance, error) {
	modelInst, err := modelManager.GetModelInstance(task.ModelContract)
	if err != nil {
		return nil, err
	}

	return &RunnerInstance{
		modelManager:     modelManager,
		task:             task,
		expectResultType: modelInst.ModelInfo.Metadata.ModelType,
	}, nil
}

// TO BE REMOVE
func (r *RunnerInstance) SetTask(task *types.TaskInfo) {
	r.task = task
}

func (r *RunnerInstance) GetResult() string {
	return r.result
}

func (r *RunnerInstance) IsDone() bool {
	return r.isDone
}

func (r *RunnerInstance) SetDone() {
	r.isDone = true
}

func (r *RunnerInstance) Run(output string) error {
	defer func() {
		r.isDone = true
	}()

	log.Println("output: ", output)

	modelInst, err := r.modelManager.GetModelInstance(r.task.ModelContract)
	if err != nil {
		return err
	}

	seed := createSeed(r.task.Params, r.task.TaskID)

	outputPath, err := modelInst.Infer(r.task.Params, output, seed)
	if err != nil {
		return err
	}
	r.result = outputPath
	return nil
}

func createSeed(params string, requestID string) uint64 {
	seed := hex.EncodeToString([]byte(params + requestID))

	h := sha256.New()

	h.Write([]byte(seed))

	bs := h.Sum(nil)

	seedHex := hex.EncodeToString(bs)

	i := new(big.Int)
	i.SetString(seedHex, 16)

	// check if the seed is too large for uint64

	if i.BitLen() > 64 {
		i = i.Mod(i, new(big.Int).SetUint64(math.MaxUint64))
	}

	return i.Uint64()
}

func (r *RunnerInstance) ToTaskRunnerInfo() types.TaskRunnerInfo {
	if r.task == nil {
		return types.TaskRunnerInfo{}
	}
	return types.TaskRunnerInfo{
		TaskID: r.task.TaskID,
		// ModelAddress: r.task.ModelAddress,
		Params: r.task.Params,
		Result: r.result,
		Error:  r.err,
		IsDone: r.isDone,
	}
}
