package watcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"eternal-infer-worker/config"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/libs/abi/base_wh_abi"
	"eternal-infer-worker/libs/db"
	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/libs/lighthouse"
	"eternal-infer-worker/libs/zkabi"
	"eternal-infer-worker/libs/zkclient"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/model_structures"
	"eternal-infer-worker/runner"
	"eternal-infer-worker/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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
	// erc20.Approve()
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
	// workerHub.RegisterMiner()
	dataBytes, err := instanceABI.Pack(
		"registerMiner", uint16(1),
	)
	if err != nil {
		return err
	}

	if tskw.ChainId() == config.HERMES_CHAIN || tskw.ChainId() == config.SYMBIOSIS {
		// hermes
		dataBytes, err = instanceABI.Pack(
			"registerMiner",
			uint16(1),
			// common.HexToAddress("0x0610132d42717Eebb6350BF9fD95fd5A41Cd9170"),
		)
		if err != nil {
			return err
		}
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
	// workerHub.JoinForMinting()
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
	// workerHub.ClaimReward()
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
	// workerHub.UnregisterMiner()
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
				// workerHub.RestakeForMiner()
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
	// workerHub.UnstakeForMiner()
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
	// workerHub.SubmitSolution()
	_assigmentId, _ := new(big.Int).SetString(assignmentID, 10)
	dataBytes, err := instanceABI.Pack(
		"submitSolution", _assigmentId, result,
	)
	if err != nil {
		return err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return err
	}
	tx, err := client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return err
	}

	receipt, err := zkClient.TransactionReceipt(context.Background(), tx.TxHash)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting tx receipt"))
	}

	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("tx failed")
	}

	for _, txLog := range receipt.Logs {
		feeLog, err := workerHub.WorkerHubFilterer.ParseTransferFee(txLog.Log)
		if err != nil {
			continue
		} else {
			_ = feeLog
			tskw.status.processedTasks++
			/*if strings.EqualFold(feeLog.Miner.Hex(), tskw.address) {
				tskw.status.currentEarning.Add(tskw.status.currentEarning, feeLog.MingingFee)
			}*/
		}
	}

	return nil
}

func (tskw *TaskWatcher) createCommitHash(nonce uint64, sender common.Address, data []byte) [32]byte {
	// uint40
	packedData := make([]byte, 5+common.AddressLength+len(data))
	packedData[0] = byte(nonce >> 32)
	packedData[1] = byte(nonce >> 24)
	packedData[2] = byte(nonce >> 16)
	packedData[3] = byte(nonce >> 8)
	packedData[4] = byte(nonce)
	// seder
	copy(packedData[5:], sender.Bytes())
	// data
	copy(packedData[5+common.AddressLength:], data)
	return crypto.Keccak256Hash(packedData[:])
}

func (tskw *TaskWatcher) Reveal(task *types.TaskInfo, data []byte) error {
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
	// workerHub.Reveal()
	_assignmentId, ok := new(big.Int).SetString(task.AssignmentID, 10)
	_ = ok
	//
	log.Info(fmt.Sprintf("TaskWatcher task Call reveal for task=%v with param _assignId=%v nonce=1 data=%v", task.TaskID, _assignmentId.String(), string(data)))
	dataBytes, err := instanceABI.Pack(
		"reveal", _assignmentId, new(big.Int).SetInt64(1), data,
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

func (tskw *TaskWatcher) getInferenceInfo(inferenceID *big.Int) (*zkabi.IWorkerHubInference, error) {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return nil, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := zkabi.NewWorkerHub(hubAddress, zkClient)
	if err != nil {
		return nil, err
	}
	info, err := workerHub.GetInferenceInfo(nil, inferenceID)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (tskw *TaskWatcher) reJoinMinting(modelAddr string) error {
	log.Info("Node is not join model address ", modelAddr)
	err := tskw.joinForMinting()
	if err != nil {
		log.Error("validator retry join for minting error: ", err)
		return err
	}
	log.Info("Node re-join model address ", modelAddr)
	return nil
}

func (tskw *TaskWatcher) isMinerOfModel(modelAddr common.Address) bool {
	info, err := tskw.getMinerAddressesOfModel(modelAddr)
	if err != nil {
		return false
	}
	for _, v := range info {
		if strings.ToLower(v.Hex()) == strings.ToLower(tskw.address) {
			// log.Info(fmt.Sprintf("Watcher: node is still miner of model %s", modelAddr.Hex()))
			return true
		}
	}
	return false
}

func (tskw *TaskWatcher) countAssignmentByMiner() uint64 {
	assignments, err := tskw.getAssignmentByMiner()
	if err != nil {
		return 0
	}
	return uint64(len(assignments))
}

func (tskw *TaskWatcher) getAssignmentByMiner() ([]zkabi.IWorkerHubAssignmentInfo, error) {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return nil, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := zkabi.NewWorkerHub(hubAddress, zkClient)
	if err != nil {
		return nil, err
	}
	info, err := workerHub.GetAssignmentByMiner(nil, common.HexToAddress(tskw.address))
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (tskw *TaskWatcher) getMinerAddressesOfModel(modelAddr common.Address) ([]common.Address, error) {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return nil, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := zkabi.NewWorkerHub(hubAddress, zkClient)
	if err != nil {
		return nil, err
	}
	info, err := workerHub.GetMinerAddressesOfModel(nil, modelAddr)
	if err != nil {
		return nil, err
	}
	return info, nil
}

const (
	ContractInferenceStatusNil = iota
	ContractInferenceStatusSolving
	ContractInferenceStatusCommit
	ContractInferenceStatusReveal
	ContractInferenceStatusProcessed
	ContractInferenceStatusKilled
)

func (tskw *TaskWatcher) Commit(task *types.TaskInfo, data []byte) error {
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
	// workerHub.Commit()
	_commitment := tskw.createCommitHash(uint64(1), common.HexToAddress(tskw.address), data)
	_assignmentId, ok := new(big.Int).SetString(task.AssignmentID, 10)
	_ = ok
	dataBytes, err := instanceABI.Pack(
		"commit", _assignmentId, _commitment,
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

func (tskw *TaskWatcher) seizeMinerRole(_assignmentId *big.Int) (*zktypes.Receipt, error) {
	client := zkclient.NewZkClient(tskw.networkCfg.RPC,
		tskw.paymasterFeeZero,
		tskw.paymasterAddr,
		tskw.paymasterToken)

	zkClient, err := client.GetZkClient()
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(tskw.taskContract)
	workerHub, err := zkabi.NewWorkerHub(contractAddress, zkClient)
	if err != nil {
		return nil, err
	}
	_ = workerHub

	instanceABI, err := abi.JSON(strings.NewReader(zkabi.WorkerHubABI))
	if err != nil {
		return nil, err
	}
	// workerHub.SeizeMinerRole()
	dataBytes, err := instanceABI.Pack(
		"seizeMinerRole", _assignmentId,
	)
	if err != nil {
		return nil, err
	}

	_, pbkHex, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return nil, err
	}
	transact, err := client.Transact(tskw.account, *pbkHex, contractAddress, big.NewInt(0), dataBytes)
	if err != nil {
		return nil, err
	}

	return transact, nil
}

func (tskw *TaskWatcher) executeWorkerTaskDefaultZk(modelInst *manager.ModelInstance, task *types.TaskInfo, ext string, newRunner *runner.RunnerInstance) error {
	if ok := tskw.isMinerOfModel(common.HexToAddress(task.ModelContract)); !ok {
		err := tskw.reJoinMinting(task.ModelContract)
		if err != nil {
			return err
		}
	}
	return tskw.executeWorkerTaskDefault(modelInst, task, ext, newRunner)
}

func (tskw *TaskWatcher) executeVerifierTaskDefaultZk(task *types.TaskInfo) error {
	if ok := tskw.isMinerOfModel(common.HexToAddress(task.ModelContract)); !ok {
		err := tskw.reJoinMinting(task.ModelContract)
		if err != nil {
			return err
		}
	}
	inferenceId, ok := new(big.Int).SetString(task.InferenceID, 10)
	if !ok {
		return errors.New("invalid task")
	}
	infer, err := tskw.getInferenceInfo(inferenceId)
	if err != nil {
		return err
	} else {
		log.Info("-------- executeVerifierTaskDefaultZk get inference info taskID: ", task.TaskID)
		switch infer.Status {
		case ContractInferenceStatusCommit:
			{
				log.Info("executeVerifierTaskDefaultZk Process as a ContractInferenceStatusCommit taskID: ", task.TaskID)
				if task.Status == ContractInferenceStatusCommit {
					if task.TaskResult != nil {
						return nil
					}
				}

				runnerInst := tskw.GetRunner(task.TaskID)
				if runnerInst == nil {
					log.Error("executeVerifierTaskDefaultZk runner not found taskID: ", task.TaskID)
					return errors.New("executeVerifierTaskDefaultZk runner not found")
				}

				modelInst, err := tskw.modelManager.GetModelInstance(task.ModelContract)
				if err != nil {
					log.Error("validator get model instance taskID: ", task.TaskID, " ,error: ", err)
					return err
				}

				// execute to get result from docker container
				ext := modelInst.GetExt()
				taskResult, err := tskw.runDockerToGetValue(modelInst, task, ext, runnerInst, !task.ZKSync)
				log.Info("executeVerifierTaskDefaultZk taskID: ", task.TaskID, " ,taskResult", taskResult)
				if err != nil {
					log.Error("executeVerifierTaskDefaultZk validator run docker taskID: ", task.TaskID, " ,get result error: ", err)
					return err
				}

				resultData, err := json.Marshal(taskResult)
				log.Info("executeVerifierTaskDefaultZk taskID: ", task.TaskID, " resultData: ", string(resultData))
				if err != nil {
					log.Error("executeVerifierTaskDefaultZk taskID: ", task.TaskID, " validator marshal result error: ", err)
					return err
				}

				err = tskw.Commit(task, resultData)
				if err != nil {
					log.Error("executeVerifierTaskDefaultZk taskID: ", task.TaskID, " ,validator commit result error: ", err)
					return err
				}

				task.TaskResult = taskResult
				task.Status = infer.Status
			}
		case ContractInferenceStatusReveal:
			{
				log.Info("executeVerifierTaskDefaultZk taskID: ", task.TaskID, " Process as a ContractInferenceStatusReveal")
				if task.Status == ContractInferenceStatusReveal {
					return nil
				}
				resultData, err := json.Marshal(task.TaskResult)
				if err != nil {
					log.Error("executeVerifierTaskDefaultZk taskID: ", task.TaskID, "  validator marshal result error: ", err)
					return err
				}
				err = tskw.Reveal(task, resultData)
				if err != nil {
					log.Error("executeVerifierTaskDefaultZk  taskID: ", task.TaskID, " validator reveal result error: ", err)
					return err
				}
				task.Status = infer.Status
			}
		default:
			{
				log.Info(fmt.Sprintf("executeVerifierTaskDefaultZk waiting %s, infer status %d", task.TaskID, infer.Status))
			}
		}
	}
	return nil
}

func (tskw *TaskWatcher) getPendingTaskFromContractZk() ([]types.TaskInfo, error) {
	ctx := context.Background()
	jobName := "getPendingTaskFromContractZk"
	state, err := tskw.GetContractSyncState(tskw.taskContract, jobName)

	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		log.Error("[getPendingTaskFromContractZk][NewEthClient] - NewEthClient, tskw.networkCfg.RPC:", tskw.networkCfg.RPC, " ,err: ", err)
		return []types.TaskInfo{}, err
	}

	currentBlock, err := ethClient.BlockNumber(ctx)
	if err != nil {
		log.Error("[getPendingTaskFromContractZk] - currentBlock: ", currentBlock, " ,err: ", err)
		return nil, err
	}

	// log.Info("[getPendingTaskFromContractZk][ContractSyncState] - state: ", state, " ,err: ", err)
	if err != nil || state == nil {
		state = &model_structures.ContractSyncState{
			Job:             jobName,
			ContractAddress: strings.ToLower(tskw.taskContract),
			LastSyncedBlock: (currentBlock - 50), // don't need to start from Block 0, use the (currentBlock - 50) instead.
			ResyncFromBlock: 0,
		}
	}

	startBlock := state.LastSyncedBlock
	var endBlock uint64
	tasks := []types.TaskInfo{}

	for {
		//(currentBlock - 50) to pass this condition
		if endBlock >= currentBlock || startBlock >= currentBlock {
			break
		}

		if currentBlock-startBlock > 1000 {
			endBlock = startBlock + 1000
		} else {
			endBlock = currentBlock
		}

		contract, err := zkabi.NewWorkerHub(common.HexToAddress(tskw.taskContract), ethClient)
		if err != nil {
			log.Error("[getPendingTaskFromContractZk] - currentBlock: ", currentBlock, " ,endBlock: ", endBlock, " startBlock: ", startBlock, " ,err: ", err)
			return nil, err
		}

		_tasks, err := tskw.filterZKEventNewInference(contract, startBlock, endBlock)
		if err != nil {
			log.Error("[getPendingTaskFromContractZk] - currentBlock: ", currentBlock, " ,endBlock: ", endBlock, " startBlock: ", startBlock, " ,err: ", err)
			return nil, err
		}

		state.LastSyncedBlock = endBlock
		states := []model_structures.ContractSyncState{}
		states = append(states, *state)
		err = tskw.UpdateContractSyncStateByAddressAndJob(states)
		if err != nil {
			log.Error("[getPendingTaskFromContractZk] - currentBlock: ", currentBlock, " ,endBlock: ", endBlock, " startBlock: ", startBlock, " ,err: ", err)
			break
		}

		startBlock = endBlock + 1
		tasks = append(tasks, _tasks...)
	}

	if len(tasks) > 0 {
		log.Info("[getPendingTaskFromContractZk] - currentBlock: ", currentBlock, " ,endBlock: ", endBlock, " startBlock: ", startBlock, " ,tasks: ", len(tasks))
	}
	return tasks, nil
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
		log.Error("[filterZKEventNewInference] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,err: ", err)
		return nil, err
	}

	models := tskw.modelManager.GetLoadeModels()
	for iter.Next() {
		requestId := iter.Event.InferenceId
		// ???
		if requestId == nil {
			err = errors.New("request_id is nil")
			log.Error("[filterZKEventNewInference][Err] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,err: ", err)
			return nil, err
		}

		requestInfo, err := whContract.GetInferenceInfo(nil, requestId)
		if err != nil {
			log.Error("[filterZKEventNewInference][GetInferenceInfo] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,err: ", err)
			return nil, err
		}
		_b, _ := json.Marshal(requestInfo)

		modelAddr := strings.ToLower(requestInfo.ModelAddress.Hex())
		if models[modelAddr] == nil {
			err := errors.New("model_address is nil")
			log.Error("[filterZKEventNewInference][Err] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,request_info: ", string(_b), "err: ", err)
			// log.Error("[filterZKEventNewInference][requestInfo.ModelAddress] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,err: ", err)
			_ = err
			continue
		}

		assignments, err := whContract.GetMintingAssignmentsOfInference(nil, requestId)
		if err != nil {
			log.Error("[filterZKEventNewInference][Err] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,request_info: ", string(_b), "err: ", err)
			return nil, err
		}

		// log.Info("[filterZKEventNewInference][assignmentInfo], requestId: ", requestId.String(), " ,assignments: ", len(assignments))
		for _, assignment := range assignments {
			assignmentInfo, err := whContract.Assignments(nil, assignment.AssignmentId)
			if err != nil {
				log.Error("[filterZKEventNewInference][Err] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,request_info: ", string(_b), "err: ", err)
				continue
			}

			if strings.ToLower(assignmentInfo.Worker.String()) == strings.ToLower(tskw.address) {
				task := types.TaskInfo{
					TaskID:         assignment.InferenceId.String(),
					AssignmentID:   assignment.AssignmentId.String(),
					ModelContract:  strings.ToLower(modelAddr),
					Params:         string(requestInfo.Input),
					Requestor:      strings.ToLower(requestInfo.Creator.Hex()),
					Value:          assignment.Value.String(),
					ZKSync:         true,
					InferenceID:    iter.Event.InferenceId.String(),
					AssignmentRole: libs.MODE_VALIDATOR,
				}

				transact, err := tskw.seizeMinerRole(assignment.AssignmentId)
				if err == nil && transact != nil {
					for _, txLog := range transact.Receipt.Logs {
						if txLog == nil {
							continue
						}

						minerRoleSeized, err := whContract.ParseMinerRoleSeized(*txLog)
						if err != nil {
							log.Error("[filterZKEventNewInference][Err] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,request_info: ", string(_b), "err: ", err)
							continue
						}
						if strings.EqualFold(tskw.address, strings.ToLower(minerRoleSeized.Miner.Hex())) {
							task.AssignmentRole = libs.MODE_MINER
						}
					}
				} else {
					log.Error("[filterZKEventNewInference][Err] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,request_info: ", string(_b), "err: ", err)
				}

				log.Info("[filterZKEventNewInference][Proccessing] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,request_info: ", string(_b), "role: ", assignmentInfo.Role)
				tasks = append(tasks, task)
				continue
			} else {
				/*log.Info("[filterZKEventNewInference][seizeMinerRole] startBlock: ",
				startBlock, " ,endBlock: ", endBlock, " ,requestId: ", requestId.String(), " ,assignment.AssignmentId ",
				assignment.AssignmentId, " ,assign for: ", assignmentInfo.Worker.String())*/
			}
		}
	}

	if len(tasks) > 0 {
		log.Info("[filterZKEventNewInference] get tasks: ", len(tasks), " ,startBlock: ", startBlock, ", toBlock: ", endBlock)
	}

	return tasks, nil
}

func (tskw *TaskWatcher) filterBaseChainEventNewInference(whContract *base_wh_abi.WorkerHub, ethClient *ethclient.Client, startBlock uint64, endBlock uint64,
) ([]types.TaskInfo, error) {
	tasks := make([]types.TaskInfo, 0)
	ctx := context.Background()
	iter, err := whContract.FilterRawSubmitted(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil, nil)
	if err != nil {
		log.Error("[filterBaseChainEventNewInference] startBlock: ", startBlock, " ,endBlock: ", endBlock, " ,err: ", err)
		return nil, err
	}

	models := tskw.modelManager.GetLoadeModels()
	for iter.Next() {
		err := tskw.ProcessBaseChainEventNewInference(ctx, iter.Event, tskw.chainCfg, common.HexToAddress(tskw.taskContract), whContract, ethClient)
		if err != nil {
			return nil, err
		}
	}

	if len(tasks) > 0 {
		log.Info("[filterZKEventNewInference] get tasks: ", len(tasks), " ,startBlock: ", startBlock, ", toBlock: ", endBlock)
	}

	_ = models
	return tasks, nil
}

func (tskw *TaskWatcher) GetContractSyncState(contractAddress string, jobName string) (*model_structures.ContractSyncState, error) {
	temp := model_structures.ContractSyncState{}
	data, err := db.Query(temp.CollectionName(),
		temp.Query, contractAddress, jobName,
		&[]model_structures.ContractSyncState{})
	if err != nil {
		return nil, err
	}

	if data != nil {
		if _v, ok := data.(model_structures.ContractSyncState); ok {
			return &_v, nil
		} else {
			return nil, errors.New("GetContractSyncState can not parse to model_structures.ContractSyncState")
		}
	}
	return nil, errors.New("GetContractSyncState not found")
}

func (tskw *TaskWatcher) UpdateContractSyncStateByAddressAndJob(state []model_structures.ContractSyncState) error {
	err := db.Update(model_structures.ContractSyncState{}.CollectionName(), state)
	if err != nil {
		return err
	}
	return nil
}

func (tskw *TaskWatcher) ProcessBaseChainEventNewInference(ctx context.Context, event *base_wh_abi.WorkerHubRawSubmitted, chainConfig *config.ChainConfig, contractAddress common.Address, whContract *base_wh_abi.WorkerHub,
	client *ethclient.Client,
) error {
	var err error
	tasks := make([]types.TaskInfo, 0)
	requestId := event.InferenceId
	requestIdStr := requestId.String()
	_ = requestIdStr
	requestInfo, err := whContract.GetInferenceInfo(nil, requestId)
	if err != nil {
		return err
	}

	var batchInfers []*model_structures.BatchInferHistory
	var externalData *model_structures.AgentInferExternalData

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
		inputBytes, _, err := lighthouse.DownloadDataSimpleWithRetry(string(requestInfo.Input))
		if err == nil {
			batchFullPrompts := []*model_structures.BatchInferHistory{}
			err = json.Unmarshal(inputBytes, &batchFullPrompts)
			if err == nil && len(batchFullPrompts) > 0 {
				batchInfers = batchFullPrompts
				isBatch = true
			}
		}
	}

	var assignments []zkabi.IWorkerHubAssignmentInfo

	assignmentIds, err := whContract.GetAssignmentsByInference(nil, requestId)
	if err != nil {
		return err
	}

	for _, assignmentId := range assignmentIds {
		assignments = append(assignments, zkabi.IWorkerHubAssignmentInfo{
			AssignmentId: assignmentId,
			InferenceId:  requestId,
		})
	}

	for _, assignment := range assignments {
		assignmentInfo, err := whContract.Assignments(nil, assignment.AssignmentId)
		if err != nil {
			continue
		}

		if strings.ToLower(assignmentInfo.Worker.String()) == strings.ToLower(tskw.address) {

			task := types.TaskInfo{
				TaskID:         assignment.InferenceId.String(),
				AssignmentID:   assignment.AssignmentId.String(),
				ModelContract:  strings.ToLower(event.Model.Hex()),
				Params:         string(requestInfo.Input), // here
				Requestor:      strings.ToLower(requestInfo.Creator.Hex()),
				Value:          assignment.Value.String(),
				ZKSync:         true,
				InferenceID:    event.InferenceId.String(),
				AssignmentRole: libs.MODE_VALIDATOR,
				IsBatch:        isBatch,
				BatchInfers:    batchInfers, // here
				ExternalData:   externalData,
			}
			// spew.Dump(task)

			transact, err := tskw.seizeMinerRole(assignment.AssignmentId) // ask contract do i have miner role?
			if err == nil && transact != nil {
				for _, txLog := range transact.Receipt.Logs {
					if txLog == nil {
						continue
					}

					minerRoleSeized, err := whContract.ParseMinerRoleSeized(*txLog)
					if err != nil {
						continue
					}
					if strings.EqualFold(tskw.address, strings.ToLower(minerRoleSeized.Miner.Hex())) {
						task.AssignmentRole = libs.MODE_MINER
					}
				}

			}

			tasks = append(tasks, task)
			continue
		}
	}

	return nil
}
