package base

import (
	"eternal-infer-worker/chains/base/contract"
	"eternal-infer-worker/chains/interfaces"
	"eternal-infer-worker/config"
	"eternal-infer-worker/libs/eth"

	"github.com/ethereum/go-ethereum/common"
)

type Base struct {
	interfaces.Chain
	StakingHub        *contract.BaseWhAbi
	StakingHubAddress string
}

func NewBaseChain(cnf *config.Config) (*Base, error) {
	b := new(Base)
	err := b.Connect(cnf.Rpc)
	if err != nil {
		return nil, err
	}

	b.Account = cnf.Account
	b.StakingHubAddress = cnf.StakingHubAddress

	sthub, err := contract.NewBaseWhAbi(common.HexToAddress(b.StakingHubAddress), b.Client)
	if err != nil {
		return nil, err
	}

	b.StakingHub = sthub
	return b, nil
}

func (b *Base) GetPendingTasks(fromblock, toBlock int64) (*interfaces.Tasks, error) {
	return nil, nil
}

func (b *Base) IsStaked() (bool, error) {
	_, address, err := eth.GetAccountInfo(b.Account)
	if err != nil {
		return false, err
	}

	workerInfo, err := b.StakingHub.Miners(nil, *address)
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

	pendingUnstake, err := b.StakingHub.MinerUnstakeRequests(nil, *address)
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

func (b *Base) StakeForWorker() (bool, error) {
	return false, nil
}

func (b *Base) JoinForMinting() (bool, error) {
	return false, nil
}
