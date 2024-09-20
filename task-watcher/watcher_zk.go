package watcher

import (
	"errors"
	"eternal-infer-worker/libs/eth"
	"eternal-infer-worker/libs/zkabi"
	"eternal-infer-worker/libs/zkclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
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
		"registerMiner", 1, minStake,
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
