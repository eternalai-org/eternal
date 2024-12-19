package base

import (
	"context"
	"eternal-infer-worker/chains/base/contract"
	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs/eth"

	"github.com/ethereum/go-ethereum/common"
)

type Base struct {
	interfaces.Chain
	StakingHub           *contract.BaseWhAbi
	StakingHubAddress    string
	Erc20contractAddress string
	Erc20contract        *contract.Abi
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
	sthub, err := contract.NewBaseWhAbi(common.HexToAddress(b.StakingHubAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.StakingHub = sthub

	b.Erc20contractAddress = "0x4b6bf1d365ea1a8d916da37fafd4ae8c86d061d7"
	erc20, err := contract.NewAbi(common.HexToAddress(b.Erc20contractAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.Erc20contract = erc20
	return b, nil
}

func (b *Base) GetPendingTasks(fromblock, toBlock int64) (*interfaces.Tasks, error) {
	return nil, nil
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

	err = eth.ApproveERC20(ctx, b.Client, b.PrivateKey, common.HexToAddress(b.StakingHubAddress), common.HexToAddress(b.Erc20contractAddress))
	if err != nil {
		return err
	}

	minStake, err := b.StakingHub.MinerMinimumStake(nil)
	if err != nil {
		return err
	}

	_ = balance
	_ = minStake
	return nil
}

func (b *Base) JoinForMinting() (bool, error) {
	return false, nil
}
