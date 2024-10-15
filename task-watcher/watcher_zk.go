package watcher

import (
	"context"
	"errors"
	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/libs/zkabi"
	"eternal-infer-worker/libs/zkclient"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"strings"
	"time"
)

func (tskw *TaskWatcher) approveErc20Zk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.paymasterToken)
	erc20, err := zkabi.NewErc20(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = erc20

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.Erc20ABI))
	if err != nil {
		return err
	}
	//erc20.Approve()
	maxBigInt := new(big.Int)
	maxBigInt.SetString("99999999999999999999999999999999999999999999999999999999", 10)
	dataBytes, err := instanceABI.Pack(
		"approve", common.HexToAddress(tskw.taskContract), maxBigInt,
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}
	return nil
}

func (tskw *TaskWatcher) stakeForWorkerZk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}

	minStake, err := workerHub.WorkerHubCaller.MinerMinimumStake(nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting minimum stake"))
	}

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.RegisterMiner()
	dataBytes, err := instanceABI.Pack(
		"registerMiner", 1,
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, minStake, dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) joinForMintingZk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.JoinForMinting()
	dataBytes, err := instanceABI.Pack(
		"joinForMinting",
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) ClaimMiningRewardZk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.ClaimReward()
	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	dataBytes, err := instanceABI.Pack(
		"claimReward", pbkHex,
	)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) UnregisterZk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.UnregisterMiner()
	dataBytes, err := instanceABI.Pack(
		"unregisterMiner",
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) RestakeZk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	unstake, err := workerHub.WorkerHubCaller.MinerUnstakeRequests(nil, *pbkHex)
	if err == nil {
		if unstake.UnlockAt != nil && unstake.Stake.Cmp(new(big.Int).SetInt64(0)) > 0 {
			unstakeAt := time.Unix(unstake.UnlockAt.Int64(), 0).UTC()
			if time.Now().Before(unstakeAt) {
				instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
				if err != nil {
					return err
				}
				//workerHub.RestakeForMiner()
				dataBytes, err := instanceABI.Pack(
					"restakeForMiner", 1,
				)
				if err != nil {
					return err
				}

				_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
				if err != nil {
					return err
				}

				return nil
			} else {
				return errors.New("can not restake: unstakeAt after")
			}
		} else {
			return errors.New("can not restake: unstake.Stake <= 0")
		}
	} else {
		return err
	}
}

func (tskw *TaskWatcher) ReclaimStakeZk() error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.UnstakeForMiner()
	dataBytes, err := instanceABI.Pack(
		"unstakeForMiner",
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) SubmitResultZk(assignmentID string, result []byte) error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.SubmitSolution()
	dataBytes, err := instanceABI.Pack(
		"submitSolution", assignmentID, result,
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) seizeMinerRole(_assignmentId *big.Int) error {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return err
	}
	//workerHub.SeizeMinerRole()
	dataBytes, err := instanceABI.Pack(
		"seizeMinerRole", _assignmentId,
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	_, err = client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	return nil
}

func (tskw *TaskWatcher) executeWorkerTaskDefaultZk(modelInst *manager.ModelInstance, task *types.TaskInfo, ext string) error {
	return nil
}

func (tskw *TaskWatcher) getPendingTaskFromContractZk() ([]types.TaskInfo, error) {
	startBlock := uint64(0)
	endBlock := startBlock
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return []types.TaskInfo{}, err
	}
	contract, err := zkabi.NewWorkerHub(common.HexToAddress(tskw.taskContract), ethClient)
	if err != nil {
		return nil, err
	}
	return tskw.filterZKEventNewInference(contract, startBlock, endBlock)
}

func (tskw *TaskWatcher) filterZKEventNewInference(whContract *zkabi.WorkerHub, startBlock uint64, endBlock uint64,
) ([]types.TaskInfo, error) {
	tasks := make([]types.TaskInfo, 0)
	ctx := context.Background()
	iter, err := whContract.FilterNewInference(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	models := tskw.modelManager.GetLoadeModels()
	for iter.Next() {
		requestId := iter.Event.InferenceId
		if err != nil {
			return nil, err
		}
		requestInfo, err := whContract.GetInferenceInfo(nil, requestId)
		if err != nil {
			return nil, err
		}
		modelAddr := strings.ToLower(requestInfo.ModelAddress.Hex())
		if models[modelAddr] == nil {
			continue
		}
		assignments, err := whContract.GetMintingAssignmentsOfInference(nil, requestId)
		if err != nil {
			return nil, err
		}
		for _, assignment := range assignments {
			assignmentInfo, err := whContract.Assignments(nil, assignment.AssignmentId)
			if err != nil {
				continue
			}
			if strings.ToLower(assignmentInfo.Worker.String()) == strings.ToLower(tskw.address) {
				task := types.TaskInfo{
					TaskID:        assignment.InferenceId.String(),
					AssignmentID:  assignment.AssignmentId.String(),
					ModelContract: strings.ToLower(modelAddr),
					Params:        string(requestInfo.Input),
					Requestor:     strings.ToLower(requestInfo.Creator.Hex()),
					Value:         assignment.Value.String(),
				}
				log.Println("task: ", task.TaskID, task.ModelContract, task.Params, task.Requestor)
				err := tskw.seizeMinerRole(assignment.AssignmentId)
				if err == nil {
					tasks = append(tasks, task)
				}
				continue
			}
		}
	}
	return tasks, nil
}
