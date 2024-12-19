package interfaces

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Tasks struct{}

type Chain struct {
	Client     *ethclient.Client
	PrivateKey string
	Address    *common.Address
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
	IStaking
	ITasks
}

type ITasks interface {
	Connect(rpc string) error
	GetPendingTasks(fromblock, toBlock int64) ([]*Tasks, error)
	SubmitTask()
}

type IStaking interface {
	IsStaked() (bool, error)
	StakeForWorker() error
	JoinForMinting() error
}
