package interfaces

import "github.com/ethereum/go-ethereum/ethclient"

type Tasks struct{}

type Chain struct {
	Client *ethclient.Client
}

func (c *Chain) Connect(rpc string) error {
	ethClient, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}

	c.Client = ethClient
	return nil
}

type IChain interface {
	Connect(rpc string) error
	GetPendingTasks(fromblock, toBlock int64) (*Tasks, error)
}

type IStaking interface {
	IsStaked() (bool, error)
	StakeForWorker() error
	JoinForMinting() error
}

type Staking struct{}

func (t *Staking) IsStaked() (bool, error) {
	return false, nil
}

func (t *Staking) StakeForWorker() error {
	return nil
}

func (t *Staking) JoinForMinting() error {
	return nil
}
