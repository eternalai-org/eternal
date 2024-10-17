package watcher

import (
	"eternal-infer-worker/libs/abi"
	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/types"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (tskw *TaskWatcher) watchHubGlobalInfo() {
	for {
		tskw.getGlobalInfo()
		time.Sleep(10 * time.Second)
	}

}

func (tskw *TaskWatcher) getGlobalInfo() {
	ethClient, err := eth.NewEthClient(tskw.networkCfg.RPC)
	if err != nil {
		log.Println("get eth client error: ", err)
		return
	}

	hubAddress := common.HexToAddress(tskw.taskContract)

	workerHub, err := abi.NewWorkerHub(hubAddress, ethClient)
	if err != nil {
		log.Println("get worker hub error: ", err)
		return
	}

	miners, err := workerHub.WorkerHubCaller.GetMiners(nil)
	if err != nil {
		log.Println("get miners error: ", err)
		return
	}

	//validators, err := workerHub.WorkerHubCaller.GetValidators(nil)
	//if err != nil {
	//	log.Println("get miners error: ", err)
	//	return
	//}

	//models, err := workerHub.WorkerHubCaller.GetModelAddresses(nil)
	//if err != nil {
	//	log.Println("get miners error: ", err)
	//	return
	//}
	//feePercent, err := workerHub.WorkerHubCaller.FeePercentage(nil)
	//if err != nil {
	//	log.Println("get fee percent error: ", err)
	//	return
	//}

	//unstakeDelay, err := workerHub.WorkerHubCaller.UnstakeDelayTime(nil)
	//if err != nil {
	//	log.Println("get unstake delay error: ", err)
	//	return
	//}
	//unstakeDelayTime, err := durafmt.ParseString(unstakeDelay.String() + "s")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// for idx, v := range miners {
	// 	log.Println("miner: ", idx, v.WorkerAddress.String())
	// }

	tskw.globalInfo = types.HubGlobalInfo{
		//TotalValidators: uint64(len(validators)),
		TotalMiners: uint64(len(miners)),
		//TotalModels:     uint64(len(models)),
		//FeePercent:   feePercent,
		//UnstakeDelay: unstakeDelayTime.String(),
	}

}

func (tskw *TaskWatcher) GetHubGlobalInfo() (*types.HubGlobalInfo, error) {
	a := tskw.globalInfo
	return &a, nil
}
