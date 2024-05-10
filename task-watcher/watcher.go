package watcher

import (
	"context"
	"errors"
	"eternal-infer-worker/coordinator"
	"eternal-infer-worker/libs/abi"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/runner"
	"eternal-infer-worker/tui"
	"eternal-infer-worker/types"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkConfig struct {
	RPC string
	WS  string
}

type TaskWatcherStatus struct {
	processedTasks         uint64
	currentEarning         *big.Int
	stakeStatus            string
	stakedAmount           *big.Int
	pendingUnstakeAmount   *big.Int
	pendingUnstakeUnlockAt time.Time
	assignModel            string
}

type TaskWatcher struct {
	modelManager *manager.ModelManager
	coordinator  *coordinator.Coordinator
	networkCfg   NetworkConfig
	modelsDir    string
	account      string
	address      string
	taskContract string
	// watcherID     int
	// numOfWorker   int
	taskQueue     chan *types.TaskInfo
	runnerLock    sync.RWMutex
	currentRunner map[string]*runner.RunnerInstance
	lighthouseAPI string

	status TaskWatcherStatus
	mode   string

	unstakeLock sync.Mutex
}

func NewTaskWatcher(networkCfg NetworkConfig, taskContract, account, modelsDir, lighthouseAPI, mode string, id, numOfWorker int, modelManager *manager.ModelManager, coordinator *coordinator.Coordinator) (*TaskWatcher, error) {

	_, address, err := eth.GenerateAddressFromPrivKey(account)
	if err != nil {
		return nil, err
	}

	return &TaskWatcher{
		networkCfg:   networkCfg,
		modelsDir:    modelsDir,
		account:      account,
		address:      strings.ToLower(address),
		taskContract: taskContract,
		// watcherID:     id,
		// numOfWorker:   numOfWorker,
		modelManager:  modelManager,
		coordinator:   coordinator,
		taskQueue:     make(chan *types.TaskInfo, 1),
		currentRunner: make(map[string]*runner.RunnerInstance),
		lighthouseAPI: lighthouseAPI,
		mode:          mode,
		status: TaskWatcherStatus{
			processedTasks:       0,
			stakeStatus:          "-",
			assignModel:          "-",
			currentEarning:       big.NewInt(0),
			stakedAmount:         big.NewInt(0),
			pendingUnstakeAmount: big.NewInt(0),
		},
	}, nil
}

func (tskw *TaskWatcher) Start() {

	log.Println("start task watcher")
	tui.UI.UpdateSectionText(tui.UIMessageData{
		Section: tui.UISectionStatusText,
		Color:   "normal",
		Text:    "starting task watcher...",
	})
	go tskw.executeTasks()

	err := tskw.checkRegisteredAndStaked()
	if err != nil {
		log.Println("check registered and staked error: ", err)
		return
	}

	err = tskw.modelManager.PreloadModels([]string{tskw.status.assignModel})
	if err != nil {
		log.Println("preload models error: ", err)
		panic(err)
		return
	}

	if tskw.mode == "miner" {
		err = tskw.joinForMinting()
		if err != nil {
			log.Println("join for minting error: ", err)
			// return
		}
	}

	tskw.watchAndAssignTask()
}

func (tskw *TaskWatcher) joinForMinting() error {
	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return err
	}

	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting chain ID"))
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting gas price"))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(workerAcc, chainID)
	if err != nil {
		return errors.Join(err, errors.New("Error while creating new keyed transactor"))
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.JoinForMinting(auth)
	if err != nil {
		return errors.Join(err, errors.New("Error while JoinForMinting"))
	}

	log.Println("JoinForMinting tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}

	log.Println("JoinForMinting success")

	return nil
}

func (tskw *TaskWatcher) GetCurrentRunningTasks() []types.TaskRunnerInfo {
	tskw.runnerLock.Lock()
	defer tskw.runnerLock.Unlock()

	runningTasks := make([]types.TaskRunnerInfo, 0)

	for _, runner := range tskw.currentRunner {
		if !runner.IsDone() {
			runningTasks = append(runningTasks, runner.ToTaskRunnerInfo())
		}
	}

	sort.SliceStable(runningTasks, func(i, j int) bool {
		taskAid, _ := new(big.Int).SetString(runningTasks[i].TaskID, 10)
		taskBid, _ := new(big.Int).SetString(runningTasks[j].TaskID, 10)
		return taskAid.Cmp(taskBid) > 0
	})

	return runningTasks
}

func (tskw *TaskWatcher) watchAndAssignTask() {
	// watch and assign task to model
	maxConcurrentTask := 1

	tui.UI.UpdateSectionText(tui.UIMessageData{
		Section: tui.UISectionStatusText,
		Text:    "running...",
		Color:   "normal",
	})
	for {
		time.Sleep(2 * time.Second)
		tskw.CleanupRunners()

		tasks, err := tskw.getPendingTaskFromContract()
		if err != nil {
			log.Println("get pending task error: ", err)
			continue
		}

		// sort tasks by task id in descending order
		sort.SliceStable(tasks, func(i, j int) bool {
			taskAid, _ := new(big.Int).SetString(tasks[i].TaskID, 10)
			taskBid, _ := new(big.Int).SetString(tasks[j].TaskID, 10)
			return taskAid.Cmp(taskBid) > 0
		})

		log.Println("pending tasks: ", len(tasks))

		for _, task := range tasks {
			if len(tskw.currentRunner) >= maxConcurrentTask {
				log.Println("worker is full")
				break
			}
			log.Println("assign task: ", task.TaskID, task.ModelContract, task.Params)
			err = tskw.AssignTask(task)
			if err != nil {
				log.Println("assign task error: ", err)
			}
		}

	}
}

func (tskw *TaskWatcher) findRequestTxhash(ctx context.Context, requestID string, ethClient *ethclient.Client, workerHub *abi.WorkerHub) (string, error) {
	currentBlock, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return "", err
	}
	requestIDInt, ok := new(big.Int).SetString(requestID, 10)
	if !ok {
		return "", errors.New("invalid request id")
	}
	for {
		opts := &bind.FilterOpts{
			Start:   currentBlock - 1000,
			End:     &currentBlock,
			Context: ctx,
		}
		iter, err := workerHub.WorkerHubFilterer.FilterNewInference(opts, []*big.Int{requestIDInt}, nil)
		if err != nil {
			return "", err
		}
		for iter.Next() {
			return iter.Event.Raw.TxHash.String(), nil
		}
		if currentBlock <= 0 {
			break
		}
		currentBlock -= 1000
	}
	return "", errors.New("txhash not found")
}

func (tskw *TaskWatcher) getPendingTaskFromContract() ([]types.TaskInfo, error) {
	// get pending tasks from contract
	log.Println("get pending task from contract: ", tskw.address)

	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return []types.TaskInfo{}, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return []types.TaskInfo{}, err
	}
	// get pending tasks from contract
	workerAddress := common.HexToAddress(tskw.address)

	tasks := make([]types.TaskInfo, 0)
	// TODO @liam get verifier task to
	if tskw.mode == "validator" {
		// requests, err := workerHub.WorkerHubCaller.GetModelUnresolvedInferences(nil, modelAddress)
		// if err != nil {
		// 	return []types.TaskInfo{}, err
		// }
		// for _, request := range requests {
		// 	task := types.TaskInfo{
		// 		TaskID:        request.RequestId.String(),
		// 		Value:         request.Value.String(),
		// 		ModelContract: strings.ToLower(request.Model.Hex()),
		// 		Params:        string(request.Data),
		// 		Requestor:     strings.ToLower(request.Creator.Hex()),
		// 	}
		// 	tasks = append(tasks, task)
		// }
	} else {
		// get unresolved claimed inference requests
		requests, err := workerHub.WorkerHubCaller.GetMiningAssignments(&bind.CallOpts{
			From: workerAddress,
		})
		if err != nil {
			return []types.TaskInfo{}, err
		}
		if len(requests) == 0 {
			return []types.TaskInfo{}, nil
		}

		for _, request := range requests {

			modelAddr := strings.ToLower(request.ModelAddress.Hex())

			inferenceInfo, err := workerHub.WorkerHubCaller.GetInferences(nil, []*big.Int{request.InferenceId})
			if err != nil {
				return []types.TaskInfo{}, err
			}
			task := types.TaskInfo{
				TaskID:        request.InferenceId.String(),
				AssignmentID:  request.AssignmentId.String(),
				ModelContract: strings.ToLower(modelAddr),
				Params:        string(request.Input),
				Requestor:     strings.ToLower(inferenceInfo[0].Creator.Hex()),
				Value:         inferenceInfo[0].Value.String(),
			}

			log.Println("task: ", task.TaskID, task.ModelContract, task.Params, task.Requestor)

			tasks = append(tasks, task)
		}

	}

	return tasks, nil
}

func (tskw *TaskWatcher) GetAllTasks() []types.TaskRunnerInfo {
	tskw.runnerLock.Lock()
	defer tskw.runnerLock.Unlock()

	pendingTasks := make([]types.TaskRunnerInfo, 0)

	for _, runner := range tskw.currentRunner {
		pendingTasks = append(pendingTasks, runner.ToTaskRunnerInfo())
	}

	return pendingTasks
}
func (tskw *TaskWatcher) GetCurrentPendingTasks() []types.TaskRunnerInfo {
	tskw.runnerLock.Lock()
	defer tskw.runnerLock.Unlock()

	pendingTasks := make([]types.TaskRunnerInfo, 0)

	for _, runner := range tskw.currentRunner {
		if !runner.IsDone() {
			pendingTasks = append(pendingTasks, runner.ToTaskRunnerInfo())
		}
	}

	return pendingTasks
}

func (tskw *TaskWatcher) assigningTask(task *types.TaskInfo) {
	tskw.taskQueue <- task
}

func (tskw *TaskWatcher) AssignTask(task types.TaskInfo) error {

	newRunner, err := runner.NewRunnerInstance(tskw.modelManager, &task)
	if err != nil {
		log.Println("create runner error: ", err)
		return err
	}

	err = tskw.AddRunner(task.TaskID, newRunner)
	if err != nil {
		log.Println("add runner error: ", err)
		return err
	}

	go tskw.assigningTask(&task)
	return nil
}

func (tskw *TaskWatcher) CheckAssignmentCompleted(assignmentID string) (bool, error) {

	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return false, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return false, err
	}

	assignmentIDBig, ok := new(big.Int).SetString(assignmentID, 10)
	if !ok {
		return false, errors.New("invalid task id")
	}

	isPending, err := workerHub.WorkerHubCaller.IsAssignmentPending(nil, assignmentIDBig)
	if err != nil {
		return false, err
	}

	if isPending {
		return false, nil
	}

	return true, nil
}

func (tskw *TaskWatcher) executeTasks() {
	for {
		task := <-tskw.taskQueue

		newRunner := tskw.GetRunner(task.TaskID)
		if newRunner == nil {
			log.Println("runner not found", task.TaskID)
			continue
		}

		err := tskw.modelManager.MakeReady(task.ModelContract)
		if err != nil {
			log.Println("make ready error: ", err)
			newRunner.SetDone()
			time.Sleep(1 * time.Second)
			continue
		}

		if tskw.mode == "miner" {
			isCompleted, err := tskw.CheckAssignmentCompleted(task.AssignmentID)
			if err != nil {
				log.Println("check task completed error: ", err)
				newRunner.SetDone()
				time.Sleep(1 * time.Second)
				continue
			}

			if isCompleted {
				newRunner.SetDone()

				log.Println("task already completed: ", task.TaskID)
				log.Println("task done: ", task.TaskID)
				continue
			}
			// assign task to worker
			err = tskw.executeWorkerTask(task)
			if err != nil {
				log.Println("execute worker task error: ", err)
				time.Sleep(10 * time.Second)
				newRunner.SetDone()
				continue
			}
		}
		if tskw.mode == "validator" {
			// assign task to validator
			err := tskw.executeVerifierTask(task)
			if err != nil {
				log.Println("execute validator task error: ", err)
				time.Sleep(10 * time.Second)
			}
		}

		// err = tskw.modelManager.PauseInstance(task.ModelContract)
		// if err != nil {
		// 	log.Println("pause instance error: ", err)
		// 	time.Sleep(1 * time.Second)
		// }

		newRunner.SetDone()

		log.Println("task done: ", task.TaskID)
		tskw.status.processedTasks++
		earning, _ := new(big.Int).SetString(task.Value, 10)
		tskw.status.currentEarning.Add(tskw.status.currentEarning, earning)
		// tskw.RemoveRunner(task.TaskID)
	}
}

func (tskw *TaskWatcher) RemoveRunner(taskID string) {
	tskw.runnerLock.Lock()
	delete(tskw.currentRunner, taskID)
	tskw.runnerLock.Unlock()
}

func (tskw *TaskWatcher) GetRunner(taskID string) *runner.RunnerInstance {
	tskw.runnerLock.RLock()
	defer tskw.runnerLock.RUnlock()
	return tskw.currentRunner[taskID]
}

func (tskw *TaskWatcher) CleanupRunners() {
	tskw.runnerLock.Lock()
	defer tskw.runnerLock.Unlock()
	for taskID, runner := range tskw.currentRunner {
		if runner.IsDone() {
			delete(tskw.currentRunner, taskID)
		}
	}
}

func (tskw *TaskWatcher) AddRunner(taskID string, runnerInst *runner.RunnerInstance) error {
	tskw.runnerLock.Lock()
	defer tskw.runnerLock.Unlock()
	if tskw.currentRunner == nil {
		tskw.currentRunner = make(map[string]*runner.RunnerInstance)
	}

	if _, ok := tskw.currentRunner[taskID]; ok {
		return errors.New("task already exists")
	}

	tskw.currentRunner[taskID] = runnerInst
	return nil
}

func (tskw *TaskWatcher) SubmitResult(assignmentID string, result []byte) error {
	// submit result to contract
	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return err
	}

	assignmentIDBig, ok := new(big.Int).SetString(assignmentID, 10)
	if !ok {
		return errors.New("invalid task id")
	}

	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting chain ID"))
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting gas price"))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(workerAcc, chainID)
	if err != nil {
		return errors.Join(err, errors.New("Error while creating new keyed transactor"))
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.SubmitSolution(auth, assignmentIDBig, result)
	if err != nil {
		return errors.Join(err, errors.New("Error while submitting result"))
	}

	log.Println("submit result tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}

	return nil
}

func (tskw *TaskWatcher) GetWorkerInfo() (*types.WorkerInfo, error) {
	var workerInfo types.WorkerInfo

	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return nil, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return nil, err
	}

	_, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return nil, errors.Join(err, errors.New("Error while getting account info"))
	}

	info, err := workerHub.WorkerHubCaller.Miners(nil, *address)
	if err != nil {
		return nil, errors.Join(err, errors.New("Error while getting worker info"))
	}

	pendingUnstake, err := workerHub.WorkerHubCaller.MinerUnstakeRequests(nil, *address)
	if err != nil {
		return nil, errors.Join(err, errors.New("Error while getting pending unstake"))
	}

	minStake, err := workerHub.WorkerHubCaller.MinerMinimumStake(nil)
	if err != nil {
		return nil, errors.Join(err, errors.New("Error while getting minimum stake"))
	}
	stakeStatus := "staked"
	if info.Stake.Cmp(minStake) < 0 {
		stakeStatus = "not enough staked"
	}
	if info.Stake.Cmp(minStake) == 0 {
		stakeStatus = "staked"
	}

	// tskw.status.pendingUnstakeAmount = pendingUnstake.Stake
	// tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)

	// tskw.status.assignModel = workerInfo.ModelAddress.Hex()
	// tskw.status.stakedAmount = workerInfo.Stake
	stakedAmount := new(big.Float).SetInt(info.Stake)
	stakedAmount = new(big.Float).Quo(stakedAmount, big.NewFloat(1e18))
	pendingUnstakeAmount := new(big.Float).SetInt(pendingUnstake.Stake)
	pendingUnstakeAmount = new(big.Float).Quo(pendingUnstakeAmount, big.NewFloat(1e18))

	workerInfo.Address = address.String()
	workerInfo.StakeStatus = stakeStatus
	workerInfo.StakedAmount = stakedAmount.String()
	workerInfo.PendingUnstakeAmount = pendingUnstakeAmount.String()
	workerInfo.PendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0).Format("2006-01-02 15:04:05")
	workerInfo.AssignModel = strings.ToLower(info.ModelAddress.Hex())

	return &workerInfo, nil
}

func (tskw *TaskWatcher) checkRegisteredAndStaked() error {

	for {
		time.Sleep(1 * time.Second)
		staked, err := tskw.isStaked()
		if err != nil {
			log.Println("isStaked error: ", err)
			continue
		}

		if !staked {
			tskw.status.stakeStatus = "not staked"
			tskw.status.assignModel = "-"
			err = tskw.stakeForWorker()
			if err != nil {
				log.Println("stakeForWorker error: ", err)
				return err
			}
		}

		if staked {
			tskw.status.stakeStatus = "staked"
			break
		}

	}

	return nil
}

func (tskw *TaskWatcher) isStaked() (bool, error) {
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return false, err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return false, err
	}

	_, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return false, errors.Join(err, errors.New("Error while getting account info"))
	}

	log.Println("check staked for: ", address.String())

	workerInfo, err := workerHub.WorkerHubCaller.Miners(nil, *address)
	if err != nil {
		return false, errors.Join(err, errors.New("Error while getting worker info"))
	}

	minStake, err := workerHub.WorkerHubCaller.MinerMinimumStake(nil)
	if err != nil {

		return false, errors.Join(err, errors.New("Error while getting minimum stake"))
	}

	if workerInfo.Stake.Cmp(minStake) < 0 {
		return false, nil
	}

	pendingUnstake, err := workerHub.WorkerHubCaller.MinerUnstakeRequests(nil, *address)
	if err != nil {
		return false, errors.Join(err, errors.New("Error while getting pending unstake"))
	}

	tskw.status.pendingUnstakeAmount = pendingUnstake.Stake
	tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)

	tskw.status.assignModel = workerInfo.ModelAddress.Hex()
	tskw.status.stakedAmount = workerInfo.Stake
	return true, nil
}

func (tskw *TaskWatcher) stakeForWorker() error {
	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return err
	}

	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	minStake, err := workerHub.WorkerHubCaller.MinerMinimumStake(nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting minimum stake"))
	}

	log.Printf("start staking for %v with value: %v\n", address.String(), minStake.String())

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting chain ID"))
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting gas price"))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(workerAcc, chainID)
	if err != nil {
		return errors.Join(err, errors.New("Error while creating new keyed transactor"))
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = minStake     // in wei
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.RegisterMiner(auth, 1)
	if err != nil {
		return errors.Join(err, errors.New("Error while staking"))
	}

	log.Println("stake tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}
	log.Println("staking success")

	return nil
}

func (tskw *TaskWatcher) SubmitTask(task *types.TaskSubmitRequest) error {
	// submit task to contract

	log.Println("submit task: ", task.ModelContract)

	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	modelContract := common.HexToAddress(task.ModelContract)

	model, err := eaimodel.NewHybridModel(modelContract, ethClient)
	if err != nil {
		return err
	}

	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting chain ID"))
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting gas price"))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(workerAcc, chainID)
	if err != nil {
		return errors.Join(err, errors.New("Error while creating new keyed transactor"))
	}

	value, ok := new(big.Int).SetString(task.Value, 10)
	if !ok {
		return errors.New("invalid model id")
	}

	//TODO: @liam check if task value is enough
	// inferenceCost, err := model.HybridModelCaller.InferenceCost(nil)
	// if err != nil {
	// 	return errors.Join(err, errors.New("Error while getting inference cost"))
	// }

	// if value.Cmp(inferenceCost) < 0 {
	// 	return errors.New("value is not enough " + value.String() + " " + inferenceCost.String())
	// }

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value        // in wei
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice = gasPrice

	data := []byte(task.Params)

	tx, err := model.HybridModelTransactor.Infer(auth, data)
	if err != nil {
		testabi, err := ethabi.JSON(strings.NewReader(abi.WorkerHubABI))
		byteCode, err := testabi.Pack("infer", data, address)
		if err != nil {
			return errors.Join(err, errors.New("Error while packing data"))
		}
		tx := ethtypes.NewTransaction(nonce, modelContract, value, 0, gasPrice, byteCode)
		signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), workerAcc)
		if err != nil {
			return errors.Join(err, errors.New("Error while signing tx"))
		}

		signedTxBytes, err := signedTx.MarshalJSON()
		if err != nil {
			return errors.Join(err, errors.New("Error while marshalling tx"))
		}

		return errors.Join(err, errors.New(fmt.Sprintf("Error while submitting task: %v", string(signedTxBytes))))
	}

	log.Println("submit task tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}

	return nil
}

func (tskw *TaskWatcher) GetWorkerAddress() string {
	return tskw.address
}

func (tskw *TaskWatcher) GetWorkerBalance() string {
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		log.Println("get eth client error: ", err)
		return ""
	}

	bal, err := ethClient.BalanceAt(context.Background(), common.HexToAddress(tskw.address), nil)
	if err != nil {
		log.Println("get balance error: ", err)
		return ""
	}

	balFloat := new(big.Float).SetInt(bal)
	balFloat = new(big.Float).Quo(balFloat, big.NewFloat(1e18))

	return balFloat.String()
}

func (tskw *TaskWatcher) ReclaimStake() error {
	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return err
	}

	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting chain ID"))
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting gas price"))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(workerAcc, chainID)
	if err != nil {
		return errors.Join(err, errors.New("Error while creating new keyed transactor"))
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.UnstakeForMiner(auth)
	if err != nil {
		return errors.Join(err, errors.New("Error while unsstaking"))
	}

	log.Println("unstake tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}

	log.Println("unstake success")

	return nil
}

func (tskw *TaskWatcher) Unregister() error {
	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return err
	}

	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting chain ID"))
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting gas price"))
	}
	auth, err := bind.NewKeyedTransactorWithChainID(workerAcc, chainID)
	if err != nil {
		return errors.Join(err, errors.New("Error while creating new keyed transactor"))
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.UnregisterMiner(auth)
	if err != nil {
		return errors.Join(err, errors.New("Error while unregister miner"))
	}

	log.Println("unregister tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}

	log.Println("unregister success")

	return nil
}

func (tskw *TaskWatcher) UnregisterAndQuit() error {
	tskw.unstakeLock.Lock()
	defer tskw.unstakeLock.Unlock()
	err := tskw.Unregister()
	if err != nil {
		return err
	}
	tskw.status.stakeStatus = "not staked"

	// for testnet only
	// time.Sleep(61 * time.Second)
	// err = tskw.ReclaimStake()
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (tskw *TaskWatcher) GetUnstakeInfo() (string, time.Time) {
	amount := new(big.Float).SetInt(tskw.status.pendingUnstakeAmount)
	amount = new(big.Float).Quo(amount, big.NewFloat(1e18))
	return amount.String(), tskw.status.pendingUnstakeUnlockAt
}

func (tskw *TaskWatcher) GetProcessedTasks() uint64 {
	return tskw.status.processedTasks
}

func (tskw *TaskWatcher) GetSessionEarning() string {
	earning := new(big.Float).SetInt(tskw.status.currentEarning)
	earning = new(big.Float).Quo(earning, big.NewFloat(1e18))
	return earning.String()
}

func (tskw *TaskWatcher) GetStakedAmount() string {
	amount := new(big.Float).SetInt(tskw.status.stakedAmount)
	amount = new(big.Float).Quo(amount, big.NewFloat(1e18))
	return amount.String()
}

func (tskw *TaskWatcher) GetAssignedModel() string {
	if tskw.status.assignModel == "" {
		return "-"
	}
	return tskw.status.assignModel
}

func (tskw *TaskWatcher) StakeStatus() string {
	return tskw.status.stakeStatus
}
