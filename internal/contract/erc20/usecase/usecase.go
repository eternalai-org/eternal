package usecase

import (
	"context"
	"math/big"

	"eternal/internal/contract/erc20"
	"eternal/internal/core/ports"
	"eternal/pkg/eth"
	"eternal/pkg/logger"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type contractErc20Usecase struct {
	contractInstance *erc20.Erc20
	contractClient   *ethclient.Client

	contractAddress common.Address
	contractRpc     string
}

func NewContractErc20Usecase(contractRpc, contractAddress string) (ports.IContractErc20Usecase, error) {
	uc := &contractErc20Usecase{
		contractRpc:     contractRpc,
		contractAddress: common.HexToAddress(contractAddress),
	}
	client, err := newClient(contractRpc)
	if err != nil {
		return nil, err
	}

	uc.contractClient = client
	ct, err := erc20.NewErc20(uc.contractAddress, client)
	if err != nil {
		return nil, err
	}

	uc.contractInstance = ct
	ctx := context.Background()
	if err := uc.healthCheck(ctx); err != nil {
		return nil, err
	}
	return uc, nil
}

func (uc *contractErc20Usecase) BalanceOfAddress(ctx context.Context, address string) (*big.Int, error) {
	balanace, err := uc.contractInstance.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		balanace, err = uc.contractClient.BalanceAt(ctx, common.HexToAddress(address), nil)
		if err != nil {
			return nil, err
		}
	}
	return balanace, nil
}

func newClient(contractRpc string) (*ethclient.Client, error) {
	client, err := eth.NewEthClient(contractRpc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (uc *contractErc20Usecase) getChainID(ctx context.Context) (uint64, error) {
	chainID, err := uc.contractClient.ChainID(ctx)
	if err != nil {
		return 0, err
	}
	return chainID.Uint64(), nil
}

func (uc *contractErc20Usecase) healthCheck(ctx context.Context) error {
	chainId, err := uc.getChainID(ctx)
	if err != nil {
		return err
	}

	logger.AtLog.Infof("contractErc20Usecase#ContractAddress %s - chainId: %d - version: %s", uc.contractAddress, chainId, "")
	return nil
}
