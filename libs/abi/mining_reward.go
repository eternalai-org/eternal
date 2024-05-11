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

// MiningRewardMetaData contains all meta data concerning the MiningReward contract.
var MiningRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MiningRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use MiningRewardMetaData.ABI instead.
var MiningRewardABI = MiningRewardMetaData.ABI

// MiningReward is an auto generated Go binding around an Ethereum contract.
type MiningReward struct {
	MiningRewardCaller     // Read-only binding to the contract
	MiningRewardTransactor // Write-only binding to the contract
	MiningRewardFilterer   // Log filterer for contract events
}

// MiningRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiningRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiningRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiningRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiningRewardSession struct {
	Contract     *MiningReward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiningRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiningRewardCallerSession struct {
	Contract *MiningRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MiningRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiningRewardTransactorSession struct {
	Contract     *MiningRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MiningRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiningRewardRaw struct {
	Contract *MiningReward // Generic contract binding to access the raw methods on
}

// MiningRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiningRewardCallerRaw struct {
	Contract *MiningRewardCaller // Generic read-only contract binding to access the raw methods on
}

// MiningRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiningRewardTransactorRaw struct {
	Contract *MiningRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiningReward creates a new instance of MiningReward, bound to a specific deployed contract.
func NewMiningReward(address common.Address, backend bind.ContractBackend) (*MiningReward, error) {
	contract, err := bindMiningReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiningReward{MiningRewardCaller: MiningRewardCaller{contract: contract}, MiningRewardTransactor: MiningRewardTransactor{contract: contract}, MiningRewardFilterer: MiningRewardFilterer{contract: contract}}, nil
}

// NewMiningRewardCaller creates a new read-only instance of MiningReward, bound to a specific deployed contract.
func NewMiningRewardCaller(address common.Address, caller bind.ContractCaller) (*MiningRewardCaller, error) {
	contract, err := bindMiningReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiningRewardCaller{contract: contract}, nil
}

// NewMiningRewardTransactor creates a new write-only instance of MiningReward, bound to a specific deployed contract.
func NewMiningRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*MiningRewardTransactor, error) {
	contract, err := bindMiningReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiningRewardTransactor{contract: contract}, nil
}

// NewMiningRewardFilterer creates a new log filterer instance of MiningReward, bound to a specific deployed contract.
func NewMiningRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*MiningRewardFilterer, error) {
	contract, err := bindMiningReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiningRewardFilterer{contract: contract}, nil
}

// bindMiningReward binds a generic wrapper to an already deployed contract.
func bindMiningReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MiningRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningReward *MiningRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiningReward.Contract.MiningRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningReward *MiningRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningReward.Contract.MiningRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningReward *MiningRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningReward.Contract.MiningRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiningReward *MiningRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiningReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiningReward *MiningRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiningReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiningReward *MiningRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiningReward.Contract.contract.Transact(opts, method, params...)
}

// RewardToClaim is a free data retrieval call binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) view returns(uint256)
func (_MiningReward *MiningRewardCaller) RewardToClaim(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MiningReward.contract.Call(opts, &out, "rewardToClaim", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardToClaim is a free data retrieval call binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) view returns(uint256)
func (_MiningReward *MiningRewardSession) RewardToClaim(_miner common.Address) (*big.Int, error) {
	return _MiningReward.Contract.RewardToClaim(&_MiningReward.CallOpts, _miner)
}

// RewardToClaim is a free data retrieval call binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) view returns(uint256)
func (_MiningReward *MiningRewardCallerSession) RewardToClaim(_miner common.Address) (*big.Int, error) {
	return _MiningReward.Contract.RewardToClaim(&_MiningReward.CallOpts, _miner)
}
