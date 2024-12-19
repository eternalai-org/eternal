package ports

import (
	"context"
	"math/big"
)

type IContractErc20Usecase interface {
	BalanceOfAddress(ctx context.Context, address string) (*big.Int, error)
}
