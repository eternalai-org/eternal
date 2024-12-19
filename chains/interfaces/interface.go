package interfaces

import "github.com/ethereum/go-ethereum/ethclient"

type Tasks struct{}

type Chain struct {
	Client  *ethclient.Client
	Account string
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
