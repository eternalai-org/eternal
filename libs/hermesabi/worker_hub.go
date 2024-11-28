// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hermesabi

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

// IWorkerHubAssignmentInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignmentInfo struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Value        *big.Int
	Input        []byte
	ModelAddress common.Address
	Creator      common.Address
	ExpiredAt    *big.Int
}

// IWorkerHubInferenceInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubInferenceInfo struct {
	InferenceId      *big.Int
	Input            []byte
	Output           []byte
	Value            *big.Int
	DisputingAddress common.Address
	ModelAddress     common.Address
	ExpiredAt        *big.Int
	Status           uint8
	Creator          common.Address
}

// IWorkerHubWorkerInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubWorkerInfo struct {
	WorkerAddress    common.Address
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}

// WorkerHubHermesMetaData contains all meta data concerning the WorkerHubHermes contract.
var WorkerHubHermesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InferMustBeSolvingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinerInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MiningSessionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MiningSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillBeingLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatingSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocks\",\"type\":\"uint256\"}],\"name\":\"BlocksPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldPercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPercent\",\"type\":\"uint16\"}],\"name\":\"FinePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"FraudulentMinerPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"}],\"name\":\"InferenceDisputation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"}],\"name\":\"MinerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"MinerUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newValue\",\"type\":\"uint40\"}],\"name\":\"MiningTimeLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"ModelTierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"oldDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newDuration\",\"type\":\"uint40\"}],\"name\":\"PenaltyDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restake\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TopUpInfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mingingFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"}],\"name\":\"TransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelayTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelayTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValidatorExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ValidatorUnstake\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"disapprovalCount\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blocksPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"disputeInfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputingTimeLimit\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disqualificationPercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"forceChangeModelForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_inferenceIds\",\"type\":\"uint256[]\"}],\"name\":\"getInferences\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"disputingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.InferenceInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"getMinerAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"workerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.WorkerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiningAssignments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getMintingAssignmentsOfInference\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNOMiner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"getValidatorAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"workerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.WorkerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"increaseMinerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"increaseValidatorStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_feePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_miningTimeLimit\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"_minerRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_blocksPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerEpochBasedOnPerf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_unstakeDelayTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_penaltyDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"isAssignmentPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForValidating\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minerUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningTimeLimit\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"resolveInference\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"restakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaskCompleted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMiner\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"}],\"name\":\"setFinePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardAmount\",\"type\":\"uint256\"}],\"name\":\"setNewRewardInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_penaltyDuration\",\"type\":\"uint40\"}],\"name\":\"setPenaltyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newUnstakeDelayTime\",\"type\":\"uint40\"}],\"name\":\"setUnstakDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isFined\",\"type\":\"bool\"}],\"name\":\"slashMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assigmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topUpAmount\",\"type\":\"uint256\"}],\"name\":\"topUpInfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeDelayTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_miningTimeLimit\",\"type\":\"uint40\"}],\"name\":\"updateMiningTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"updateModelMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_tier\",\"type\":\"uint32\"}],\"name\":\"updateModelTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatingTimeLimit\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorDisputed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WorkerHubHermesABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkerHubHermesMetaData.ABI instead.
var WorkerHubHermesABI = WorkerHubHermesMetaData.ABI

// WorkerHubHermes is an auto generated Go binding around an Ethereum contract.
type WorkerHubHermes struct {
	WorkerHubHermesCaller     // Read-only binding to the contract
	WorkerHubHermesTransactor // Write-only binding to the contract
	WorkerHubHermesFilterer   // Log filterer for contract events
}

// WorkerHubHermesCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkerHubHermesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubHermesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkerHubHermesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubHermesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkerHubHermesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkerHubHermesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkerHubHermesSession struct {
	Contract     *WorkerHubHermes  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkerHubHermesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkerHubHermesCallerSession struct {
	Contract *WorkerHubHermesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WorkerHubHermesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkerHubHermesTransactorSession struct {
	Contract     *WorkerHubHermesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WorkerHubHermesRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkerHubHermesRaw struct {
	Contract *WorkerHubHermes // Generic contract binding to access the raw methods on
}

// WorkerHubHermesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkerHubHermesCallerRaw struct {
	Contract *WorkerHubHermesCaller // Generic read-only contract binding to access the raw methods on
}

// WorkerHubHermesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkerHubHermesTransactorRaw struct {
	Contract *WorkerHubHermesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkerHubHermes creates a new instance of WorkerHubHermes, bound to a specific deployed contract.
func NewWorkerHubHermes(address common.Address, backend bind.ContractBackend) (*WorkerHubHermes, error) {
	contract, err := bindWorkerHubHermes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermes{WorkerHubHermesCaller: WorkerHubHermesCaller{contract: contract}, WorkerHubHermesTransactor: WorkerHubHermesTransactor{contract: contract}, WorkerHubHermesFilterer: WorkerHubHermesFilterer{contract: contract}}, nil
}

// NewWorkerHubHermesCaller creates a new read-only instance of WorkerHubHermes, bound to a specific deployed contract.
func NewWorkerHubHermesCaller(address common.Address, caller bind.ContractCaller) (*WorkerHubHermesCaller, error) {
	contract, err := bindWorkerHubHermes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesCaller{contract: contract}, nil
}

// NewWorkerHubHermesTransactor creates a new write-only instance of WorkerHubHermes, bound to a specific deployed contract.
func NewWorkerHubHermesTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkerHubHermesTransactor, error) {
	contract, err := bindWorkerHubHermes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesTransactor{contract: contract}, nil
}

// NewWorkerHubHermesFilterer creates a new log filterer instance of WorkerHubHermes, bound to a specific deployed contract.
func NewWorkerHubHermesFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkerHubHermesFilterer, error) {
	contract, err := bindWorkerHubHermes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesFilterer{contract: contract}, nil
}

// bindWorkerHubHermes binds a generic wrapper to an already deployed contract.
func bindWorkerHubHermes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkerHubHermesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHubHermes *WorkerHubHermesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHubHermes.Contract.WorkerHubHermesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHubHermes *WorkerHubHermesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.WorkerHubHermesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHubHermes *WorkerHubHermesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.WorkerHubHermesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkerHubHermes *WorkerHubHermesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkerHubHermes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkerHubHermes *WorkerHubHermesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkerHubHermes *WorkerHubHermesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.contract.Transact(opts, method, params...)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) AssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "assignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHubHermes.Contract.AssignmentNumber(&_WorkerHubHermes.CallOpts)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHubHermes.Contract.AssignmentNumber(&_WorkerHubHermes.CallOpts)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes output, address worker, uint8 disapprovalCount)
func (_WorkerHubHermes *WorkerHubHermesCaller) Assignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId      *big.Int
	Output           []byte
	Worker           common.Address
	DisapprovalCount uint8
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "assignments", arg0)

	outstruct := new(struct {
		InferenceId      *big.Int
		Output           []byte
		Worker           common.Address
		DisapprovalCount uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Output = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Worker = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.DisapprovalCount = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes output, address worker, uint8 disapprovalCount)
func (_WorkerHubHermes *WorkerHubHermesSession) Assignments(arg0 *big.Int) (struct {
	InferenceId      *big.Int
	Output           []byte
	Worker           common.Address
	DisapprovalCount uint8
}, error) {
	return _WorkerHubHermes.Contract.Assignments(&_WorkerHubHermes.CallOpts, arg0)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes output, address worker, uint8 disapprovalCount)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Assignments(arg0 *big.Int) (struct {
	InferenceId      *big.Int
	Output           []byte
	Worker           common.Address
	DisapprovalCount uint8
}, error) {
	return _WorkerHubHermes.Contract.Assignments(&_WorkerHubHermes.CallOpts, arg0)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) BlocksPerEpoch() (*big.Int, error) {
	return _WorkerHubHermes.Contract.BlocksPerEpoch(&_WorkerHubHermes.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _WorkerHubHermes.Contract.BlocksPerEpoch(&_WorkerHubHermes.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesSession) CurrentEpoch() (*big.Int, error) {
	return _WorkerHubHermes.Contract.CurrentEpoch(&_WorkerHubHermes.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) CurrentEpoch() (*big.Int, error) {
	return _WorkerHubHermes.Contract.CurrentEpoch(&_WorkerHubHermes.CallOpts)
}

// DisputingTimeLimit is a free data retrieval call binding the contract method 0xd64f1f5f.
//
// Solidity: function disputingTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCaller) DisputingTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "disputingTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputingTimeLimit is a free data retrieval call binding the contract method 0xd64f1f5f.
//
// Solidity: function disputingTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesSession) DisputingTimeLimit() (*big.Int, error) {
	return _WorkerHubHermes.Contract.DisputingTimeLimit(&_WorkerHubHermes.CallOpts)
}

// DisputingTimeLimit is a free data retrieval call binding the contract method 0xd64f1f5f.
//
// Solidity: function disputingTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) DisputingTimeLimit() (*big.Int, error) {
	return _WorkerHubHermes.Contract.DisputingTimeLimit(&_WorkerHubHermes.CallOpts)
}

// DisqualificationPercentage is a free data retrieval call binding the contract method 0x4670ca47.
//
// Solidity: function disqualificationPercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCaller) DisqualificationPercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "disqualificationPercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DisqualificationPercentage is a free data retrieval call binding the contract method 0x4670ca47.
//
// Solidity: function disqualificationPercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesSession) DisqualificationPercentage() (uint16, error) {
	return _WorkerHubHermes.Contract.DisqualificationPercentage(&_WorkerHubHermes.CallOpts)
}

// DisqualificationPercentage is a free data retrieval call binding the contract method 0x4670ca47.
//
// Solidity: function disqualificationPercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) DisqualificationPercentage() (uint16, error) {
	return _WorkerHubHermes.Contract.DisqualificationPercentage(&_WorkerHubHermes.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCaller) FeePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesSession) FeePercentage() (uint16, error) {
	return _WorkerHubHermes.Contract.FeePercentage(&_WorkerHubHermes.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) FeePercentage() (uint16, error) {
	return _WorkerHubHermes.Contract.FeePercentage(&_WorkerHubHermes.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesSession) FinePercentage() (uint16, error) {
	return _WorkerHubHermes.Contract.FinePercentage(&_WorkerHubHermes.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) FinePercentage() (uint16, error) {
	return _WorkerHubHermes.Contract.FinePercentage(&_WorkerHubHermes.CallOpts)
}

// GetInferences is a free data retrieval call binding the contract method 0xa01253d8.
//
// Solidity: function getInferences(uint256[] _inferenceIds) view returns((uint256,bytes,bytes,uint256,address,address,uint40,uint8,address)[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetInferences(opts *bind.CallOpts, _inferenceIds []*big.Int) ([]IWorkerHubInferenceInfo, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getInferences", _inferenceIds)

	if err != nil {
		return *new([]IWorkerHubInferenceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubInferenceInfo)).(*[]IWorkerHubInferenceInfo)

	return out0, err

}

// GetInferences is a free data retrieval call binding the contract method 0xa01253d8.
//
// Solidity: function getInferences(uint256[] _inferenceIds) view returns((uint256,bytes,bytes,uint256,address,address,uint40,uint8,address)[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetInferences(_inferenceIds []*big.Int) ([]IWorkerHubInferenceInfo, error) {
	return _WorkerHubHermes.Contract.GetInferences(&_WorkerHubHermes.CallOpts, _inferenceIds)
}

// GetInferences is a free data retrieval call binding the contract method 0xa01253d8.
//
// Solidity: function getInferences(uint256[] _inferenceIds) view returns((uint256,bytes,bytes,uint256,address,address,uint40,uint8,address)[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetInferences(_inferenceIds []*big.Int) ([]IWorkerHubInferenceInfo, error) {
	return _WorkerHubHermes.Contract.GetInferences(&_WorkerHubHermes.CallOpts, _inferenceIds)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetMinerAddresses() ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetMinerAddresses(&_WorkerHubHermes.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetMinerAddresses(&_WorkerHubHermes.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getMinerAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetMinerAddressesOfModel(&_WorkerHubHermes.CallOpts, _model)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetMinerAddressesOfModel(&_WorkerHubHermes.CallOpts, _model)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetMiners(opts *bind.CallOpts) ([]IWorkerHubWorkerInfo, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getMiners")

	if err != nil {
		return *new([]IWorkerHubWorkerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorkerInfo)).(*[]IWorkerHubWorkerInfo)

	return out0, err

}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHubHermes.Contract.GetMiners(&_WorkerHubHermes.CallOpts)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHubHermes.Contract.GetMiners(&_WorkerHubHermes.CallOpts)
}

// GetMiningAssignments is a free data retrieval call binding the contract method 0x37504f57.
//
// Solidity: function getMiningAssignments() view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetMiningAssignments(opts *bind.CallOpts) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getMiningAssignments")

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetMiningAssignments is a free data retrieval call binding the contract method 0x37504f57.
//
// Solidity: function getMiningAssignments() view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetMiningAssignments() ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHubHermes.Contract.GetMiningAssignments(&_WorkerHubHermes.CallOpts)
}

// GetMiningAssignments is a free data retrieval call binding the contract method 0x37504f57.
//
// Solidity: function getMiningAssignments() view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetMiningAssignments() ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHubHermes.Contract.GetMiningAssignments(&_WorkerHubHermes.CallOpts)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetMintingAssignmentsOfInference(opts *bind.CallOpts, _inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getMintingAssignmentsOfInference", _inferenceId)

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHubHermes.Contract.GetMintingAssignmentsOfInference(&_WorkerHubHermes.CallOpts, _inferenceId)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40)[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHubHermes.Contract.GetMintingAssignmentsOfInference(&_WorkerHubHermes.CallOpts, _inferenceId)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetModelAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getModelAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetModelAddresses() ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetModelAddresses(&_WorkerHubHermes.CallOpts)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetModelAddresses() ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetModelAddresses(&_WorkerHubHermes.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) GetNOMiner() (*big.Int, error) {
	return _WorkerHubHermes.Contract.GetNOMiner(&_WorkerHubHermes.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetNOMiner() (*big.Int, error) {
	return _WorkerHubHermes.Contract.GetNOMiner(&_WorkerHubHermes.CallOpts)
}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetValidatorAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getValidatorAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetValidatorAddresses() ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetValidatorAddresses(&_WorkerHubHermes.CallOpts)
}

// GetValidatorAddresses is a free data retrieval call binding the contract method 0xf74e921f.
//
// Solidity: function getValidatorAddresses() view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetValidatorAddresses() ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetValidatorAddresses(&_WorkerHubHermes.CallOpts)
}

// GetValidatorAddressesOfModel is a free data retrieval call binding the contract method 0xcbaaf438.
//
// Solidity: function getValidatorAddressesOfModel(address _model) view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetValidatorAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getValidatorAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorAddressesOfModel is a free data retrieval call binding the contract method 0xcbaaf438.
//
// Solidity: function getValidatorAddressesOfModel(address _model) view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetValidatorAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetValidatorAddressesOfModel(&_WorkerHubHermes.CallOpts, _model)
}

// GetValidatorAddressesOfModel is a free data retrieval call binding the contract method 0xcbaaf438.
//
// Solidity: function getValidatorAddressesOfModel(address _model) view returns(address[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetValidatorAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHubHermes.Contract.GetValidatorAddressesOfModel(&_WorkerHubHermes.CallOpts, _model)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHubHermes *WorkerHubHermesCaller) GetValidators(opts *bind.CallOpts) ([]IWorkerHubWorkerInfo, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]IWorkerHubWorkerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorkerInfo)).(*[]IWorkerHubWorkerInfo)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHubHermes *WorkerHubHermesSession) GetValidators() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHubHermes.Contract.GetValidators(&_WorkerHubHermes.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHubHermes *WorkerHubHermesCallerSession) GetValidators() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHubHermes.Contract.GetValidators(&_WorkerHubHermes.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) InferenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "inferenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHubHermes.Contract.InferenceNumber(&_WorkerHubHermes.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) InferenceNumber() (*big.Int, error) {
	return _WorkerHubHermes.Contract.InferenceNumber(&_WorkerHubHermes.CallOpts)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesCaller) IsAssignmentPending(opts *bind.CallOpts, _assignmentId *big.Int) (bool, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "isAssignmentPending", _assignmentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _WorkerHubHermes.Contract.IsAssignmentPending(&_WorkerHubHermes.CallOpts, _assignmentId)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _WorkerHubHermes.Contract.IsAssignmentPending(&_WorkerHubHermes.CallOpts, _assignmentId)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) LastBlock() (*big.Int, error) {
	return _WorkerHubHermes.Contract.LastBlock(&_WorkerHubHermes.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) LastBlock() (*big.Int, error) {
	return _WorkerHubHermes.Contract.LastBlock(&_WorkerHubHermes.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesSession) MaximumTier() (uint16, error) {
	return _WorkerHubHermes.Contract.MaximumTier(&_WorkerHubHermes.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) MaximumTier() (uint16, error) {
	return _WorkerHubHermes.Contract.MaximumTier(&_WorkerHubHermes.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) MinerMinimumStake() (*big.Int, error) {
	return _WorkerHubHermes.Contract.MinerMinimumStake(&_WorkerHubHermes.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _WorkerHubHermes.Contract.MinerMinimumStake(&_WorkerHubHermes.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHubHermes *WorkerHubHermesCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHubHermes *WorkerHubHermesSession) MinerRequirement() (uint8, error) {
	return _WorkerHubHermes.Contract.MinerRequirement(&_WorkerHubHermes.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) MinerRequirement() (uint8, error) {
	return _WorkerHubHermes.Contract.MinerRequirement(&_WorkerHubHermes.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHubHermes *WorkerHubHermesCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "minerUnstakeRequests", arg0)

	outstruct := new(struct {
		Stake    *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHubHermes *WorkerHubHermesSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHubHermes.Contract.MinerUnstakeRequests(&_WorkerHubHermes.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHubHermes.Contract.MinerUnstakeRequests(&_WorkerHubHermes.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHubHermes *WorkerHubHermesCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "miners", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		Commitment       *big.Int
		ModelAddress     common.Address
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModelAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHubHermes *WorkerHubHermesSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHubHermes.Contract.Miners(&_WorkerHubHermes.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHubHermes.Contract.Miners(&_WorkerHubHermes.CallOpts, arg0)
}

// MiningTimeLimit is a free data retrieval call binding the contract method 0x691ff7ef.
//
// Solidity: function miningTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCaller) MiningTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "miningTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningTimeLimit is a free data retrieval call binding the contract method 0x691ff7ef.
//
// Solidity: function miningTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesSession) MiningTimeLimit() (*big.Int, error) {
	return _WorkerHubHermes.Contract.MiningTimeLimit(&_WorkerHubHermes.CallOpts)
}

// MiningTimeLimit is a free data retrieval call binding the contract method 0x691ff7ef.
//
// Solidity: function miningTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) MiningTimeLimit() (*big.Int, error) {
	return _WorkerHubHermes.Contract.MiningTimeLimit(&_WorkerHubHermes.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHubHermes *WorkerHubHermesCaller) Models(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "models", arg0)

	outstruct := new(struct {
		MinimumFee *big.Int
		Tier       uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHubHermes *WorkerHubHermesSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHubHermes.Contract.Models(&_WorkerHubHermes.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHubHermes.Contract.Models(&_WorkerHubHermes.CallOpts, arg0)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) Multiplier(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "multiplier", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _WorkerHubHermes.Contract.Multiplier(&_WorkerHubHermes.CallOpts, _miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _WorkerHubHermes.Contract.Multiplier(&_WorkerHubHermes.CallOpts, _miner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesSession) Owner() (common.Address, error) {
	return _WorkerHubHermes.Contract.Owner(&_WorkerHubHermes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Owner() (common.Address, error) {
	return _WorkerHubHermes.Contract.Owner(&_WorkerHubHermes.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesSession) Paused() (bool, error) {
	return _WorkerHubHermes.Contract.Paused(&_WorkerHubHermes.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Paused() (bool, error) {
	return _WorkerHubHermes.Contract.Paused(&_WorkerHubHermes.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesSession) PenaltyDuration() (*big.Int, error) {
	return _WorkerHubHermes.Contract.PenaltyDuration(&_WorkerHubHermes.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) PenaltyDuration() (*big.Int, error) {
	return _WorkerHubHermes.Contract.PenaltyDuration(&_WorkerHubHermes.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHubHermes *WorkerHubHermesCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "rewardInEpoch", arg0)

	outstruct := new(struct {
		PerfReward         *big.Int
		EpochReward        *big.Int
		TotalTaskCompleted *big.Int
		TotalMiner         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerfReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalTaskCompleted = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalMiner = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHubHermes *WorkerHubHermesSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _WorkerHubHermes.Contract.RewardInEpoch(&_WorkerHubHermes.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _WorkerHubHermes.Contract.RewardInEpoch(&_WorkerHubHermes.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) RewardPerEpoch() (*big.Int, error) {
	return _WorkerHubHermes.Contract.RewardPerEpoch(&_WorkerHubHermes.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _WorkerHubHermes.Contract.RewardPerEpoch(&_WorkerHubHermes.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesSession) StakeToken() (common.Address, error) {
	return _WorkerHubHermes.Contract.StakeToken(&_WorkerHubHermes.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) StakeToken() (common.Address, error) {
	return _WorkerHubHermes.Contract.StakeToken(&_WorkerHubHermes.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesSession) Treasury() (common.Address, error) {
	return _WorkerHubHermes.Contract.Treasury(&_WorkerHubHermes.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Treasury() (common.Address, error) {
	return _WorkerHubHermes.Contract.Treasury(&_WorkerHubHermes.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesSession) UnstakeDelayTime() (*big.Int, error) {
	return _WorkerHubHermes.Contract.UnstakeDelayTime(&_WorkerHubHermes.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _WorkerHubHermes.Contract.UnstakeDelayTime(&_WorkerHubHermes.CallOpts)
}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCaller) ValidatingTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "validatingTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesSession) ValidatingTimeLimit() (*big.Int, error) {
	return _WorkerHubHermes.Contract.ValidatingTimeLimit(&_WorkerHubHermes.CallOpts)
}

// ValidatingTimeLimit is a free data retrieval call binding the contract method 0x56f2eca3.
//
// Solidity: function validatingTimeLimit() view returns(uint40)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) ValidatingTimeLimit() (*big.Int, error) {
	return _WorkerHubHermes.Contract.ValidatingTimeLimit(&_WorkerHubHermes.CallOpts)
}

// ValidatorDisputed is a free data retrieval call binding the contract method 0x7146ab67.
//
// Solidity: function validatorDisputed(address , uint256 ) view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesCaller) ValidatorDisputed(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "validatorDisputed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatorDisputed is a free data retrieval call binding the contract method 0x7146ab67.
//
// Solidity: function validatorDisputed(address , uint256 ) view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesSession) ValidatorDisputed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _WorkerHubHermes.Contract.ValidatorDisputed(&_WorkerHubHermes.CallOpts, arg0, arg1)
}

// ValidatorDisputed is a free data retrieval call binding the contract method 0x7146ab67.
//
// Solidity: function validatorDisputed(address , uint256 ) view returns(bool)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) ValidatorDisputed(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _WorkerHubHermes.Contract.ValidatorDisputed(&_WorkerHubHermes.CallOpts, arg0, arg1)
}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCaller) ValidatorMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "validatorMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) ValidatorMinimumStake() (*big.Int, error) {
	return _WorkerHubHermes.Contract.ValidatorMinimumStake(&_WorkerHubHermes.CallOpts)
}

// ValidatorMinimumStake is a free data retrieval call binding the contract method 0x412f9775.
//
// Solidity: function validatorMinimumStake() view returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) ValidatorMinimumStake() (*big.Int, error) {
	return _WorkerHubHermes.Contract.ValidatorMinimumStake(&_WorkerHubHermes.CallOpts)
}

// ValidatorUnstakeRequests is a free data retrieval call binding the contract method 0xfd3623f7.
//
// Solidity: function validatorUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHubHermes *WorkerHubHermesCaller) ValidatorUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "validatorUnstakeRequests", arg0)

	outstruct := new(struct {
		Stake    *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorUnstakeRequests is a free data retrieval call binding the contract method 0xfd3623f7.
//
// Solidity: function validatorUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHubHermes *WorkerHubHermesSession) ValidatorUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHubHermes.Contract.ValidatorUnstakeRequests(&_WorkerHubHermes.CallOpts, arg0)
}

// ValidatorUnstakeRequests is a free data retrieval call binding the contract method 0xfd3623f7.
//
// Solidity: function validatorUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) ValidatorUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHubHermes.Contract.ValidatorUnstakeRequests(&_WorkerHubHermes.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHubHermes *WorkerHubHermesCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		Commitment       *big.Int
		ModelAddress     common.Address
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModelAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHubHermes *WorkerHubHermesSession) Validators(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHubHermes.Contract.Validators(&_WorkerHubHermes.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Validators(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHubHermes.Contract.Validators(&_WorkerHubHermes.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHubHermes *WorkerHubHermesCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkerHubHermes.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHubHermes *WorkerHubHermesSession) Version() (string, error) {
	return _WorkerHubHermes.Contract.Version(&_WorkerHubHermes.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_WorkerHubHermes *WorkerHubHermesCallerSession) Version() (string, error) {
	return _WorkerHubHermes.Contract.Version(&_WorkerHubHermes.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) ClaimReward(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "claimReward", _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.ClaimReward(&_WorkerHubHermes.TransactOpts, _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.ClaimReward(&_WorkerHubHermes.TransactOpts, _miner)
}

// DisputeInfer is a paid mutator transaction binding the contract method 0x72422748.
//
// Solidity: function disputeInfer(uint256 _assignmentId) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) DisputeInfer(opts *bind.TransactOpts, _assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "disputeInfer", _assignmentId)
}

// DisputeInfer is a paid mutator transaction binding the contract method 0x72422748.
//
// Solidity: function disputeInfer(uint256 _assignmentId) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) DisputeInfer(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.DisputeInfer(&_WorkerHubHermes.TransactOpts, _assignmentId)
}

// DisputeInfer is a paid mutator transaction binding the contract method 0x72422748.
//
// Solidity: function disputeInfer(uint256 _assignmentId) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) DisputeInfer(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.DisputeInfer(&_WorkerHubHermes.TransactOpts, _assignmentId)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) ForceChangeModelForMiner(opts *bind.TransactOpts, _miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "forceChangeModelForMiner", _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.ForceChangeModelForMiner(&_WorkerHubHermes.TransactOpts, _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.ForceChangeModelForMiner(&_WorkerHubHermes.TransactOpts, _miner, _modelAddress)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) IncreaseMinerStake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "increaseMinerStake", stakeAmount)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) IncreaseMinerStake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.IncreaseMinerStake(&_WorkerHubHermes.TransactOpts, stakeAmount)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) IncreaseMinerStake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.IncreaseMinerStake(&_WorkerHubHermes.TransactOpts, stakeAmount)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0xc8cb2903.
//
// Solidity: function increaseValidatorStake(uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) IncreaseValidatorStake(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "increaseValidatorStake", stakeAmount)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0xc8cb2903.
//
// Solidity: function increaseValidatorStake(uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) IncreaseValidatorStake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.IncreaseValidatorStake(&_WorkerHubHermes.TransactOpts, stakeAmount)
}

// IncreaseValidatorStake is a paid mutator transaction binding the contract method 0xc8cb2903.
//
// Solidity: function increaseValidatorStake(uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) IncreaseValidatorStake(stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.IncreaseValidatorStake(&_WorkerHubHermes.TransactOpts, stakeAmount)
}

// Infer is a paid mutator transaction binding the contract method 0xbd4815f5.
//
// Solidity: function infer(bytes _input, address _creator, uint256 cost) returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesTransactor) Infer(opts *bind.TransactOpts, _input []byte, _creator common.Address, cost *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "infer", _input, _creator, cost)
}

// Infer is a paid mutator transaction binding the contract method 0xbd4815f5.
//
// Solidity: function infer(bytes _input, address _creator, uint256 cost) returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) Infer(_input []byte, _creator common.Address, cost *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Infer(&_WorkerHubHermes.TransactOpts, _input, _creator, cost)
}

// Infer is a paid mutator transaction binding the contract method 0xbd4815f5.
//
// Solidity: function infer(bytes _input, address _creator, uint256 cost) returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) Infer(_input []byte, _creator common.Address, cost *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Infer(&_WorkerHubHermes.TransactOpts, _input, _creator, cost)
}

// Initialize is a paid mutator transaction binding the contract method 0xaa3f5300.
//
// Solidity: function initialize(address _treasury, uint16 _feePercentage, uint256 _minerMinimumStake, uint256 _validatorMinimumStake, uint40 _miningTimeLimit, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpochBasedOnPerf, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage, address _stakeToken) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) Initialize(opts *bind.TransactOpts, _treasury common.Address, _feePercentage uint16, _minerMinimumStake *big.Int, _validatorMinimumStake *big.Int, _miningTimeLimit *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpochBasedOnPerf *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16, _stakeToken common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "initialize", _treasury, _feePercentage, _minerMinimumStake, _validatorMinimumStake, _miningTimeLimit, _minerRequirement, _blocksPerEpoch, _rewardPerEpochBasedOnPerf, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage, _stakeToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xaa3f5300.
//
// Solidity: function initialize(address _treasury, uint16 _feePercentage, uint256 _minerMinimumStake, uint256 _validatorMinimumStake, uint40 _miningTimeLimit, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpochBasedOnPerf, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage, address _stakeToken) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) Initialize(_treasury common.Address, _feePercentage uint16, _minerMinimumStake *big.Int, _validatorMinimumStake *big.Int, _miningTimeLimit *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpochBasedOnPerf *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16, _stakeToken common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Initialize(&_WorkerHubHermes.TransactOpts, _treasury, _feePercentage, _minerMinimumStake, _validatorMinimumStake, _miningTimeLimit, _minerRequirement, _blocksPerEpoch, _rewardPerEpochBasedOnPerf, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage, _stakeToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xaa3f5300.
//
// Solidity: function initialize(address _treasury, uint16 _feePercentage, uint256 _minerMinimumStake, uint256 _validatorMinimumStake, uint40 _miningTimeLimit, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpochBasedOnPerf, uint256 _rewardPerEpoch, uint40 _unstakeDelayTime, uint40 _penaltyDuration, uint16 _finePercentage, address _stakeToken) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) Initialize(_treasury common.Address, _feePercentage uint16, _minerMinimumStake *big.Int, _validatorMinimumStake *big.Int, _miningTimeLimit *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpochBasedOnPerf *big.Int, _rewardPerEpoch *big.Int, _unstakeDelayTime *big.Int, _penaltyDuration *big.Int, _finePercentage uint16, _stakeToken common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Initialize(&_WorkerHubHermes.TransactOpts, _treasury, _feePercentage, _minerMinimumStake, _validatorMinimumStake, _miningTimeLimit, _minerRequirement, _blocksPerEpoch, _rewardPerEpochBasedOnPerf, _rewardPerEpoch, _unstakeDelayTime, _penaltyDuration, _finePercentage, _stakeToken)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) JoinForMinting() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.JoinForMinting(&_WorkerHubHermes.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.JoinForMinting(&_WorkerHubHermes.TransactOpts)
}

// JoinForValidating is a paid mutator transaction binding the contract method 0xea4a2cac.
//
// Solidity: function joinForValidating() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) JoinForValidating(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "joinForValidating")
}

// JoinForValidating is a paid mutator transaction binding the contract method 0xea4a2cac.
//
// Solidity: function joinForValidating() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) JoinForValidating() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.JoinForValidating(&_WorkerHubHermes.TransactOpts)
}

// JoinForValidating is a paid mutator transaction binding the contract method 0xea4a2cac.
//
// Solidity: function joinForValidating() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) JoinForValidating() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.JoinForValidating(&_WorkerHubHermes.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) Pause() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Pause(&_WorkerHubHermes.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) Pause() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Pause(&_WorkerHubHermes.TransactOpts)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x668133e3.
//
// Solidity: function registerMiner(uint16 tier, uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "registerMiner", tier, stakeAmount)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x668133e3.
//
// Solidity: function registerMiner(uint16 tier, uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) RegisterMiner(tier uint16, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RegisterMiner(&_WorkerHubHermes.TransactOpts, tier, stakeAmount)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x668133e3.
//
// Solidity: function registerMiner(uint16 tier, uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) RegisterMiner(tier uint16, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RegisterMiner(&_WorkerHubHermes.TransactOpts, tier, stakeAmount)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) RegisterModel(opts *bind.TransactOpts, _model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "registerModel", _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RegisterModel(&_WorkerHubHermes.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RegisterModel(&_WorkerHubHermes.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xc64a5695.
//
// Solidity: function registerValidator(uint16 tier, uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) RegisterValidator(opts *bind.TransactOpts, tier uint16, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "registerValidator", tier, stakeAmount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xc64a5695.
//
// Solidity: function registerValidator(uint16 tier, uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) RegisterValidator(tier uint16, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RegisterValidator(&_WorkerHubHermes.TransactOpts, tier, stakeAmount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xc64a5695.
//
// Solidity: function registerValidator(uint16 tier, uint256 stakeAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) RegisterValidator(tier uint16, stakeAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RegisterValidator(&_WorkerHubHermes.TransactOpts, tier, stakeAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RenounceOwnership(&_WorkerHubHermes.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RenounceOwnership(&_WorkerHubHermes.TransactOpts)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) ResolveInference(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "resolveInference", _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.ResolveInference(&_WorkerHubHermes.TransactOpts, _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.ResolveInference(&_WorkerHubHermes.TransactOpts, _inferenceId)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RestakeForMiner(&_WorkerHubHermes.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RestakeForMiner(&_WorkerHubHermes.TransactOpts, tier)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesTransactor) RewardToClaim(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "rewardToClaim", _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RewardToClaim(&_WorkerHubHermes.TransactOpts, _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.RewardToClaim(&_WorkerHubHermes.TransactOpts, _miner)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "setBlocksPerEpoch", _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetBlocksPerEpoch(&_WorkerHubHermes.TransactOpts, _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetBlocksPerEpoch(&_WorkerHubHermes.TransactOpts, _blocks)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SetFinePercentage(opts *bind.TransactOpts, _finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "setFinePercentage", _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetFinePercentage(&_WorkerHubHermes.TransactOpts, _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetFinePercentage(&_WorkerHubHermes.TransactOpts, _finePercentage)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, _newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "setNewRewardInEpoch", _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetNewRewardInEpoch(&_WorkerHubHermes.TransactOpts, _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetNewRewardInEpoch(&_WorkerHubHermes.TransactOpts, _newRewardAmount)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SetPenaltyDuration(opts *bind.TransactOpts, _penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "setPenaltyDuration", _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetPenaltyDuration(&_WorkerHubHermes.TransactOpts, _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetPenaltyDuration(&_WorkerHubHermes.TransactOpts, _penaltyDuration)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SetUnstakDelayTime(opts *bind.TransactOpts, _newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "setUnstakDelayTime", _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetUnstakDelayTime(&_WorkerHubHermes.TransactOpts, _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SetUnstakDelayTime(&_WorkerHubHermes.TransactOpts, _newUnstakeDelayTime)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SlashMiner(opts *bind.TransactOpts, _miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "slashMiner", _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SlashMiner(&_WorkerHubHermes.TransactOpts, _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SlashMiner(&_WorkerHubHermes.TransactOpts, _miner, _isFined)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) SubmitSolution(opts *bind.TransactOpts, _assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "submitSolution", _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SubmitSolution(&_WorkerHubHermes.TransactOpts, _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.SubmitSolution(&_WorkerHubHermes.TransactOpts, _assigmentId, _data)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xf6bba943.
//
// Solidity: function topUpInfer(uint256 _inferenceId, uint256 topUpAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) TopUpInfer(opts *bind.TransactOpts, _inferenceId *big.Int, topUpAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "topUpInfer", _inferenceId, topUpAmount)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xf6bba943.
//
// Solidity: function topUpInfer(uint256 _inferenceId, uint256 topUpAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) TopUpInfer(_inferenceId *big.Int, topUpAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.TopUpInfer(&_WorkerHubHermes.TransactOpts, _inferenceId, topUpAmount)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xf6bba943.
//
// Solidity: function topUpInfer(uint256 _inferenceId, uint256 topUpAmount) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) TopUpInfer(_inferenceId *big.Int, topUpAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.TopUpInfer(&_WorkerHubHermes.TransactOpts, _inferenceId, topUpAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.TransferOwnership(&_WorkerHubHermes.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.TransferOwnership(&_WorkerHubHermes.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) Unpause() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Unpause(&_WorkerHubHermes.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) Unpause() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Unpause(&_WorkerHubHermes.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UnregisterMiner() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnregisterMiner(&_WorkerHubHermes.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnregisterMiner(&_WorkerHubHermes.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UnregisterModel(opts *bind.TransactOpts, _model common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "unregisterModel", _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnregisterModel(&_WorkerHubHermes.TransactOpts, _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnregisterModel(&_WorkerHubHermes.TransactOpts, _model)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UnregisterValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "unregisterValidator")
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UnregisterValidator() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnregisterValidator(&_WorkerHubHermes.TransactOpts)
}

// UnregisterValidator is a paid mutator transaction binding the contract method 0x6ca56267.
//
// Solidity: function unregisterValidator() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UnregisterValidator() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnregisterValidator(&_WorkerHubHermes.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UnstakeForMiner() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnstakeForMiner(&_WorkerHubHermes.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnstakeForMiner(&_WorkerHubHermes.TransactOpts)
}

// UnstakeForValidator is a paid mutator transaction binding the contract method 0x42b093cf.
//
// Solidity: function unstakeForValidator() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UnstakeForValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "unstakeForValidator")
}

// UnstakeForValidator is a paid mutator transaction binding the contract method 0x42b093cf.
//
// Solidity: function unstakeForValidator() returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UnstakeForValidator() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnstakeForValidator(&_WorkerHubHermes.TransactOpts)
}

// UnstakeForValidator is a paid mutator transaction binding the contract method 0x42b093cf.
//
// Solidity: function unstakeForValidator() returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UnstakeForValidator() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UnstakeForValidator(&_WorkerHubHermes.TransactOpts)
}

// UpdateMiningTimeLimit is a paid mutator transaction binding the contract method 0x34193a4a.
//
// Solidity: function updateMiningTimeLimit(uint40 _miningTimeLimit) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UpdateMiningTimeLimit(opts *bind.TransactOpts, _miningTimeLimit *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "updateMiningTimeLimit", _miningTimeLimit)
}

// UpdateMiningTimeLimit is a paid mutator transaction binding the contract method 0x34193a4a.
//
// Solidity: function updateMiningTimeLimit(uint40 _miningTimeLimit) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UpdateMiningTimeLimit(_miningTimeLimit *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UpdateMiningTimeLimit(&_WorkerHubHermes.TransactOpts, _miningTimeLimit)
}

// UpdateMiningTimeLimit is a paid mutator transaction binding the contract method 0x34193a4a.
//
// Solidity: function updateMiningTimeLimit(uint40 _miningTimeLimit) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UpdateMiningTimeLimit(_miningTimeLimit *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UpdateMiningTimeLimit(&_WorkerHubHermes.TransactOpts, _miningTimeLimit)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, _model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "updateModelMinimumFee", _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UpdateModelMinimumFee(&_WorkerHubHermes.TransactOpts, _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UpdateModelMinimumFee(&_WorkerHubHermes.TransactOpts, _model, _minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) UpdateModelTier(opts *bind.TransactOpts, _model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.Transact(opts, "updateModelTier", _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHubHermes *WorkerHubHermesSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UpdateModelTier(&_WorkerHubHermes.TransactOpts, _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.UpdateModelTier(&_WorkerHubHermes.TransactOpts, _model, _tier)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHubHermes *WorkerHubHermesTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHubHermes.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHubHermes *WorkerHubHermesSession) Receive() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Receive(&_WorkerHubHermes.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WorkerHubHermes *WorkerHubHermesTransactorSession) Receive() (*types.Transaction, error) {
	return _WorkerHubHermes.Contract.Receive(&_WorkerHubHermes.TransactOpts)
}

// WorkerHubHermesBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the WorkerHubHermes contract.
type WorkerHubHermesBlocksPerEpochIterator struct {
	Event *WorkerHubHermesBlocksPerEpoch // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesBlocksPerEpoch)
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
		it.Event = new(WorkerHubHermesBlocksPerEpoch)
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
func (it *WorkerHubHermesBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesBlocksPerEpoch represents a BlocksPerEpoch event raised by the WorkerHubHermes contract.
type WorkerHubHermesBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*WorkerHubHermesBlocksPerEpochIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesBlocksPerEpochIterator{contract: _WorkerHubHermes.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesBlocksPerEpoch)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
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

// ParseBlocksPerEpoch is a log parse operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseBlocksPerEpoch(log types.Log) (*WorkerHubHermesBlocksPerEpoch, error) {
	event := new(WorkerHubHermesBlocksPerEpoch)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the WorkerHubHermes contract.
type WorkerHubHermesFinePercentageUpdatedIterator struct {
	Event *WorkerHubHermesFinePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesFinePercentageUpdated)
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
		it.Event = new(WorkerHubHermesFinePercentageUpdated)
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
func (it *WorkerHubHermesFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesFinePercentageUpdated represents a FinePercentageUpdated event raised by the WorkerHubHermes contract.
type WorkerHubHermesFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*WorkerHubHermesFinePercentageUpdatedIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesFinePercentageUpdatedIterator{contract: _WorkerHubHermes.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesFinePercentageUpdated)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
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

// ParseFinePercentageUpdated is a log parse operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseFinePercentageUpdated(log types.Log) (*WorkerHubHermesFinePercentageUpdated, error) {
	event := new(WorkerHubHermesFinePercentageUpdated)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the WorkerHubHermes contract.
type WorkerHubHermesFraudulentMinerPenalizedIterator struct {
	Event *WorkerHubHermesFraudulentMinerPenalized // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesFraudulentMinerPenalized)
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
		it.Event = new(WorkerHubHermesFraudulentMinerPenalized)
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
func (it *WorkerHubHermesFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the WorkerHubHermes contract.
type WorkerHubHermesFraudulentMinerPenalized struct {
	Miner        common.Address
	ModelAddress common.Address
	Treasury     common.Address
	Fine         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (*WorkerHubHermesFraudulentMinerPenalizedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesFraudulentMinerPenalizedIterator{contract: _WorkerHubHermes.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesFraudulentMinerPenalized, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesFraudulentMinerPenalized)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
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

// ParseFraudulentMinerPenalized is a log parse operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseFraudulentMinerPenalized(log types.Log) (*WorkerHubHermesFraudulentMinerPenalized, error) {
	event := new(WorkerHubHermesFraudulentMinerPenalized)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesInferenceDisputationIterator is returned from FilterInferenceDisputation and is used to iterate over the raw logs and unpacked data for InferenceDisputation events raised by the WorkerHubHermes contract.
type WorkerHubHermesInferenceDisputationIterator struct {
	Event *WorkerHubHermesInferenceDisputation // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesInferenceDisputationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesInferenceDisputation)
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
		it.Event = new(WorkerHubHermesInferenceDisputation)
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
func (it *WorkerHubHermesInferenceDisputationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesInferenceDisputationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesInferenceDisputation represents a InferenceDisputation event raised by the WorkerHubHermes contract.
type WorkerHubHermesInferenceDisputation struct {
	Validator   common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceDisputation is a free log retrieval operation binding the contract event 0x7a35142f149dfe3f5cd7125e68104edcde63a29e539db29b23cf0823512dec9c.
//
// Solidity: event InferenceDisputation(address indexed validator, uint256 indexed assigmentId)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterInferenceDisputation(opts *bind.FilterOpts, validator []common.Address, assigmentId []*big.Int) (*WorkerHubHermesInferenceDisputationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "InferenceDisputation", validatorRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesInferenceDisputationIterator{contract: _WorkerHubHermes.contract, event: "InferenceDisputation", logs: logs, sub: sub}, nil
}

// WatchInferenceDisputation is a free log subscription operation binding the contract event 0x7a35142f149dfe3f5cd7125e68104edcde63a29e539db29b23cf0823512dec9c.
//
// Solidity: event InferenceDisputation(address indexed validator, uint256 indexed assigmentId)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchInferenceDisputation(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesInferenceDisputation, validator []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "InferenceDisputation", validatorRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesInferenceDisputation)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "InferenceDisputation", log); err != nil {
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

// ParseInferenceDisputation is a log parse operation binding the contract event 0x7a35142f149dfe3f5cd7125e68104edcde63a29e539db29b23cf0823512dec9c.
//
// Solidity: event InferenceDisputation(address indexed validator, uint256 indexed assigmentId)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseInferenceDisputation(log types.Log) (*WorkerHubHermesInferenceDisputation, error) {
	event := new(WorkerHubHermesInferenceDisputation)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "InferenceDisputation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the WorkerHubHermes contract.
type WorkerHubHermesInferenceStatusUpdateIterator struct {
	Event *WorkerHubHermesInferenceStatusUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesInferenceStatusUpdate)
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
		it.Event = new(WorkerHubHermesInferenceStatusUpdate)
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
func (it *WorkerHubHermesInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the WorkerHubHermes contract.
type WorkerHubHermesInferenceStatusUpdate struct {
	InferenceId *big.Int
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []*big.Int) (*WorkerHubHermesInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesInferenceStatusUpdateIterator{contract: _WorkerHubHermes.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesInferenceStatusUpdate, inferenceId []*big.Int) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesInferenceStatusUpdate)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
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

// ParseInferenceStatusUpdate is a log parse operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseInferenceStatusUpdate(log types.Log) (*WorkerHubHermesInferenceStatusUpdate, error) {
	event := new(WorkerHubHermesInferenceStatusUpdate)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkerHubHermes contract.
type WorkerHubHermesInitializedIterator struct {
	Event *WorkerHubHermesInitialized // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesInitialized)
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
		it.Event = new(WorkerHubHermesInitialized)
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
func (it *WorkerHubHermesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesInitialized represents a Initialized event raised by the WorkerHubHermes contract.
type WorkerHubHermesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkerHubHermesInitializedIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesInitializedIterator{contract: _WorkerHubHermes.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesInitialized)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseInitialized(log types.Log) (*WorkerHubHermesInitialized, error) {
	event := new(WorkerHubHermesInitialized)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerDeactivatedIterator struct {
	Event *WorkerHubHermesMinerDeactivated // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMinerDeactivated)
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
		it.Event = new(WorkerHubHermesMinerDeactivated)
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
func (it *WorkerHubHermesMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMinerDeactivated represents a MinerDeactivated event raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerDeactivated struct {
	Miner        common.Address
	ModelAddress common.Address
	ActiveTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address) (*WorkerHubHermesMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMinerDeactivatedIterator{contract: _WorkerHubHermes.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMinerDeactivated, miner []common.Address, modelAddress []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMinerDeactivated)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
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

// ParseMinerDeactivated is a log parse operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMinerDeactivated(log types.Log) (*WorkerHubHermesMinerDeactivated, error) {
	event := new(WorkerHubHermesMinerDeactivated)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerExtraStakeIterator struct {
	Event *WorkerHubHermesMinerExtraStake // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMinerExtraStake)
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
		it.Event = new(WorkerHubHermesMinerExtraStake)
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
func (it *WorkerHubHermesMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMinerExtraStake represents a MinerExtraStake event raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubHermesMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMinerExtraStakeIterator{contract: _WorkerHubHermes.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMinerExtraStake)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
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

// ParseMinerExtraStake is a log parse operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMinerExtraStake(log types.Log) (*WorkerHubHermesMinerExtraStake, error) {
	event := new(WorkerHubHermesMinerExtraStake)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerJoinIterator struct {
	Event *WorkerHubHermesMinerJoin // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMinerJoin)
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
		it.Event = new(WorkerHubHermesMinerJoin)
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
func (it *WorkerHubHermesMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMinerJoin represents a MinerJoin event raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubHermesMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMinerJoinIterator{contract: _WorkerHubHermes.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMinerJoin)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerJoin", log); err != nil {
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

// ParseMinerJoin is a log parse operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMinerJoin(log types.Log) (*WorkerHubHermesMinerJoin, error) {
	event := new(WorkerHubHermesMinerJoin)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerRegistrationIterator struct {
	Event *WorkerHubHermesMinerRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMinerRegistration)
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
		it.Event = new(WorkerHubHermesMinerRegistration)
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
func (it *WorkerHubHermesMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMinerRegistration represents a MinerRegistration event raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*WorkerHubHermesMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMinerRegistrationIterator{contract: _WorkerHubHermes.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMinerRegistration)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
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

// ParseMinerRegistration is a log parse operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMinerRegistration(log types.Log) (*WorkerHubHermesMinerRegistration, error) {
	event := new(WorkerHubHermesMinerRegistration)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerUnregistrationIterator struct {
	Event *WorkerHubHermesMinerUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMinerUnregistration)
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
		it.Event = new(WorkerHubHermesMinerUnregistration)
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
func (it *WorkerHubHermesMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMinerUnregistration represents a MinerUnregistration event raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubHermesMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMinerUnregistrationIterator{contract: _WorkerHubHermes.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMinerUnregistration)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
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

// ParseMinerUnregistration is a log parse operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMinerUnregistration(log types.Log) (*WorkerHubHermesMinerUnregistration, error) {
	event := new(WorkerHubHermesMinerUnregistration)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerUnstakeIterator struct {
	Event *WorkerHubHermesMinerUnstake // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMinerUnstake)
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
		it.Event = new(WorkerHubHermesMinerUnstake)
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
func (it *WorkerHubHermesMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMinerUnstake represents a MinerUnstake event raised by the WorkerHubHermes contract.
type WorkerHubHermesMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubHermesMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMinerUnstakeIterator{contract: _WorkerHubHermes.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMinerUnstake)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
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

// ParseMinerUnstake is a log parse operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMinerUnstake(log types.Log) (*WorkerHubHermesMinerUnstake, error) {
	event := new(WorkerHubHermesMinerUnstake)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesMiningTimeLimitUpdateIterator is returned from FilterMiningTimeLimitUpdate and is used to iterate over the raw logs and unpacked data for MiningTimeLimitUpdate events raised by the WorkerHubHermes contract.
type WorkerHubHermesMiningTimeLimitUpdateIterator struct {
	Event *WorkerHubHermesMiningTimeLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesMiningTimeLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesMiningTimeLimitUpdate)
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
		it.Event = new(WorkerHubHermesMiningTimeLimitUpdate)
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
func (it *WorkerHubHermesMiningTimeLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesMiningTimeLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesMiningTimeLimitUpdate represents a MiningTimeLimitUpdate event raised by the WorkerHubHermes contract.
type WorkerHubHermesMiningTimeLimitUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMiningTimeLimitUpdate is a free log retrieval operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterMiningTimeLimitUpdate(opts *bind.FilterOpts) (*WorkerHubHermesMiningTimeLimitUpdateIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesMiningTimeLimitUpdateIterator{contract: _WorkerHubHermes.contract, event: "MiningTimeLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchMiningTimeLimitUpdate is a free log subscription operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchMiningTimeLimitUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesMiningTimeLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesMiningTimeLimitUpdate)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
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

// ParseMiningTimeLimitUpdate is a log parse operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseMiningTimeLimitUpdate(log types.Log) (*WorkerHubHermesMiningTimeLimitUpdate, error) {
	event := new(WorkerHubHermesMiningTimeLimitUpdate)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the WorkerHubHermes contract.
type WorkerHubHermesModelMinimumFeeUpdateIterator struct {
	Event *WorkerHubHermesModelMinimumFeeUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesModelMinimumFeeUpdate)
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
		it.Event = new(WorkerHubHermesModelMinimumFeeUpdate)
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
func (it *WorkerHubHermesModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the WorkerHubHermes contract.
type WorkerHubHermesModelMinimumFeeUpdate struct {
	Model      common.Address
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, model []common.Address) (*WorkerHubHermesModelMinimumFeeUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesModelMinimumFeeUpdateIterator{contract: _WorkerHubHermes.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesModelMinimumFeeUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesModelMinimumFeeUpdate)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
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

// ParseModelMinimumFeeUpdate is a log parse operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*WorkerHubHermesModelMinimumFeeUpdate, error) {
	event := new(WorkerHubHermesModelMinimumFeeUpdate)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the WorkerHubHermes contract.
type WorkerHubHermesModelRegistrationIterator struct {
	Event *WorkerHubHermesModelRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesModelRegistration)
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
		it.Event = new(WorkerHubHermesModelRegistration)
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
func (it *WorkerHubHermesModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesModelRegistration represents a ModelRegistration event raised by the WorkerHubHermes contract.
type WorkerHubHermesModelRegistration struct {
	Model      common.Address
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterModelRegistration(opts *bind.FilterOpts, model []common.Address, tier []uint16) (*WorkerHubHermesModelRegistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesModelRegistrationIterator{contract: _WorkerHubHermes.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesModelRegistration, model []common.Address, tier []uint16) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesModelRegistration)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
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

// ParseModelRegistration is a log parse operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseModelRegistration(log types.Log) (*WorkerHubHermesModelRegistration, error) {
	event := new(WorkerHubHermesModelRegistration)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the WorkerHubHermes contract.
type WorkerHubHermesModelTierUpdateIterator struct {
	Event *WorkerHubHermesModelTierUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesModelTierUpdate)
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
		it.Event = new(WorkerHubHermesModelTierUpdate)
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
func (it *WorkerHubHermesModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesModelTierUpdate represents a ModelTierUpdate event raised by the WorkerHubHermes contract.
type WorkerHubHermesModelTierUpdate struct {
	Model common.Address
	Tier  uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, model []common.Address) (*WorkerHubHermesModelTierUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesModelTierUpdateIterator{contract: _WorkerHubHermes.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesModelTierUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesModelTierUpdate)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
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

// ParseModelTierUpdate is a log parse operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseModelTierUpdate(log types.Log) (*WorkerHubHermesModelTierUpdate, error) {
	event := new(WorkerHubHermesModelTierUpdate)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the WorkerHubHermes contract.
type WorkerHubHermesModelUnregistrationIterator struct {
	Event *WorkerHubHermesModelUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesModelUnregistration)
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
		it.Event = new(WorkerHubHermesModelUnregistration)
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
func (it *WorkerHubHermesModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesModelUnregistration represents a ModelUnregistration event raised by the WorkerHubHermes contract.
type WorkerHubHermesModelUnregistration struct {
	Model common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterModelUnregistration(opts *bind.FilterOpts, model []common.Address) (*WorkerHubHermesModelUnregistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesModelUnregistrationIterator{contract: _WorkerHubHermes.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesModelUnregistration, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesModelUnregistration)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseModelUnregistration(log types.Log) (*WorkerHubHermesModelUnregistration, error) {
	event := new(WorkerHubHermesModelUnregistration)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the WorkerHubHermes contract.
type WorkerHubHermesNewAssignmentIterator struct {
	Event *WorkerHubHermesNewAssignment // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesNewAssignment)
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
		it.Event = new(WorkerHubHermesNewAssignment)
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
func (it *WorkerHubHermesNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesNewAssignment represents a NewAssignment event raised by the WorkerHubHermes contract.
type WorkerHubHermesNewAssignment struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	ExpiredAt    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterNewAssignment(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*WorkerHubHermesNewAssignmentIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesNewAssignmentIterator{contract: _WorkerHubHermes.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesNewAssignment, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesNewAssignment)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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

// ParseNewAssignment is a log parse operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseNewAssignment(log types.Log) (*WorkerHubHermesNewAssignment, error) {
	event := new(WorkerHubHermesNewAssignment)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the WorkerHubHermes contract.
type WorkerHubHermesNewInferenceIterator struct {
	Event *WorkerHubHermesNewInference // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesNewInference)
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
		it.Event = new(WorkerHubHermesNewInference)
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
func (it *WorkerHubHermesNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesNewInference represents a NewInference event raised by the WorkerHubHermes contract.
type WorkerHubHermesNewInference struct {
	InferenceId *big.Int
	Model       common.Address
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x77f5630b01c86fd0283f51543024803955bf2c66ddc644ef0ebd1dd193d032ee.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*WorkerHubHermesNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesNewInferenceIterator{contract: _WorkerHubHermes.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x77f5630b01c86fd0283f51543024803955bf2c66ddc644ef0ebd1dd193d032ee.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesNewInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesNewInference)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "NewInference", log); err != nil {
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

// ParseNewInference is a log parse operation binding the contract event 0x77f5630b01c86fd0283f51543024803955bf2c66ddc644ef0ebd1dd193d032ee.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseNewInference(log types.Log) (*WorkerHubHermesNewInference, error) {
	event := new(WorkerHubHermesNewInference)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkerHubHermes contract.
type WorkerHubHermesOwnershipTransferredIterator struct {
	Event *WorkerHubHermesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesOwnershipTransferred)
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
		it.Event = new(WorkerHubHermesOwnershipTransferred)
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
func (it *WorkerHubHermesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesOwnershipTransferred represents a OwnershipTransferred event raised by the WorkerHubHermes contract.
type WorkerHubHermesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkerHubHermesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesOwnershipTransferredIterator{contract: _WorkerHubHermes.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesOwnershipTransferred)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseOwnershipTransferred(log types.Log) (*WorkerHubHermesOwnershipTransferred, error) {
	event := new(WorkerHubHermesOwnershipTransferred)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WorkerHubHermes contract.
type WorkerHubHermesPausedIterator struct {
	Event *WorkerHubHermesPaused // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesPaused)
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
		it.Event = new(WorkerHubHermesPaused)
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
func (it *WorkerHubHermesPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesPaused represents a Paused event raised by the WorkerHubHermes contract.
type WorkerHubHermesPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterPaused(opts *bind.FilterOpts) (*WorkerHubHermesPausedIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesPausedIterator{contract: _WorkerHubHermes.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesPaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesPaused)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParsePaused(log types.Log) (*WorkerHubHermesPaused, error) {
	event := new(WorkerHubHermesPaused)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the WorkerHubHermes contract.
type WorkerHubHermesPenaltyDurationUpdatedIterator struct {
	Event *WorkerHubHermesPenaltyDurationUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesPenaltyDurationUpdated)
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
		it.Event = new(WorkerHubHermesPenaltyDurationUpdated)
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
func (it *WorkerHubHermesPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the WorkerHubHermes contract.
type WorkerHubHermesPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*WorkerHubHermesPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesPenaltyDurationUpdatedIterator{contract: _WorkerHubHermes.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesPenaltyDurationUpdated)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
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

// ParsePenaltyDurationUpdated is a log parse operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParsePenaltyDurationUpdated(log types.Log) (*WorkerHubHermesPenaltyDurationUpdated, error) {
	event := new(WorkerHubHermesPenaltyDurationUpdated)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the WorkerHubHermes contract.
type WorkerHubHermesRestakeIterator struct {
	Event *WorkerHubHermesRestake // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesRestake)
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
		it.Event = new(WorkerHubHermesRestake)
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
func (it *WorkerHubHermesRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesRestake represents a Restake event raised by the WorkerHubHermes contract.
type WorkerHubHermesRestake struct {
	Miner   common.Address
	Restake *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, model []common.Address) (*WorkerHubHermesRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesRestakeIterator{contract: _WorkerHubHermes.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesRestake, miner []common.Address, model []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesRestake)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "Restake", log); err != nil {
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

// ParseRestake is a log parse operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseRestake(log types.Log) (*WorkerHubHermesRestake, error) {
	event := new(WorkerHubHermesRestake)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the WorkerHubHermes contract.
type WorkerHubHermesRewardClaimIterator struct {
	Event *WorkerHubHermesRewardClaim // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesRewardClaim)
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
		it.Event = new(WorkerHubHermesRewardClaim)
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
func (it *WorkerHubHermesRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesRewardClaim represents a RewardClaim event raised by the WorkerHubHermes contract.
type WorkerHubHermesRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*WorkerHubHermesRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesRewardClaimIterator{contract: _WorkerHubHermes.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesRewardClaim)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "RewardClaim", log); err != nil {
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

// ParseRewardClaim is a log parse operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseRewardClaim(log types.Log) (*WorkerHubHermesRewardClaim, error) {
	event := new(WorkerHubHermesRewardClaim)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the WorkerHubHermes contract.
type WorkerHubHermesRewardPerEpochIterator struct {
	Event *WorkerHubHermesRewardPerEpoch // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesRewardPerEpoch)
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
		it.Event = new(WorkerHubHermesRewardPerEpoch)
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
func (it *WorkerHubHermesRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesRewardPerEpoch represents a RewardPerEpoch event raised by the WorkerHubHermes contract.
type WorkerHubHermesRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*WorkerHubHermesRewardPerEpochIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesRewardPerEpochIterator{contract: _WorkerHubHermes.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesRewardPerEpoch)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
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

// ParseRewardPerEpoch is a log parse operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseRewardPerEpoch(log types.Log) (*WorkerHubHermesRewardPerEpoch, error) {
	event := new(WorkerHubHermesRewardPerEpoch)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the WorkerHubHermes contract.
type WorkerHubHermesSolutionSubmissionIterator struct {
	Event *WorkerHubHermesSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesSolutionSubmission)
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
		it.Event = new(WorkerHubHermesSolutionSubmission)
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
func (it *WorkerHubHermesSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesSolutionSubmission represents a SolutionSubmission event raised by the WorkerHubHermes contract.
type WorkerHubHermesSolutionSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubHermesSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesSolutionSubmissionIterator{contract: _WorkerHubHermes.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesSolutionSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesSolutionSubmission)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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

// ParseSolutionSubmission is a log parse operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseSolutionSubmission(log types.Log) (*WorkerHubHermesSolutionSubmission, error) {
	event := new(WorkerHubHermesSolutionSubmission)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesTopUpInferIterator is returned from FilterTopUpInfer and is used to iterate over the raw logs and unpacked data for TopUpInfer events raised by the WorkerHubHermes contract.
type WorkerHubHermesTopUpInferIterator struct {
	Event *WorkerHubHermesTopUpInfer // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesTopUpInferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesTopUpInfer)
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
		it.Event = new(WorkerHubHermesTopUpInfer)
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
func (it *WorkerHubHermesTopUpInferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesTopUpInferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesTopUpInfer represents a TopUpInfer event raised by the WorkerHubHermes contract.
type WorkerHubHermesTopUpInfer struct {
	InferenceId *big.Int
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTopUpInfer is a free log retrieval operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterTopUpInfer(opts *bind.FilterOpts, inferenceId []*big.Int, creator []common.Address) (*WorkerHubHermesTopUpInferIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesTopUpInferIterator{contract: _WorkerHubHermes.contract, event: "TopUpInfer", logs: logs, sub: sub}, nil
}

// WatchTopUpInfer is a free log subscription operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchTopUpInfer(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesTopUpInfer, inferenceId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesTopUpInfer)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
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

// ParseTopUpInfer is a log parse operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseTopUpInfer(log types.Log) (*WorkerHubHermesTopUpInfer, error) {
	event := new(WorkerHubHermesTopUpInfer)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesTransferFeeIterator is returned from FilterTransferFee and is used to iterate over the raw logs and unpacked data for TransferFee events raised by the WorkerHubHermes contract.
type WorkerHubHermesTransferFeeIterator struct {
	Event *WorkerHubHermesTransferFee // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesTransferFee)
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
		it.Event = new(WorkerHubHermesTransferFee)
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
func (it *WorkerHubHermesTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesTransferFee represents a TransferFee event raised by the WorkerHubHermes contract.
type WorkerHubHermesTransferFee struct {
	Miner       common.Address
	MingingFee  *big.Int
	Treasury    common.Address
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferFee is a free log retrieval operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed miner, uint256 mingingFee, address indexed treasury, uint256 protocolFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterTransferFee(opts *bind.FilterOpts, miner []common.Address, treasury []common.Address) (*WorkerHubHermesTransferFeeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "TransferFee", minerRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesTransferFeeIterator{contract: _WorkerHubHermes.contract, event: "TransferFee", logs: logs, sub: sub}, nil
}

// WatchTransferFee is a free log subscription operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed miner, uint256 mingingFee, address indexed treasury, uint256 protocolFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchTransferFee(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesTransferFee, miner []common.Address, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "TransferFee", minerRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesTransferFee)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "TransferFee", log); err != nil {
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

// ParseTransferFee is a log parse operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed miner, uint256 mingingFee, address indexed treasury, uint256 protocolFee)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseTransferFee(log types.Log) (*WorkerHubHermesTransferFee, error) {
	event := new(WorkerHubHermesTransferFee)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "TransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WorkerHubHermes contract.
type WorkerHubHermesUnpausedIterator struct {
	Event *WorkerHubHermesUnpaused // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesUnpaused)
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
		it.Event = new(WorkerHubHermesUnpaused)
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
func (it *WorkerHubHermesUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesUnpaused represents a Unpaused event raised by the WorkerHubHermes contract.
type WorkerHubHermesUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WorkerHubHermesUnpausedIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesUnpausedIterator{contract: _WorkerHubHermes.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesUnpaused) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesUnpaused)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseUnpaused(log types.Log) (*WorkerHubHermesUnpaused, error) {
	event := new(WorkerHubHermesUnpaused)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the WorkerHubHermes contract.
type WorkerHubHermesUnstakeDelayTimeIterator struct {
	Event *WorkerHubHermesUnstakeDelayTime // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesUnstakeDelayTime)
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
		it.Event = new(WorkerHubHermesUnstakeDelayTime)
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
func (it *WorkerHubHermesUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesUnstakeDelayTime represents a UnstakeDelayTime event raised by the WorkerHubHermes contract.
type WorkerHubHermesUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*WorkerHubHermesUnstakeDelayTimeIterator, error) {

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesUnstakeDelayTimeIterator{contract: _WorkerHubHermes.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesUnstakeDelayTime)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
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

// ParseUnstakeDelayTime is a log parse operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseUnstakeDelayTime(log types.Log) (*WorkerHubHermesUnstakeDelayTime, error) {
	event := new(WorkerHubHermesUnstakeDelayTime)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesValidatorExtraStakeIterator is returned from FilterValidatorExtraStake and is used to iterate over the raw logs and unpacked data for ValidatorExtraStake events raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorExtraStakeIterator struct {
	Event *WorkerHubHermesValidatorExtraStake // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesValidatorExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesValidatorExtraStake)
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
		it.Event = new(WorkerHubHermesValidatorExtraStake)
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
func (it *WorkerHubHermesValidatorExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesValidatorExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesValidatorExtraStake represents a ValidatorExtraStake event raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorExtraStake struct {
	Validator common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorExtraStake is a free log retrieval operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterValidatorExtraStake(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubHermesValidatorExtraStakeIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ValidatorExtraStake", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesValidatorExtraStakeIterator{contract: _WorkerHubHermes.contract, event: "ValidatorExtraStake", logs: logs, sub: sub}, nil
}

// WatchValidatorExtraStake is a free log subscription operation binding the contract event 0x34922005bddea1820fe67d4e0d79b91845321a99fc0d43fe025b74ac23e1063d.
//
// Solidity: event ValidatorExtraStake(address indexed validator, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchValidatorExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesValidatorExtraStake, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ValidatorExtraStake", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesValidatorExtraStake)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorExtraStake", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseValidatorExtraStake(log types.Log) (*WorkerHubHermesValidatorExtraStake, error) {
	event := new(WorkerHubHermesValidatorExtraStake)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesValidatorJoinIterator is returned from FilterValidatorJoin and is used to iterate over the raw logs and unpacked data for ValidatorJoin events raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorJoinIterator struct {
	Event *WorkerHubHermesValidatorJoin // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesValidatorJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesValidatorJoin)
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
		it.Event = new(WorkerHubHermesValidatorJoin)
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
func (it *WorkerHubHermesValidatorJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesValidatorJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesValidatorJoin represents a ValidatorJoin event raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorJoin struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJoin is a free log retrieval operation binding the contract event 0x8c97974b44dade54a2967b9de94b1229332187b50127027ff91c2e9b28d69d75.
//
// Solidity: event ValidatorJoin(address indexed validator)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterValidatorJoin(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubHermesValidatorJoinIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ValidatorJoin", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesValidatorJoinIterator{contract: _WorkerHubHermes.contract, event: "ValidatorJoin", logs: logs, sub: sub}, nil
}

// WatchValidatorJoin is a free log subscription operation binding the contract event 0x8c97974b44dade54a2967b9de94b1229332187b50127027ff91c2e9b28d69d75.
//
// Solidity: event ValidatorJoin(address indexed validator)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchValidatorJoin(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesValidatorJoin, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ValidatorJoin", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesValidatorJoin)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorJoin", log); err != nil {
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

// ParseValidatorJoin is a log parse operation binding the contract event 0x8c97974b44dade54a2967b9de94b1229332187b50127027ff91c2e9b28d69d75.
//
// Solidity: event ValidatorJoin(address indexed validator)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseValidatorJoin(log types.Log) (*WorkerHubHermesValidatorJoin, error) {
	event := new(WorkerHubHermesValidatorJoin)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesValidatorRegistrationIterator is returned from FilterValidatorRegistration and is used to iterate over the raw logs and unpacked data for ValidatorRegistration events raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorRegistrationIterator struct {
	Event *WorkerHubHermesValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesValidatorRegistration)
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
		it.Event = new(WorkerHubHermesValidatorRegistration)
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
func (it *WorkerHubHermesValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesValidatorRegistration represents a ValidatorRegistration event raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorRegistration struct {
	Validator common.Address
	Tier      uint16
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistration is a free log retrieval operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterValidatorRegistration(opts *bind.FilterOpts, validator []common.Address, tier []uint16) (*WorkerHubHermesValidatorRegistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ValidatorRegistration", validatorRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesValidatorRegistrationIterator{contract: _WorkerHubHermes.contract, event: "ValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistration is a free log subscription operation binding the contract event 0xcc39e21667abb04befce1bb972c8b03a1b15e1f4b84a3db2535c2fcd6179bb6f.
//
// Solidity: event ValidatorRegistration(address indexed validator, uint16 indexed tier, uint256 value)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchValidatorRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesValidatorRegistration, validator []common.Address, tier []uint16) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ValidatorRegistration", validatorRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesValidatorRegistration)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseValidatorRegistration(log types.Log) (*WorkerHubHermesValidatorRegistration, error) {
	event := new(WorkerHubHermesValidatorRegistration)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesValidatorUnregistrationIterator is returned from FilterValidatorUnregistration and is used to iterate over the raw logs and unpacked data for ValidatorUnregistration events raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorUnregistrationIterator struct {
	Event *WorkerHubHermesValidatorUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesValidatorUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesValidatorUnregistration)
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
		it.Event = new(WorkerHubHermesValidatorUnregistration)
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
func (it *WorkerHubHermesValidatorUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesValidatorUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesValidatorUnregistration represents a ValidatorUnregistration event raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorUnregistration struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnregistration is a free log retrieval operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterValidatorUnregistration(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubHermesValidatorUnregistrationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ValidatorUnregistration", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesValidatorUnregistrationIterator{contract: _WorkerHubHermes.contract, event: "ValidatorUnregistration", logs: logs, sub: sub}, nil
}

// WatchValidatorUnregistration is a free log subscription operation binding the contract event 0x3b4cb6e47f5990bd17a69f36b4e7e5b9eca32c5fbc0b09ffa37149fb77348816.
//
// Solidity: event ValidatorUnregistration(address indexed validator)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchValidatorUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesValidatorUnregistration, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ValidatorUnregistration", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesValidatorUnregistration)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorUnregistration", log); err != nil {
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
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseValidatorUnregistration(log types.Log) (*WorkerHubHermesValidatorUnregistration, error) {
	event := new(WorkerHubHermesValidatorUnregistration)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubHermesValidatorUnstakeIterator is returned from FilterValidatorUnstake and is used to iterate over the raw logs and unpacked data for ValidatorUnstake events raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorUnstakeIterator struct {
	Event *WorkerHubHermesValidatorUnstake // Event containing the contract specifics and raw log

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
func (it *WorkerHubHermesValidatorUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubHermesValidatorUnstake)
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
		it.Event = new(WorkerHubHermesValidatorUnstake)
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
func (it *WorkerHubHermesValidatorUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubHermesValidatorUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubHermesValidatorUnstake represents a ValidatorUnstake event raised by the WorkerHubHermes contract.
type WorkerHubHermesValidatorUnstake struct {
	Validator common.Address
	Stake     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnstake is a free log retrieval operation binding the contract event 0xfa8e6fe2f340a3c8da106ab24f4645cdc4f034a0c0514cabfa764531f13798b3.
//
// Solidity: event ValidatorUnstake(address indexed validator, uint256 stake)
func (_WorkerHubHermes *WorkerHubHermesFilterer) FilterValidatorUnstake(opts *bind.FilterOpts, validator []common.Address) (*WorkerHubHermesValidatorUnstakeIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.FilterLogs(opts, "ValidatorUnstake", validatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubHermesValidatorUnstakeIterator{contract: _WorkerHubHermes.contract, event: "ValidatorUnstake", logs: logs, sub: sub}, nil
}

// WatchValidatorUnstake is a free log subscription operation binding the contract event 0xfa8e6fe2f340a3c8da106ab24f4645cdc4f034a0c0514cabfa764531f13798b3.
//
// Solidity: event ValidatorUnstake(address indexed validator, uint256 stake)
func (_WorkerHubHermes *WorkerHubHermesFilterer) WatchValidatorUnstake(opts *bind.WatchOpts, sink chan<- *WorkerHubHermesValidatorUnstake, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _WorkerHubHermes.contract.WatchLogs(opts, "ValidatorUnstake", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubHermesValidatorUnstake)
				if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorUnstake", log); err != nil {
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

// ParseValidatorUnstake is a log parse operation binding the contract event 0xfa8e6fe2f340a3c8da106ab24f4645cdc4f034a0c0514cabfa764531f13798b3.
//
// Solidity: event ValidatorUnstake(address indexed validator, uint256 stake)
func (_WorkerHubHermes *WorkerHubHermesFilterer) ParseValidatorUnstake(log types.Log) (*WorkerHubHermesValidatorUnstake, error) {
	event := new(WorkerHubHermesValidatorUnstake)
	if err := _WorkerHubHermes.contract.UnpackLog(event, "ValidatorUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
