// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ModelCollectionMetaData contains all meta data concerning the ModelCollection contract.
var ModelCollectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MintPriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"RoyaltyPortionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"TokenModelUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"TokenURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getHashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nextModelId\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"mintBySignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"modelAddressOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextModelId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"updateMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"}],\"name\":\"updateRoyaltyPortion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"}],\"name\":\"updateRoyaltyReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"updateTokenModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ModelCollectionABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelCollectionMetaData.ABI instead.
var ModelCollectionABI = ModelCollectionMetaData.ABI

// ModelCollection is an auto generated Go binding around an Ethereum contract.
type ModelCollection struct {
	ModelCollectionCaller     // Read-only binding to the contract
	ModelCollectionTransactor // Write-only binding to the contract
	ModelCollectionFilterer   // Log filterer for contract events
}

// ModelCollectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelCollectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelCollectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelCollectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelCollectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelCollectionSession struct {
	Contract     *ModelCollection  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelCollectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelCollectionCallerSession struct {
	Contract *ModelCollectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ModelCollectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelCollectionTransactorSession struct {
	Contract     *ModelCollectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ModelCollectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelCollectionRaw struct {
	Contract *ModelCollection // Generic contract binding to access the raw methods on
}

// ModelCollectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelCollectionCallerRaw struct {
	Contract *ModelCollectionCaller // Generic read-only contract binding to access the raw methods on
}

// ModelCollectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelCollectionTransactorRaw struct {
	Contract *ModelCollectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModelCollection creates a new instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollection(address common.Address, backend bind.ContractBackend) (*ModelCollection, error) {
	contract, err := bindModelCollection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModelCollection{ModelCollectionCaller: ModelCollectionCaller{contract: contract}, ModelCollectionTransactor: ModelCollectionTransactor{contract: contract}, ModelCollectionFilterer: ModelCollectionFilterer{contract: contract}}, nil
}

// NewModelCollectionCaller creates a new read-only instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollectionCaller(address common.Address, caller bind.ContractCaller) (*ModelCollectionCaller, error) {
	contract, err := bindModelCollection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionCaller{contract: contract}, nil
}

// NewModelCollectionTransactor creates a new write-only instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollectionTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelCollectionTransactor, error) {
	contract, err := bindModelCollection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionTransactor{contract: contract}, nil
}

// NewModelCollectionFilterer creates a new log filterer instance of ModelCollection, bound to a specific deployed contract.
func NewModelCollectionFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelCollectionFilterer, error) {
	contract, err := bindModelCollection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionFilterer{contract: contract}, nil
}

// bindModelCollection binds a generic wrapper to an already deployed contract.
func bindModelCollection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelCollectionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelCollection *ModelCollectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelCollection.Contract.ModelCollectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelCollection *ModelCollectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.Contract.ModelCollectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelCollection *ModelCollectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelCollection.Contract.ModelCollectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelCollection *ModelCollectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelCollection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelCollection *ModelCollectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelCollection *ModelCollectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelCollection.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollection *ModelCollectionSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ModelCollection.Contract.BalanceOf(&_ModelCollection.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ModelCollection.Contract.BalanceOf(&_ModelCollection.CallOpts, owner)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollection *ModelCollectionCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollection *ModelCollectionSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ModelCollection.Contract.Eip712Domain(&_ModelCollection.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ModelCollection *ModelCollectionCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ModelCollection.Contract.Eip712Domain(&_ModelCollection.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.GetApproved(&_ModelCollection.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.GetApproved(&_ModelCollection.CallOpts, tokenId)
}

// GetHashToSign is a free data retrieval call binding the contract method 0xab53af17.
//
// Solidity: function getHashToSign(address _to, string _uri, address _model, uint256 _tokenId) view returns(bytes32)
func (_ModelCollection *ModelCollectionCaller) GetHashToSign(opts *bind.CallOpts, _to common.Address, _uri string, _model common.Address, _tokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "getHashToSign", _to, _uri, _model, _tokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashToSign is a free data retrieval call binding the contract method 0xab53af17.
//
// Solidity: function getHashToSign(address _to, string _uri, address _model, uint256 _tokenId) view returns(bytes32)
func (_ModelCollection *ModelCollectionSession) GetHashToSign(_to common.Address, _uri string, _model common.Address, _tokenId *big.Int) ([32]byte, error) {
	return _ModelCollection.Contract.GetHashToSign(&_ModelCollection.CallOpts, _to, _uri, _model, _tokenId)
}

// GetHashToSign is a free data retrieval call binding the contract method 0xab53af17.
//
// Solidity: function getHashToSign(address _to, string _uri, address _model, uint256 _tokenId) view returns(bytes32)
func (_ModelCollection *ModelCollectionCallerSession) GetHashToSign(_to common.Address, _uri string, _model common.Address, _tokenId *big.Int) ([32]byte, error) {
	return _ModelCollection.Contract.GetHashToSign(&_ModelCollection.CallOpts, _to, _uri, _model, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollection *ModelCollectionSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ModelCollection.Contract.IsApprovedForAll(&_ModelCollection.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ModelCollection.Contract.IsApprovedForAll(&_ModelCollection.CallOpts, owner, operator)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) MintPrice() (*big.Int, error) {
	return _ModelCollection.Contract.MintPrice(&_ModelCollection.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) MintPrice() (*big.Int, error) {
	return _ModelCollection.Contract.MintPrice(&_ModelCollection.CallOpts)
}

// ModelAddressOf is a free data retrieval call binding the contract method 0x7f006226.
//
// Solidity: function modelAddressOf(uint256 _tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCaller) ModelAddressOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "modelAddressOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelAddressOf is a free data retrieval call binding the contract method 0x7f006226.
//
// Solidity: function modelAddressOf(uint256 _tokenId) view returns(address)
func (_ModelCollection *ModelCollectionSession) ModelAddressOf(_tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.ModelAddressOf(&_ModelCollection.CallOpts, _tokenId)
}

// ModelAddressOf is a free data retrieval call binding the contract method 0x7f006226.
//
// Solidity: function modelAddressOf(uint256 _tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) ModelAddressOf(_tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.ModelAddressOf(&_ModelCollection.CallOpts, _tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollection *ModelCollectionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollection *ModelCollectionSession) Name() (string, error) {
	return _ModelCollection.Contract.Name(&_ModelCollection.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ModelCollection *ModelCollectionCallerSession) Name() (string, error) {
	return _ModelCollection.Contract.Name(&_ModelCollection.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) NextModelId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "nextModelId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) NextModelId() (*big.Int, error) {
	return _ModelCollection.Contract.NextModelId(&_ModelCollection.CallOpts)
}

// NextModelId is a free data retrieval call binding the contract method 0xe472ae8b.
//
// Solidity: function nextModelId() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) NextModelId() (*big.Int, error) {
	return _ModelCollection.Contract.NextModelId(&_ModelCollection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollection *ModelCollectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollection *ModelCollectionSession) Owner() (common.Address, error) {
	return _ModelCollection.Contract.Owner(&_ModelCollection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) Owner() (common.Address, error) {
	return _ModelCollection.Contract.Owner(&_ModelCollection.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.OwnerOf(&_ModelCollection.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ModelCollection.Contract.OwnerOf(&_ModelCollection.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollection *ModelCollectionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollection *ModelCollectionSession) Paused() (bool, error) {
	return _ModelCollection.Contract.Paused(&_ModelCollection.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) Paused() (bool, error) {
	return _ModelCollection.Contract.Paused(&_ModelCollection.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ModelCollection *ModelCollectionCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ModelCollection *ModelCollectionSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ModelCollection.Contract.RoyaltyInfo(&_ModelCollection.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _ModelCollection.Contract.RoyaltyInfo(&_ModelCollection.CallOpts, _tokenId, _salePrice)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionCaller) RoyaltyPortion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionSession) RoyaltyPortion() (uint16, error) {
	return _ModelCollection.Contract.RoyaltyPortion(&_ModelCollection.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyPortion() (uint16, error) {
	return _ModelCollection.Contract.RoyaltyPortion(&_ModelCollection.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionCaller) RoyaltyReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionSession) RoyaltyReceiver() (common.Address, error) {
	return _ModelCollection.Contract.RoyaltyReceiver(&_ModelCollection.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_ModelCollection *ModelCollectionCallerSession) RoyaltyReceiver() (common.Address, error) {
	return _ModelCollection.Contract.RoyaltyReceiver(&_ModelCollection.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ModelCollection *ModelCollectionCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ModelCollection *ModelCollectionSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ModelCollection.Contract.SupportsInterface(&_ModelCollection.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_ModelCollection *ModelCollectionCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ModelCollection.Contract.SupportsInterface(&_ModelCollection.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollection *ModelCollectionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollection *ModelCollectionSession) Symbol() (string, error) {
	return _ModelCollection.Contract.Symbol(&_ModelCollection.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ModelCollection *ModelCollectionCallerSession) Symbol() (string, error) {
	return _ModelCollection.Contract.Symbol(&_ModelCollection.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenByIndex(&_ModelCollection.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenByIndex(&_ModelCollection.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenOfOwnerByIndex(&_ModelCollection.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ModelCollection.Contract.TokenOfOwnerByIndex(&_ModelCollection.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ModelCollection *ModelCollectionCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ModelCollection *ModelCollectionSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ModelCollection.Contract.TokenURI(&_ModelCollection.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_ModelCollection *ModelCollectionCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ModelCollection.Contract.TokenURI(&_ModelCollection.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollection *ModelCollectionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollection *ModelCollectionSession) TotalSupply() (*big.Int, error) {
	return _ModelCollection.Contract.TotalSupply(&_ModelCollection.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ModelCollection *ModelCollectionCallerSession) TotalSupply() (*big.Int, error) {
	return _ModelCollection.Contract.TotalSupply(&_ModelCollection.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollection *ModelCollectionCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ModelCollection.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollection *ModelCollectionSession) Version() (string, error) {
	return _ModelCollection.Contract.Version(&_ModelCollection.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ModelCollection *ModelCollectionCallerSession) Version() (string, error) {
	return _ModelCollection.Contract.Version(&_ModelCollection.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Approve(&_ModelCollection.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Approve(&_ModelCollection.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe973764.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextModelId) returns()
func (_ModelCollection *ModelCollectionTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextModelId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "initialize", _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextModelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe973764.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextModelId) returns()
func (_ModelCollection *ModelCollectionSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextModelId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Initialize(&_ModelCollection.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextModelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe973764.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextModelId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextModelId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Initialize(&_ModelCollection.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextModelId)
}

// Mint is a paid mutator transaction binding the contract method 0xfa8509c8.
//
// Solidity: function mint(address _to, string _uri, address _model) payable returns(uint256)
func (_ModelCollection *ModelCollectionTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _uri string, _model common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "mint", _to, _uri, _model)
}

// Mint is a paid mutator transaction binding the contract method 0xfa8509c8.
//
// Solidity: function mint(address _to, string _uri, address _model) payable returns(uint256)
func (_ModelCollection *ModelCollectionSession) Mint(_to common.Address, _uri string, _model common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.Mint(&_ModelCollection.TransactOpts, _to, _uri, _model)
}

// Mint is a paid mutator transaction binding the contract method 0xfa8509c8.
//
// Solidity: function mint(address _to, string _uri, address _model) payable returns(uint256)
func (_ModelCollection *ModelCollectionTransactorSession) Mint(_to common.Address, _uri string, _model common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.Mint(&_ModelCollection.TransactOpts, _to, _uri, _model)
}

// MintBySignature is a paid mutator transaction binding the contract method 0x5f730245.
//
// Solidity: function mintBySignature(address _to, string _uri, address _model, uint256 _tokenId, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_ModelCollection *ModelCollectionTransactor) MintBySignature(opts *bind.TransactOpts, _to common.Address, _uri string, _model common.Address, _tokenId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "mintBySignature", _to, _uri, _model, _tokenId, v, r, s)
}

// MintBySignature is a paid mutator transaction binding the contract method 0x5f730245.
//
// Solidity: function mintBySignature(address _to, string _uri, address _model, uint256 _tokenId, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_ModelCollection *ModelCollectionSession) MintBySignature(_to common.Address, _uri string, _model common.Address, _tokenId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ModelCollection.Contract.MintBySignature(&_ModelCollection.TransactOpts, _to, _uri, _model, _tokenId, v, r, s)
}

// MintBySignature is a paid mutator transaction binding the contract method 0x5f730245.
//
// Solidity: function mintBySignature(address _to, string _uri, address _model, uint256 _tokenId, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_ModelCollection *ModelCollectionTransactorSession) MintBySignature(_to common.Address, _uri string, _model common.Address, _tokenId *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ModelCollection.Contract.MintBySignature(&_ModelCollection.TransactOpts, _to, _uri, _model, _tokenId, v, r, s)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollection *ModelCollectionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollection *ModelCollectionSession) Pause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Pause(&_ModelCollection.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ModelCollection *ModelCollectionTransactorSession) Pause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Pause(&_ModelCollection.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollection *ModelCollectionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollection *ModelCollectionSession) RenounceOwnership() (*types.Transaction, error) {
	return _ModelCollection.Contract.RenounceOwnership(&_ModelCollection.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ModelCollection *ModelCollectionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ModelCollection.Contract.RenounceOwnership(&_ModelCollection.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollection *ModelCollectionTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollection *ModelCollectionSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom0(&_ModelCollection.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ModelCollection *ModelCollectionTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ModelCollection.Contract.SafeTransferFrom0(&_ModelCollection.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollection *ModelCollectionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollection *ModelCollectionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollection.Contract.SetApprovalForAll(&_ModelCollection.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ModelCollection *ModelCollectionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ModelCollection.Contract.SetApprovalForAll(&_ModelCollection.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ModelCollection *ModelCollectionTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferFrom(&_ModelCollection.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollection *ModelCollectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollection *ModelCollectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferOwnership(&_ModelCollection.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ModelCollection *ModelCollectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.TransferOwnership(&_ModelCollection.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollection *ModelCollectionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollection *ModelCollectionSession) Unpause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Unpause(&_ModelCollection.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ModelCollection *ModelCollectionTransactorSession) Unpause() (*types.Transaction, error) {
	return _ModelCollection.Contract.Unpause(&_ModelCollection.TransactOpts)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateMintPrice", _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_ModelCollection *ModelCollectionSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateMintPrice(&_ModelCollection.TransactOpts, _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateMintPrice(&_ModelCollection.TransactOpts, _mintPrice)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateRoyaltyPortion(opts *bind.TransactOpts, _royaltyPortion uint16) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateRoyaltyPortion", _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_ModelCollection *ModelCollectionSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyPortion(&_ModelCollection.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyPortion(&_ModelCollection.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateRoyaltyReceiver(opts *bind.TransactOpts, _royaltyReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateRoyaltyReceiver", _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_ModelCollection *ModelCollectionSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyReceiver(&_ModelCollection.TransactOpts, _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateRoyaltyReceiver(&_ModelCollection.TransactOpts, _royaltyReceiver)
}

// UpdateTokenModel is a paid mutator transaction binding the contract method 0x5e68842a.
//
// Solidity: function updateTokenModel(uint256 _tokenId, address _model) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateTokenModel(opts *bind.TransactOpts, _tokenId *big.Int, _model common.Address) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateTokenModel", _tokenId, _model)
}

// UpdateTokenModel is a paid mutator transaction binding the contract method 0x5e68842a.
//
// Solidity: function updateTokenModel(uint256 _tokenId, address _model) returns()
func (_ModelCollection *ModelCollectionSession) UpdateTokenModel(_tokenId *big.Int, _model common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateTokenModel(&_ModelCollection.TransactOpts, _tokenId, _model)
}

// UpdateTokenModel is a paid mutator transaction binding the contract method 0x5e68842a.
//
// Solidity: function updateTokenModel(uint256 _tokenId, address _model) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateTokenModel(_tokenId *big.Int, _model common.Address) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateTokenModel(&_ModelCollection.TransactOpts, _tokenId, _model)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 _tokenId, string _uri) returns()
func (_ModelCollection *ModelCollectionTransactor) UpdateTokenURI(opts *bind.TransactOpts, _tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "updateTokenURI", _tokenId, _uri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 _tokenId, string _uri) returns()
func (_ModelCollection *ModelCollectionSession) UpdateTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateTokenURI(&_ModelCollection.TransactOpts, _tokenId, _uri)
}

// UpdateTokenURI is a paid mutator transaction binding the contract method 0x18e97fd1.
//
// Solidity: function updateTokenURI(uint256 _tokenId, string _uri) returns()
func (_ModelCollection *ModelCollectionTransactorSession) UpdateTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _ModelCollection.Contract.UpdateTokenURI(&_ModelCollection.TransactOpts, _tokenId, _uri)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_ModelCollection *ModelCollectionTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ModelCollection.contract.Transact(opts, "withdraw", _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_ModelCollection *ModelCollectionSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Withdraw(&_ModelCollection.TransactOpts, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_ModelCollection *ModelCollectionTransactorSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ModelCollection.Contract.Withdraw(&_ModelCollection.TransactOpts, _to, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollection *ModelCollectionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelCollection.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollection *ModelCollectionSession) Receive() (*types.Transaction, error) {
	return _ModelCollection.Contract.Receive(&_ModelCollection.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ModelCollection *ModelCollectionTransactorSession) Receive() (*types.Transaction, error) {
	return _ModelCollection.Contract.Receive(&_ModelCollection.TransactOpts)
}

// ModelCollectionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ModelCollection contract.
type ModelCollectionApprovalIterator struct {
	Event *ModelCollectionApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionApproval represents a Approval event raised by the ModelCollection contract.
type ModelCollectionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ModelCollectionApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionApprovalIterator{contract: _ModelCollection.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ModelCollectionApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionApproval)
				if err := _ModelCollection.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseApproval(log types.Log) (*ModelCollectionApproval, error) {
	event := new(ModelCollectionApproval)
	if err := _ModelCollection.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ModelCollection contract.
type ModelCollectionApprovalForAllIterator struct {
	Event *ModelCollectionApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionApprovalForAll represents a ApprovalForAll event raised by the ModelCollection contract.
type ModelCollectionApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollection *ModelCollectionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ModelCollectionApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionApprovalForAllIterator{contract: _ModelCollection.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollection *ModelCollectionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ModelCollectionApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionApprovalForAll)
				if err := _ModelCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ModelCollection *ModelCollectionFilterer) ParseApprovalForAll(log types.Log) (*ModelCollectionApprovalForAll, error) {
	event := new(ModelCollectionApprovalForAll)
	if err := _ModelCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ModelCollection contract.
type ModelCollectionBatchMetadataUpdateIterator struct {
	Event *ModelCollectionBatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionBatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionBatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ModelCollection contract.
type ModelCollectionBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ModelCollectionBatchMetadataUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionBatchMetadataUpdateIterator{contract: _ModelCollection.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionBatchMetadataUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseBatchMetadataUpdate(log types.Log) (*ModelCollectionBatchMetadataUpdate, error) {
	event := new(ModelCollectionBatchMetadataUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ModelCollection contract.
type ModelCollectionEIP712DomainChangedIterator struct {
	Event *ModelCollectionEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionEIP712DomainChanged represents a EIP712DomainChanged event raised by the ModelCollection contract.
type ModelCollectionEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollection *ModelCollectionFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ModelCollectionEIP712DomainChangedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionEIP712DomainChangedIterator{contract: _ModelCollection.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollection *ModelCollectionFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ModelCollectionEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionEIP712DomainChanged)
				if err := _ModelCollection.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ModelCollection *ModelCollectionFilterer) ParseEIP712DomainChanged(log types.Log) (*ModelCollectionEIP712DomainChanged, error) {
	event := new(ModelCollectionEIP712DomainChanged)
	if err := _ModelCollection.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ModelCollection contract.
type ModelCollectionInitializedIterator struct {
	Event *ModelCollectionInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionInitialized represents a Initialized event raised by the ModelCollection contract.
type ModelCollectionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollection *ModelCollectionFilterer) FilterInitialized(opts *bind.FilterOpts) (*ModelCollectionInitializedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionInitializedIterator{contract: _ModelCollection.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollection *ModelCollectionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ModelCollectionInitialized) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionInitialized)
				if err := _ModelCollection.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ModelCollection *ModelCollectionFilterer) ParseInitialized(log types.Log) (*ModelCollectionInitialized, error) {
	event := new(ModelCollectionInitialized)
	if err := _ModelCollection.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ModelCollection contract.
type ModelCollectionMetadataUpdateIterator struct {
	Event *ModelCollectionMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionMetadataUpdate represents a MetadataUpdate event raised by the ModelCollection contract.
type ModelCollectionMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ModelCollectionMetadataUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionMetadataUpdateIterator{contract: _ModelCollection.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionMetadataUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseMetadataUpdate(log types.Log) (*ModelCollectionMetadataUpdate, error) {
	event := new(ModelCollectionMetadataUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionMintPriceUpdateIterator is returned from FilterMintPriceUpdate and is used to iterate over the raw logs and unpacked data for MintPriceUpdate events raised by the ModelCollection contract.
type ModelCollectionMintPriceUpdateIterator struct {
	Event *ModelCollectionMintPriceUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionMintPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionMintPriceUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionMintPriceUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionMintPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionMintPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionMintPriceUpdate represents a MintPriceUpdate event raised by the ModelCollection contract.
type ModelCollectionMintPriceUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPriceUpdate is a free log retrieval operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollection *ModelCollectionFilterer) FilterMintPriceUpdate(opts *bind.FilterOpts) (*ModelCollectionMintPriceUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionMintPriceUpdateIterator{contract: _ModelCollection.contract, event: "MintPriceUpdate", logs: logs, sub: sub}, nil
}

// WatchMintPriceUpdate is a free log subscription operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollection *ModelCollectionFilterer) WatchMintPriceUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionMintPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionMintPriceUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintPriceUpdate is a log parse operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_ModelCollection *ModelCollectionFilterer) ParseMintPriceUpdate(log types.Log) (*ModelCollectionMintPriceUpdate, error) {
	event := new(ModelCollectionMintPriceUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionNewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the ModelCollection contract.
type ModelCollectionNewTokenIterator struct {
	Event *ModelCollectionNewToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionNewToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionNewToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionNewToken represents a NewToken event raised by the ModelCollection contract.
type ModelCollectionNewToken struct {
	TokenId *big.Int
	Uri     string
	Model   common.Address
	Minter  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewToken is a free log retrieval operation binding the contract event 0x3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef74.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, address model, address indexed minter)
func (_ModelCollection *ModelCollectionFilterer) FilterNewToken(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address) (*ModelCollectionNewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionNewTokenIterator{contract: _ModelCollection.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef74.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, address model, address indexed minter)
func (_ModelCollection *ModelCollectionFilterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *ModelCollectionNewToken, tokenId []*big.Int, minter []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionNewToken)
				if err := _ModelCollection.contract.UnpackLog(event, "NewToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewToken is a log parse operation binding the contract event 0x3a434d4cd39d7a80e9d0fa54f10ba7b7e1aa16cbd063df3cc05523ac81adef74.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, address model, address indexed minter)
func (_ModelCollection *ModelCollectionFilterer) ParseNewToken(log types.Log) (*ModelCollectionNewToken, error) {
	event := new(ModelCollectionNewToken)
	if err := _ModelCollection.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ModelCollection contract.
type ModelCollectionOwnershipTransferredIterator struct {
	Event *ModelCollectionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionOwnershipTransferred represents a OwnershipTransferred event raised by the ModelCollection contract.
type ModelCollectionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollection *ModelCollectionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ModelCollectionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionOwnershipTransferredIterator{contract: _ModelCollection.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollection *ModelCollectionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ModelCollectionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionOwnershipTransferred)
				if err := _ModelCollection.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ModelCollection *ModelCollectionFilterer) ParseOwnershipTransferred(log types.Log) (*ModelCollectionOwnershipTransferred, error) {
	event := new(ModelCollectionOwnershipTransferred)
	if err := _ModelCollection.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ModelCollection contract.
type ModelCollectionPausedIterator struct {
	Event *ModelCollectionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionPaused represents a Paused event raised by the ModelCollection contract.
type ModelCollectionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollection *ModelCollectionFilterer) FilterPaused(opts *bind.FilterOpts) (*ModelCollectionPausedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionPausedIterator{contract: _ModelCollection.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollection *ModelCollectionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ModelCollectionPaused) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionPaused)
				if err := _ModelCollection.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ModelCollection *ModelCollectionFilterer) ParsePaused(log types.Log) (*ModelCollectionPaused, error) {
	event := new(ModelCollectionPaused)
	if err := _ModelCollection.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionRoyaltyPortionUpdateIterator is returned from FilterRoyaltyPortionUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyPortionUpdate events raised by the ModelCollection contract.
type ModelCollectionRoyaltyPortionUpdateIterator struct {
	Event *ModelCollectionRoyaltyPortionUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionRoyaltyPortionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionRoyaltyPortionUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionRoyaltyPortionUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionRoyaltyPortionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionRoyaltyPortionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionRoyaltyPortionUpdate represents a RoyaltyPortionUpdate event raised by the ModelCollection contract.
type ModelCollectionRoyaltyPortionUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPortionUpdate is a free log retrieval operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollection *ModelCollectionFilterer) FilterRoyaltyPortionUpdate(opts *bind.FilterOpts) (*ModelCollectionRoyaltyPortionUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionRoyaltyPortionUpdateIterator{contract: _ModelCollection.contract, event: "RoyaltyPortionUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPortionUpdate is a free log subscription operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollection *ModelCollectionFilterer) WatchRoyaltyPortionUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionRoyaltyPortionUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionRoyaltyPortionUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyPortionUpdate is a log parse operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_ModelCollection *ModelCollectionFilterer) ParseRoyaltyPortionUpdate(log types.Log) (*ModelCollectionRoyaltyPortionUpdate, error) {
	event := new(ModelCollectionRoyaltyPortionUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionRoyaltyReceiverUpdateIterator is returned from FilterRoyaltyReceiverUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyReceiverUpdate events raised by the ModelCollection contract.
type ModelCollectionRoyaltyReceiverUpdateIterator struct {
	Event *ModelCollectionRoyaltyReceiverUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionRoyaltyReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionRoyaltyReceiverUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionRoyaltyReceiverUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionRoyaltyReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionRoyaltyReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionRoyaltyReceiverUpdate represents a RoyaltyReceiverUpdate event raised by the ModelCollection contract.
type ModelCollectionRoyaltyReceiverUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyReceiverUpdate is a free log retrieval operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollection *ModelCollectionFilterer) FilterRoyaltyReceiverUpdate(opts *bind.FilterOpts) (*ModelCollectionRoyaltyReceiverUpdateIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionRoyaltyReceiverUpdateIterator{contract: _ModelCollection.contract, event: "RoyaltyReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyReceiverUpdate is a free log subscription operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollection *ModelCollectionFilterer) WatchRoyaltyReceiverUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionRoyaltyReceiverUpdate) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionRoyaltyReceiverUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyReceiverUpdate is a log parse operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_ModelCollection *ModelCollectionFilterer) ParseRoyaltyReceiverUpdate(log types.Log) (*ModelCollectionRoyaltyReceiverUpdate, error) {
	event := new(ModelCollectionRoyaltyReceiverUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionTokenModelUpdateIterator is returned from FilterTokenModelUpdate and is used to iterate over the raw logs and unpacked data for TokenModelUpdate events raised by the ModelCollection contract.
type ModelCollectionTokenModelUpdateIterator struct {
	Event *ModelCollectionTokenModelUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionTokenModelUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionTokenModelUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionTokenModelUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionTokenModelUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionTokenModelUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionTokenModelUpdate represents a TokenModelUpdate event raised by the ModelCollection contract.
type ModelCollectionTokenModelUpdate struct {
	TokenId *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenModelUpdate is a free log retrieval operation binding the contract event 0xa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d.
//
// Solidity: event TokenModelUpdate(uint256 indexed tokenId, address model)
func (_ModelCollection *ModelCollectionFilterer) FilterTokenModelUpdate(opts *bind.FilterOpts, tokenId []*big.Int) (*ModelCollectionTokenModelUpdateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "TokenModelUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionTokenModelUpdateIterator{contract: _ModelCollection.contract, event: "TokenModelUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenModelUpdate is a free log subscription operation binding the contract event 0xa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d.
//
// Solidity: event TokenModelUpdate(uint256 indexed tokenId, address model)
func (_ModelCollection *ModelCollectionFilterer) WatchTokenModelUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionTokenModelUpdate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "TokenModelUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionTokenModelUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "TokenModelUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenModelUpdate is a log parse operation binding the contract event 0xa0e7c03adff356c553e53dfec7043edb3e476fab3bdd27e5ef42955b92fb3e0d.
//
// Solidity: event TokenModelUpdate(uint256 indexed tokenId, address model)
func (_ModelCollection *ModelCollectionFilterer) ParseTokenModelUpdate(log types.Log) (*ModelCollectionTokenModelUpdate, error) {
	event := new(ModelCollectionTokenModelUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "TokenModelUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionTokenURIUpdateIterator is returned from FilterTokenURIUpdate and is used to iterate over the raw logs and unpacked data for TokenURIUpdate events raised by the ModelCollection contract.
type ModelCollectionTokenURIUpdateIterator struct {
	Event *ModelCollectionTokenURIUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionTokenURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionTokenURIUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionTokenURIUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionTokenURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionTokenURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionTokenURIUpdate represents a TokenURIUpdate event raised by the ModelCollection contract.
type ModelCollectionTokenURIUpdate struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenURIUpdate is a free log retrieval operation binding the contract event 0xc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe5.
//
// Solidity: event TokenURIUpdate(uint256 indexed tokenId, string uri)
func (_ModelCollection *ModelCollectionFilterer) FilterTokenURIUpdate(opts *bind.FilterOpts, tokenId []*big.Int) (*ModelCollectionTokenURIUpdateIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "TokenURIUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionTokenURIUpdateIterator{contract: _ModelCollection.contract, event: "TokenURIUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenURIUpdate is a free log subscription operation binding the contract event 0xc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe5.
//
// Solidity: event TokenURIUpdate(uint256 indexed tokenId, string uri)
func (_ModelCollection *ModelCollectionFilterer) WatchTokenURIUpdate(opts *bind.WatchOpts, sink chan<- *ModelCollectionTokenURIUpdate, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "TokenURIUpdate", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionTokenURIUpdate)
				if err := _ModelCollection.contract.UnpackLog(event, "TokenURIUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenURIUpdate is a log parse operation binding the contract event 0xc9e4a39d461f7a039fb05e3e4695cba6be812449c380b885df430abf38c19fe5.
//
// Solidity: event TokenURIUpdate(uint256 indexed tokenId, string uri)
func (_ModelCollection *ModelCollectionFilterer) ParseTokenURIUpdate(log types.Log) (*ModelCollectionTokenURIUpdate, error) {
	event := new(ModelCollectionTokenURIUpdate)
	if err := _ModelCollection.contract.UnpackLog(event, "TokenURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ModelCollection contract.
type ModelCollectionTransferIterator struct {
	Event *ModelCollectionTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionTransfer represents a Transfer event raised by the ModelCollection contract.
type ModelCollectionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ModelCollectionTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelCollectionTransferIterator{contract: _ModelCollection.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ModelCollectionTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionTransfer)
				if err := _ModelCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ModelCollection *ModelCollectionFilterer) ParseTransfer(log types.Log) (*ModelCollectionTransfer, error) {
	event := new(ModelCollectionTransfer)
	if err := _ModelCollection.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelCollectionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ModelCollection contract.
type ModelCollectionUnpausedIterator struct {
	Event *ModelCollectionUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ModelCollectionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelCollectionUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ModelCollectionUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ModelCollectionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelCollectionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelCollectionUnpaused represents a Unpaused event raised by the ModelCollection contract.
type ModelCollectionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollection *ModelCollectionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ModelCollectionUnpausedIterator, error) {

	logs, sub, err := _ModelCollection.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ModelCollectionUnpausedIterator{contract: _ModelCollection.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollection *ModelCollectionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ModelCollectionUnpaused) (event.Subscription, error) {

	logs, sub, err := _ModelCollection.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelCollectionUnpaused)
				if err := _ModelCollection.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ModelCollection *ModelCollectionFilterer) ParseUnpaused(log types.Log) (*ModelCollectionUnpaused, error) {
	event := new(ModelCollectionUnpaused)
	if err := _ModelCollection.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
