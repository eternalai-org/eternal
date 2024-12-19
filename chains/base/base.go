package base

import "eternal-infer-worker/chains/interfaces"

type Base struct {
	interfaces.Chain
	interfaces.Staking
}

func NewBaseChain(rpc string) (*Base, error) {
	b := new(Base)
	err := b.Connect(rpc)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Base) GetPendingTasks(fromblock, toBlock int64) (*interfaces.Tasks, error) {
	return nil, nil
}

func (b *Base) IsStaked() (bool, error) {
	return false, nil
}

func (b *Base) StakeForWorker() (bool, error) {
	return false, nil
}

func (b *Base) JoinForMinting() (bool, error) {
	return false, nil
}
