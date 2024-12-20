package task_watcher

import (
	"context"
	"encoding/json"
	"math/big"
	"sync"
	"time"

	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/pkg/logger"

	"go.uber.org/zap"
)

type TaskWatcher struct {
	taskQueue  chan *interfaces.Task
	runnerLock sync.RWMutex
	chain      interfaces.IChain
	cnf        *config.Config
}

func NewTasksWatcher(base interfaces.IChain, cnf *config.Config) *TaskWatcher {
	return &TaskWatcher{
		chain: base,
		cnf:   cnf,
	}
}

func (t *TaskWatcher) GetPendingTasks(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	logger.AtLog.Info("Waiting task...")

	for {
		time.Sleep(time.Second * libs.TimeToWating)
		fBlock := t.chain.FromBlock()
		tBlock := t.chain.ToBlock()

		tasks, err := t.chain.GetPendingTasks(ctx, fBlock, tBlock)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks",
				zap.Uint64("from_block", fBlock),
				zap.Uint64("to_block", tBlock),
				zap.Error(err),
			)
			return
		}
		logger.GetLoggerInstanceFromContext(ctx).Info("GetPendingTasks",
			zap.Uint64("from_block", fBlock),
			zap.Uint64("to_block", tBlock),
			zap.Int("tasks", len(tasks)),
		)

		if len(tasks) == 0 {
			continue
		}

		t.taskQueue = make(chan *interfaces.Task, len(tasks))
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
			t.taskQueue <- v
		}
		return
	}
}

func (t *TaskWatcher) ExecueteTasks(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	logger.AtLog.Info("Execuete task...")
	for task := range t.taskQueue {
		assigmentID, ok := big.NewInt(0).SetString(task.AssignmentID, 10)
		if !ok {
			continue
		}

		// TODO - execute and get this taskResult
		// 1. batch -> promt output
		// 1. no batch -> goi
		taskResult := interfaces.TaskResult{}

		resultData, err := json.Marshal(taskResult)
		if err != nil {
			continue
		}

		tx, err := t.chain.SubmitTask(assigmentID, resultData)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("ExecueteTasks",
				zap.Any("assigment_id", task.AssignmentID),
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
}

func (t *TaskWatcher) Verify() bool {
	isStake, _ := t.chain.IsStaked()
	isStake = true //
	return isStake
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
