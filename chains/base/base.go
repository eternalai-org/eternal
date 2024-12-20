package base

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"eternal-infer-worker/chains/base/contract/erc20"
	"eternal-infer-worker/chains/base/contract/staking_hub"
	"eternal-infer-worker/chains/base/contract/worker_hub"
	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"

	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/libs/lighthouse"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	b := new(Base)
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
	b.Erc20contractAddress = "0x4b6bf1d365ea1a8d916da37fafd4ae8c86d061d7"
	erc20, err := erc20.NewErc20(common.HexToAddress(b.Erc20contractAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.WorkerHubAddress = cnf.WorkerHubAddress
	wkHub, err := worker_hub.NewWorkerHub(common.HexToAddress(b.StakingHubAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.WorkerHub = wkHub
	b.Erc20contract = erc20
	b.GasLimit = 200_000
	return b, nil
}

func (b *Base) GetPendingTasks(startBlock, endBlock uint64) ([]*interfaces.Tasks, error) {
	ctx := context.Background()
	tasks := []*interfaces.Tasks{}
	iter, err := b.WorkerHub.FilterRawSubmitted(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}

	for iter.Next() {
		tsk, err := b.ProcessBaseChainEventNewInference(ctx, iter.Event)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, tsk...)
	}

	return nil, nil
}

func (b *Base) ProcessBaseChainEventNewInference(ctx context.Context, event *worker_hub.WorkerHubRawSubmitted) ([]*interfaces.Tasks, error) {
	var err error
	tasks := make([]*interfaces.Tasks, 0)
	requestId := event.InferenceId
	requestIdStr := requestId.String()
	_ = requestIdStr
	requestInfo, err := b.WorkerHub.GetInferenceInfo(nil, requestId)
	if err != nil {
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
				// logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
			} else if len(batchFullPrompts) > 0 {
				batchInfers = batchFullPrompts
				isBatch = true
			}
		} else {
			// logger.GetLoggerInstanceFromContext(ctx).Error("DownloadDataSimpleWithRetry", zap.Error(err))
		}

	}

	assignmentIds, err := b.WorkerHub.GetAssignmentsByInference(nil, requestId)
	if err != nil {
		return nil, err
	}
	// here
	for _, assignmentId := range assignmentIds {

		assignment, err := b.WorkerHub.Assignments(nil, assignmentId)
		if err != nil {
			continue
		}

		// fmt.Println("--->", strings.ToLower(assignmentInfo.Worker.String()), "------", strings.ToLower(tskw.address))
		if strings.EqualFold(assignment.Worker.String(), b.Address.Hex()) {
			task := &interfaces.Tasks{
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
			// spew.Dump(task)

			auth, err := eth.CreateBindTransactionOpts(ctx, b.Client, b.PrivateKey, 200_000)
			if err != nil {
				continue
			}

			transact, err := b.WorkerHub.SeizeMinerRole(auth, assignmentId)
			if err != nil {
				continue
			}

			//TODO - transact.Receipt.Logs ???
			/*if err == nil && transact != nil {
				for _, txLog := range transact.Receipt.Logs {
					if txLog == nil {
						continue
					}

					minerRoleSeized, err := b.WorkerHub.ParseMinerRoleSeized(*txLog)
					if err != nil {
						continue
					}
					if strings.EqualFold(b.Address.Hex(), minerRoleSeized.Miner.Hex()) {
						task.AssignmentRole = libs.MODE_MINER
					}
				}
			}*/

			/*
				logger.GetLoggerInstanceFromContext(ctx).Info("ProcessBaseChainEventNewInference",
					zap.String("inferenceId", event.InferenceId.String()),
					zap.String("task", task.TaskID),
					zap.String("worker_address", strings.ToLower(assignmentInfo.Worker.String())),
					zap.String("role", task.AssignmentRole),
					zap.Bool("is_batch", isBatch),
				)*/
			_ = transact

			task.AssignmentRole = libs.MODE_MINER
			tasks = append(tasks, task)
			continue
		}
	}

	return tasks, nil
}

func (b *Base) SubmitTask() {
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
