package zkclient

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
	"math/big"
	"time"
)

type Client struct {
	BaseURL          string
	chainID          uint64
	PaymasterFeeZero bool
	PaymasterAddress string
	PaymasterToken   string
}

func NewZkClient(rpc string, paymentFeeZero bool, paymentMasterAddress string, paymentMasterToken string) *Client {
	return &Client{
		BaseURL:          rpc,
		PaymasterFeeZero: paymentFeeZero,
		PaymasterAddress: paymentMasterAddress,
		PaymasterToken:   paymentMasterToken,
	}
}

func (c *Client) GetZkClient() (clients.Client, error) {
	zkClient, err := clients.Dial(c.BaseURL)
	if err != nil {
		return nil, err
	}
	return zkClient, nil
}

func (c *Client) GetChainID() (uint64, error) {
	if c.chainID > 0 {
		return c.chainID, nil
	}
	client, err := c.GetZkClient()
	if err != nil {
		return 0, err
	}
	defer client.Close()
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return 0, err
	}
	c.chainID = chainID.Uint64()
	return c.chainID, nil
}

func (c *Client) GetGasPrice() (*big.Int, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func (c *Client) PopulateTransaction(ctx context.Context, address common.Address, tx accounts.Transaction) (*zktypes.Transaction712, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}
	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return nil, err
	}
	if tx.ChainID == nil {
		tx.ChainID = big.NewInt(int64(chainID))
	}
	tx.Nonce = new(big.Int).SetUint64(nonce)
	tx.GasFeeCap = gasPrice
	if tx.GasTipCap == nil {
		tx.GasTipCap = big.NewInt(0)
	}
	if tx.Meta == nil {
		tx.Meta = &zktypes.Eip712Meta{GasPerPubdata: utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64())}
	} else if tx.Meta.GasPerPubdata == nil {
		tx.Meta.GasPerPubdata = utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64())
	}
	if tx.Gas == 0 {
		gas, err := client.EstimateGasL2(context.Background(), tx.ToCallMsg(address))
		if err != nil {
			return nil, fmt.Errorf("failed to EstimateGasL2: %w", err)
		}
		tx.Gas = gas
	}
	if tx.Data == nil {
		tx.Data = hexutil.Bytes{}
	}
	if tx.Meta.PaymasterParams != nil {
		paymasterParams, err := c.GetPaymasterParamsWithFee(big.NewInt(0).Mul(tx.GasFeeCap, big.NewInt(int64(tx.Gas))))
		if err != nil {
			panic(err)
		}
		tx.Meta.PaymasterParams = paymasterParams
	}
	return tx.ToTransaction712(address), nil
}

func (c *Client) SignTransaction(signer accounts.Signer, tx *zktypes.Transaction712) ([]byte, error) {
	var gas uint64 = 0
	if tx.Gas != nil {
		gas = tx.Gas.Uint64()
	}
	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		signer.Address(),
		accounts.Transaction{
			To:         tx.To,
			Data:       tx.Data,
			Value:      tx.Value,
			Nonce:      tx.Nonce,
			GasTipCap:  tx.GasTipCap,
			GasFeeCap:  tx.GasFeeCap,
			Gas:        gas,
			AccessList: tx.AccessList,
			ChainID:    tx.ChainID,
			Meta:       tx.Meta,
		},
	)
	if err != nil {
		return nil, err
	}
	signature, err := signer.SignTypedData(signer.Domain(), preparedTx)
	if err != nil {
		return nil, err
	}
	return preparedTx.RLPValues(signature)
}

func (c *Client) GetPaymasterParams() (*zktypes.PaymasterParams, error) {
	if c.PaymasterAddress == "" || c.PaymasterToken == "" {
		return nil, nil
	}
	fee := big.NewInt(1000000000000)
	if c.PaymasterFeeZero {
		fee = big.NewInt(0)
	}
	paymasterParams, err := utils.GetPaymasterParams(
		common.HexToAddress(c.PaymasterAddress),
		&zktypes.ApprovalBasedPaymasterInput{
			Token:            common.HexToAddress(c.PaymasterToken),
			MinimalAllowance: fee,
			InnerInput:       []byte{},
		})
	if err != nil {
		return nil, err
	}
	return paymasterParams, nil
}

func (c *Client) GetPaymasterParamsWithFee(fee *big.Int) (*zktypes.PaymasterParams, error) {
	if c.PaymasterAddress == "" || c.PaymasterToken == "" {
		return nil, nil
	}
	if c.PaymasterFeeZero {
		fee = big.NewInt(0)
	}
	paymasterParams, err := utils.GetPaymasterParams(
		common.HexToAddress(c.PaymasterAddress),
		&zktypes.ApprovalBasedPaymasterInput{
			Token:            common.HexToAddress(c.PaymasterToken),
			MinimalAllowance: fee,
			InnerInput:       []byte{},
		})
	if err != nil {
		return nil, err
	}
	return paymasterParams, nil
}

func (c *Client) SetPaymasterParams(tx *accounts.Transaction) error {
	paymasterParams, err := c.GetPaymasterParams()
	if err != nil {
		return err
	}
	tx.Meta.PaymasterParams = paymasterParams
	return nil
}

func (c *Client) PaymasterParams() *zktypes.PaymasterParams {
	paymasterParams, err := c.GetPaymasterParams()
	if err != nil {
		panic(err)
	}
	return paymasterParams
}

func (c *Client) Transact(prkHex string, from common.Address, to common.Address, value *big.Int, input []byte) (*zktypes.Receipt, error) {
	transact, err := c.createTransact(from, to, value, input)
	if err != nil {
		return nil, err
	}
	tx, err := c.signAndSendTx(prkHex, from, transact)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Client) createTransact(from common.Address, to common.Address, value *big.Int, input []byte) (*accounts.Transaction, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}

	gasNumber, err := client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From: from,
				To:   &to,
				Data: input,
			},
			Meta: &zktypes.Eip712Meta{
				GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
				PaymasterParams: c.PaymasterParams(),
			},
		},
	)

	if err != nil {
		return nil, err
	}

	gasPrice, err := c.GetGasPrice()
	if err != nil {
		return nil, err
	}

	return &accounts.Transaction{
		GasFeeCap: gasPrice,
		GasTipCap: gasPrice,
		Gas:       gasNumber,
		To:        &to,
		Value:     value,
		Data:      input,
		Meta: &zktypes.Eip712Meta{
			GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
			PaymasterParams: c.PaymasterParams(),
		},
	}, err
}

func (c *Client) signAndSendTx(prkHex string, pbkHex common.Address, transact *accounts.Transaction) (*zktypes.Receipt, error) {
	chainID, err := c.GetChainID()
	if err != nil {
		return nil, err
	}

	zkClient, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}

	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		pbkHex,
		*transact,
	)
	if err != nil {
		return nil, err
	}

	prkBytes, err := hex.DecodeString(prkHex)
	if err != nil {
		return nil, err
	}
	baseSigner, err := accounts.NewBaseSignerFromRawPrivateKey(prkBytes, int64(chainID))
	if err != nil {
		return nil, err
	}
	signer := accounts.Signer(baseSigner)
	rawTx, err := c.SignTransaction(signer, preparedTx)
	if err != nil {
		return nil, err
	}
	hash, err := zkClient.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return nil, err
	}
	tx, err := zkClient.WaitMined(context.Background(), hash)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *Client) WaitMined(ctx context.Context, txHash common.Hash) (*zktypes.Receipt, error) {
	zkClient, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()
	retry := 0
	for {
		retry++
		receipt, err := zkClient.TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil && receipt.BlockNumber != nil {
			return receipt, nil
		}
		if retry > 30 {
			return nil, fmt.Errorf("err:%v when find tx :%v at rpc :%v after %v (times) retries ", err, txHash.Hex(), c.BaseURL, retry)
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (c *Client) BlockByNumber(ctx context.Context, blockNumber uint64) (*zktypes.Block, error) {
	client, err := c.GetZkClient()
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, err
	}
	return block, nil
}
