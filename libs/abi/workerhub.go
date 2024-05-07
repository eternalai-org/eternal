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

// WorkerHubMetaData contains all meta data concerning the WorkerHub contract.
var WorkerHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSetValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"DuplicatedAddressSetValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DuplicatedUint256SetValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintingSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256SetValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatingSessionNotEnded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinterExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinterRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValidatorExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUnregistration\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getInferenceInput\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getInferenceOutput\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinterAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinterAssignment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseMinterStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseValidatorStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inferences\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minterMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minterRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"_mintingTimeLimit\",\"type\":\"uint40\"},{\"internalType\":\"uint256\",\"name\":\"_validatorMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_validatorRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"_validatingTimeLimit\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"_maximumTier\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minterMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minterPivot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minterRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTaskId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingAssignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintingAssignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"accomplished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingTimeLimit\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"modelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"registerMinter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_output\",\"type\":\"bytes\"}],\"name\":\"submitOutput\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatingAssignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatingAssignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"accomplished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatingTimeLimit\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorPivot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTaskId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WorkerHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerHubMetaData.ABI instead.
var WorkerHubABI = WorkerHubMetaData.ABI

// WorkerHub is an auto generated Go binding around an Ethereum contract.
type WorkerHub struct {
	WorkerHubCaller     // Read-only binding to the contract
	WorkerHubTransactor // Write-only binding to the contract
	WorkerHubFilterer   // Log filterer for contract events
}

// WorkerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerHubSession struct {
	Contract     *WorkerHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerHubCallerSession struct {
	Contract *WorkerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WorkerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerHubTransactorSession struct {
	Contract     *WorkerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WorkerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerHubRaw struct {
	Contract *WorkerHub // Generic contract binding to access the raw methods on
}

// WorkerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerHubCallerRaw struct {
	Contract *WorkerHubCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerHubTransactorRaw struct {
	Contract *WorkerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerHub creates a new instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHub(address common.Address, backend bind.ContractBackend) (*WorkerHub, error) {
	contract, err := bindWorkerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerHub{WorkerHubCaller: WorkerHubCaller{contract: contract}, WorkerHubTransactor: WorkerHubTransactor{contract: contract}, WorkerHubFilterer: WorkerHubFilterer{contract: contract}}, nil
}

// NewWorkerHubCaller creates a new read-only instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubCaller(address common.Address, caller bind.ContractCaller) (*WorkerHubCaller, error) {
	contract, err := bindWorkerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubCaller{contract: contract}, nil
}

// NewWorkerHubTransactor creates a new write-only instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerHubTransactor, error) {
	contract, err := bindWorkerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTransactor{contract: contract}, nil
}

// NewWorkerHubFilterer creates a new log filterer instance of WorkerHub, bound to a specific deployed contract.
func NewWorkerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerHubFilterer, error) {
	contract, err := bindWorkerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerHubFilterer{contract: contract}, nil
}

// bindWorkerHub binds a generic wrapper to an already deployed contract.
func bindWorkerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkerHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHub *WorkerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHub.Contract.WorkerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHub *WorkerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.Contract.WorkerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHub *WorkerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHub.Contract.WorkerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHub *WorkerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHub *WorkerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHub *WorkerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHub.Contract.contract.Transact(opts, method, params...)
}

// GetInferenceInput is a free data retrieval call binding the contract method 0xab973d5d.
//
// Solidity: function getInferenceInput(uint256 _inferenceId) view returns(bytes)
func (_WorkerHub *WorkerHubCaller) GetInferenceInput(opts *bind.CallOpts, _inferenceId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getInferenceInput", _inferenceId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInferenceInput is a free data retrieval call binding the contract method 0xab973d5d.
//
// Solidity: function getInferenceInput(uint256 _inferenceId) view returns(bytes)
func (_WorkerHub *WorkerHubSession) GetInferenceInput(_inferenceId *big.Int) ([]byte, error) {
	return _WorkerHub.Contract.GetInferenceInput(&_WorkerHub.CallOpts, _inferenceId)
}

// GetInferenceInput is a free data retrieval call binding the contract method 0xab973d5d.
//
// Solidity: function getInferenceInput(uint256 _inferenceId) view returns(bytes)
func (_WorkerHub *WorkerHubCallerSession) GetInferenceInput(_inferenceId *big.Int) ([]byte, error) {
	return _WorkerHub.Contract.GetInferenceInput(&_WorkerHub.CallOpts, _inferenceId)
}

// GetInferenceOutput is a free data retrieval call binding the contract method 0xbc0e1f84.
//
// Solidity: function getInferenceOutput(uint256 _inferenceId) view returns(bytes)
func (_WorkerHub *WorkerHubCaller) GetInferenceOutput(opts *bind.CallOpts, _inferenceId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getInferenceOutput", _inferenceId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInferenceOutput is a free data retrieval call binding the contract method 0xbc0e1f84.
//
// Solidity: function getInferenceOutput(uint256 _inferenceId) view returns(bytes)
func (_WorkerHub *WorkerHubSession) GetInferenceOutput(_inferenceId *big.Int) ([]byte, error) {
	return _WorkerHub.Contract.GetInferenceOutput(&_WorkerHub.CallOpts, _inferenceId)
}

// GetInferenceOutput is a free data retrieval call binding the contract method 0xbc0e1f84.
//
// Solidity: function getInferenceOutput(uint256 _inferenceId) view returns(bytes)
func (_WorkerHub *WorkerHubCallerSession) GetInferenceOutput(_inferenceId *big.Int) ([]byte, error) {
	return _WorkerHub.Contract.GetInferenceOutput(&_WorkerHub.CallOpts, _inferenceId)
}

// GetMinterAddresses is a free data retrieval call binding the contract method 0x22543d63.
//
// Solidity: function getMinterAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetMinterAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinterAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinterAddresses is a free data retrieval call binding the contract method 0x22543d63.
//
// Solidity: function getMinterAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetMinterAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinterAddresses(&_WorkerHub.CallOpts)
}

// GetMinterAddresses is a free data retrieval call binding the contract method 0x22543d63.
//
// Solidity: function getMinterAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetMinterAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinterAddresses(&_WorkerHub.CallOpts)
}

// GetMinterAssignment is a free data retrieval call binding the contract method 0x766cf103.
//
// Solidity: function getMinterAssignment() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) GetMinterAssignment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinterAssignment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinterAssignment is a free data retrieval call binding the contract method 0x766cf103.
//
// Solidity: function getMinterAssignment() view returns(uint256)
func (_WorkerHub *WorkerHubSession) GetMinterAssignment() (*big.Int, error) {
	return _WorkerHub.Contract.GetMinterAssignment(&_WorkerHub.CallOpts)
}

// GetMinterAssignment is a free data retrieval call binding the contract method 0x766cf103.
//
// Solidity: function getMinterAssignment() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) GetMinterAssignment() (*big.Int, error) {
	return _WorkerHub.Contract.GetMinterAssignment(&_WorkerHub.CallOpts)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetModelAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getModelAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetModelAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetModelAddresses(&_WorkerHub.CallOpts)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetModelAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetModelAddresses(&_WorkerHub.CallOpts)
}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetValidatorAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getValidatorAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetValidatorAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetValidatorAddresses(&_WorkerHub.CallOpts)
}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetValidatorAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetValidatorAddresses(&_WorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) InferenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "inferenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHub.Contract.InferenceNumber(&_WorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHub.Contract.InferenceNumber(&_WorkerHub.CallOpts)
}

// Inferences is a free data retrieval call binding the contract method 0xfa8e8c69.
//
// Solidity: function inferences(uint256 ) view returns(address modelAddress, uint256 modelId, uint256 value, bytes input, address creator)
func (_WorkerHub *WorkerHubCaller) Inferences(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ModelAddress common.Address
	ModelId      *big.Int
	Value        *big.Int
	Input        []byte
	Creator      common.Address
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "inferences", arg0)

	outstruct := new(struct {
		ModelAddress common.Address
		ModelId      *big.Int
		Value        *big.Int
		Input        []byte
		Creator      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ModelId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Input = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.Creator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Inferences is a free data retrieval call binding the contract method 0xfa8e8c69.
//
// Solidity: function inferences(uint256 ) view returns(address modelAddress, uint256 modelId, uint256 value, bytes input, address creator)
func (_WorkerHub *WorkerHubSession) Inferences(arg0 *big.Int) (struct {
	ModelAddress common.Address
	ModelId      *big.Int
	Value        *big.Int
	Input        []byte
	Creator      common.Address
}, error) {
	return _WorkerHub.Contract.Inferences(&_WorkerHub.CallOpts, arg0)
}

// Inferences is a free data retrieval call binding the contract method 0xfa8e8c69.
//
// Solidity: function inferences(uint256 ) view returns(address modelAddress, uint256 modelId, uint256 value, bytes input, address creator)
func (_WorkerHub *WorkerHubCallerSession) Inferences(arg0 *big.Int) (struct {
	ModelAddress common.Address
	ModelId      *big.Int
	Value        *big.Int
	Input        []byte
	Creator      common.Address
}, error) {
	return _WorkerHub.Contract.Inferences(&_WorkerHub.CallOpts, arg0)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHub *WorkerHubSession) MaximumTier() (uint16, error) {
	return _WorkerHub.Contract.MaximumTier(&_WorkerHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) MaximumTier() (uint16, error) {
	return _WorkerHub.Contract.MaximumTier(&_WorkerHub.CallOpts)
}

// MinterMinimumStake is a free data retrieval call binding the contract method 0xba9d88c4.
//
// Solidity: function minterMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) MinterMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minterMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterMinimumStake is a free data retrieval call binding the contract method 0xba9d88c4.
//
// Solidity: function minterMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubSession) MinterMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.MinterMinimumStake(&_WorkerHub.CallOpts)
}

// MinterMinimumStake is a free data retrieval call binding the contract method 0xba9d88c4.
//
// Solidity: function minterMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) MinterMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.MinterMinimumStake(&_WorkerHub.CallOpts)
}

// MinterPivot is a free data retrieval call binding the contract method 0xed14e12f.
//
// Solidity: function minterPivot() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) MinterPivot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minterPivot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterPivot is a free data retrieval call binding the contract method 0xed14e12f.
//
// Solidity: function minterPivot() view returns(uint256)
func (_WorkerHub *WorkerHubSession) MinterPivot() (*big.Int, error) {
	return _WorkerHub.Contract.MinterPivot(&_WorkerHub.CallOpts)
}

// MinterPivot is a free data retrieval call binding the contract method 0xed14e12f.
//
// Solidity: function minterPivot() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) MinterPivot() (*big.Int, error) {
	return _WorkerHub.Contract.MinterPivot(&_WorkerHub.CallOpts)
}

// MinterRequirement is a free data retrieval call binding the contract method 0x09efefe7.
//
// Solidity: function minterRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCaller) MinterRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minterRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinterRequirement is a free data retrieval call binding the contract method 0x09efefe7.
//
// Solidity: function minterRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubSession) MinterRequirement() (uint8, error) {
	return _WorkerHub.Contract.MinterRequirement(&_WorkerHub.CallOpts)
}

// MinterRequirement is a free data retrieval call binding the contract method 0x09efefe7.
//
// Solidity: function minterRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCallerSession) MinterRequirement() (uint8, error) {
	return _WorkerHub.Contract.MinterRequirement(&_WorkerHub.CallOpts)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(uint256 stake, uint256 currentTaskId, uint256 commission, uint16 tier)
func (_WorkerHub *WorkerHubCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake         *big.Int
	CurrentTaskId *big.Int
	Commission    *big.Int
	Tier          uint16
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minters", arg0)

	outstruct := new(struct {
		Stake         *big.Int
		CurrentTaskId *big.Int
		Commission    *big.Int
		Tier          uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentTaskId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Commission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(uint256 stake, uint256 currentTaskId, uint256 commission, uint16 tier)
func (_WorkerHub *WorkerHubSession) Minters(arg0 common.Address) (struct {
	Stake         *big.Int
	CurrentTaskId *big.Int
	Commission    *big.Int
	Tier          uint16
}, error) {
	return _WorkerHub.Contract.Minters(&_WorkerHub.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(uint256 stake, uint256 currentTaskId, uint256 commission, uint16 tier)
func (_WorkerHub *WorkerHubCallerSession) Minters(arg0 common.Address) (struct {
	Stake         *big.Int
	CurrentTaskId *big.Int
	Commission    *big.Int
	Tier          uint16
}, error) {
	return _WorkerHub.Contract.Minters(&_WorkerHub.CallOpts, arg0)
}

// MintingAssignmentNumber is a free data retrieval call binding the contract method 0x460ae7f3.
//
// Solidity: function mintingAssignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) MintingAssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "mintingAssignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintingAssignmentNumber is a free data retrieval call binding the contract method 0x460ae7f3.
//
// Solidity: function mintingAssignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) MintingAssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.MintingAssignmentNumber(&_WorkerHub.CallOpts)
}

// MintingAssignmentNumber is a free data retrieval call binding the contract method 0x460ae7f3.
//
// Solidity: function mintingAssignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) MintingAssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.MintingAssignmentNumber(&_WorkerHub.CallOpts)
}

// MintingAssignments is a free data retrieval call binding the contract method 0x11d87854.
//
// Solidity: function mintingAssignments(uint256 ) view returns(uint256 inferenceId, address worker, uint40 expiredAt, uint8 index, bool accomplished)
func (_WorkerHub *WorkerHubCaller) MintingAssignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId  *big.Int
	Worker       common.Address
	ExpiredAt    *big.Int
	Index        uint8
	Accomplished bool
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "mintingAssignments", arg0)

	outstruct := new(struct {
		InferenceId  *big.Int
		Worker       common.Address
		ExpiredAt    *big.Int
		Index        uint8
		Accomplished bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Worker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ExpiredAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Accomplished = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// MintingAssignments is a free data retrieval call binding the contract method 0x11d87854.
//
// Solidity: function mintingAssignments(uint256 ) view returns(uint256 inferenceId, address worker, uint40 expiredAt, uint8 index, bool accomplished)
func (_WorkerHub *WorkerHubSession) MintingAssignments(arg0 *big.Int) (struct {
	InferenceId  *big.Int
	Worker       common.Address
	ExpiredAt    *big.Int
	Index        uint8
	Accomplished bool
}, error) {
	return _WorkerHub.Contract.MintingAssignments(&_WorkerHub.CallOpts, arg0)
}

// MintingAssignments is a free data retrieval call binding the contract method 0x11d87854.
//
// Solidity: function mintingAssignments(uint256 ) view returns(uint256 inferenceId, address worker, uint40 expiredAt, uint8 index, bool accomplished)
func (_WorkerHub *WorkerHubCallerSession) MintingAssignments(arg0 *big.Int) (struct {
	InferenceId  *big.Int
	Worker       common.Address
	ExpiredAt    *big.Int
	Index        uint8
	Accomplished bool
}, error) {
	return _WorkerHub.Contract.MintingAssignments(&_WorkerHub.CallOpts, arg0)
}

// MintingTimeLimit is a free data retrieval call binding the contract method 0x8da0445f.
//
// Solidity: function mintingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) MintingTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "mintingTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintingTimeLimit is a free data retrieval call binding the contract method 0x8da0445f.
//
// Solidity: function mintingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubSession) MintingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.MintingTimeLimit(&_WorkerHub.CallOpts)
}

// MintingTimeLimit is a free data retrieval call binding the contract method 0x8da0445f.
//
// Solidity: function mintingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) MintingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.MintingTimeLimit(&_WorkerHub.CallOpts)
}

// ModelNumber is a free data retrieval call binding the contract method 0x79065621.
//
// Solidity: function modelNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) ModelNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "modelNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelNumber is a free data retrieval call binding the contract method 0x79065621.
//
// Solidity: function modelNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) ModelNumber() (*big.Int, error) {
	return _WorkerHub.Contract.ModelNumber(&_WorkerHub.CallOpts)
}

// ModelNumber is a free data retrieval call binding the contract method 0x79065621.
//
// Solidity: function modelNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) ModelNumber() (*big.Int, error) {
	return _WorkerHub.Contract.ModelNumber(&_WorkerHub.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 modelId, uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubCaller) Models(opts *bind.CallOpts, arg0 common.Address) (struct {
	ModelId    *big.Int
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "models", arg0)

	outstruct := new(struct {
		ModelId    *big.Int
		MinimumFee *big.Int
		Tier       uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinimumFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 modelId, uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubSession) Models(arg0 common.Address) (struct {
	ModelId    *big.Int
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHub.Contract.Models(&_WorkerHub.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 modelId, uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubCallerSession) Models(arg0 common.Address) (struct {
	ModelId    *big.Int
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHub.Contract.Models(&_WorkerHub.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubSession) Owner() (common.Address, error) {
	return _WorkerHub.Contract.Owner(&_WorkerHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) Owner() (common.Address, error) {
	return _WorkerHub.Contract.Owner(&_WorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubSession) Paused() (bool, error) {
	return _WorkerHub.Contract.Paused(&_WorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) Paused() (bool, error) {
	return _WorkerHub.Contract.Paused(&_WorkerHub.CallOpts)
}

// ValidatingAssignmentNumber is a free data retrieval call binding the contract method 0x26d9bd4b.
//
// Solidity: function validatingAssignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) ValidatingAssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatingAssignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatingAssignmentNumber is a free data retrieval call binding the contract method 0x26d9bd4b.
//
// Solidity: function validatingAssignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) ValidatingAssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatingAssignmentNumber(&_WorkerHub.CallOpts)
}

// ValidatingAssignmentNumber is a free data retrieval call binding the contract method 0x26d9bd4b.
//
// Solidity: function validatingAssignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) ValidatingAssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatingAssignmentNumber(&_WorkerHub.CallOpts)
}

// ValidatingAssignments is a free data retrieval call binding the contract method 0x02596788.
//
// Solidity: function validatingAssignments(uint256 ) view returns(uint256 inferenceId, address worker, uint40 expiredAt, uint8 index, bool accomplished)
func (_WorkerHub *WorkerHubCaller) ValidatingAssignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId  *big.Int
	Worker       common.Address
	ExpiredAt    *big.Int
	Index        uint8
	Accomplished bool
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatingAssignments", arg0)

	outstruct := new(struct {
		InferenceId  *big.Int
		Worker       common.Address
		ExpiredAt    *big.Int
		Index        uint8
		Accomplished bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Worker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ExpiredAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Accomplished = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ValidatingAssignments is a free data retrieval call binding the contract method 0x02596788.
//
// Solidity: function validatingAssignments(uint256 ) view returns(uint256 inferenceId, address worker, uint40 expiredAt, uint8 index, bool accomplished)
func (_WorkerHub *WorkerHubSession) ValidatingAssignments(arg0 *big.Int) (struct {
	InferenceId  *big.Int
	Worker       common.Address
	ExpiredAt    *big.Int
	Index        uint8
	Accomplished bool
}, error) {
	return _WorkerHub.Contract.ValidatingAssignments(&_WorkerHub.CallOpts, arg0)
}

// ValidatingAssignments is a free data retrieval call binding the contract method 0x02596788.
//
// Solidity: function validatingAssignments(uint256 ) view returns(uint256 inferenceId, address worker, uint40 expiredAt, uint8 index, bool accomplished)
func (_WorkerHub *WorkerHubCallerSession) ValidatingAssignments(arg0 *big.Int) (struct {
	InferenceId  *big.Int
	Worker       common.Address
	ExpiredAt    *big.Int
	Index        uint8
	Accomplished bool
}, error) {
	return _WorkerHub.Contract.ValidatingAssignments(&_WorkerHub.CallOpts, arg0)
}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) ValidatingTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatingTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubSession) ValidatingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatingTimeLimit(&_WorkerHub.CallOpts)
}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) ValidatingTimeLimit() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatingTimeLimit(&_WorkerHub.CallOpts)
}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) ValidatorMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatorMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubSession) ValidatorMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatorMinimumStake(&_WorkerHub.CallOpts)
}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) ValidatorMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatorMinimumStake(&_WorkerHub.CallOpts)
}

// ValidatorPivot is a free data retrieval call binding the contract method 0x34a1eb27.
//
// Solidity: function validatorPivot() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) ValidatorPivot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatorPivot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorPivot is a free data retrieval call binding the contract method 0x34a1eb27.
//
// Solidity: function validatorPivot() view returns(uint256)
func (_WorkerHub *WorkerHubSession) ValidatorPivot() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatorPivot(&_WorkerHub.CallOpts)
}

// ValidatorPivot is a free data retrieval call binding the contract method 0x34a1eb27.
//
// Solidity: function validatorPivot() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) ValidatorPivot() (*big.Int, error) {
	return _WorkerHub.Contract.ValidatorPivot(&_WorkerHub.CallOpts)
}

// ValidatorRequirement is a free data retrieval call binding the contract method 0x4f57cc80.
//
// Solidity: function validatorRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCaller) ValidatorRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validatorRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorRequirement is a free data retrieval call binding the contract method 0x4f57cc80.
//
// Solidity: function validatorRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubSession) ValidatorRequirement() (uint8, error) {
	return _WorkerHub.Contract.ValidatorRequirement(&_WorkerHub.CallOpts)
}

// ValidatorRequirement is a free data retrieval call binding the contract method 0x4f57cc80.
//
// Solidity: function validatorRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCallerSession) ValidatorRequirement() (uint8, error) {
	return _WorkerHub.Contract.ValidatorRequirement(&_WorkerHub.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 currentTaskId, uint256 commission, uint16 tier)
func (_WorkerHub *WorkerHubCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake         *big.Int
	CurrentTaskId *big.Int
	Commission    *big.Int
	Tier          uint16
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Stake         *big.Int
		CurrentTaskId *big.Int
		Commission    *big.Int
		Tier          uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentTaskId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Commission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 currentTaskId, uint256 commission, uint16 tier)
func (_WorkerHub *WorkerHubSession) Validators(arg0 common.Address) (struct {
	Stake         *big.Int
	CurrentTaskId *big.Int
	Commission    *big.Int
	Tier          uint16
}, error) {
	return _WorkerHub.Contract.Validators(&_WorkerHub.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 currentTaskId, uint256 commission, uint16 tier)
func (_WorkerHub *WorkerHubCallerSession) Validators(arg0 common.Address) (struct {
	Stake         *big.Int
	CurrentTaskId *big.Int
	Commission    *big.Int
	Tier          uint16
}, error) {
	return _WorkerHub.Contract.Validators(&_WorkerHub.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubSession) Version() (string, error) {
	return _WorkerHub.Contract.Version(&_WorkerHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHub *WorkerHubCallerSession) Version() (string, error) {
	return _WorkerHub.Contract.Version(&_WorkerHub.CallOpts)
}

// IncreaseMinterStake is a paid mutator transaction binding the contract method 0xba22db75.
//
// Solidity: function increaseMinterStake() payable returns()
func (_WorkerHub *WorkerHubTransactor) IncreaseMinterStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "increaseMinterStake")
}

// IncreaseMinterStake is a paid mutator transaction binding the contract method 0xba22db75.
//
// Solidity: function increaseMinterStake() payable returns()
func (_WorkerHub *WorkerHubSession) IncreaseMinterStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseMinterStake(&_WorkerHub.TransactOpts)
}

// IncreaseMinterStake is a paid mutator transaction binding the contract method 0xba22db75.
//
// Solidity: function increaseMinterStake() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) IncreaseMinterStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseMinterStake(&_WorkerHub.TransactOpts)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0x62ed5c12.
//
// Solidity: function increaseValidatorStake() payable returns()
func (_WorkerHub *WorkerHubTransactor) IncreaseValidatorStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "increaseValidatorStake")
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0x62ed5c12.
//
// Solidity: function increaseValidatorStake() payable returns()
func (_WorkerHub *WorkerHubSession) IncreaseValidatorStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseValidatorStake(&_WorkerHub.TransactOpts)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0x62ed5c12.
//
// Solidity: function increaseValidatorStake() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) IncreaseValidatorStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseValidatorStake(&_WorkerHub.TransactOpts)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactor) Infer(opts *bind.TransactOpts, _input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "infer", _input, _creator)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubSession) Infer(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer(&_WorkerHub.TransactOpts, _input, _creator)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) Infer(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.Infer(&_WorkerHub.TransactOpts, _input, _creator)
}

// Initialize is a paid mutator transaction binding the contract method 0xd3cdaa3d.
//
// Solidity: function initialize(uint256 _minterMinimumStake, uint8 _minterRequirement, uint40 _mintingTimeLimit, uint256 _validatorMinimumStake, uint8 _validatorRequirement, uint40 _validatingTimeLimit, uint16 _maximumTier) returns()
func (_WorkerHub *WorkerHubTransactor) Initialize(opts *bind.TransactOpts, _minterMinimumStake *big.Int, _minterRequirement uint8, _mintingTimeLimit *big.Int, _validatorMinimumStake *big.Int, _validatorRequirement uint8, _validatingTimeLimit *big.Int, _maximumTier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "initialize", _minterMinimumStake, _minterRequirement, _mintingTimeLimit, _validatorMinimumStake, _validatorRequirement, _validatingTimeLimit, _maximumTier)
}

// Initialize is a paid mutator transaction binding the contract method 0xd3cdaa3d.
//
// Solidity: function initialize(uint256 _minterMinimumStake, uint8 _minterRequirement, uint40 _mintingTimeLimit, uint256 _validatorMinimumStake, uint8 _validatorRequirement, uint40 _validatingTimeLimit, uint16 _maximumTier) returns()
func (_WorkerHub *WorkerHubSession) Initialize(_minterMinimumStake *big.Int, _minterRequirement uint8, _mintingTimeLimit *big.Int, _validatorMinimumStake *big.Int, _validatorRequirement uint8, _validatingTimeLimit *big.Int, _maximumTier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _minterMinimumStake, _minterRequirement, _mintingTimeLimit, _validatorMinimumStake, _validatorRequirement, _validatingTimeLimit, _maximumTier)
}

// Initialize is a paid mutator transaction binding the contract method 0xd3cdaa3d.
//
// Solidity: function initialize(uint256 _minterMinimumStake, uint8 _minterRequirement, uint40 _mintingTimeLimit, uint256 _validatorMinimumStake, uint8 _validatorRequirement, uint40 _validatingTimeLimit, uint16 _maximumTier) returns()
func (_WorkerHub *WorkerHubTransactorSession) Initialize(_minterMinimumStake *big.Int, _minterRequirement uint8, _mintingTimeLimit *big.Int, _validatorMinimumStake *big.Int, _validatorRequirement uint8, _validatingTimeLimit *big.Int, _maximumTier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _minterMinimumStake, _minterRequirement, _mintingTimeLimit, _validatorMinimumStake, _validatorRequirement, _validatingTimeLimit, _maximumTier)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubSession) Pause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Pause(&_WorkerHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHub *WorkerHubTransactorSession) Pause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Pause(&_WorkerHub.TransactOpts)
}

// RegisterMinter is a paid mutator transaction binding the contract method 0x38ffc859.
//
// Solidity: function registerMinter(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactor) RegisterMinter(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerMinter", tier)
}

// RegisterMinter is a paid mutator transaction binding the contract method 0x38ffc859.
//
// Solidity: function registerMinter(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubSession) RegisterMinter(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterMinter(&_WorkerHub.TransactOpts, tier)
}

// RegisterMinter is a paid mutator transaction binding the contract method 0x38ffc859.
//
// Solidity: function registerMinter(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterMinter(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterMinter(&_WorkerHub.TransactOpts, tier)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns(uint256)
func (_WorkerHub *WorkerHubTransactor) RegisterModel(opts *bind.TransactOpts, _model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerModel", _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns(uint256)
func (_WorkerHub *WorkerHubSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterModel(&_WorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterModel(&_WorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x47f809e7.
//
// Solidity: function registerValidator(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactor) RegisterValidator(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerValidator", tier)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x47f809e7.
//
// Solidity: function registerValidator(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubSession) RegisterValidator(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterValidator(&_WorkerHub.TransactOpts, tier)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x47f809e7.
//
// Solidity: function registerValidator(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterValidator(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterValidator(&_WorkerHub.TransactOpts, tier)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHub.Contract.RenounceOwnership(&_WorkerHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHub *WorkerHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHub.Contract.RenounceOwnership(&_WorkerHub.TransactOpts)
}

// SubmitOutput is a paid mutator transaction binding the contract method 0x41d24baa.
//
// Solidity: function submitOutput(uint256 _assignmentId, bytes _output) returns()
func (_WorkerHub *WorkerHubTransactor) SubmitOutput(opts *bind.TransactOpts, _assignmentId *big.Int, _output []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "submitOutput", _assignmentId, _output)
}

// SubmitOutput is a paid mutator transaction binding the contract method 0x41d24baa.
//
// Solidity: function submitOutput(uint256 _assignmentId, bytes _output) returns()
func (_WorkerHub *WorkerHubSession) SubmitOutput(_assignmentId *big.Int, _output []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitOutput(&_WorkerHub.TransactOpts, _assignmentId, _output)
}

// SubmitOutput is a paid mutator transaction binding the contract method 0x41d24baa.
//
// Solidity: function submitOutput(uint256 _assignmentId, bytes _output) returns()
func (_WorkerHub *WorkerHubTransactorSession) SubmitOutput(_assignmentId *big.Int, _output []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitOutput(&_WorkerHub.TransactOpts, _assignmentId, _output)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.TransferOwnership(&_WorkerHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHub *WorkerHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.TransferOwnership(&_WorkerHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubSession) Unpause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Unpause(&_WorkerHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHub *WorkerHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _WorkerHub.Contract.Unpause(&_WorkerHub.TransactOpts)
}

// UnregisterMinter is a paid mutator transaction binding the contract method 0xfb8dac3b.
//
// Solidity: function unregisterMinter() returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterMinter")
}

// UnregisterMinter is a paid mutator transaction binding the contract method 0xfb8dac3b.
//
// Solidity: function unregisterMinter() returns()
func (_WorkerHub *WorkerHubSession) UnregisterMinter() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterMinter(&_WorkerHub.TransactOpts)
}

// UnregisterMinter is a paid mutator transaction binding the contract method 0xfb8dac3b.
//
// Solidity: function unregisterMinter() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterMinter() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterMinter(&_WorkerHub.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterModel(opts *bind.TransactOpts, _model common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterModel", _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHub *WorkerHubSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterModel(&_WorkerHub.TransactOpts, _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterModel(&_WorkerHub.TransactOpts, _model)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterValidator")
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHub *WorkerHubSession) UnregisterValidator() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterValidator(&_WorkerHub.TransactOpts)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterValidator() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterValidator(&_WorkerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubSession) Receive() (*types.Transaction, error) {
	return _WorkerHub.Contract.Receive(&_WorkerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) Receive() (*types.Transaction, error) {
	return _WorkerHub.Contract.Receive(&_WorkerHub.TransactOpts)
}

// WorkerHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkerHub contract.
type WorkerHubInitializedIterator struct {
	Event *WorkerHubInitialized // Event containing the contract specifics and raw log

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
func (it *WorkerHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInitialized)
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
		it.Event = new(WorkerHubInitialized)
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
func (it *WorkerHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInitialized represents a Initialized event raised by the WorkerHub contract.
type WorkerHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkerHubInitializedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkerHubInitializedIterator{contract: _WorkerHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHub *WorkerHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkerHubInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInitialized)
				if err := _WorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseInitialized(log types.Log) (*WorkerHubInitialized, error) {
	event := new(WorkerHubInitialized)
	if err := _WorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinterExtraStakeIterator is returned from FilterMinterExtraStake and is used to iterate over the raw logs and unpacked data for MinterExtraStake events raised by the WorkerHub contract.
type WorkerHubMinterExtraStakeIterator struct {
	Event *WorkerHubMinterExtraStake // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinterExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinterExtraStake)
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
		it.Event = new(WorkerHubMinterExtraStake)
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
func (it *WorkerHubMinterExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinterExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinterExtraStake represents a MinterExtraStake event raised by the WorkerHub contract.
type WorkerHubMinterExtraStake struct {
	Minter common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterExtraStake is a free log retrieval operation binding the contract event 0xba9e9d97d4bc7273959e32b5720b06dcc22e3cf5c05ab05ec9c5b8c0f633201b.
//
// Solidity: event MinterExtraStake(address indexed minter, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterMinterExtraStake(opts *bind.FilterOpts, minter []common.Address) (*WorkerHubMinterExtraStakeIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinterExtraStake", minterRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinterExtraStakeIterator{contract: _WorkerHub.contract, event: "MinterExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinterExtraStake is a free log subscription operation binding the contract event 0xba9e9d97d4bc7273959e32b5720b06dcc22e3cf5c05ab05ec9c5b8c0f633201b.
//
// Solidity: event MinterExtraStake(address indexed minter, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchMinterExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubMinterExtraStake, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinterExtraStake", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinterExtraStake)
				if err := _WorkerHub.contract.UnpackLog(event, "MinterExtraStake", log); err != nil {
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

// ParseMinterExtraStake is a log parse operation binding the contract event 0xba9e9d97d4bc7273959e32b5720b06dcc22e3cf5c05ab05ec9c5b8c0f633201b.
//
// Solidity: event MinterExtraStake(address indexed minter, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseMinterExtraStake(log types.Log) (*WorkerHubMinterExtraStake, error) {
	event := new(WorkerHubMinterExtraStake)
	if err := _WorkerHub.contract.UnpackLog(event, "MinterExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinterRegistrationIterator is returned from FilterMinterRegistration and is used to iterate over the raw logs and unpacked data for MinterRegistration events raised by the WorkerHub contract.
type WorkerHubMinterRegistrationIterator struct {
	Event *WorkerHubMinterRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinterRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinterRegistration)
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
		it.Event = new(WorkerHubMinterRegistration)
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
func (it *WorkerHubMinterRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinterRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinterRegistration represents a MinterRegistration event raised by the WorkerHub contract.
type WorkerHubMinterRegistration struct {
	Minter common.Address
	Tier   uint16
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRegistration is a free log retrieval operation binding the contract event 0x64233ee02fdcf4e2a2d20c26ee9c6eb0e10b3f229302b778551e568c5420b58a.
//
// Solidity: event MinterRegistration(address indexed minter, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterMinterRegistration(opts *bind.FilterOpts, minter []common.Address, tier []uint16) (*WorkerHubMinterRegistrationIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinterRegistration", minterRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinterRegistrationIterator{contract: _WorkerHub.contract, event: "MinterRegistration", logs: logs, sub: sub}, nil
}

// WatchMinterRegistration is a free log subscription operation binding the contract event 0x64233ee02fdcf4e2a2d20c26ee9c6eb0e10b3f229302b778551e568c5420b58a.
//
// Solidity: event MinterRegistration(address indexed minter, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchMinterRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubMinterRegistration, minter []common.Address, tier []uint16) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinterRegistration", minterRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinterRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "MinterRegistration", log); err != nil {
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

// ParseMinterRegistration is a log parse operation binding the contract event 0x64233ee02fdcf4e2a2d20c26ee9c6eb0e10b3f229302b778551e568c5420b58a.
//
// Solidity: event MinterRegistration(address indexed minter, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseMinterRegistration(log types.Log) (*WorkerHubMinterRegistration, error) {
	event := new(WorkerHubMinterRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "MinterRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinterUnregistrationIterator is returned from FilterMinterUnregistration and is used to iterate over the raw logs and unpacked data for MinterUnregistration events raised by the WorkerHub contract.
type WorkerHubMinterUnregistrationIterator struct {
	Event *WorkerHubMinterUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinterUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinterUnregistration)
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
		it.Event = new(WorkerHubMinterUnregistration)
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
func (it *WorkerHubMinterUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinterUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinterUnregistration represents a MinterUnregistration event raised by the WorkerHub contract.
type WorkerHubMinterUnregistration struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterUnregistration is a free log retrieval operation binding the contract event 0xf74a5115ea4662a0c4fbc19a38d7d9d1fa864223d3e52ef55ad2c44cd1b1e6ef.
//
// Solidity: event MinterUnregistration(address indexed minter)
func (_WorkerHub *WorkerHubFilterer) FilterMinterUnregistration(opts *bind.FilterOpts, minter []common.Address) (*WorkerHubMinterUnregistrationIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinterUnregistration", minterRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinterUnregistrationIterator{contract: _WorkerHub.contract, event: "MinterUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinterUnregistration is a free log subscription operation binding the contract event 0xf74a5115ea4662a0c4fbc19a38d7d9d1fa864223d3e52ef55ad2c44cd1b1e6ef.
//
// Solidity: event MinterUnregistration(address indexed minter)
func (_WorkerHub *WorkerHubFilterer) WatchMinterUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubMinterUnregistration, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinterUnregistration", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinterUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "MinterUnregistration", log); err != nil {
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

// ParseMinterUnregistration is a log parse operation binding the contract event 0xf74a5115ea4662a0c4fbc19a38d7d9d1fa864223d3e52ef55ad2c44cd1b1e6ef.
//
// Solidity: event MinterUnregistration(address indexed minter)
func (_WorkerHub *WorkerHubFilterer) ParseMinterUnregistration(log types.Log) (*WorkerHubMinterUnregistration, error) {
	event := new(WorkerHubMinterUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "MinterUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the WorkerHub contract.
type WorkerHubModelRegistrationIterator struct {
	Event *WorkerHubModelRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelRegistration)
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
		it.Event = new(WorkerHubModelRegistration)
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
func (it *WorkerHubModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelRegistration represents a ModelRegistration event raised by the WorkerHub contract.
type WorkerHubModelRegistration struct {
	Model      common.Address
	ModelId    *big.Int
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0xc866a92eb047ee02b8ba71f5ac25523fda8737ce206f248fa675d2e7c9645ffd.
//
// Solidity: event ModelRegistration(address indexed model, uint256 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) FilterModelRegistration(opts *bind.FilterOpts, model []common.Address, modelId []*big.Int, tier []uint16) (*WorkerHubModelRegistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelRegistration", modelRule, modelIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelRegistrationIterator{contract: _WorkerHub.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0xc866a92eb047ee02b8ba71f5ac25523fda8737ce206f248fa675d2e7c9645ffd.
//
// Solidity: event ModelRegistration(address indexed model, uint256 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubModelRegistration, model []common.Address, modelId []*big.Int, tier []uint16) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelRegistration", modelRule, modelIdRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
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

// ParseModelRegistration is a log parse operation binding the contract event 0xc866a92eb047ee02b8ba71f5ac25523fda8737ce206f248fa675d2e7c9645ffd.
//
// Solidity: event ModelRegistration(address indexed model, uint256 indexed modelId, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) ParseModelRegistration(log types.Log) (*WorkerHubModelRegistration, error) {
	event := new(WorkerHubModelRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the WorkerHub contract.
type WorkerHubModelUnregistrationIterator struct {
	Event *WorkerHubModelUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelUnregistration)
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
		it.Event = new(WorkerHubModelUnregistration)
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
func (it *WorkerHubModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelUnregistration represents a ModelUnregistration event raised by the WorkerHub contract.
type WorkerHubModelUnregistration struct {
	Model common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHub *WorkerHubFilterer) FilterModelUnregistration(opts *bind.FilterOpts, model []common.Address) (*WorkerHubModelUnregistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelUnregistrationIterator{contract: _WorkerHub.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHub *WorkerHubFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubModelUnregistration, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
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

// ParseModelUnregistration is a log parse operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHub *WorkerHubFilterer) ParseModelUnregistration(log types.Log) (*WorkerHubModelUnregistration, error) {
	event := new(WorkerHubModelUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the WorkerHub contract.
type WorkerHubNewInferenceIterator struct {
	Event *WorkerHubNewInference // Event containing the contract specifics and raw log

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
func (it *WorkerHubNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewInference)
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
		it.Event = new(WorkerHubNewInference)
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
func (it *WorkerHubNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewInference represents a NewInference event raised by the WorkerHub contract.
type WorkerHubNewInference struct {
	InferenceId *big.Int
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x38ddd254e73473217830b34f391798374ee63b7335fbc718901f8c47ac16932b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []*big.Int, creator []common.Address) (*WorkerHubNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewInference", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewInferenceIterator{contract: _WorkerHub.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x38ddd254e73473217830b34f391798374ee63b7335fbc718901f8c47ac16932b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *WorkerHubNewInference, inferenceId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewInference", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewInference)
				if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
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

// ParseNewInference is a log parse operation binding the contract event 0x38ddd254e73473217830b34f391798374ee63b7335fbc718901f8c47ac16932b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseNewInference(log types.Log) (*WorkerHubNewInference, error) {
	event := new(WorkerHubNewInference)
	if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkerHub contract.
type WorkerHubOwnershipTransferredIterator struct {
	Event *WorkerHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorkerHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubOwnershipTransferred)
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
		it.Event = new(WorkerHubOwnershipTransferred)
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
func (it *WorkerHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubOwnershipTransferred represents a OwnershipTransferred event raised by the WorkerHub contract.
type WorkerHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkerHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubOwnershipTransferredIterator{contract: _WorkerHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHub *WorkerHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkerHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubOwnershipTransferred)
				if err := _WorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseOwnershipTransferred(log types.Log) (*WorkerHubOwnershipTransferred, error) {
	event := new(WorkerHubOwnershipTransferred)
	if err := _WorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WorkerHub contract.
type WorkerHubPausedIterator struct {
	Event *WorkerHubPaused // Event containing the contract specifics and raw log

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
func (it *WorkerHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubPaused)
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
		it.Event = new(WorkerHubPaused)
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
func (it *WorkerHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubPaused represents a Paused event raised by the WorkerHub contract.
type WorkerHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) FilterPaused(opts *bind.FilterOpts) (*WorkerHubPausedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubPausedIterator{contract: _WorkerHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHub *WorkerHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WorkerHubPaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubPaused)
				if err := _WorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParsePaused(log types.Log) (*WorkerHubPaused, error) {
	event := new(WorkerHubPaused)
	if err := _WorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WorkerHub contract.
type WorkerHubUnpausedIterator struct {
	Event *WorkerHubUnpaused // Event containing the contract specifics and raw log

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
func (it *WorkerHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubUnpaused)
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
		it.Event = new(WorkerHubUnpaused)
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
func (it *WorkerHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubUnpaused represents a Unpaused event raised by the WorkerHub contract.
type WorkerHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WorkerHubUnpausedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubUnpausedIterator{contract: _WorkerHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHub *WorkerHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WorkerHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubUnpaused)
				if err := _WorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseUnpaused(log types.Log) (*WorkerHubUnpaused, error) {
	event := new(WorkerHubUnpaused)
	if err := _WorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorExtraStakeIterator is returned from FilterValidatorExtraStake and is used to iterate over the raw logs and unpacked data for ValidatorExtraStake events raised by the WorkerHub contract.
type WorkerHubValidatorExtraStakeIterator struct {
	Event *WorkerHubValidatorExtraStake // Event containing the contract specifics and raw log

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
func (it *WorkerHubValidatorExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorExtraStake)
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
		it.Event = new(WorkerHubValidatorExtraStake)
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
func (it *WorkerHubValidatorExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorExtraStake represents a ValidatorExtraStake event raised by the WorkerHub contract.
type WorkerHubValidatorExtraStake struct {
	Validator common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorExtraStake is a free log retrieval operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorExtraStake(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubValidatorExtraStakeIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorExtraStake", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorExtraStakeIterator{contract: _WorkerHub.contract, event: "ValidatorExtraStake", logs: logs, sub: sub}, nil
}

// WatchValidatorExtraStake is a free log subscription operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorExtraStake, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorExtraStake", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorExtraStake)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorExtraStake", log); err != nil {
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

// ParseValidatorExtraStake is a log parse operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorExtraStake(log types.Log) (*WorkerHubValidatorExtraStake, error) {
	event := new(WorkerHubValidatorExtraStake)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorRegistrationIterator is returned from FilterValidatorRegistration and is used to iterate over the raw logs and unpacked data for ValidatorRegistration events raised by the WorkerHub contract.
type WorkerHubValidatorRegistrationIterator struct {
	Event *WorkerHubValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorRegistration)
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
		it.Event = new(WorkerHubValidatorRegistration)
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
func (it *WorkerHubValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorRegistration represents a ValidatorRegistration event raised by the WorkerHub contract.
type WorkerHubValidatorRegistration struct {
	Validator common.Address
	Tier      uint16
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistration is a free log retrieval operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorRegistration(opts *bind.FilterOpts, validator []common.Address, tier []uint16) (*WorkerHubValidatorRegistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorRegistration", validatorRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorRegistrationIterator{contract: _WorkerHub.contract, event: "ValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistration is a free log subscription operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorRegistration, validator []common.Address, tier []uint16) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorRegistration", validatorRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
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

// ParseValidatorRegistration is a log parse operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorRegistration(log types.Log) (*WorkerHubValidatorRegistration, error) {
	event := new(WorkerHubValidatorRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubValidatorUnregistrationIterator is returned from FilterValidatorUnregistration and is used to iterate over the raw logs and unpacked data for ValidatorUnregistration events raised by the WorkerHub contract.
type WorkerHubValidatorUnregistrationIterator struct {
	Event *WorkerHubValidatorUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubValidatorUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubValidatorUnregistration)
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
		it.Event = new(WorkerHubValidatorUnregistration)
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
func (it *WorkerHubValidatorUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubValidatorUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubValidatorUnregistration represents a ValidatorUnregistration event raised by the WorkerHub contract.
type WorkerHubValidatorUnregistration struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnregistration is a free log retrieval operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) FilterValidatorUnregistration(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubValidatorUnregistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ValidatorUnregistration", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubValidatorUnregistrationIterator{contract: _WorkerHub.contract, event: "ValidatorUnregistration", logs: logs, sub: sub}, nil
}

// WatchValidatorUnregistration is a free log subscription operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) WatchValidatorUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubValidatorUnregistration, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ValidatorUnregistration", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubValidatorUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "ValidatorUnregistration", log); err != nil {
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

// ParseValidatorUnregistration is a log parse operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHub *WorkerHubFilterer) ParseValidatorUnregistration(log types.Log) (*WorkerHubValidatorUnregistration, error) {
	event := new(WorkerHubValidatorUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ValidatorUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
