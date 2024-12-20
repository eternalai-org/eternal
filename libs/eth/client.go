package eth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ethsecp "github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient(rpc string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewEthWsClient(ws string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(ws)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func WaitForTx(client *ethclient.Client, tx common.Hash) error {
	i := 0
	for {
		time.Sleep(2 * time.Second)
		if i > 20 {
			return errors.New("timeout")
		}
		i++
		_, isPending, err := client.TransactionByHash(context.Background(), tx)
		if err != nil {
			continue
		}
		if !isPending {
			break
		}
	}
	return nil
}

func WalletAddressFromCompressedPublicKey(publicKeyStr string) (string, error) {
	pubBytes, err := hex.DecodeString(publicKeyStr)
	if err != nil {
		return "", err
	}

	x, y := ethsecp.DecompressPubkey(pubBytes)

	pubkey := elliptic.Marshal(ethsecp.S256(), x, y)

	ecdsaPub, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		return "", err
	}
	ethAddress := crypto.PubkeyToAddress(*ecdsaPub).String()
	return ethAddress, nil
}

func GetAccountInfo(privKey string) (*ecdsa.PrivateKey, *common.Address, error) {
	masterWallet, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey := masterWallet.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	promptFeeAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return masterWallet, &promptFeeAddress, nil
}

func GenerateKeyFromSeedOld(seed string) (string, string, string, error) {
	seedHex := hex.EncodeToString([]byte(seed))
	master, ch := hd.ComputeMastersFromSeed([]byte(seedHex))
	path := "m/44'/1022'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", "", "", err
	}
	privateKey := secp256k1.GenPrivKeyFromSecret(priv)

	publicKey := privateKey.PubKey()

	privKey := hex.EncodeToString(priv)

	pubKey := hex.EncodeToString(publicKey.Bytes())

	address := "0x" + hex.EncodeToString(publicKey.Address().Bytes())

	return privKey, pubKey, address, nil
}

func GenerateKeyFromSeedNew(seed string) (string, string, string, error) {
	seedHex := hex.EncodeToString([]byte(seed))
	master, ch := hd.ComputeMastersFromSeed([]byte(seedHex))
	path := "m/44'/1022'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", "", "", err
	}

	privateKey, err := crypto.ToECDSA(priv)
	if err != nil {
		return "", "", "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	promptFeeAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	privKey := hex.EncodeToString(priv)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	pubKey := hex.EncodeToString(publicKeyBytes)

	return privKey, pubKey, strings.ToLower(promptFeeAddress.String()), nil
}

func GenerateAddress() (privKey, pubKey, address string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("failed to cast public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

func GenerateAddressFromPrivKey(privKey string) (pubKey, address string, err error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("failed to cast public key to ECDSA")
		return
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

type Client struct {
	eth *ethclient.Client
}

func NewClient(eth *ethclient.Client) *Client {
	return &Client{eth}
}

func (c *Client) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.eth.PendingNonceAt(ctx, address)
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.eth.SuggestGasPrice(ctx)
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.eth.NetworkID(ctx)
}

func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.eth.SendTransaction(ctx, tx)
}

// transfer:
func (c *Client) Transfer(senderPrivKey, receiverAddress string, amount, gasPrice *big.Int, gasLimit, nonce uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(senderPrivKey)
	if err != nil {
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if nonce <= 0 {
		nonce, err = c.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return "", err
		}
	}

	if gasLimit == 0 {
		gasLimit = uint64(21000)
	}

	if gasPrice == nil {
		gasPrice, err = c.SuggestGasPrice(context.Background())
		if err != nil {
			return "", err
		}
		// gasPrice = big.NewInt(0).Mul(gasPrice, big.NewInt(2))
		gasPrice = new(big.Int).Add(gasPrice, big.NewInt(2000000000))
	}

	fee := new(big.Int)
	fee.Mul(big.NewInt(int64(gasLimit)), gasPrice)

	toAddress := common.HexToAddress(receiverAddress)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func SignMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}
