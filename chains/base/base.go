package base

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"

	"eternal-infer-worker/chains/base/contract/erc20"
	"eternal-infer-worker/chains/base/contract/staking_hub"
	"eternal-infer-worker/chains/base/contract/worker_hub"
	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/pkg/logger"

	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/libs/lighthouse"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type Base struct {
	interfaces.Chain
	StakingHubAddress    string
	Erc20contractAddress string
	ModelAddress         string
	GasLimit             uint64

	StakingHub    *staking_hub.StakingHub
	WorkerHub     *worker_hub.WorkerHub
	Erc20contract *erc20.Erc20

	WorkerHubAddress string
}

func NewBaseChain(cnf *config.Config) (*Base, error) {
	b := &Base{}
	err := b.Connect(cnf.Rpc)
	if err != nil {
		return nil, err
	}

	b.PrivateKey = cnf.Account
	b.StakingHubAddress = cnf.StakingHubAddress

	_, address, err := eth.GetAccountInfo(b.PrivateKey)
	if err != nil {
		return nil, err
	}

	b.Address = address
	sthub, err := staking_hub.NewStakingHub(common.HexToAddress(b.StakingHubAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.StakingHub = sthub
	errors.Join(err, errors.New("Error while BaseWhAbiTransactor JoinForMinting"))
	// b.Erc20contractAddress = "0x4b6bf1d365ea1a8d916da37fafd4ae8c86d061d7"
	erc20, err := erc20.NewErc20(common.HexToAddress(b.Erc20contractAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.WorkerHubAddress = cnf.WorkerHubAddress
	wkHub, err := worker_hub.NewWorkerHub(common.HexToAddress(b.WorkerHubAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.WorkerHub = wkHub
	b.Erc20contract = erc20
	b.GasLimit = 200_000
	return b, nil
}

func (b *Base) GetPendingTasks(ctx context.Context, startBlock, endBlock uint64) ([]*interfaces.Task, error) {
	tasks := []*interfaces.Task{}
	iter, err := b.WorkerHub.FilterRawSubmitted(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetPendingTasks#error", zap.Error(err))
		return nil, err
	}

	for iter.Next() {
		tsk, err := b.ProcessBaseChainEventNewInference(ctx, iter.Event)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("ProcessBaseChainEventNewInference#error", zap.Error(err))
			return nil, err
		}
		tasks = append(tasks, tsk...)
	}

	return tasks, nil
}

func (b *Base) ProcessBaseChainEventNewInference(ctx context.Context, event *worker_hub.WorkerHubRawSubmitted) ([]*interfaces.Task, error) {
	var err error
	tasks := make([]*interfaces.Task, 0)
	requestId := event.InferenceId
	requestIdStr := requestId.String()
	_ = requestIdStr
	requestInfo, err := b.WorkerHub.GetInferenceInfo(nil, requestId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetInferenceInfo", zap.Error(err))
		return nil, err
	}

	var batchInfers []*interfaces.BatchInferHistory
	var externalData *interfaces.AgentInferExternalData

	/*
		TODO - chainConfig.AgentContractAddress ???
		if chainConfig.AgentContractAddress != "" {
			isAgentInfer, batchInfers, externalData, err = s.handleNewInferIsAgentInfer(ctx, modelInfo.ModelID.String(), chainConfig, ethClient, event.Raw.TxHash, aiZKClient)
			if err != nil {
				return err
			}
		}*/

	// Detect if  is batch
	isBatch := false
	if strings.HasPrefix(string(requestInfo.Input), config.IPFSPrefix) {
		// TODO - HERE
		inputBytes, _, err := lighthouse.DownloadDataSimpleWithRetry(string(requestInfo.Input))
		if err == nil {
			batchFullPrompts := []*interfaces.BatchInferHistory{}
			err = json.Unmarshal(inputBytes, &batchFullPrompts)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
			} else if len(batchFullPrompts) > 0 {
				batchInfers = batchFullPrompts
				isBatch = true
			}
		} else {
			logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
		}

	}

	assignmentIds, err := b.WorkerHub.GetAssignmentsByInference(nil, requestId)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("GetAssignmentsByInference", zap.Error(err))
		return nil, err
	}
	// here
	for _, assignmentId := range assignmentIds {
		assignment, err := b.WorkerHub.Assignments(nil, assignmentId)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("WorkerHub.Assignments", zap.Error(err))
			continue
		}

		if !strings.EqualFold(assignment.Worker.String(), b.Address.Hex()) {
			continue
		}

		task := &interfaces.Task{
			TaskID:         assignment.InferenceId.String(),
			AssignmentID:   assignmentId.String(),
			ModelContract:  strings.ToLower(event.Model.Hex()),
			Params:         string(requestInfo.Input), // here
			Requestor:      strings.ToLower(requestInfo.Creator.Hex()),
			ZKSync:         true,
			InferenceID:    event.InferenceId.String(),
			AssignmentRole: libs.MODE_VALIDATOR,
			IsBatch:        isBatch,
			BatchInfers:    batchInfers, // here
			ExternalData:   externalData,
		}

		auth, err := eth.CreateBindTransactionOpts(ctx, b.Client, b.PrivateKey, 200_000)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("CreateBindTransactionOpts", zap.Error(err))
			continue
		}

		transact, err := b.WorkerHub.SeizeMinerRole(auth, assignmentId)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("SeizeMinerRole", zap.Error(err))
			continue
		}
		_ = transact

		// for _, txLog := range transact {
		// 	if txLog == nil {
		// 		continue
		// 	}

		// 	minerRoleSeized, err := b.WorkerHub.ParseMinerRoleSeized(*txLog)
		// 	if err != nil {
		// 		continue
		// 	}

		// 	if strings.EqualFold(b.Address.Hex(), minerRoleSeized.Miner.Hex()) {
		// 		task.AssignmentRole = libs.MODE_MINER
		// 	}
		// }
		task.AssignmentRole = libs.MODE_MINER
		logger.GetLoggerInstanceFromContext(ctx).Info("task", zap.Any("id", task.TaskID), zap.Any("assignment_id", task.AssignmentID))
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (b *Base) SubmitTask(ctx context.Context, assigmentID *big.Int, result []byte) (*types.Transaction, error) {
	auth, err := eth.CreateBindTransactionOpts(ctx, b.Client, b.PrivateKey, int64(b.GasLimit))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#CreateBindTransactionOpts", zap.Error(err))
		return nil, err
	}

	tx, err := b.WorkerHub.SubmitSolution(auth, assigmentID, result)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#SubmitSolution", zap.Error(err), zap.Any("assigmentID", assigmentID), zap.Any("result", string(result)))
		return nil, err
	}

	err = eth.WaitForTx(b.Client, tx.Hash())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#WaitForTx", zap.Error(err))
		return nil, err
	}

	receipt, err := b.Client.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("SubmitTask#WaitForTxTransactionReceipt", zap.Error(err))
		return nil, err
	}

	//TODO - check this
	/*if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("tx failed")
	}

	for _, txLog := range receipt.Logs {
		feeLog, err := workerHub.WorkerHubFilterer.ParseTransferFee(*txLog)
		if err != nil {
			continue
		} else {
			if strings.EqualFold(feeLog.Miner.Hex(), tskw.address) {
				tskw.status.processedTasks++
				tskw.status.currentEarning.Add(tskw.status.currentEarning, feeLog.MingingFee)
			}
		}
	}*/
	_ = receipt

	return tx, nil
}

func (b *Base) IsStaked() (bool, error) {
	workerInfo, err := b.StakingHub.Miners(nil, *b.Address)
	if err != nil {
		return false, err
	}

	minStake, err := b.StakingHub.MinerMinimumStake(nil)
	if err != nil {
		return false, err
	}

	if workerInfo.Stake.Cmp(minStake) < 0 {
		return false, nil
	}

	pendingUnstake, err := b.StakingHub.MinerUnstakeRequests(nil, *b.Address)
	if err != nil {
		return false, err
	}

	_ = pendingUnstake
	/*
		tskw.status.pendingUnstakeAmount = pendingUnstake.Stake
		if tskw.status.pendingUnstakeAmount.Cmp(new(big.Int).SetUint64(0)) > 0 {
			tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)
		}

		tskw.status.assignModel = workerInfo.ModelAddress.Hex()
		tskw.status.stakedAmount = workerInfo.Stake*/

	return true, nil
}

func (b *Base) StakeForWorker() error {
	ctx := context.Background()

	balance, err := b.Erc20contract.BalanceOf(nil, *b.Address)
	if err != nil {
		return err
	}

	err = eth.ApproveERC20(ctx, b.Client, b.PrivateKey, common.HexToAddress(b.StakingHubAddress), common.HexToAddress(b.Erc20contractAddress), int64(b.GasLimit))
	if err != nil {
		return err
	}

	minStake, err := b.StakingHub.MinerMinimumStake(nil)
	if err != nil {
		return err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, b.Client, b.PrivateKey, int64(b.GasLimit))
	if err != nil {
		return err
	}

	auth.Value = minStake
	tx := new(types.Transaction)
	if b.ModelAddress != "" {
		tx, err = b.StakingHub.RegisterMiner0(auth, 1, common.HexToAddress(b.ModelAddress))
		if err != nil {
			return err
		}
	} else {
		tx, err = b.StakingHub.RegisterMiner(auth, 1)
		if err != nil {
			return err
		}
	}

	// TODO - here
	_ = tx
	_ = balance
	_ = minStake
	return nil
}

func (b *Base) JoinForMinting() error {
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, b.Client, b.PrivateKey, int64(b.GasLimit))
	if err != nil {
		return err
	}

	tx, err := b.StakingHub.StakingHubTransactor.JoinForMinting(auth)
	if err != nil {
		return err
	}

	_ = tx
	return nil
}
