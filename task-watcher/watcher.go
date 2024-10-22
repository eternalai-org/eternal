package watcher

import (
	"context"
	"errors"
	"eternal-infer-worker/coordinator"
	"eternal-infer-worker/libs"
	"eternal-infer-worker/libs/abi"
	"eternal-infer-worker/libs/eaimodel"
	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/libs/github"
	"eternal-infer-worker/manager"
	"eternal-infer-worker/runner"
	"eternal-infer-worker/tui"
	"eternal-infer-worker/types"
	"fmt"
	log "github.com/sirupsen/logrus"
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

const DEFAULT_GAS_LIMIT = 200_000_000

type NetworkConfig struct {
	RPC string
	WS  string
}

type TaskWatcherStatus struct {
	processedTasks         uint64
	balance                *big.Int
	currentEarning         *big.Int
	stakeStatus            string
	stakedAmount           *big.Int
	pendingUnstakeAmount   *big.Int
	pendingUnstakeUnlockAt time.Time
	assignModel            string
	miningRewardAmount     *big.Int
	modelStatus            string
}

type TaskWatcher struct {
	version      string
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

	status     TaskWatcherStatus
	globalInfo types.HubGlobalInfo
	mode       string

	unstakeLock sync.Mutex

	hasNewVersion bool
	newVersion    string

	zkSync           bool
	paymasterAddr    string
	paymasterToken   string
	paymasterFeeZero bool
}

func NewTaskWatcher(networkCfg NetworkConfig, version, taskContract, account,
	modelsDir, lighthouseAPI, mode string, id, numOfWorker int,
	modelManager *manager.ModelManager,
	coordinator *coordinator.Coordinator,
	zkSycn bool,
	paymasterAddr string,
	paymasterToken string,
	paymasterZeroFee bool) (*TaskWatcher, error) {

	_, address, err := eth.GenerateAddressFromPrivKey(account)
	if err != nil {
		return nil, err
	}

	return &TaskWatcher{
		zkSync:           zkSycn,
		paymasterAddr:    paymasterAddr,
		paymasterToken:   paymasterToken,
		paymasterFeeZero: paymasterZeroFee,
		networkCfg:       networkCfg,
		modelsDir:        modelsDir,
		version:          version,
		account:          account,
		address:          strings.ToLower(address),
		taskContract:     taskContract,
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
			balance:              big.NewInt(0),
			currentEarning:       big.NewInt(0),
			stakedAmount:         big.NewInt(0),
			pendingUnstakeAmount: big.NewInt(0),
			miningRewardAmount:   big.NewInt(0),
		},
	}, nil
}

func (tskw *TaskWatcher) Start() {

	log.Println("start task watcher")
	if tui.UI != nil {
		tui.UI.UpdateSectionText(tui.UIMessageData{
			Section: tui.UISectionStatusText,
			Color:   "normal",
			Text:    "starting task watcher...",
		})
	}
	go tskw.executeTasks()
	go tskw.watchWorkerInfo()
	go tskw.watchHubGlobalInfo()

	//there is no action in this function, turn off it. Use the infinity loop in main instead.
	//go tskw.watchNewVersion()

	tskw.GetWorkerInfo()
	go tskw.watchAssignedModel()
	var err error
	// err = tskw.modelManager.PreloadModels([]string{tskw.status.assignModel})
	// if err != nil {
	// 	log.Println("preload models error: ", tskw.status.assignModel, err)
	// 	panic(err)
	// }

	if tskw.mode == libs.MODE_MINER {
		staked, _ := tskw.isStaked()
		if !staked {
			err = tskw.stakeForWorker()
			if err != nil {
				log.Error("register error: ", err)
			}
		}
		err = tskw.joinForMinting()
		if err != nil {
			log.Error("join for minting error: ", err)
		}
	}

	tskw.watchAndAssignTask()
}

func (tskw *TaskWatcher) watchAssignedModel() {
	for {
		time.Sleep(2 * time.Second)
		if tskw.status.assignModel == "" || tskw.status.assignModel == "-" || tskw.status.assignModel == "0x0000000000000000000000000000000000000000" {
			continue
		}

		log.Println("[watchAssignedModel].watchAssignedModel - assignModel ", tskw.status.assignModel)
		currentLoadedModels := tskw.modelManager.GetLoadeModels()
		if _, ok := currentLoadedModels[tskw.status.assignModel]; ok {
			time.Sleep(5 * time.Second)
			continue
		}

		err := tskw.modelManager.PreloadModels([]string{tskw.status.assignModel})
		if err != nil {
			log.Println("[watchAssignedModel][Err] preload models error: ", tskw.status.assignModel, err)
			continue
		}
	}
}

func (tskw *TaskWatcher) watchWorkerInfo() {
	for {
		time.Sleep(2 * time.Second)
		_, err := tskw.GetWorkerInfo()
		if err != nil {
			log.Println("get worker info error: ", err)
			continue
		}
	}
}

func (tskw *TaskWatcher) joinForMinting() error {
	if tskw.zkSync {
		return tskw.joinForMintingZk()
	}
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
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
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
	maxConcurrentTask := 10

	if tui.UI != nil {
		tui.UI.UpdateSectionText(tui.UIMessageData{
			Section: tui.UISectionStatusText,
			Text:    "running...",
			Color:   "normal",
		})
	}
	for {
		time.Sleep(1 * time.Second)
		tskw.CleanupRunners()
		if tskw.modelManager.GetStatus() != "ready" {

		}

		var tasks []types.TaskInfo
		var err error
		if tskw.zkSync {
			tasks, err = tskw.getPendingTaskFromContractZk()
			if err != nil {
				log.Println("[watchAndAssignTask][ERR] get pending task error: ", err)
				continue
			}
		} else {
			tasks, err = tskw.getPendingTaskFromContract()
			if err != nil {
				log.Println("[watchAndAssignTask][ERR] get pending task error: ", err)
				continue
			}
		}

		// sort tasks by task id in descending order
		sort.SliceStable(tasks, func(i, j int) bool {
			taskAid, _ := new(big.Int).SetString(tasks[i].TaskID, 10)
			taskBid, _ := new(big.Int).SetString(tasks[j].TaskID, 10)
			return taskAid.Cmp(taskBid) > 0
		})

		log.Println("[watchAndAssignTask] pending tasks: ", len(tasks))

		for _, task := range tasks {
			if len(tskw.currentRunner) >= maxConcurrentTask {
				log.Println("[watchAndAssignTask] worker is full, current tasks: ", len(tskw.currentRunner))
				break
			}

			log.Println("[watchAndAssignTask] assign task: ", task.TaskID, task.ModelContract, task.Params)
			err = tskw.AssignTask(task)
			if err != nil {
				log.Println("[watchAndAssignTask] assign task error: ", err)
			}
		}

		if len(tasks) == 0 {
			modelAddr := tskw.GetAssignedModel()
			//log.Warning("Watcher: Maybe node is slashed for model ", modelAddr)
			if ok := tskw.isMinerOfModel(common.HexToAddress(modelAddr)); !ok {
				log.Info("Watcher: node need to be reJoinMinting for model ", modelAddr)
				err := tskw.reJoinMinting(modelAddr)
				if err != nil {
					log.Error("Watcher:reJoinMinting error", err)
				}
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
		iter, err := workerHub.WorkerHubFilterer.FilterNewInference(opts, []*big.Int{requestIDInt}, nil, nil)
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
	if tskw.mode == libs.MODE_VALIDATOR {
		return nil, errors.New("not support validator")
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
	log.Println("[TaskWatcher].assigningTask - received task ,ModelContract: ", task.ModelContract, " ,TaskID: ", task.TaskID, " ,TaskMode: ", tskw.mode)
	tskw.taskQueue <- task
}

func (tskw *TaskWatcher) AssignTask(task types.TaskInfo) error {
	log.Println("[TaskWatcher].AssignTask")
	newRunner, err := runner.NewRunnerInstance(tskw.modelManager, &task)
	if err != nil {
		log.Error("create runner error: ", err)
		return err
	}

	log.Println("[TaskWatcher].AssignTask - task.TaskID: ", task.TaskID)
	err = tskw.AddRunner(task.TaskID, newRunner)
	if err != nil {
		log.Error("[TaskWatcher].AssignTask - add runner error: ", err)
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
	log.Println("[TaskWatcher].executeTasks")
	for {
		task := <-tskw.taskQueue
		log.Println("[TaskWatcher].executeTasks - received task")
		newRunner := tskw.GetRunner(task.TaskID)
		if newRunner == nil {
			log.Println("runner not found", task.TaskID)
			continue
		}
		if newRunner.IsDone() {
			log.Println("runner is done", task.TaskID)
			continue
		}

		log.Println("[TaskWatcher].executeTasks - received task ,ModelContract: ", task.ModelContract, " ,TaskID: ", task.TaskID)
		err := tskw.modelManager.MakeReady(task.ModelContract)
		if err != nil {
			log.Println("make ready error: ", err)
			newRunner.SetDone()
			time.Sleep(1 * time.Second)
			continue
		}

		mode := tskw.mode
		if task.ZKSync {
			mode = task.AssignmentRole
		}

		log.Info(fmt.Sprintf("[TaskWatcher].executeTasks ------ Process task %s with mode %s", task.TaskID, mode))

		switch mode {
		case libs.MODE_MINER:
			{
				isCompleted, err := tskw.CheckAssignmentCompleted(task.AssignmentID)
				if err != nil {
					log.Println("[TaskWatcher].executeTasks check task completed error: ", err)
					newRunner.SetDone()
					time.Sleep(1 * time.Second)
					continue
				}

				if isCompleted {
					newRunner.SetDone()

					log.Println("[TaskWatcher].executeTasks - task already completed: ", task.TaskID)
					log.Println("[TaskWatcher].executeTasks - task done: ", task.TaskID)
					continue
				}
				// assign task to worker
				err = tskw.executeWorkerTask(task)
				if err != nil {
					log.Println("[TaskWatcher].executeTasks - execute worker task error: ", err)
					time.Sleep(10 * time.Second)
					newRunner.SetDone()
					continue
				}
			}
		case libs.MODE_VALIDATOR:
			{
				// assign task to validator
				err := tskw.executeVerifierTask(task)
				if !task.ZKSync {
					if err != nil {
						log.Error("validator [TaskWatcher].executeTasks - execute validator task error: ", err)
						time.Sleep(10 * time.Second)
					}
					newRunner.SetDone()
					log.Println("validator [TaskWatcher].executeTasks - task done: ", task.TaskID)
				} else {
					switch task.Status {
					case ContractInferenceStatusReveal:
						{
							if err == nil {
								newRunner.SetDone()
								log.Info("validator [TaskWatcher].executeTasks - task send reveal done: ", task.TaskID)
							} else {
								log.Info("validator [TaskWatcher].executeTasks - task send reveal error: ", task.TaskID, err)
							}
						}
					case ContractInferenceStatusCommit:
						{
							if err == nil {
								log.Info(fmt.Sprintf("validator [TaskWatcher].executeTasks - task send commit done: %s, newRunner.IsDone=%v", task.TaskID, newRunner.IsDone()))
								newRunner.SetNotDone()
								task.Retry = 0
							} else {
								log.Info("validator [TaskWatcher].executeTasks - task send commit error: ", task.TaskID, err)
							}
						}
					}
					if !newRunner.IsDone() {
						if task.Retry <= 500 {
							task.Retry++
							log.Info(fmt.Sprintf("validator [TaskWatcher].executeTasks set task %s status %d into queue again retry %d", task.TaskID, task.Status, task.Retry))
							tskw.taskQueue <- task // set again
						}
					}
				}
			}
		}
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
	if tskw.zkSync {
		log.Info(fmt.Sprintf("SubmitResult ZK for assigmentID %s result %s", assignmentID, string(result)))
		return tskw.SubmitResultZk(assignmentID, result)
	}
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
	auth.Value = big.NewInt(DEFAULT_GAS_LIMIT) // in wei
	auth.GasLimit = uint64(0)                  // in units
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

	receipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while getting tx receipt"))
	}

	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
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

	miningReward, err := abi.NewMiningReward(common.HexToAddress(tskw.taskContract), ethClient)
	if err != nil {
		return nil, err
	}

	rewardToClaim, err := miningReward.MiningRewardCaller.RewardToClaim(nil, *address)
	if err != nil {
		return nil, err
	}

	bal, err := ethClient.BalanceAt(context.Background(), common.HexToAddress(tskw.address), nil)
	if err != nil {
		log.Println("get balance error: ", err)
		return nil, err
	}

	balFloat := new(big.Float).SetInt(bal)
	balFloat = new(big.Float).Quo(balFloat, big.NewFloat(1e18))

	stakedAmount := new(big.Float).SetInt(info.Stake)
	stakedAmount = new(big.Float).Quo(stakedAmount, big.NewFloat(1e18))
	pendingUnstakeAmount := new(big.Float).SetInt(pendingUnstake.Stake)
	pendingUnstakeAmount = new(big.Float).Quo(pendingUnstakeAmount, big.NewFloat(1e18))
	rewardToClaimAmount := new(big.Float).SetInt(rewardToClaim)
	rewardToClaimAmount = new(big.Float).Quo(rewardToClaimAmount, big.NewFloat(1e18))
	balanceAmount := new(big.Float).SetInt(bal)
	balanceAmount = new(big.Float).Quo(balanceAmount, big.NewFloat(1e18))

	workerInfo.Address = address.String()
	workerInfo.StakeStatus = stakeStatus
	workerInfo.StakedAmount = stakedAmount.String()
	workerInfo.PendingUnstakeAmount = pendingUnstakeAmount.String()
	workerInfo.PendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0).Format("2006-01-02 15:04:05")
	workerInfo.AssignModel = strings.ToLower(info.ModelAddress.Hex())
	workerInfo.MiningReward = rewardToClaimAmount.String()
	workerInfo.Balance = balanceAmount.String()

	tskw.status.stakeStatus = stakeStatus
	tskw.status.stakedAmount = info.Stake
	tskw.status.pendingUnstakeAmount = pendingUnstake.Stake
	if tskw.status.pendingUnstakeAmount.Cmp(new(big.Int).SetUint64(0)) > 0 {
		tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)
	}
	tskw.status.assignModel = strings.ToLower(info.ModelAddress.Hex())
	tskw.status.miningRewardAmount = rewardToClaim
	tskw.status.balance = bal

	return &workerInfo, nil
}

func (tskw *TaskWatcher) GetMiningReward() string {
	rewardAmount := new(big.Float).SetInt(tskw.status.miningRewardAmount)
	rewardAmount = new(big.Float).Quo(rewardAmount, big.NewFloat(1e18))
	return rewardAmount.String()
}

func (tskw *TaskWatcher) GetUnskateAmount() (string, error) {

	pendingUnstakeAmount := new(big.Float).SetInt(tskw.status.pendingUnstakeAmount)
	pendingUnstakeAmount = new(big.Float).Quo(pendingUnstakeAmount, big.NewFloat(1e18))

	return pendingUnstakeAmount.String(), nil

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
	if tskw.status.pendingUnstakeAmount.Cmp(new(big.Int).SetUint64(0)) > 0 {
		tskw.status.pendingUnstakeUnlockAt = time.Unix(pendingUnstake.UnlockAt.Int64(), 0)
	}

	tskw.status.assignModel = workerInfo.ModelAddress.Hex()
	tskw.status.stakedAmount = workerInfo.Stake
	return true, nil
}

func (tskw *TaskWatcher) stakeForWorker() error {
	if tskw.zkSync {
		//err := tskw.approveErc20Zk()
		//if err != nil {
		//	log.Println("register error: ", err)
		//	return err
		//}
		return tskw.stakeForWorkerZk()
	}
	ctx := context.Background()
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		return err
	}

	hubAddress := common.HexToAddress(tskw.taskContract)
	log.Printf("workerhub contract addr: %v \n", tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		return err
	}

	log.Printf("private key: %v \n", tskw.account)
	workerAcc, address, err := eth.GetAccountInfo(tskw.account)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting account info"))
	}

	minStake, err := workerHub.WorkerHubCaller.MinerMinimumStake(nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting minimum stake"))
	}
	// Get the balance
	balance, err := ethClient.BalanceAt(context.Background(), *address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	log.Printf("Start staking for address %v - Balance %v with value: %v\n", address.String(), balance.String(), minStake.String())

	nonce, err := ethClient.NonceAt(ctx, *address, nil)
	if err != nil {
		return errors.Join(err, errors.New("Error while getting nonce"))
	}

	chainID, err := ethClient.NetworkID(context.Background())
	log.Printf("RPC: %v ---- ChainID: %v\n", tskw.networkCfg.RPC, chainID)
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
	auth.Value = minStake                     // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
	auth.GasPrice = gasPrice
	log.Printf("GasPrice:%v Nonce:%v Value:%v \n", auth.GasPrice.String(), auth.Nonce, auth.Value)

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

func (tskw *TaskWatcher) Restake() error {
	if tskw.zkSync {
		return tskw.RestakeZk()
	}
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
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.RestakeForMiner(auth, 1)
	if err != nil {
		return errors.Join(err, errors.New("Error while re-staking"))
	}

	log.Println("restake tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}
	log.Println("restaking success")

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
	auth.Value = value                        // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
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
	amount := new(big.Float).SetInt(tskw.status.balance)
	amount = new(big.Float).Quo(amount, big.NewFloat(1e18))
	return amount.String()
}

func (tskw *TaskWatcher) ReclaimStake() error {
	if tskw.zkSync {
		return tskw.ReclaimStakeZk()
	}
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
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
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
	if tskw.zkSync {
		return tskw.UnregisterZk()
	}
	tskw.unstakeLock.Lock()
	defer tskw.unstakeLock.Unlock()
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
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
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
	tskw.status.stakeStatus = "not staked"

	return nil
}

func (tskw *TaskWatcher) ClaimMiningReward() error {
	if tskw.zkSync {
		return tskw.ClaimMiningRewardZk()
	}
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
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(DEFAULT_GAS_LIMIT) // in units
	auth.GasPrice = gasPrice

	tx, err := workerHub.WorkerHubTransactor.ClaimReward(auth, *address)
	if err != nil {
		return errors.Join(err, errors.New("Error while claim reward miner"))
	}

	log.Println("claim reward tx: ", tx.Hash().Hex())

	err = eth.WaitForTx(ethClient, tx.Hash())
	if err != nil {
		return errors.Join(err, errors.New("Error while waiting for tx"))
	}

	log.Println("claim reward success")

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

func (tskw *TaskWatcher) GetAssignedTasks() uint64 {
	if tskw.zkSync {
		return tskw.countAssignmentByMiner()
	}
	return 0
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

func (tskw *TaskWatcher) GetStakeStatus() (string, error) {
	_, err := tskw.isStaked()
	if err != nil {
		return "", err
	}
	return tskw.status.stakeStatus, nil

}

func (tskw *TaskWatcher) GetAssignedModel() string {
	if tskw.status.assignModel == "" {
		return "-"
	}
	return tskw.status.assignModel
}

func (tskw *TaskWatcher) GetModelStatus() string {
	return tskw.modelManager.GetStatus()
}

func (tskw *TaskWatcher) StakeStatus() string {
	return tskw.status.stakeStatus
}

func (tskw *TaskWatcher) watchNewVersion() {
	err := tskw.checkNewVersion()
	if err != nil {
		log.Println("check new version error: ", err)
	}
	for {
		time.Sleep(120 * time.Second)
		err := tskw.checkNewVersion()
		if err != nil {
			log.Println("check new version error: ", err)
		}
	}
}

func (tskw *TaskWatcher) checkNewVersion() error {
	releaseInfo, err := github.GetLatestRelease()
	if err != nil {
		return err
	}

	if releaseInfo.TagName != tskw.version {
		tskw.hasNewVersion = true
		tskw.newVersion = releaseInfo.TagName
	} else {
		tskw.hasNewVersion = false
		tskw.newVersion = ""
	}

	return nil
}

func (tskw *TaskWatcher) HasNewVersion() (bool, string) {
	return tskw.hasNewVersion, tskw.newVersion
}
