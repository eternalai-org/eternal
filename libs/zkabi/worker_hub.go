// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkabi

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

// IWorkerHubAssignment is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignment struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}

// IWorkerHubAssignmentInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignmentInfo struct {
	AssignmentId  *big.Int
	InferenceId   *big.Int
	Value         *big.Int
	Input         []byte
	ModelAddress  common.Address
	Creator       common.Address
	SubmitTimeout *big.Int
	CommitTimeout *big.Int
	RevealTimeout *big.Int
}

// IWorkerHubDAOTokenPercentage is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubDAOTokenPercentage struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}

// IWorkerHubDAOTokenReceiverInfor is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubDAOTokenReceiverInfor struct {
	Receiver common.Address
	Amount   *big.Int
	Role     uint8
}

// IWorkerHubInference is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubInference struct {
	Assignments    []*big.Int
	Input          []byte
	Value          *big.Int
	FeeL2          *big.Int
	FeeTreasury    *big.Int
	ModelAddress   common.Address
	SubmitTimeout  *big.Int
	CommitTimeout  *big.Int
	RevealTimeout  *big.Int
	Status         uint8
	Creator        common.Address
	ProcessedMiner common.Address
	Referrer       common.Address
}

// IWorkerHubUnstakeRequest is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubUnstakeRequest struct {
	Stake    *big.Int
	UnlockAt *big.Int
}

// IWorkerHubWorker is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubWorker struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
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

// WorkerHubMetaData contains all meta data concerning the WorkerHub contract.
var WorkerHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRevealed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"Bytes32Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InferMustBeSolvingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinerInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MiningSessionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MiningSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevealTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillBeingLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatingSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocks\",\"type\":\"uint256\"}],\"name\":\"BlocksPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"CommitDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIWorkerHub.DAOTokenReceiverRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenReceiverInfor[]\",\"name\":\"receivers\",\"type\":\"tuple[]\"}],\"name\":\"DAOTokenMintedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"oldValue\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"DAOTokenPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DAOTokenRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"DAOTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldPercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPercent\",\"type\":\"uint16\"}],\"name\":\"FinePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"FraudulentMinerPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"L2OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinFeeToUseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"}],\"name\":\"MinerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerRoleSeized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"MinerUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newValue\",\"type\":\"uint40\"}],\"name\":\"MiningTimeLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"ModelTierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"}],\"name\":\"NewScoringInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"oldDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newDuration\",\"type\":\"uint40\"}],\"name\":\"PenaltyDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restake\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"RevealDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"nonce\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"name\":\"RevealSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"SubmitDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TopUpInfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"L2OwnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"L2OwnerFee\",\"type\":\"uint256\"}],\"name\":\"TransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TreasuryAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelayTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelayTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blocksPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoTokenPercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoTokenReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeL2Percentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRatioMinerValidator\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTreasuryPercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"forceChangeModelForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getAllAssignments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIWorkerHub.Assignment[]\",\"name\":\"assignmentData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getAllInferences\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"assignments\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeL2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTreasury\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.Inference[]\",\"name\":\"inferenceData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMinerUnstakeRequests\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"unstakeAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.UnstakeRequest[]\",\"name\":\"unstakeRequests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMiners\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.Worker[]\",\"name\":\"minerData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferId\",\"type\":\"uint256\"}],\"name\":\"getAssignmentByInferenceId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIWorkerHub.Assignment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minerAddr\",\"type\":\"address\"}],\"name\":\"getAssignmentByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"assignments\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeL2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTreasury\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"getMinerAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"workerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.WorkerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getMintingAssignmentsOfInference\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNOMiner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"getRoleByAssigmentId\",\"outputs\":[{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseMinerStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originInferId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"inferWithCallback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoToken\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_feeL2Percentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeTreasuryPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minerRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_blocksPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeRatioMinerValidor\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minFeeToUse\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_daoTokenReward\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"_daoTokenPercentage\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"isAssignmentPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isReferrer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minerUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelScoring\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"registerReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_referrers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_referees\",\"type\":\"address[]\"}],\"name\":\"registerReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"resolveInference\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"restakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"resultReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_originInferId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"resultReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_nonce\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaskCompleted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMiner\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"seizeMinerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newCommitDuration\",\"type\":\"uint40\"}],\"name\":\"setCommitDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoTokenAddress\",\"type\":\"address\"}],\"name\":\"setDAOToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"_daoTokenPercentage\",\"type\":\"tuple\"}],\"name\":\"setDAOTokenPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDAOTokenReward\",\"type\":\"uint256\"}],\"name\":\"setDAOTokenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newRatio\",\"type\":\"uint16\"}],\"name\":\"setFeeRatioMinerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"}],\"name\":\"setFinePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2OwnerAddress\",\"type\":\"address\"}],\"name\":\"setL2Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minFeeToUse\",\"type\":\"uint256\"}],\"name\":\"setMinFeeToUse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardAmount\",\"type\":\"uint256\"}],\"name\":\"setNewRewardInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_penaltyDuration\",\"type\":\"uint40\"}],\"name\":\"setPenaltyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newRevealDuration\",\"type\":\"uint40\"}],\"name\":\"setRevealDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_workerHubScoring\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelScoring\",\"type\":\"address\"}],\"name\":\"setScoringInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newSubmitDuration\",\"type\":\"uint40\"}],\"name\":\"setSubmitDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newUnstakeDelayTime\",\"type\":\"uint40\"}],\"name\":\"setUnstakDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isFined\",\"type\":\"bool\"}],\"name\":\"slashMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"streamData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assigmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"topUpInfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeDelayTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"updateModelMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_tier\",\"type\":\"uint32\"}],\"name\":\"updateModelTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inferId\",\"type\":\"uint256\"}],\"name\":\"votingInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"totalCommit\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"totalReveal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workerHubScoring\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) AssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "assignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.AssignmentNumber(&_WorkerHub.CallOpts)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) AssignmentNumber() (*big.Int, error) {
	return _WorkerHub.Contract.AssignmentNumber(&_WorkerHub.CallOpts)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_WorkerHub *WorkerHubCaller) Assignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "assignments", arg0)

	outstruct := new(struct {
		InferenceId *big.Int
		Commitment  [32]byte
		Digest      [32]byte
		RevealNonce *big.Int
		Worker      common.Address
		Role        uint8
		Vote        uint8
		Output      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Digest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.RevealNonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Worker = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Role = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Vote = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Output = *abi.ConvertType(out[7], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_WorkerHub *WorkerHubSession) Assignments(arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	return _WorkerHub.Contract.Assignments(&_WorkerHub.CallOpts, arg0)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_WorkerHub *WorkerHubCallerSession) Assignments(arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	return _WorkerHub.Contract.Assignments(&_WorkerHub.CallOpts, arg0)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubSession) BlocksPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.BlocksPerEpoch(&_WorkerHub.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.BlocksPerEpoch(&_WorkerHub.CallOpts)
}

// CommitDuration is a free data retrieval call binding the contract method 0x6f833811.
//
// Solidity: function commitDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) CommitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "commitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitDuration is a free data retrieval call binding the contract method 0x6f833811.
//
// Solidity: function commitDuration() view returns(uint40)
func (_WorkerHub *WorkerHubSession) CommitDuration() (*big.Int, error) {
	return _WorkerHub.Contract.CommitDuration(&_WorkerHub.CallOpts)
}

// CommitDuration is a free data retrieval call binding the contract method 0x6f833811.
//
// Solidity: function commitDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) CommitDuration() (*big.Int, error) {
	return _WorkerHub.Contract.CommitDuration(&_WorkerHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHub *WorkerHubSession) CurrentEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.CurrentEpoch(&_WorkerHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) CurrentEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.CurrentEpoch(&_WorkerHub.CallOpts)
}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_WorkerHub *WorkerHubCaller) DaoToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "daoToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_WorkerHub *WorkerHubSession) DaoToken() (common.Address, error) {
	return _WorkerHub.Contract.DaoToken(&_WorkerHub.CallOpts)
}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) DaoToken() (common.Address, error) {
	return _WorkerHub.Contract.DaoToken(&_WorkerHub.CallOpts)
}

// DaoTokenPercentage is a free data retrieval call binding the contract method 0xff5db406.
//
// Solidity: function daoTokenPercentage() view returns(uint16 minerPercentage, uint16 userPercentage, uint16 referrerPercentage, uint16 refereePercentage, uint16 l2OwnerPercentage)
func (_WorkerHub *WorkerHubCaller) DaoTokenPercentage(opts *bind.CallOpts) (struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "daoTokenPercentage")

	outstruct := new(struct {
		MinerPercentage    uint16
		UserPercentage     uint16
		ReferrerPercentage uint16
		RefereePercentage  uint16
		L2OwnerPercentage  uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinerPercentage = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.UserPercentage = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.ReferrerPercentage = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.RefereePercentage = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.L2OwnerPercentage = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// DaoTokenPercentage is a free data retrieval call binding the contract method 0xff5db406.
//
// Solidity: function daoTokenPercentage() view returns(uint16 minerPercentage, uint16 userPercentage, uint16 referrerPercentage, uint16 refereePercentage, uint16 l2OwnerPercentage)
func (_WorkerHub *WorkerHubSession) DaoTokenPercentage() (struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}, error) {
	return _WorkerHub.Contract.DaoTokenPercentage(&_WorkerHub.CallOpts)
}

// DaoTokenPercentage is a free data retrieval call binding the contract method 0xff5db406.
//
// Solidity: function daoTokenPercentage() view returns(uint16 minerPercentage, uint16 userPercentage, uint16 referrerPercentage, uint16 refereePercentage, uint16 l2OwnerPercentage)
func (_WorkerHub *WorkerHubCallerSession) DaoTokenPercentage() (struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}, error) {
	return _WorkerHub.Contract.DaoTokenPercentage(&_WorkerHub.CallOpts)
}

// DaoTokenReward is a free data retrieval call binding the contract method 0x0940c392.
//
// Solidity: function daoTokenReward() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) DaoTokenReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "daoTokenReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaoTokenReward is a free data retrieval call binding the contract method 0x0940c392.
//
// Solidity: function daoTokenReward() view returns(uint256)
func (_WorkerHub *WorkerHubSession) DaoTokenReward() (*big.Int, error) {
	return _WorkerHub.Contract.DaoTokenReward(&_WorkerHub.CallOpts)
}

// DaoTokenReward is a free data retrieval call binding the contract method 0x0940c392.
//
// Solidity: function daoTokenReward() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) DaoTokenReward() (*big.Int, error) {
	return _WorkerHub.Contract.DaoTokenReward(&_WorkerHub.CallOpts)
}

// FeeL2Percentage is a free data retrieval call binding the contract method 0x39d2e296.
//
// Solidity: function feeL2Percentage() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) FeeL2Percentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "feeL2Percentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeL2Percentage is a free data retrieval call binding the contract method 0x39d2e296.
//
// Solidity: function feeL2Percentage() view returns(uint16)
func (_WorkerHub *WorkerHubSession) FeeL2Percentage() (uint16, error) {
	return _WorkerHub.Contract.FeeL2Percentage(&_WorkerHub.CallOpts)
}

// FeeL2Percentage is a free data retrieval call binding the contract method 0x39d2e296.
//
// Solidity: function feeL2Percentage() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) FeeL2Percentage() (uint16, error) {
	return _WorkerHub.Contract.FeeL2Percentage(&_WorkerHub.CallOpts)
}

// FeeRatioMinerValidator is a free data retrieval call binding the contract method 0x50eac7c8.
//
// Solidity: function feeRatioMinerValidator() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) FeeRatioMinerValidator(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "feeRatioMinerValidator")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeRatioMinerValidator is a free data retrieval call binding the contract method 0x50eac7c8.
//
// Solidity: function feeRatioMinerValidator() view returns(uint16)
func (_WorkerHub *WorkerHubSession) FeeRatioMinerValidator() (uint16, error) {
	return _WorkerHub.Contract.FeeRatioMinerValidator(&_WorkerHub.CallOpts)
}

// FeeRatioMinerValidator is a free data retrieval call binding the contract method 0x50eac7c8.
//
// Solidity: function feeRatioMinerValidator() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) FeeRatioMinerValidator() (uint16, error) {
	return _WorkerHub.Contract.FeeRatioMinerValidator(&_WorkerHub.CallOpts)
}

// FeeTreasuryPercentage is a free data retrieval call binding the contract method 0x09c83b4f.
//
// Solidity: function feeTreasuryPercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) FeeTreasuryPercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "feeTreasuryPercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeTreasuryPercentage is a free data retrieval call binding the contract method 0x09c83b4f.
//
// Solidity: function feeTreasuryPercentage() view returns(uint16)
func (_WorkerHub *WorkerHubSession) FeeTreasuryPercentage() (uint16, error) {
	return _WorkerHub.Contract.FeeTreasuryPercentage(&_WorkerHub.CallOpts)
}

// FeeTreasuryPercentage is a free data retrieval call binding the contract method 0x09c83b4f.
//
// Solidity: function feeTreasuryPercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) FeeTreasuryPercentage() (uint16, error) {
	return _WorkerHub.Contract.FeeTreasuryPercentage(&_WorkerHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubSession) FinePercentage() (uint16, error) {
	return _WorkerHub.Contract.FinePercentage(&_WorkerHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_WorkerHub *WorkerHubCallerSession) FinePercentage() (uint16, error) {
	return _WorkerHub.Contract.FinePercentage(&_WorkerHub.CallOpts)
}

// GetAllAssignments is a free data retrieval call binding the contract method 0x16d0a88f.
//
// Solidity: function getAllAssignments(uint256 startId, uint256 count) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[] assignmentData)
func (_WorkerHub *WorkerHubCaller) GetAllAssignments(opts *bind.CallOpts, startId *big.Int, count *big.Int) ([]IWorkerHubAssignment, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAllAssignments", startId, count)

	if err != nil {
		return *new([]IWorkerHubAssignment), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignment)).(*[]IWorkerHubAssignment)

	return out0, err

}

// GetAllAssignments is a free data retrieval call binding the contract method 0x16d0a88f.
//
// Solidity: function getAllAssignments(uint256 startId, uint256 count) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[] assignmentData)
func (_WorkerHub *WorkerHubSession) GetAllAssignments(startId *big.Int, count *big.Int) ([]IWorkerHubAssignment, error) {
	return _WorkerHub.Contract.GetAllAssignments(&_WorkerHub.CallOpts, startId, count)
}

// GetAllAssignments is a free data retrieval call binding the contract method 0x16d0a88f.
//
// Solidity: function getAllAssignments(uint256 startId, uint256 count) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[] assignmentData)
func (_WorkerHub *WorkerHubCallerSession) GetAllAssignments(startId *big.Int, count *big.Int) ([]IWorkerHubAssignment, error) {
	return _WorkerHub.Contract.GetAllAssignments(&_WorkerHub.CallOpts, startId, count)
}

// GetAllInferences is a free data retrieval call binding the contract method 0xf1ea45e3.
//
// Solidity: function getAllInferences(uint256 startId, uint256 count) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address)[] inferenceData)
func (_WorkerHub *WorkerHubCaller) GetAllInferences(opts *bind.CallOpts, startId *big.Int, count *big.Int) ([]IWorkerHubInference, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAllInferences", startId, count)

	if err != nil {
		return *new([]IWorkerHubInference), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubInference)).(*[]IWorkerHubInference)

	return out0, err

}

// GetAllInferences is a free data retrieval call binding the contract method 0xf1ea45e3.
//
// Solidity: function getAllInferences(uint256 startId, uint256 count) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address)[] inferenceData)
func (_WorkerHub *WorkerHubSession) GetAllInferences(startId *big.Int, count *big.Int) ([]IWorkerHubInference, error) {
	return _WorkerHub.Contract.GetAllInferences(&_WorkerHub.CallOpts, startId, count)
}

// GetAllInferences is a free data retrieval call binding the contract method 0xf1ea45e3.
//
// Solidity: function getAllInferences(uint256 startId, uint256 count) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address)[] inferenceData)
func (_WorkerHub *WorkerHubCallerSession) GetAllInferences(startId *big.Int, count *big.Int) ([]IWorkerHubInference, error) {
	return _WorkerHub.Contract.GetAllInferences(&_WorkerHub.CallOpts, startId, count)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_WorkerHub *WorkerHubCaller) GetAllMinerUnstakeRequests(opts *bind.CallOpts) (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IWorkerHubUnstakeRequest
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAllMinerUnstakeRequests")

	outstruct := new(struct {
		UnstakeAddresses []common.Address
		UnstakeRequests  []IWorkerHubUnstakeRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnstakeAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UnstakeRequests = *abi.ConvertType(out[1], new([]IWorkerHubUnstakeRequest)).(*[]IWorkerHubUnstakeRequest)

	return *outstruct, err

}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_WorkerHub *WorkerHubSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IWorkerHubUnstakeRequest
}, error) {
	return _WorkerHub.Contract.GetAllMinerUnstakeRequests(&_WorkerHub.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_WorkerHub *WorkerHubCallerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IWorkerHubUnstakeRequest
}, error) {
	return _WorkerHub.Contract.GetAllMinerUnstakeRequests(&_WorkerHub.CallOpts)
}

// GetAllMiners is a free data retrieval call binding the contract method 0x4b17bf30.
//
// Solidity: function getAllMiners() view returns((uint256,uint256,address,uint40,uint40,uint16)[] minerData)
func (_WorkerHub *WorkerHubCaller) GetAllMiners(opts *bind.CallOpts) ([]IWorkerHubWorker, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAllMiners")

	if err != nil {
		return *new([]IWorkerHubWorker), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorker)).(*[]IWorkerHubWorker)

	return out0, err

}

// GetAllMiners is a free data retrieval call binding the contract method 0x4b17bf30.
//
// Solidity: function getAllMiners() view returns((uint256,uint256,address,uint40,uint40,uint16)[] minerData)
func (_WorkerHub *WorkerHubSession) GetAllMiners() ([]IWorkerHubWorker, error) {
	return _WorkerHub.Contract.GetAllMiners(&_WorkerHub.CallOpts)
}

// GetAllMiners is a free data retrieval call binding the contract method 0x4b17bf30.
//
// Solidity: function getAllMiners() view returns((uint256,uint256,address,uint40,uint40,uint16)[] minerData)
func (_WorkerHub *WorkerHubCallerSession) GetAllMiners() ([]IWorkerHubWorker, error) {
	return _WorkerHub.Contract.GetAllMiners(&_WorkerHub.CallOpts)
}

// GetAssignmentByInferenceId is a free data retrieval call binding the contract method 0x19a9dc71.
//
// Solidity: function getAssignmentByInferenceId(uint256 _inferId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[])
func (_WorkerHub *WorkerHubCaller) GetAssignmentByInferenceId(opts *bind.CallOpts, _inferId *big.Int) ([]IWorkerHubAssignment, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAssignmentByInferenceId", _inferId)

	if err != nil {
		return *new([]IWorkerHubAssignment), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignment)).(*[]IWorkerHubAssignment)

	return out0, err

}

// GetAssignmentByInferenceId is a free data retrieval call binding the contract method 0x19a9dc71.
//
// Solidity: function getAssignmentByInferenceId(uint256 _inferId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[])
func (_WorkerHub *WorkerHubSession) GetAssignmentByInferenceId(_inferId *big.Int) ([]IWorkerHubAssignment, error) {
	return _WorkerHub.Contract.GetAssignmentByInferenceId(&_WorkerHub.CallOpts, _inferId)
}

// GetAssignmentByInferenceId is a free data retrieval call binding the contract method 0x19a9dc71.
//
// Solidity: function getAssignmentByInferenceId(uint256 _inferId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[])
func (_WorkerHub *WorkerHubCallerSession) GetAssignmentByInferenceId(_inferId *big.Int) ([]IWorkerHubAssignment, error) {
	return _WorkerHub.Contract.GetAssignmentByInferenceId(&_WorkerHub.CallOpts, _inferId)
}

// GetAssignmentByMiner is a free data retrieval call binding the contract method 0x5937e5ed.
//
// Solidity: function getAssignmentByMiner(address _minerAddr) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_WorkerHub *WorkerHubCaller) GetAssignmentByMiner(opts *bind.CallOpts, _minerAddr common.Address) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getAssignmentByMiner", _minerAddr)

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetAssignmentByMiner is a free data retrieval call binding the contract method 0x5937e5ed.
//
// Solidity: function getAssignmentByMiner(address _minerAddr) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_WorkerHub *WorkerHubSession) GetAssignmentByMiner(_minerAddr common.Address) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetAssignmentByMiner(&_WorkerHub.CallOpts, _minerAddr)
}

// GetAssignmentByMiner is a free data retrieval call binding the contract method 0x5937e5ed.
//
// Solidity: function getAssignmentByMiner(address _minerAddr) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_WorkerHub *WorkerHubCallerSession) GetAssignmentByMiner(_minerAddr common.Address) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetAssignmentByMiner(&_WorkerHub.CallOpts, _minerAddr)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_WorkerHub *WorkerHubCaller) GetInferenceInfo(opts *bind.CallOpts, _inferenceId *big.Int) (IWorkerHubInference, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getInferenceInfo", _inferenceId)

	if err != nil {
		return *new(IWorkerHubInference), err
	}

	out0 := *abi.ConvertType(out[0], new(IWorkerHubInference)).(*IWorkerHubInference)

	return out0, err

}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_WorkerHub *WorkerHubSession) GetInferenceInfo(_inferenceId *big.Int) (IWorkerHubInference, error) {
	return _WorkerHub.Contract.GetInferenceInfo(&_WorkerHub.CallOpts, _inferenceId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_WorkerHub *WorkerHubCallerSession) GetInferenceInfo(_inferenceId *big.Int) (IWorkerHubInference, error) {
	return _WorkerHub.Contract.GetInferenceInfo(&_WorkerHub.CallOpts, _inferenceId)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_WorkerHub *WorkerHubCaller) GetMinFeeToUse(opts *bind.CallOpts, _modelAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinFeeToUse", _modelAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_WorkerHub *WorkerHubSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.GetMinFeeToUse(&_WorkerHub.CallOpts, _modelAddress)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.GetMinFeeToUse(&_WorkerHub.CallOpts, _modelAddress)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHub *WorkerHubSession) GetMinerAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddresses(&_WorkerHub.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddresses(&_WorkerHub.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMinerAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddressesOfModel(&_WorkerHub.CallOpts, _model)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_WorkerHub *WorkerHubCallerSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _WorkerHub.Contract.GetMinerAddressesOfModel(&_WorkerHub.CallOpts, _model)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubCaller) GetMiners(opts *bind.CallOpts) ([]IWorkerHubWorkerInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMiners")

	if err != nil {
		return *new([]IWorkerHubWorkerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorkerInfo)).(*[]IWorkerHubWorkerInfo)

	return out0, err

}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHub.Contract.GetMiners(&_WorkerHub.CallOpts)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_WorkerHub *WorkerHubCallerSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _WorkerHub.Contract.GetMiners(&_WorkerHub.CallOpts)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_WorkerHub *WorkerHubCaller) GetMintingAssignmentsOfInference(opts *bind.CallOpts, _inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getMintingAssignmentsOfInference", _inferenceId)

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_WorkerHub *WorkerHubSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetMintingAssignmentsOfInference(&_WorkerHub.CallOpts, _inferenceId)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_WorkerHub *WorkerHubCallerSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _WorkerHub.Contract.GetMintingAssignmentsOfInference(&_WorkerHub.CallOpts, _inferenceId)
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

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHub *WorkerHubSession) GetNOMiner() (*big.Int, error) {
	return _WorkerHub.Contract.GetNOMiner(&_WorkerHub.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) GetNOMiner() (*big.Int, error) {
	return _WorkerHub.Contract.GetNOMiner(&_WorkerHub.CallOpts)
}

// GetRoleByAssigmentId is a free data retrieval call binding the contract method 0xca0c80fc.
//
// Solidity: function getRoleByAssigmentId(uint256 _assignmentId) view returns(uint8)
func (_WorkerHub *WorkerHubCaller) GetRoleByAssigmentId(opts *bind.CallOpts, _assignmentId *big.Int) (uint8, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "getRoleByAssigmentId", _assignmentId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRoleByAssigmentId is a free data retrieval call binding the contract method 0xca0c80fc.
//
// Solidity: function getRoleByAssigmentId(uint256 _assignmentId) view returns(uint8)
func (_WorkerHub *WorkerHubSession) GetRoleByAssigmentId(_assignmentId *big.Int) (uint8, error) {
	return _WorkerHub.Contract.GetRoleByAssigmentId(&_WorkerHub.CallOpts, _assignmentId)
}

// GetRoleByAssigmentId is a free data retrieval call binding the contract method 0xca0c80fc.
//
// Solidity: function getRoleByAssigmentId(uint256 _assignmentId) view returns(uint8)
func (_WorkerHub *WorkerHubCallerSession) GetRoleByAssigmentId(_assignmentId *big.Int) (uint8, error) {
	return _WorkerHub.Contract.GetRoleByAssigmentId(&_WorkerHub.CallOpts, _assignmentId)
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

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHub *WorkerHubCaller) IsAssignmentPending(opts *bind.CallOpts, _assignmentId *big.Int) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "isAssignmentPending", _assignmentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHub *WorkerHubSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _WorkerHub.Contract.IsAssignmentPending(&_WorkerHub.CallOpts, _assignmentId)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _WorkerHub.Contract.IsAssignmentPending(&_WorkerHub.CallOpts, _assignmentId)
}

// IsReferrer is a free data retrieval call binding the contract method 0xd64d6968.
//
// Solidity: function isReferrer(address ) view returns(bool)
func (_WorkerHub *WorkerHubCaller) IsReferrer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "isReferrer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferrer is a free data retrieval call binding the contract method 0xd64d6968.
//
// Solidity: function isReferrer(address ) view returns(bool)
func (_WorkerHub *WorkerHubSession) IsReferrer(arg0 common.Address) (bool, error) {
	return _WorkerHub.Contract.IsReferrer(&_WorkerHub.CallOpts, arg0)
}

// IsReferrer is a free data retrieval call binding the contract method 0xd64d6968.
//
// Solidity: function isReferrer(address ) view returns(bool)
func (_WorkerHub *WorkerHubCallerSession) IsReferrer(arg0 common.Address) (bool, error) {
	return _WorkerHub.Contract.IsReferrer(&_WorkerHub.CallOpts, arg0)
}

// L2Owner is a free data retrieval call binding the contract method 0xf003a0c5.
//
// Solidity: function l2Owner() view returns(address)
func (_WorkerHub *WorkerHubCaller) L2Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "l2Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Owner is a free data retrieval call binding the contract method 0xf003a0c5.
//
// Solidity: function l2Owner() view returns(address)
func (_WorkerHub *WorkerHubSession) L2Owner() (common.Address, error) {
	return _WorkerHub.Contract.L2Owner(&_WorkerHub.CallOpts)
}

// L2Owner is a free data retrieval call binding the contract method 0xf003a0c5.
//
// Solidity: function l2Owner() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) L2Owner() (common.Address, error) {
	return _WorkerHub.Contract.L2Owner(&_WorkerHub.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHub *WorkerHubSession) LastBlock() (*big.Int, error) {
	return _WorkerHub.Contract.LastBlock(&_WorkerHub.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) LastBlock() (*big.Int, error) {
	return _WorkerHub.Contract.LastBlock(&_WorkerHub.CallOpts)
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

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) MinFeeToUse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minFeeToUse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_WorkerHub *WorkerHubSession) MinFeeToUse() (*big.Int, error) {
	return _WorkerHub.Contract.MinFeeToUse(&_WorkerHub.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) MinFeeToUse() (*big.Int, error) {
	return _WorkerHub.Contract.MinFeeToUse(&_WorkerHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubSession) MinerMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.MinerMinimumStake(&_WorkerHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _WorkerHub.Contract.MinerMinimumStake(&_WorkerHub.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubSession) MinerRequirement() (uint8, error) {
	return _WorkerHub.Contract.MinerRequirement(&_WorkerHub.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_WorkerHub *WorkerHubCallerSession) MinerRequirement() (uint8, error) {
	return _WorkerHub.Contract.MinerRequirement(&_WorkerHub.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "minerUnstakeRequests", arg0)

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
func (_WorkerHub *WorkerHubSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHub.Contract.MinerUnstakeRequests(&_WorkerHub.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_WorkerHub *WorkerHubCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _WorkerHub.Contract.MinerUnstakeRequests(&_WorkerHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "miners", arg0)

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
func (_WorkerHub *WorkerHubSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHub.Contract.Miners(&_WorkerHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_WorkerHub *WorkerHubCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _WorkerHub.Contract.Miners(&_WorkerHub.CallOpts, arg0)
}

// ModelScoring is a free data retrieval call binding the contract method 0xfe0503c0.
//
// Solidity: function modelScoring() view returns(address)
func (_WorkerHub *WorkerHubCaller) ModelScoring(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "modelScoring")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelScoring is a free data retrieval call binding the contract method 0xfe0503c0.
//
// Solidity: function modelScoring() view returns(address)
func (_WorkerHub *WorkerHubSession) ModelScoring() (common.Address, error) {
	return _WorkerHub.Contract.ModelScoring(&_WorkerHub.CallOpts)
}

// ModelScoring is a free data retrieval call binding the contract method 0xfe0503c0.
//
// Solidity: function modelScoring() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) ModelScoring() (common.Address, error) {
	return _WorkerHub.Contract.ModelScoring(&_WorkerHub.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubCaller) Models(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "models", arg0)

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
func (_WorkerHub *WorkerHubSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHub.Contract.Models(&_WorkerHub.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_WorkerHub *WorkerHubCallerSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _WorkerHub.Contract.Models(&_WorkerHub.CallOpts, arg0)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHub *WorkerHubCaller) Multiplier(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "multiplier", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHub *WorkerHubSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.Multiplier(&_WorkerHub.CallOpts, _miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _WorkerHub.Contract.Multiplier(&_WorkerHub.CallOpts, _miner)
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

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHub *WorkerHubSession) PenaltyDuration() (*big.Int, error) {
	return _WorkerHub.Contract.PenaltyDuration(&_WorkerHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) PenaltyDuration() (*big.Int, error) {
	return _WorkerHub.Contract.PenaltyDuration(&_WorkerHub.CallOpts)
}

// RevealDuration is a free data retrieval call binding the contract method 0x886a6de1.
//
// Solidity: function revealDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) RevealDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "revealDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealDuration is a free data retrieval call binding the contract method 0x886a6de1.
//
// Solidity: function revealDuration() view returns(uint40)
func (_WorkerHub *WorkerHubSession) RevealDuration() (*big.Int, error) {
	return _WorkerHub.Contract.RevealDuration(&_WorkerHub.CallOpts)
}

// RevealDuration is a free data retrieval call binding the contract method 0x886a6de1.
//
// Solidity: function revealDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) RevealDuration() (*big.Int, error) {
	return _WorkerHub.Contract.RevealDuration(&_WorkerHub.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHub *WorkerHubCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "rewardInEpoch", arg0)

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
func (_WorkerHub *WorkerHubSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _WorkerHub.Contract.RewardInEpoch(&_WorkerHub.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_WorkerHub *WorkerHubCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _WorkerHub.Contract.RewardInEpoch(&_WorkerHub.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubSession) RewardPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.RewardPerEpoch(&_WorkerHub.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_WorkerHub *WorkerHubCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _WorkerHub.Contract.RewardPerEpoch(&_WorkerHub.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0xcc56b6f8.
//
// Solidity: function submitDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) SubmitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "submitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitDuration is a free data retrieval call binding the contract method 0xcc56b6f8.
//
// Solidity: function submitDuration() view returns(uint40)
func (_WorkerHub *WorkerHubSession) SubmitDuration() (*big.Int, error) {
	return _WorkerHub.Contract.SubmitDuration(&_WorkerHub.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0xcc56b6f8.
//
// Solidity: function submitDuration() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) SubmitDuration() (*big.Int, error) {
	return _WorkerHub.Contract.SubmitDuration(&_WorkerHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHub *WorkerHubCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHub *WorkerHubSession) Treasury() (common.Address, error) {
	return _WorkerHub.Contract.Treasury(&_WorkerHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) Treasury() (common.Address, error) {
	return _WorkerHub.Contract.Treasury(&_WorkerHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHub *WorkerHubCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHub *WorkerHubSession) UnstakeDelayTime() (*big.Int, error) {
	return _WorkerHub.Contract.UnstakeDelayTime(&_WorkerHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_WorkerHub *WorkerHubCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _WorkerHub.Contract.UnstakeDelayTime(&_WorkerHub.CallOpts)
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

// VotingInfo is a free data retrieval call binding the contract method 0xe5309a66.
//
// Solidity: function votingInfo(uint256 inferId) view returns(uint8 totalCommit, uint8 totalReveal)
func (_WorkerHub *WorkerHubCaller) VotingInfo(opts *bind.CallOpts, inferId *big.Int) (struct {
	TotalCommit uint8
	TotalReveal uint8
}, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "votingInfo", inferId)

	outstruct := new(struct {
		TotalCommit uint8
		TotalReveal uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCommit = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.TotalReveal = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// VotingInfo is a free data retrieval call binding the contract method 0xe5309a66.
//
// Solidity: function votingInfo(uint256 inferId) view returns(uint8 totalCommit, uint8 totalReveal)
func (_WorkerHub *WorkerHubSession) VotingInfo(inferId *big.Int) (struct {
	TotalCommit uint8
	TotalReveal uint8
}, error) {
	return _WorkerHub.Contract.VotingInfo(&_WorkerHub.CallOpts, inferId)
}

// VotingInfo is a free data retrieval call binding the contract method 0xe5309a66.
//
// Solidity: function votingInfo(uint256 inferId) view returns(uint8 totalCommit, uint8 totalReveal)
func (_WorkerHub *WorkerHubCallerSession) VotingInfo(inferId *big.Int) (struct {
	TotalCommit uint8
	TotalReveal uint8
}, error) {
	return _WorkerHub.Contract.VotingInfo(&_WorkerHub.CallOpts, inferId)
}

// WorkerHubScoring is a free data retrieval call binding the contract method 0x2b426301.
//
// Solidity: function workerHubScoring() view returns(address)
func (_WorkerHub *WorkerHubCaller) WorkerHubScoring(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkerHub.contract.Call(opts, &out, "workerHubScoring")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkerHubScoring is a free data retrieval call binding the contract method 0x2b426301.
//
// Solidity: function workerHubScoring() view returns(address)
func (_WorkerHub *WorkerHubSession) WorkerHubScoring() (common.Address, error) {
	return _WorkerHub.Contract.WorkerHubScoring(&_WorkerHub.CallOpts)
}

// WorkerHubScoring is a free data retrieval call binding the contract method 0x2b426301.
//
// Solidity: function workerHubScoring() view returns(address)
func (_WorkerHub *WorkerHubCallerSession) WorkerHubScoring() (common.Address, error) {
	return _WorkerHub.Contract.WorkerHubScoring(&_WorkerHub.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHub *WorkerHubTransactor) ClaimReward(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "claimReward", _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHub *WorkerHubSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.ClaimReward(&_WorkerHub.TransactOpts, _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_WorkerHub *WorkerHubTransactorSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.ClaimReward(&_WorkerHub.TransactOpts, _miner)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_WorkerHub *WorkerHubTransactor) Commit(opts *bind.TransactOpts, _assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "commit", _assignId, _commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_WorkerHub *WorkerHubSession) Commit(_assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Commit(&_WorkerHub.TransactOpts, _assignId, _commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_WorkerHub *WorkerHubTransactorSession) Commit(_assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Commit(&_WorkerHub.TransactOpts, _assignId, _commitment)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_WorkerHub *WorkerHubTransactor) ForceChangeModelForMiner(opts *bind.TransactOpts, _miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "forceChangeModelForMiner", _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_WorkerHub *WorkerHubSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.ForceChangeModelForMiner(&_WorkerHub.TransactOpts, _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_WorkerHub *WorkerHubTransactorSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.ForceChangeModelForMiner(&_WorkerHub.TransactOpts, _miner, _modelAddress)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0x9fb40bac.
//
// Solidity: function increaseMinerStake() payable returns()
func (_WorkerHub *WorkerHubTransactor) IncreaseMinerStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "increaseMinerStake")
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0x9fb40bac.
//
// Solidity: function increaseMinerStake() payable returns()
func (_WorkerHub *WorkerHubSession) IncreaseMinerStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseMinerStake(&_WorkerHub.TransactOpts)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0x9fb40bac.
//
// Solidity: function increaseMinerStake() payable returns()
func (_WorkerHub *WorkerHubTransactorSession) IncreaseMinerStake() (*types.Transaction, error) {
	return _WorkerHub.Contract.IncreaseMinerStake(&_WorkerHub.TransactOpts)
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

// InferWithCallback is a paid mutator transaction binding the contract method 0xb8cfec3d.
//
// Solidity: function inferWithCallback(uint256 originInferId, bytes _input, address _creator, address callback) payable returns(uint256 inferenceId)
func (_WorkerHub *WorkerHubTransactor) InferWithCallback(opts *bind.TransactOpts, originInferId *big.Int, _input []byte, _creator common.Address, callback common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "inferWithCallback", originInferId, _input, _creator, callback)
}

// InferWithCallback is a paid mutator transaction binding the contract method 0xb8cfec3d.
//
// Solidity: function inferWithCallback(uint256 originInferId, bytes _input, address _creator, address callback) payable returns(uint256 inferenceId)
func (_WorkerHub *WorkerHubSession) InferWithCallback(originInferId *big.Int, _input []byte, _creator common.Address, callback common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.InferWithCallback(&_WorkerHub.TransactOpts, originInferId, _input, _creator, callback)
}

// InferWithCallback is a paid mutator transaction binding the contract method 0xb8cfec3d.
//
// Solidity: function inferWithCallback(uint256 originInferId, bytes _input, address _creator, address callback) payable returns(uint256 inferenceId)
func (_WorkerHub *WorkerHubTransactorSession) InferWithCallback(originInferId *big.Int, _input []byte, _creator common.Address, callback common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.InferWithCallback(&_WorkerHub.TransactOpts, originInferId, _input, _creator, callback)
}

// Initialize is a paid mutator transaction binding the contract method 0xe2f32c82.
//
// Solidity: function initialize(address _l2Owner, address _treasury, address _daoToken, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint256 _minerMinimumStake, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint256 _duration, uint16 _finePercentage, uint16 _feeRatioMinerValidor, uint256 _minFeeToUse, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubTransactor) Initialize(opts *bind.TransactOpts, _l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerMinimumStake *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _duration *big.Int, _finePercentage uint16, _feeRatioMinerValidor uint16, _minFeeToUse *big.Int, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "initialize", _l2Owner, _treasury, _daoToken, _feeL2Percentage, _feeTreasuryPercentage, _minerMinimumStake, _minerRequirement, _blocksPerEpoch, _rewardPerEpoch, _duration, _finePercentage, _feeRatioMinerValidor, _minFeeToUse, _daoTokenReward, _daoTokenPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xe2f32c82.
//
// Solidity: function initialize(address _l2Owner, address _treasury, address _daoToken, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint256 _minerMinimumStake, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint256 _duration, uint16 _finePercentage, uint16 _feeRatioMinerValidor, uint256 _minFeeToUse, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubSession) Initialize(_l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerMinimumStake *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _duration *big.Int, _finePercentage uint16, _feeRatioMinerValidor uint16, _minFeeToUse *big.Int, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _l2Owner, _treasury, _daoToken, _feeL2Percentage, _feeTreasuryPercentage, _minerMinimumStake, _minerRequirement, _blocksPerEpoch, _rewardPerEpoch, _duration, _finePercentage, _feeRatioMinerValidor, _minFeeToUse, _daoTokenReward, _daoTokenPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xe2f32c82.
//
// Solidity: function initialize(address _l2Owner, address _treasury, address _daoToken, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint256 _minerMinimumStake, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint256 _duration, uint16 _finePercentage, uint16 _feeRatioMinerValidor, uint256 _minFeeToUse, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubTransactorSession) Initialize(_l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerMinimumStake *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _duration *big.Int, _finePercentage uint16, _feeRatioMinerValidor uint16, _minFeeToUse *big.Int, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.Contract.Initialize(&_WorkerHub.TransactOpts, _l2Owner, _treasury, _daoToken, _feeL2Percentage, _feeTreasuryPercentage, _minerMinimumStake, _minerRequirement, _blocksPerEpoch, _rewardPerEpoch, _duration, _finePercentage, _feeRatioMinerValidor, _minFeeToUse, _daoTokenReward, _daoTokenPercentage)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHub *WorkerHubTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHub *WorkerHubSession) JoinForMinting() (*types.Transaction, error) {
	return _WorkerHub.Contract.JoinForMinting(&_WorkerHub.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_WorkerHub *WorkerHubTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _WorkerHub.Contract.JoinForMinting(&_WorkerHub.TransactOpts)
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

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerMiner", tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterMiner(&_WorkerHub.TransactOpts, tier)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x1fdadcb7.
//
// Solidity: function registerMiner(uint16 tier) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterMiner(&_WorkerHub.TransactOpts, tier)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactor) RegisterModel(opts *bind.TransactOpts, _model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerModel", _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterModel(&_WorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterModel(&_WorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0x9ea7685a.
//
// Solidity: function registerReferrer(address _referrer) returns()
func (_WorkerHub *WorkerHubTransactor) RegisterReferrer(opts *bind.TransactOpts, _referrer common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerReferrer", _referrer)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0x9ea7685a.
//
// Solidity: function registerReferrer(address _referrer) returns()
func (_WorkerHub *WorkerHubSession) RegisterReferrer(_referrer common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterReferrer(&_WorkerHub.TransactOpts, _referrer)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0x9ea7685a.
//
// Solidity: function registerReferrer(address _referrer) returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterReferrer(_referrer common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterReferrer(&_WorkerHub.TransactOpts, _referrer)
}

// RegisterReferrer0 is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_WorkerHub *WorkerHubTransactor) RegisterReferrer0(opts *bind.TransactOpts, _referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "registerReferrer0", _referrers, _referees)
}

// RegisterReferrer0 is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_WorkerHub *WorkerHubSession) RegisterReferrer0(_referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterReferrer0(&_WorkerHub.TransactOpts, _referrers, _referees)
}

// RegisterReferrer0 is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_WorkerHub *WorkerHubTransactorSession) RegisterReferrer0(_referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RegisterReferrer0(&_WorkerHub.TransactOpts, _referrers, _referees)
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

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubTransactor) ResolveInference(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "resolveInference", _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResolveInference(&_WorkerHub.TransactOpts, _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_WorkerHub *WorkerHubTransactorSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResolveInference(&_WorkerHub.TransactOpts, _inferenceId)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHub *WorkerHubTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHub *WorkerHubSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RestakeForMiner(&_WorkerHub.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_WorkerHub *WorkerHubTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.RestakeForMiner(&_WorkerHub.TransactOpts, tier)
}

// ResultReceived is a paid mutator transaction binding the contract method 0xc3477018.
//
// Solidity: function resultReceived(bytes result) returns()
func (_WorkerHub *WorkerHubTransactor) ResultReceived(opts *bind.TransactOpts, result []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "resultReceived", result)
}

// ResultReceived is a paid mutator transaction binding the contract method 0xc3477018.
//
// Solidity: function resultReceived(bytes result) returns()
func (_WorkerHub *WorkerHubSession) ResultReceived(result []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResultReceived(&_WorkerHub.TransactOpts, result)
}

// ResultReceived is a paid mutator transaction binding the contract method 0xc3477018.
//
// Solidity: function resultReceived(bytes result) returns()
func (_WorkerHub *WorkerHubTransactorSession) ResultReceived(result []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResultReceived(&_WorkerHub.TransactOpts, result)
}

// ResultReceived0 is a paid mutator transaction binding the contract method 0xd2a554e7.
//
// Solidity: function resultReceived(uint256 _originInferId, bytes _result) returns()
func (_WorkerHub *WorkerHubTransactor) ResultReceived0(opts *bind.TransactOpts, _originInferId *big.Int, _result []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "resultReceived0", _originInferId, _result)
}

// ResultReceived0 is a paid mutator transaction binding the contract method 0xd2a554e7.
//
// Solidity: function resultReceived(uint256 _originInferId, bytes _result) returns()
func (_WorkerHub *WorkerHubSession) ResultReceived0(_originInferId *big.Int, _result []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResultReceived0(&_WorkerHub.TransactOpts, _originInferId, _result)
}

// ResultReceived0 is a paid mutator transaction binding the contract method 0xd2a554e7.
//
// Solidity: function resultReceived(uint256 _originInferId, bytes _result) returns()
func (_WorkerHub *WorkerHubTransactorSession) ResultReceived0(_originInferId *big.Int, _result []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.ResultReceived0(&_WorkerHub.TransactOpts, _originInferId, _result)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactor) Reveal(opts *bind.TransactOpts, _assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "reveal", _assignId, _nonce, _data)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_WorkerHub *WorkerHubSession) Reveal(_assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Reveal(&_WorkerHub.TransactOpts, _assignId, _nonce, _data)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactorSession) Reveal(_assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.Reveal(&_WorkerHub.TransactOpts, _assignId, _nonce, _data)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHub *WorkerHubTransactor) RewardToClaim(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "rewardToClaim", _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHub *WorkerHubSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RewardToClaim(&_WorkerHub.TransactOpts, _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_WorkerHub *WorkerHubTransactorSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.RewardToClaim(&_WorkerHub.TransactOpts, _miner)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubTransactor) SeizeMinerRole(opts *bind.TransactOpts, _assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "seizeMinerRole", _assignmentId)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubSession) SeizeMinerRole(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SeizeMinerRole(&_WorkerHub.TransactOpts, _assignmentId)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_WorkerHub *WorkerHubTransactorSession) SeizeMinerRole(_assignmentId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SeizeMinerRole(&_WorkerHub.TransactOpts, _assignmentId)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHub *WorkerHubTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setBlocksPerEpoch", _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHub *WorkerHubSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetBlocksPerEpoch(&_WorkerHub.TransactOpts, _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetBlocksPerEpoch(&_WorkerHub.TransactOpts, _blocks)
}

// SetCommitDuration is a paid mutator transaction binding the contract method 0x54b18651.
//
// Solidity: function setCommitDuration(uint40 _newCommitDuration) returns()
func (_WorkerHub *WorkerHubTransactor) SetCommitDuration(opts *bind.TransactOpts, _newCommitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setCommitDuration", _newCommitDuration)
}

// SetCommitDuration is a paid mutator transaction binding the contract method 0x54b18651.
//
// Solidity: function setCommitDuration(uint40 _newCommitDuration) returns()
func (_WorkerHub *WorkerHubSession) SetCommitDuration(_newCommitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetCommitDuration(&_WorkerHub.TransactOpts, _newCommitDuration)
}

// SetCommitDuration is a paid mutator transaction binding the contract method 0x54b18651.
//
// Solidity: function setCommitDuration(uint40 _newCommitDuration) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetCommitDuration(_newCommitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetCommitDuration(&_WorkerHub.TransactOpts, _newCommitDuration)
}

// SetDAOToken is a paid mutator transaction binding the contract method 0x70a52354.
//
// Solidity: function setDAOToken(address _daoTokenAddress) returns()
func (_WorkerHub *WorkerHubTransactor) SetDAOToken(opts *bind.TransactOpts, _daoTokenAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setDAOToken", _daoTokenAddress)
}

// SetDAOToken is a paid mutator transaction binding the contract method 0x70a52354.
//
// Solidity: function setDAOToken(address _daoTokenAddress) returns()
func (_WorkerHub *WorkerHubSession) SetDAOToken(_daoTokenAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOToken(&_WorkerHub.TransactOpts, _daoTokenAddress)
}

// SetDAOToken is a paid mutator transaction binding the contract method 0x70a52354.
//
// Solidity: function setDAOToken(address _daoTokenAddress) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetDAOToken(_daoTokenAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOToken(&_WorkerHub.TransactOpts, _daoTokenAddress)
}

// SetDAOTokenPercentage is a paid mutator transaction binding the contract method 0x3860ce68.
//
// Solidity: function setDAOTokenPercentage((uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubTransactor) SetDAOTokenPercentage(opts *bind.TransactOpts, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setDAOTokenPercentage", _daoTokenPercentage)
}

// SetDAOTokenPercentage is a paid mutator transaction binding the contract method 0x3860ce68.
//
// Solidity: function setDAOTokenPercentage((uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubSession) SetDAOTokenPercentage(_daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOTokenPercentage(&_WorkerHub.TransactOpts, _daoTokenPercentage)
}

// SetDAOTokenPercentage is a paid mutator transaction binding the contract method 0x3860ce68.
//
// Solidity: function setDAOTokenPercentage((uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetDAOTokenPercentage(_daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOTokenPercentage(&_WorkerHub.TransactOpts, _daoTokenPercentage)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_WorkerHub *WorkerHubTransactor) SetDAOTokenReward(opts *bind.TransactOpts, _newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setDAOTokenReward", _newDAOTokenReward)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_WorkerHub *WorkerHubSession) SetDAOTokenReward(_newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOTokenReward(&_WorkerHub.TransactOpts, _newDAOTokenReward)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetDAOTokenReward(_newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetDAOTokenReward(&_WorkerHub.TransactOpts, _newDAOTokenReward)
}

// SetFeeRatioMinerValidator is a paid mutator transaction binding the contract method 0xafa82609.
//
// Solidity: function setFeeRatioMinerValidator(uint16 _newRatio) returns()
func (_WorkerHub *WorkerHubTransactor) SetFeeRatioMinerValidator(opts *bind.TransactOpts, _newRatio uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setFeeRatioMinerValidator", _newRatio)
}

// SetFeeRatioMinerValidator is a paid mutator transaction binding the contract method 0xafa82609.
//
// Solidity: function setFeeRatioMinerValidator(uint16 _newRatio) returns()
func (_WorkerHub *WorkerHubSession) SetFeeRatioMinerValidator(_newRatio uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetFeeRatioMinerValidator(&_WorkerHub.TransactOpts, _newRatio)
}

// SetFeeRatioMinerValidator is a paid mutator transaction binding the contract method 0xafa82609.
//
// Solidity: function setFeeRatioMinerValidator(uint16 _newRatio) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetFeeRatioMinerValidator(_newRatio uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetFeeRatioMinerValidator(&_WorkerHub.TransactOpts, _newRatio)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubTransactor) SetFinePercentage(opts *bind.TransactOpts, _finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setFinePercentage", _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetFinePercentage(&_WorkerHub.TransactOpts, _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetFinePercentage(&_WorkerHub.TransactOpts, _finePercentage)
}

// SetL2Owner is a paid mutator transaction binding the contract method 0xb530c110.
//
// Solidity: function setL2Owner(address _l2OwnerAddress) returns()
func (_WorkerHub *WorkerHubTransactor) SetL2Owner(opts *bind.TransactOpts, _l2OwnerAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setL2Owner", _l2OwnerAddress)
}

// SetL2Owner is a paid mutator transaction binding the contract method 0xb530c110.
//
// Solidity: function setL2Owner(address _l2OwnerAddress) returns()
func (_WorkerHub *WorkerHubSession) SetL2Owner(_l2OwnerAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetL2Owner(&_WorkerHub.TransactOpts, _l2OwnerAddress)
}

// SetL2Owner is a paid mutator transaction binding the contract method 0xb530c110.
//
// Solidity: function setL2Owner(address _l2OwnerAddress) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetL2Owner(_l2OwnerAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetL2Owner(&_WorkerHub.TransactOpts, _l2OwnerAddress)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_WorkerHub *WorkerHubTransactor) SetMinFeeToUse(opts *bind.TransactOpts, _minFeeToUse *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setMinFeeToUse", _minFeeToUse)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_WorkerHub *WorkerHubSession) SetMinFeeToUse(_minFeeToUse *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetMinFeeToUse(&_WorkerHub.TransactOpts, _minFeeToUse)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetMinFeeToUse(_minFeeToUse *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetMinFeeToUse(&_WorkerHub.TransactOpts, _minFeeToUse)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHub *WorkerHubTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, _newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setNewRewardInEpoch", _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHub *WorkerHubSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetNewRewardInEpoch(&_WorkerHub.TransactOpts, _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetNewRewardInEpoch(&_WorkerHub.TransactOpts, _newRewardAmount)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHub *WorkerHubTransactor) SetPenaltyDuration(opts *bind.TransactOpts, _penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setPenaltyDuration", _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHub *WorkerHubSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetPenaltyDuration(&_WorkerHub.TransactOpts, _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetPenaltyDuration(&_WorkerHub.TransactOpts, _penaltyDuration)
}

// SetRevealDuration is a paid mutator transaction binding the contract method 0x1eb9a99a.
//
// Solidity: function setRevealDuration(uint40 _newRevealDuration) returns()
func (_WorkerHub *WorkerHubTransactor) SetRevealDuration(opts *bind.TransactOpts, _newRevealDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setRevealDuration", _newRevealDuration)
}

// SetRevealDuration is a paid mutator transaction binding the contract method 0x1eb9a99a.
//
// Solidity: function setRevealDuration(uint40 _newRevealDuration) returns()
func (_WorkerHub *WorkerHubSession) SetRevealDuration(_newRevealDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetRevealDuration(&_WorkerHub.TransactOpts, _newRevealDuration)
}

// SetRevealDuration is a paid mutator transaction binding the contract method 0x1eb9a99a.
//
// Solidity: function setRevealDuration(uint40 _newRevealDuration) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetRevealDuration(_newRevealDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetRevealDuration(&_WorkerHub.TransactOpts, _newRevealDuration)
}

// SetScoringInfo is a paid mutator transaction binding the contract method 0x0d425ea5.
//
// Solidity: function setScoringInfo(address _workerHubScoring, address _modelScoring) returns()
func (_WorkerHub *WorkerHubTransactor) SetScoringInfo(opts *bind.TransactOpts, _workerHubScoring common.Address, _modelScoring common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setScoringInfo", _workerHubScoring, _modelScoring)
}

// SetScoringInfo is a paid mutator transaction binding the contract method 0x0d425ea5.
//
// Solidity: function setScoringInfo(address _workerHubScoring, address _modelScoring) returns()
func (_WorkerHub *WorkerHubSession) SetScoringInfo(_workerHubScoring common.Address, _modelScoring common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetScoringInfo(&_WorkerHub.TransactOpts, _workerHubScoring, _modelScoring)
}

// SetScoringInfo is a paid mutator transaction binding the contract method 0x0d425ea5.
//
// Solidity: function setScoringInfo(address _workerHubScoring, address _modelScoring) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetScoringInfo(_workerHubScoring common.Address, _modelScoring common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetScoringInfo(&_WorkerHub.TransactOpts, _workerHubScoring, _modelScoring)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 _newSubmitDuration) returns()
func (_WorkerHub *WorkerHubTransactor) SetSubmitDuration(opts *bind.TransactOpts, _newSubmitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setSubmitDuration", _newSubmitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 _newSubmitDuration) returns()
func (_WorkerHub *WorkerHubSession) SetSubmitDuration(_newSubmitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetSubmitDuration(&_WorkerHub.TransactOpts, _newSubmitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 _newSubmitDuration) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetSubmitDuration(_newSubmitDuration *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetSubmitDuration(&_WorkerHub.TransactOpts, _newSubmitDuration)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_WorkerHub *WorkerHubTransactor) SetTreasuryAddress(opts *bind.TransactOpts, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setTreasuryAddress", _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_WorkerHub *WorkerHubSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetTreasuryAddress(&_WorkerHub.TransactOpts, _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetTreasuryAddress(&_WorkerHub.TransactOpts, _treasuryAddress)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHub *WorkerHubTransactor) SetUnstakDelayTime(opts *bind.TransactOpts, _newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "setUnstakDelayTime", _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHub *WorkerHubSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetUnstakDelayTime(&_WorkerHub.TransactOpts, _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_WorkerHub *WorkerHubTransactorSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.SetUnstakDelayTime(&_WorkerHub.TransactOpts, _newUnstakeDelayTime)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHub *WorkerHubTransactor) SlashMiner(opts *bind.TransactOpts, _miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "slashMiner", _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHub *WorkerHubSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHub.Contract.SlashMiner(&_WorkerHub.TransactOpts, _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_WorkerHub *WorkerHubTransactorSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _WorkerHub.Contract.SlashMiner(&_WorkerHub.TransactOpts, _miner, _isFined)
}

// StreamData is a paid mutator transaction binding the contract method 0x020e3011.
//
// Solidity: function streamData(uint256 _assignmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactor) StreamData(opts *bind.TransactOpts, _assignmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "streamData", _assignmentId, _data)
}

// StreamData is a paid mutator transaction binding the contract method 0x020e3011.
//
// Solidity: function streamData(uint256 _assignmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubSession) StreamData(_assignmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.StreamData(&_WorkerHub.TransactOpts, _assignmentId, _data)
}

// StreamData is a paid mutator transaction binding the contract method 0x020e3011.
//
// Solidity: function streamData(uint256 _assignmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactorSession) StreamData(_assignmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.StreamData(&_WorkerHub.TransactOpts, _assignmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactor) SubmitSolution(opts *bind.TransactOpts, _assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "submitSolution", _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitSolution(&_WorkerHub.TransactOpts, _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_WorkerHub *WorkerHubTransactorSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _WorkerHub.Contract.SubmitSolution(&_WorkerHub.TransactOpts, _assigmentId, _data)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_WorkerHub *WorkerHubTransactor) TopUpInfer(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "topUpInfer", _inferenceId)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_WorkerHub *WorkerHubSession) TopUpInfer(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.TopUpInfer(&_WorkerHub.TransactOpts, _inferenceId)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_WorkerHub *WorkerHubTransactorSession) TopUpInfer(_inferenceId *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.TopUpInfer(&_WorkerHub.TransactOpts, _inferenceId)
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

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHub *WorkerHubTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHub *WorkerHubSession) UnregisterMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterMiner(&_WorkerHub.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnregisterMiner(&_WorkerHub.TransactOpts)
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

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHub *WorkerHubTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHub *WorkerHubSession) UnstakeForMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnstakeForMiner(&_WorkerHub.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_WorkerHub *WorkerHubTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _WorkerHub.Contract.UnstakeForMiner(&_WorkerHub.TransactOpts)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, _model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "updateModelMinimumFee", _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelMinimumFee(&_WorkerHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_WorkerHub *WorkerHubTransactorSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelMinimumFee(&_WorkerHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHub *WorkerHubTransactor) UpdateModelTier(opts *bind.TransactOpts, _model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHub.contract.Transact(opts, "updateModelTier", _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHub *WorkerHubSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelTier(&_WorkerHub.TransactOpts, _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_WorkerHub *WorkerHubTransactorSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _WorkerHub.Contract.UpdateModelTier(&_WorkerHub.TransactOpts, _model, _tier)
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

// WorkerHubBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the WorkerHub contract.
type WorkerHubBlocksPerEpochIterator struct {
	Event *WorkerHubBlocksPerEpoch // Event containing the contract specifics and raw log

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
func (it *WorkerHubBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubBlocksPerEpoch)
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
		it.Event = new(WorkerHubBlocksPerEpoch)
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
func (it *WorkerHubBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubBlocksPerEpoch represents a BlocksPerEpoch event raised by the WorkerHub contract.
type WorkerHubBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHub *WorkerHubFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*WorkerHubBlocksPerEpochIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &WorkerHubBlocksPerEpochIterator{contract: _WorkerHub.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_WorkerHub *WorkerHubFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *WorkerHubBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubBlocksPerEpoch)
				if err := _WorkerHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseBlocksPerEpoch(log types.Log) (*WorkerHubBlocksPerEpoch, error) {
	event := new(WorkerHubBlocksPerEpoch)
	if err := _WorkerHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubCommitDurationIterator is returned from FilterCommitDuration and is used to iterate over the raw logs and unpacked data for CommitDuration events raised by the WorkerHub contract.
type WorkerHubCommitDurationIterator struct {
	Event *WorkerHubCommitDuration // Event containing the contract specifics and raw log

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
func (it *WorkerHubCommitDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubCommitDuration)
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
		it.Event = new(WorkerHubCommitDuration)
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
func (it *WorkerHubCommitDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubCommitDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubCommitDuration represents a CommitDuration event raised by the WorkerHub contract.
type WorkerHubCommitDuration struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCommitDuration is a free log retrieval operation binding the contract event 0xc9bc20c9ff07142c58c480090e116ebe561a42316260069d619782bb38faf619.
//
// Solidity: event CommitDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) FilterCommitDuration(opts *bind.FilterOpts) (*WorkerHubCommitDurationIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "CommitDuration")
	if err != nil {
		return nil, err
	}
	return &WorkerHubCommitDurationIterator{contract: _WorkerHub.contract, event: "CommitDuration", logs: logs, sub: sub}, nil
}

// WatchCommitDuration is a free log subscription operation binding the contract event 0xc9bc20c9ff07142c58c480090e116ebe561a42316260069d619782bb38faf619.
//
// Solidity: event CommitDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) WatchCommitDuration(opts *bind.WatchOpts, sink chan<- *WorkerHubCommitDuration) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "CommitDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubCommitDuration)
				if err := _WorkerHub.contract.UnpackLog(event, "CommitDuration", log); err != nil {
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

// ParseCommitDuration is a log parse operation binding the contract event 0xc9bc20c9ff07142c58c480090e116ebe561a42316260069d619782bb38faf619.
//
// Solidity: event CommitDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) ParseCommitDuration(log types.Log) (*WorkerHubCommitDuration, error) {
	event := new(WorkerHubCommitDuration)
	if err := _WorkerHub.contract.UnpackLog(event, "CommitDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubCommitmentSubmissionIterator is returned from FilterCommitmentSubmission and is used to iterate over the raw logs and unpacked data for CommitmentSubmission events raised by the WorkerHub contract.
type WorkerHubCommitmentSubmissionIterator struct {
	Event *WorkerHubCommitmentSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubCommitmentSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubCommitmentSubmission)
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
		it.Event = new(WorkerHubCommitmentSubmission)
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
func (it *WorkerHubCommitmentSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubCommitmentSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubCommitmentSubmission represents a CommitmentSubmission event raised by the WorkerHub contract.
type WorkerHubCommitmentSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommitmentSubmission is a free log retrieval operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_WorkerHub *WorkerHubFilterer) FilterCommitmentSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubCommitmentSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "CommitmentSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubCommitmentSubmissionIterator{contract: _WorkerHub.contract, event: "CommitmentSubmission", logs: logs, sub: sub}, nil
}

// WatchCommitmentSubmission is a free log subscription operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_WorkerHub *WorkerHubFilterer) WatchCommitmentSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubCommitmentSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "CommitmentSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubCommitmentSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "CommitmentSubmission", log); err != nil {
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

// ParseCommitmentSubmission is a log parse operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_WorkerHub *WorkerHubFilterer) ParseCommitmentSubmission(log types.Log) (*WorkerHubCommitmentSubmission, error) {
	event := new(WorkerHubCommitmentSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "CommitmentSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubDAOTokenMintedV2Iterator is returned from FilterDAOTokenMintedV2 and is used to iterate over the raw logs and unpacked data for DAOTokenMintedV2 events raised by the WorkerHub contract.
type WorkerHubDAOTokenMintedV2Iterator struct {
	Event *WorkerHubDAOTokenMintedV2 // Event containing the contract specifics and raw log

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
func (it *WorkerHubDAOTokenMintedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubDAOTokenMintedV2)
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
		it.Event = new(WorkerHubDAOTokenMintedV2)
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
func (it *WorkerHubDAOTokenMintedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubDAOTokenMintedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubDAOTokenMintedV2 represents a DAOTokenMintedV2 event raised by the WorkerHub contract.
type WorkerHubDAOTokenMintedV2 struct {
	ChainId      *big.Int
	InferenceId  *big.Int
	ModelAddress common.Address
	Receivers    []IWorkerHubDAOTokenReceiverInfor
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenMintedV2 is a free log retrieval operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_WorkerHub *WorkerHubFilterer) FilterDAOTokenMintedV2(opts *bind.FilterOpts) (*WorkerHubDAOTokenMintedV2Iterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "DAOTokenMintedV2")
	if err != nil {
		return nil, err
	}
	return &WorkerHubDAOTokenMintedV2Iterator{contract: _WorkerHub.contract, event: "DAOTokenMintedV2", logs: logs, sub: sub}, nil
}

// WatchDAOTokenMintedV2 is a free log subscription operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_WorkerHub *WorkerHubFilterer) WatchDAOTokenMintedV2(opts *bind.WatchOpts, sink chan<- *WorkerHubDAOTokenMintedV2) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "DAOTokenMintedV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubDAOTokenMintedV2)
				if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenMintedV2", log); err != nil {
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

// ParseDAOTokenMintedV2 is a log parse operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_WorkerHub *WorkerHubFilterer) ParseDAOTokenMintedV2(log types.Log) (*WorkerHubDAOTokenMintedV2, error) {
	event := new(WorkerHubDAOTokenMintedV2)
	if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenMintedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubDAOTokenPercentageUpdatedIterator is returned from FilterDAOTokenPercentageUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenPercentageUpdated events raised by the WorkerHub contract.
type WorkerHubDAOTokenPercentageUpdatedIterator struct {
	Event *WorkerHubDAOTokenPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubDAOTokenPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubDAOTokenPercentageUpdated)
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
		it.Event = new(WorkerHubDAOTokenPercentageUpdated)
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
func (it *WorkerHubDAOTokenPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubDAOTokenPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubDAOTokenPercentageUpdated represents a DAOTokenPercentageUpdated event raised by the WorkerHub contract.
type WorkerHubDAOTokenPercentageUpdated struct {
	OldValue IWorkerHubDAOTokenPercentage
	NewValue IWorkerHubDAOTokenPercentage
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenPercentageUpdated is a free log retrieval operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_WorkerHub *WorkerHubFilterer) FilterDAOTokenPercentageUpdated(opts *bind.FilterOpts) (*WorkerHubDAOTokenPercentageUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "DAOTokenPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubDAOTokenPercentageUpdatedIterator{contract: _WorkerHub.contract, event: "DAOTokenPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenPercentageUpdated is a free log subscription operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_WorkerHub *WorkerHubFilterer) WatchDAOTokenPercentageUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubDAOTokenPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "DAOTokenPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubDAOTokenPercentageUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenPercentageUpdated", log); err != nil {
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

// ParseDAOTokenPercentageUpdated is a log parse operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_WorkerHub *WorkerHubFilterer) ParseDAOTokenPercentageUpdated(log types.Log) (*WorkerHubDAOTokenPercentageUpdated, error) {
	event := new(WorkerHubDAOTokenPercentageUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubDAOTokenRewardUpdatedIterator is returned from FilterDAOTokenRewardUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenRewardUpdated events raised by the WorkerHub contract.
type WorkerHubDAOTokenRewardUpdatedIterator struct {
	Event *WorkerHubDAOTokenRewardUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubDAOTokenRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubDAOTokenRewardUpdated)
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
		it.Event = new(WorkerHubDAOTokenRewardUpdated)
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
func (it *WorkerHubDAOTokenRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubDAOTokenRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubDAOTokenRewardUpdated represents a DAOTokenRewardUpdated event raised by the WorkerHub contract.
type WorkerHubDAOTokenRewardUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenRewardUpdated is a free log retrieval operation binding the contract event 0x454d79b61f30800ce19615c79c4f9a1eb892ed9372cf95ba71cbd2345f8fa9aa.
//
// Solidity: event DAOTokenRewardUpdated(uint256 oldValue, uint256 newValue)
func (_WorkerHub *WorkerHubFilterer) FilterDAOTokenRewardUpdated(opts *bind.FilterOpts) (*WorkerHubDAOTokenRewardUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "DAOTokenRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubDAOTokenRewardUpdatedIterator{contract: _WorkerHub.contract, event: "DAOTokenRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenRewardUpdated is a free log subscription operation binding the contract event 0x454d79b61f30800ce19615c79c4f9a1eb892ed9372cf95ba71cbd2345f8fa9aa.
//
// Solidity: event DAOTokenRewardUpdated(uint256 oldValue, uint256 newValue)
func (_WorkerHub *WorkerHubFilterer) WatchDAOTokenRewardUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubDAOTokenRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "DAOTokenRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubDAOTokenRewardUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenRewardUpdated", log); err != nil {
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

// ParseDAOTokenRewardUpdated is a log parse operation binding the contract event 0x454d79b61f30800ce19615c79c4f9a1eb892ed9372cf95ba71cbd2345f8fa9aa.
//
// Solidity: event DAOTokenRewardUpdated(uint256 oldValue, uint256 newValue)
func (_WorkerHub *WorkerHubFilterer) ParseDAOTokenRewardUpdated(log types.Log) (*WorkerHubDAOTokenRewardUpdated, error) {
	event := new(WorkerHubDAOTokenRewardUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubDAOTokenUpdatedIterator is returned from FilterDAOTokenUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenUpdated events raised by the WorkerHub contract.
type WorkerHubDAOTokenUpdatedIterator struct {
	Event *WorkerHubDAOTokenUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubDAOTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubDAOTokenUpdated)
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
		it.Event = new(WorkerHubDAOTokenUpdated)
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
func (it *WorkerHubDAOTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubDAOTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubDAOTokenUpdated represents a DAOTokenUpdated event raised by the WorkerHub contract.
type WorkerHubDAOTokenUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenUpdated is a free log retrieval operation binding the contract event 0x518cc1a1508767ac2e92e88727dbf2ace68f44768b3684e0ad2305f6db0cd8da.
//
// Solidity: event DAOTokenUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) FilterDAOTokenUpdated(opts *bind.FilterOpts) (*WorkerHubDAOTokenUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "DAOTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubDAOTokenUpdatedIterator{contract: _WorkerHub.contract, event: "DAOTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenUpdated is a free log subscription operation binding the contract event 0x518cc1a1508767ac2e92e88727dbf2ace68f44768b3684e0ad2305f6db0cd8da.
//
// Solidity: event DAOTokenUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) WatchDAOTokenUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubDAOTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "DAOTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubDAOTokenUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenUpdated", log); err != nil {
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

// ParseDAOTokenUpdated is a log parse operation binding the contract event 0x518cc1a1508767ac2e92e88727dbf2ace68f44768b3684e0ad2305f6db0cd8da.
//
// Solidity: event DAOTokenUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) ParseDAOTokenUpdated(log types.Log) (*WorkerHubDAOTokenUpdated, error) {
	event := new(WorkerHubDAOTokenUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "DAOTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the WorkerHub contract.
type WorkerHubFinePercentageUpdatedIterator struct {
	Event *WorkerHubFinePercentageUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubFinePercentageUpdated)
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
		it.Event = new(WorkerHubFinePercentageUpdated)
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
func (it *WorkerHubFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubFinePercentageUpdated represents a FinePercentageUpdated event raised by the WorkerHub contract.
type WorkerHubFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHub *WorkerHubFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*WorkerHubFinePercentageUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubFinePercentageUpdatedIterator{contract: _WorkerHub.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_WorkerHub *WorkerHubFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubFinePercentageUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseFinePercentageUpdated(log types.Log) (*WorkerHubFinePercentageUpdated, error) {
	event := new(WorkerHubFinePercentageUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the WorkerHub contract.
type WorkerHubFraudulentMinerPenalizedIterator struct {
	Event *WorkerHubFraudulentMinerPenalized // Event containing the contract specifics and raw log

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
func (it *WorkerHubFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubFraudulentMinerPenalized)
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
		it.Event = new(WorkerHubFraudulentMinerPenalized)
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
func (it *WorkerHubFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the WorkerHub contract.
type WorkerHubFraudulentMinerPenalized struct {
	Miner        common.Address
	ModelAddress common.Address
	Treasury     common.Address
	Fine         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHub *WorkerHubFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (*WorkerHubFraudulentMinerPenalizedIterator, error) {

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

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubFraudulentMinerPenalizedIterator{contract: _WorkerHub.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_WorkerHub *WorkerHubFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *WorkerHubFraudulentMinerPenalized, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubFraudulentMinerPenalized)
				if err := _WorkerHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseFraudulentMinerPenalized(log types.Log) (*WorkerHubFraudulentMinerPenalized, error) {
	event := new(WorkerHubFraudulentMinerPenalized)
	if err := _WorkerHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the WorkerHub contract.
type WorkerHubInferenceStatusUpdateIterator struct {
	Event *WorkerHubInferenceStatusUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubInferenceStatusUpdate)
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
		it.Event = new(WorkerHubInferenceStatusUpdate)
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
func (it *WorkerHubInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the WorkerHub contract.
type WorkerHubInferenceStatusUpdate struct {
	InferenceId *big.Int
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []*big.Int) (*WorkerHubInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubInferenceStatusUpdateIterator{contract: _WorkerHub.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_WorkerHub *WorkerHubFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubInferenceStatusUpdate, inferenceId []*big.Int) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubInferenceStatusUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseInferenceStatusUpdate(log types.Log) (*WorkerHubInferenceStatusUpdate, error) {
	event := new(WorkerHubInferenceStatusUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// WorkerHubL2OwnerUpdatedIterator is returned from FilterL2OwnerUpdated and is used to iterate over the raw logs and unpacked data for L2OwnerUpdated events raised by the WorkerHub contract.
type WorkerHubL2OwnerUpdatedIterator struct {
	Event *WorkerHubL2OwnerUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubL2OwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubL2OwnerUpdated)
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
		it.Event = new(WorkerHubL2OwnerUpdated)
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
func (it *WorkerHubL2OwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubL2OwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubL2OwnerUpdated represents a L2OwnerUpdated event raised by the WorkerHub contract.
type WorkerHubL2OwnerUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2OwnerUpdated is a free log retrieval operation binding the contract event 0x3cfa9fea14972d7cbbd0fddda517d4467bd2863f1d28e76fa4e0fe230a7bf274.
//
// Solidity: event L2OwnerUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) FilterL2OwnerUpdated(opts *bind.FilterOpts) (*WorkerHubL2OwnerUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "L2OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubL2OwnerUpdatedIterator{contract: _WorkerHub.contract, event: "L2OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchL2OwnerUpdated is a free log subscription operation binding the contract event 0x3cfa9fea14972d7cbbd0fddda517d4467bd2863f1d28e76fa4e0fe230a7bf274.
//
// Solidity: event L2OwnerUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) WatchL2OwnerUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubL2OwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "L2OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubL2OwnerUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "L2OwnerUpdated", log); err != nil {
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

// ParseL2OwnerUpdated is a log parse operation binding the contract event 0x3cfa9fea14972d7cbbd0fddda517d4467bd2863f1d28e76fa4e0fe230a7bf274.
//
// Solidity: event L2OwnerUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) ParseL2OwnerUpdated(log types.Log) (*WorkerHubL2OwnerUpdated, error) {
	event := new(WorkerHubL2OwnerUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "L2OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinFeeToUseUpdatedIterator is returned from FilterMinFeeToUseUpdated and is used to iterate over the raw logs and unpacked data for MinFeeToUseUpdated events raised by the WorkerHub contract.
type WorkerHubMinFeeToUseUpdatedIterator struct {
	Event *WorkerHubMinFeeToUseUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinFeeToUseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinFeeToUseUpdated)
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
		it.Event = new(WorkerHubMinFeeToUseUpdated)
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
func (it *WorkerHubMinFeeToUseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinFeeToUseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinFeeToUseUpdated represents a MinFeeToUseUpdated event raised by the WorkerHub contract.
type WorkerHubMinFeeToUseUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeToUseUpdated is a free log retrieval operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_WorkerHub *WorkerHubFilterer) FilterMinFeeToUseUpdated(opts *bind.FilterOpts) (*WorkerHubMinFeeToUseUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinFeeToUseUpdatedIterator{contract: _WorkerHub.contract, event: "MinFeeToUseUpdated", logs: logs, sub: sub}, nil
}

// WatchMinFeeToUseUpdated is a free log subscription operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_WorkerHub *WorkerHubFilterer) WatchMinFeeToUseUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubMinFeeToUseUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinFeeToUseUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
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

// ParseMinFeeToUseUpdated is a log parse operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_WorkerHub *WorkerHubFilterer) ParseMinFeeToUseUpdated(log types.Log) (*WorkerHubMinFeeToUseUpdated, error) {
	event := new(WorkerHubMinFeeToUseUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the WorkerHub contract.
type WorkerHubMinerDeactivatedIterator struct {
	Event *WorkerHubMinerDeactivated // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerDeactivated)
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
		it.Event = new(WorkerHubMinerDeactivated)
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
func (it *WorkerHubMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerDeactivated represents a MinerDeactivated event raised by the WorkerHub contract.
type WorkerHubMinerDeactivated struct {
	Miner        common.Address
	ModelAddress common.Address
	ActiveTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHub *WorkerHubFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address) (*WorkerHubMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerDeactivatedIterator{contract: _WorkerHub.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_WorkerHub *WorkerHubFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerDeactivated, miner []common.Address, modelAddress []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerDeactivated)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMinerDeactivated(log types.Log) (*WorkerHubMinerDeactivated, error) {
	event := new(WorkerHubMinerDeactivated)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the WorkerHub contract.
type WorkerHubMinerExtraStakeIterator struct {
	Event *WorkerHubMinerExtraStake // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerExtraStake)
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
		it.Event = new(WorkerHubMinerExtraStake)
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
func (it *WorkerHubMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerExtraStake represents a MinerExtraStake event raised by the WorkerHub contract.
type WorkerHubMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerExtraStakeIterator{contract: _WorkerHub.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerExtraStake)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMinerExtraStake(log types.Log) (*WorkerHubMinerExtraStake, error) {
	event := new(WorkerHubMinerExtraStake)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the WorkerHub contract.
type WorkerHubMinerJoinIterator struct {
	Event *WorkerHubMinerJoin // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerJoin)
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
		it.Event = new(WorkerHubMinerJoin)
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
func (it *WorkerHubMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerJoin represents a MinerJoin event raised by the WorkerHub contract.
type WorkerHubMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerJoinIterator{contract: _WorkerHub.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerJoin)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMinerJoin(log types.Log) (*WorkerHubMinerJoin, error) {
	event := new(WorkerHubMinerJoin)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the WorkerHub contract.
type WorkerHubMinerRegistrationIterator struct {
	Event *WorkerHubMinerRegistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerRegistration)
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
		it.Event = new(WorkerHubMinerRegistration)
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
func (it *WorkerHubMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerRegistration represents a MinerRegistration event raised by the WorkerHub contract.
type WorkerHubMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*WorkerHubMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerRegistrationIterator{contract: _WorkerHub.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerRegistration)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMinerRegistration(log types.Log) (*WorkerHubMinerRegistration, error) {
	event := new(WorkerHubMinerRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerRoleSeizedIterator is returned from FilterMinerRoleSeized and is used to iterate over the raw logs and unpacked data for MinerRoleSeized events raised by the WorkerHub contract.
type WorkerHubMinerRoleSeizedIterator struct {
	Event *WorkerHubMinerRoleSeized // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerRoleSeizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerRoleSeized)
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
		it.Event = new(WorkerHubMinerRoleSeized)
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
func (it *WorkerHubMinerRoleSeizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerRoleSeizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerRoleSeized represents a MinerRoleSeized event raised by the WorkerHub contract.
type WorkerHubMinerRoleSeized struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerRoleSeized is a free log retrieval operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_WorkerHub *WorkerHubFilterer) FilterMinerRoleSeized(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*WorkerHubMinerRoleSeizedIterator, error) {

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

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerRoleSeized", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerRoleSeizedIterator{contract: _WorkerHub.contract, event: "MinerRoleSeized", logs: logs, sub: sub}, nil
}

// WatchMinerRoleSeized is a free log subscription operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_WorkerHub *WorkerHubFilterer) WatchMinerRoleSeized(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerRoleSeized, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerRoleSeized", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerRoleSeized)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerRoleSeized", log); err != nil {
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

// ParseMinerRoleSeized is a log parse operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_WorkerHub *WorkerHubFilterer) ParseMinerRoleSeized(log types.Log) (*WorkerHubMinerRoleSeized, error) {
	event := new(WorkerHubMinerRoleSeized)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerRoleSeized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the WorkerHub contract.
type WorkerHubMinerUnregistrationIterator struct {
	Event *WorkerHubMinerUnregistration // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerUnregistration)
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
		it.Event = new(WorkerHubMinerUnregistration)
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
func (it *WorkerHubMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerUnregistration represents a MinerUnregistration event raised by the WorkerHub contract.
type WorkerHubMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerUnregistrationIterator{contract: _WorkerHub.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_WorkerHub *WorkerHubFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerUnregistration)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMinerUnregistration(log types.Log) (*WorkerHubMinerUnregistration, error) {
	event := new(WorkerHubMinerUnregistration)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the WorkerHub contract.
type WorkerHubMinerUnstakeIterator struct {
	Event *WorkerHubMinerUnstake // Event containing the contract specifics and raw log

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
func (it *WorkerHubMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMinerUnstake)
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
		it.Event = new(WorkerHubMinerUnstake)
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
func (it *WorkerHubMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMinerUnstake represents a MinerUnstake event raised by the WorkerHub contract.
type WorkerHubMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*WorkerHubMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubMinerUnstakeIterator{contract: _WorkerHub.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_WorkerHub *WorkerHubFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *WorkerHubMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMinerUnstake)
				if err := _WorkerHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMinerUnstake(log types.Log) (*WorkerHubMinerUnstake, error) {
	event := new(WorkerHubMinerUnstake)
	if err := _WorkerHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubMiningTimeLimitUpdateIterator is returned from FilterMiningTimeLimitUpdate and is used to iterate over the raw logs and unpacked data for MiningTimeLimitUpdate events raised by the WorkerHub contract.
type WorkerHubMiningTimeLimitUpdateIterator struct {
	Event *WorkerHubMiningTimeLimitUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubMiningTimeLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubMiningTimeLimitUpdate)
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
		it.Event = new(WorkerHubMiningTimeLimitUpdate)
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
func (it *WorkerHubMiningTimeLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubMiningTimeLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubMiningTimeLimitUpdate represents a MiningTimeLimitUpdate event raised by the WorkerHub contract.
type WorkerHubMiningTimeLimitUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMiningTimeLimitUpdate is a free log retrieval operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHub *WorkerHubFilterer) FilterMiningTimeLimitUpdate(opts *bind.FilterOpts) (*WorkerHubMiningTimeLimitUpdateIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &WorkerHubMiningTimeLimitUpdateIterator{contract: _WorkerHub.contract, event: "MiningTimeLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchMiningTimeLimitUpdate is a free log subscription operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_WorkerHub *WorkerHubFilterer) WatchMiningTimeLimitUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubMiningTimeLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubMiningTimeLimitUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseMiningTimeLimitUpdate(log types.Log) (*WorkerHubMiningTimeLimitUpdate, error) {
	event := new(WorkerHubMiningTimeLimitUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the WorkerHub contract.
type WorkerHubModelMinimumFeeUpdateIterator struct {
	Event *WorkerHubModelMinimumFeeUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelMinimumFeeUpdate)
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
		it.Event = new(WorkerHubModelMinimumFeeUpdate)
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
func (it *WorkerHubModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the WorkerHub contract.
type WorkerHubModelMinimumFeeUpdate struct {
	Model      common.Address
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, model []common.Address) (*WorkerHubModelMinimumFeeUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelMinimumFeeUpdateIterator{contract: _WorkerHub.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubModelMinimumFeeUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelMinimumFeeUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*WorkerHubModelMinimumFeeUpdate, error) {
	event := new(WorkerHubModelMinimumFeeUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
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
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) FilterModelRegistration(opts *bind.FilterOpts, model []common.Address, tier []uint16) (*WorkerHubModelRegistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelRegistrationIterator{contract: _WorkerHub.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *WorkerHubModelRegistration, model []common.Address, tier []uint16) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelRegistration", modelRule, tierRule)
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

// ParseModelRegistration is a log parse operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_WorkerHub *WorkerHubFilterer) ParseModelRegistration(log types.Log) (*WorkerHubModelRegistration, error) {
	event := new(WorkerHubModelRegistration)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the WorkerHub contract.
type WorkerHubModelTierUpdateIterator struct {
	Event *WorkerHubModelTierUpdate // Event containing the contract specifics and raw log

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
func (it *WorkerHubModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubModelTierUpdate)
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
		it.Event = new(WorkerHubModelTierUpdate)
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
func (it *WorkerHubModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubModelTierUpdate represents a ModelTierUpdate event raised by the WorkerHub contract.
type WorkerHubModelTierUpdate struct {
	Model common.Address
	Tier  uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHub *WorkerHubFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, model []common.Address) (*WorkerHubModelTierUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubModelTierUpdateIterator{contract: _WorkerHub.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_WorkerHub *WorkerHubFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *WorkerHubModelTierUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubModelTierUpdate)
				if err := _WorkerHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseModelTierUpdate(log types.Log) (*WorkerHubModelTierUpdate, error) {
	event := new(WorkerHubModelTierUpdate)
	if err := _WorkerHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
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

// WorkerHubNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the WorkerHub contract.
type WorkerHubNewAssignmentIterator struct {
	Event *WorkerHubNewAssignment // Event containing the contract specifics and raw log

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
func (it *WorkerHubNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewAssignment)
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
		it.Event = new(WorkerHubNewAssignment)
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
func (it *WorkerHubNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewAssignment represents a NewAssignment event raised by the WorkerHub contract.
type WorkerHubNewAssignment struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	ExpiredAt    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) FilterNewAssignment(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*WorkerHubNewAssignmentIterator, error) {

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

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewAssignmentIterator{contract: _WorkerHub.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_WorkerHub *WorkerHubFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *WorkerHubNewAssignment, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewAssignment)
				if err := _WorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseNewAssignment(log types.Log) (*WorkerHubNewAssignment, error) {
	event := new(WorkerHubNewAssignment)
	if err := _WorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*WorkerHubNewInferenceIterator, error) {

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

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewInferenceIterator{contract: _WorkerHub.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *WorkerHubNewInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
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

// ParseNewInference is a log parse operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) ParseNewInference(log types.Log) (*WorkerHubNewInference, error) {
	event := new(WorkerHubNewInference)
	if err := _WorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubNewScoringInferenceIterator is returned from FilterNewScoringInference and is used to iterate over the raw logs and unpacked data for NewScoringInference events raised by the WorkerHub contract.
type WorkerHubNewScoringInferenceIterator struct {
	Event *WorkerHubNewScoringInference // Event containing the contract specifics and raw log

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
func (it *WorkerHubNewScoringInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubNewScoringInference)
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
		it.Event = new(WorkerHubNewScoringInference)
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
func (it *WorkerHubNewScoringInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubNewScoringInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubNewScoringInference represents a NewScoringInference event raised by the WorkerHub contract.
type WorkerHubNewScoringInference struct {
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewScoringInference is a free log retrieval operation binding the contract event 0x3ec54c04f8c304e8caa7314d1ac4d34bff1c57151f207745b19e6d8f0a579ea9.
//
// Solidity: event NewScoringInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) FilterNewScoringInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*WorkerHubNewScoringInferenceIterator, error) {

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

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "NewScoringInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubNewScoringInferenceIterator{contract: _WorkerHub.contract, event: "NewScoringInference", logs: logs, sub: sub}, nil
}

// WatchNewScoringInference is a free log subscription operation binding the contract event 0x3ec54c04f8c304e8caa7314d1ac4d34bff1c57151f207745b19e6d8f0a579ea9.
//
// Solidity: event NewScoringInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) WatchNewScoringInference(opts *bind.WatchOpts, sink chan<- *WorkerHubNewScoringInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "NewScoringInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubNewScoringInference)
				if err := _WorkerHub.contract.UnpackLog(event, "NewScoringInference", log); err != nil {
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

// ParseNewScoringInference is a log parse operation binding the contract event 0x3ec54c04f8c304e8caa7314d1ac4d34bff1c57151f207745b19e6d8f0a579ea9.
//
// Solidity: event NewScoringInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_WorkerHub *WorkerHubFilterer) ParseNewScoringInference(log types.Log) (*WorkerHubNewScoringInference, error) {
	event := new(WorkerHubNewScoringInference)
	if err := _WorkerHub.contract.UnpackLog(event, "NewScoringInference", log); err != nil {
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

// WorkerHubPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the WorkerHub contract.
type WorkerHubPenaltyDurationUpdatedIterator struct {
	Event *WorkerHubPenaltyDurationUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubPenaltyDurationUpdated)
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
		it.Event = new(WorkerHubPenaltyDurationUpdated)
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
func (it *WorkerHubPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the WorkerHub contract.
type WorkerHubPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHub *WorkerHubFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*WorkerHubPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubPenaltyDurationUpdatedIterator{contract: _WorkerHub.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_WorkerHub *WorkerHubFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubPenaltyDurationUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParsePenaltyDurationUpdated(log types.Log) (*WorkerHubPenaltyDurationUpdated, error) {
	event := new(WorkerHubPenaltyDurationUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the WorkerHub contract.
type WorkerHubRestakeIterator struct {
	Event *WorkerHubRestake // Event containing the contract specifics and raw log

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
func (it *WorkerHubRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRestake)
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
		it.Event = new(WorkerHubRestake)
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
func (it *WorkerHubRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRestake represents a Restake event raised by the WorkerHub contract.
type WorkerHubRestake struct {
	Miner   common.Address
	Restake *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHub *WorkerHubFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, model []common.Address) (*WorkerHubRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRestakeIterator{contract: _WorkerHub.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_WorkerHub *WorkerHubFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *WorkerHubRestake, miner []common.Address, model []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRestake)
				if err := _WorkerHub.contract.UnpackLog(event, "Restake", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseRestake(log types.Log) (*WorkerHubRestake, error) {
	event := new(WorkerHubRestake)
	if err := _WorkerHub.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRevealDurationIterator is returned from FilterRevealDuration and is used to iterate over the raw logs and unpacked data for RevealDuration events raised by the WorkerHub contract.
type WorkerHubRevealDurationIterator struct {
	Event *WorkerHubRevealDuration // Event containing the contract specifics and raw log

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
func (it *WorkerHubRevealDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRevealDuration)
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
		it.Event = new(WorkerHubRevealDuration)
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
func (it *WorkerHubRevealDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRevealDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRevealDuration represents a RevealDuration event raised by the WorkerHub contract.
type WorkerHubRevealDuration struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealDuration is a free log retrieval operation binding the contract event 0xacb24019039b4d00193b2be5c85ea8ed6bd6747ed79f7d1e5a6d9384282b4a9d.
//
// Solidity: event RevealDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) FilterRevealDuration(opts *bind.FilterOpts) (*WorkerHubRevealDurationIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RevealDuration")
	if err != nil {
		return nil, err
	}
	return &WorkerHubRevealDurationIterator{contract: _WorkerHub.contract, event: "RevealDuration", logs: logs, sub: sub}, nil
}

// WatchRevealDuration is a free log subscription operation binding the contract event 0xacb24019039b4d00193b2be5c85ea8ed6bd6747ed79f7d1e5a6d9384282b4a9d.
//
// Solidity: event RevealDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) WatchRevealDuration(opts *bind.WatchOpts, sink chan<- *WorkerHubRevealDuration) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RevealDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRevealDuration)
				if err := _WorkerHub.contract.UnpackLog(event, "RevealDuration", log); err != nil {
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

// ParseRevealDuration is a log parse operation binding the contract event 0xacb24019039b4d00193b2be5c85ea8ed6bd6747ed79f7d1e5a6d9384282b4a9d.
//
// Solidity: event RevealDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) ParseRevealDuration(log types.Log) (*WorkerHubRevealDuration, error) {
	event := new(WorkerHubRevealDuration)
	if err := _WorkerHub.contract.UnpackLog(event, "RevealDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRevealSubmissionIterator is returned from FilterRevealSubmission and is used to iterate over the raw logs and unpacked data for RevealSubmission events raised by the WorkerHub contract.
type WorkerHubRevealSubmissionIterator struct {
	Event *WorkerHubRevealSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubRevealSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRevealSubmission)
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
		it.Event = new(WorkerHubRevealSubmission)
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
func (it *WorkerHubRevealSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRevealSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRevealSubmission represents a RevealSubmission event raised by the WorkerHub contract.
type WorkerHubRevealSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Nonce       *big.Int
	Output      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealSubmission is a free log retrieval operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_WorkerHub *WorkerHubFilterer) FilterRevealSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubRevealSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RevealSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRevealSubmissionIterator{contract: _WorkerHub.contract, event: "RevealSubmission", logs: logs, sub: sub}, nil
}

// WatchRevealSubmission is a free log subscription operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_WorkerHub *WorkerHubFilterer) WatchRevealSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubRevealSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RevealSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRevealSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "RevealSubmission", log); err != nil {
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

// ParseRevealSubmission is a log parse operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_WorkerHub *WorkerHubFilterer) ParseRevealSubmission(log types.Log) (*WorkerHubRevealSubmission, error) {
	event := new(WorkerHubRevealSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "RevealSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the WorkerHub contract.
type WorkerHubRewardClaimIterator struct {
	Event *WorkerHubRewardClaim // Event containing the contract specifics and raw log

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
func (it *WorkerHubRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRewardClaim)
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
		it.Event = new(WorkerHubRewardClaim)
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
func (it *WorkerHubRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRewardClaim represents a RewardClaim event raised by the WorkerHub contract.
type WorkerHubRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*WorkerHubRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubRewardClaimIterator{contract: _WorkerHub.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *WorkerHubRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRewardClaim)
				if err := _WorkerHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseRewardClaim(log types.Log) (*WorkerHubRewardClaim, error) {
	event := new(WorkerHubRewardClaim)
	if err := _WorkerHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the WorkerHub contract.
type WorkerHubRewardPerEpochIterator struct {
	Event *WorkerHubRewardPerEpoch // Event containing the contract specifics and raw log

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
func (it *WorkerHubRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubRewardPerEpoch)
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
		it.Event = new(WorkerHubRewardPerEpoch)
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
func (it *WorkerHubRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubRewardPerEpoch represents a RewardPerEpoch event raised by the WorkerHub contract.
type WorkerHubRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHub *WorkerHubFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*WorkerHubRewardPerEpochIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &WorkerHubRewardPerEpochIterator{contract: _WorkerHub.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_WorkerHub *WorkerHubFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *WorkerHubRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubRewardPerEpoch)
				if err := _WorkerHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseRewardPerEpoch(log types.Log) (*WorkerHubRewardPerEpoch, error) {
	event := new(WorkerHubRewardPerEpoch)
	if err := _WorkerHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the WorkerHub contract.
type WorkerHubSolutionSubmissionIterator struct {
	Event *WorkerHubSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *WorkerHubSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubSolutionSubmission)
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
		it.Event = new(WorkerHubSolutionSubmission)
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
func (it *WorkerHubSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubSolutionSubmission represents a SolutionSubmission event raised by the WorkerHub contract.
type WorkerHubSolutionSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*WorkerHubSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubSolutionSubmissionIterator{contract: _WorkerHub.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_WorkerHub *WorkerHubFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *WorkerHubSolutionSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubSolutionSubmission)
				if err := _WorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseSolutionSubmission(log types.Log) (*WorkerHubSolutionSubmission, error) {
	event := new(WorkerHubSolutionSubmission)
	if err := _WorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubStreamedDataIterator is returned from FilterStreamedData and is used to iterate over the raw logs and unpacked data for StreamedData events raised by the WorkerHub contract.
type WorkerHubStreamedDataIterator struct {
	Event *WorkerHubStreamedData // Event containing the contract specifics and raw log

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
func (it *WorkerHubStreamedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubStreamedData)
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
		it.Event = new(WorkerHubStreamedData)
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
func (it *WorkerHubStreamedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubStreamedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubStreamedData represents a StreamedData event raised by the WorkerHub contract.
type WorkerHubStreamedData struct {
	AssignmentId *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStreamedData is a free log retrieval operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerHub *WorkerHubFilterer) FilterStreamedData(opts *bind.FilterOpts, assignmentId []*big.Int) (*WorkerHubStreamedDataIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubStreamedDataIterator{contract: _WorkerHub.contract, event: "StreamedData", logs: logs, sub: sub}, nil
}

// WatchStreamedData is a free log subscription operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerHub *WorkerHubFilterer) WatchStreamedData(opts *bind.WatchOpts, sink chan<- *WorkerHubStreamedData, assignmentId []*big.Int) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubStreamedData)
				if err := _WorkerHub.contract.UnpackLog(event, "StreamedData", log); err != nil {
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

// ParseStreamedData is a log parse operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_WorkerHub *WorkerHubFilterer) ParseStreamedData(log types.Log) (*WorkerHubStreamedData, error) {
	event := new(WorkerHubStreamedData)
	if err := _WorkerHub.contract.UnpackLog(event, "StreamedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubSubmitDurationIterator is returned from FilterSubmitDuration and is used to iterate over the raw logs and unpacked data for SubmitDuration events raised by the WorkerHub contract.
type WorkerHubSubmitDurationIterator struct {
	Event *WorkerHubSubmitDuration // Event containing the contract specifics and raw log

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
func (it *WorkerHubSubmitDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubSubmitDuration)
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
		it.Event = new(WorkerHubSubmitDuration)
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
func (it *WorkerHubSubmitDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubSubmitDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubSubmitDuration represents a SubmitDuration event raised by the WorkerHub contract.
type WorkerHubSubmitDuration struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitDuration is a free log retrieval operation binding the contract event 0x8c0ac957fb32132ec541e9495c4fe8f1d9fdb4dd19a02e7144659d4b382064f3.
//
// Solidity: event SubmitDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) FilterSubmitDuration(opts *bind.FilterOpts) (*WorkerHubSubmitDurationIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "SubmitDuration")
	if err != nil {
		return nil, err
	}
	return &WorkerHubSubmitDurationIterator{contract: _WorkerHub.contract, event: "SubmitDuration", logs: logs, sub: sub}, nil
}

// WatchSubmitDuration is a free log subscription operation binding the contract event 0x8c0ac957fb32132ec541e9495c4fe8f1d9fdb4dd19a02e7144659d4b382064f3.
//
// Solidity: event SubmitDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) WatchSubmitDuration(opts *bind.WatchOpts, sink chan<- *WorkerHubSubmitDuration) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "SubmitDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubSubmitDuration)
				if err := _WorkerHub.contract.UnpackLog(event, "SubmitDuration", log); err != nil {
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

// ParseSubmitDuration is a log parse operation binding the contract event 0x8c0ac957fb32132ec541e9495c4fe8f1d9fdb4dd19a02e7144659d4b382064f3.
//
// Solidity: event SubmitDuration(uint256 oldTime, uint256 newTime)
func (_WorkerHub *WorkerHubFilterer) ParseSubmitDuration(log types.Log) (*WorkerHubSubmitDuration, error) {
	event := new(WorkerHubSubmitDuration)
	if err := _WorkerHub.contract.UnpackLog(event, "SubmitDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubTopUpInferIterator is returned from FilterTopUpInfer and is used to iterate over the raw logs and unpacked data for TopUpInfer events raised by the WorkerHub contract.
type WorkerHubTopUpInferIterator struct {
	Event *WorkerHubTopUpInfer // Event containing the contract specifics and raw log

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
func (it *WorkerHubTopUpInferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubTopUpInfer)
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
		it.Event = new(WorkerHubTopUpInfer)
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
func (it *WorkerHubTopUpInferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubTopUpInferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubTopUpInfer represents a TopUpInfer event raised by the WorkerHub contract.
type WorkerHubTopUpInfer struct {
	InferenceId *big.Int
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTopUpInfer is a free log retrieval operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) FilterTopUpInfer(opts *bind.FilterOpts, inferenceId []*big.Int, creator []common.Address) (*WorkerHubTopUpInferIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTopUpInferIterator{contract: _WorkerHub.contract, event: "TopUpInfer", logs: logs, sub: sub}, nil
}

// WatchTopUpInfer is a free log subscription operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_WorkerHub *WorkerHubFilterer) WatchTopUpInfer(opts *bind.WatchOpts, sink chan<- *WorkerHubTopUpInfer, inferenceId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubTopUpInfer)
				if err := _WorkerHub.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseTopUpInfer(log types.Log) (*WorkerHubTopUpInfer, error) {
	event := new(WorkerHubTopUpInfer)
	if err := _WorkerHub.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubTransferFeeIterator is returned from FilterTransferFee and is used to iterate over the raw logs and unpacked data for TransferFee events raised by the WorkerHub contract.
type WorkerHubTransferFeeIterator struct {
	Event *WorkerHubTransferFee // Event containing the contract specifics and raw log

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
func (it *WorkerHubTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubTransferFee)
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
		it.Event = new(WorkerHubTransferFee)
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
func (it *WorkerHubTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubTransferFee represents a TransferFee event raised by the WorkerHub contract.
type WorkerHubTransferFee struct {
	Treasury       common.Address
	TreasuryFee    *big.Int
	L2OwnerAddress common.Address
	L2OwnerFee     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferFee is a free log retrieval operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed treasury, uint256 treasuryFee, address L2OwnerAddress, uint256 L2OwnerFee)
func (_WorkerHub *WorkerHubFilterer) FilterTransferFee(opts *bind.FilterOpts, treasury []common.Address) (*WorkerHubTransferFeeIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "TransferFee", treasuryRule)
	if err != nil {
		return nil, err
	}
	return &WorkerHubTransferFeeIterator{contract: _WorkerHub.contract, event: "TransferFee", logs: logs, sub: sub}, nil
}

// WatchTransferFee is a free log subscription operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed treasury, uint256 treasuryFee, address L2OwnerAddress, uint256 L2OwnerFee)
func (_WorkerHub *WorkerHubFilterer) WatchTransferFee(opts *bind.WatchOpts, sink chan<- *WorkerHubTransferFee, treasury []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "TransferFee", treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubTransferFee)
				if err := _WorkerHub.contract.UnpackLog(event, "TransferFee", log); err != nil {
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
// Solidity: event TransferFee(address indexed treasury, uint256 treasuryFee, address L2OwnerAddress, uint256 L2OwnerFee)
func (_WorkerHub *WorkerHubFilterer) ParseTransferFee(log types.Log) (*WorkerHubTransferFee, error) {
	event := new(WorkerHubTransferFee)
	if err := _WorkerHub.contract.UnpackLog(event, "TransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkerHubTreasuryAddressUpdatedIterator is returned from FilterTreasuryAddressUpdated and is used to iterate over the raw logs and unpacked data for TreasuryAddressUpdated events raised by the WorkerHub contract.
type WorkerHubTreasuryAddressUpdatedIterator struct {
	Event *WorkerHubTreasuryAddressUpdated // Event containing the contract specifics and raw log

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
func (it *WorkerHubTreasuryAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubTreasuryAddressUpdated)
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
		it.Event = new(WorkerHubTreasuryAddressUpdated)
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
func (it *WorkerHubTreasuryAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubTreasuryAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubTreasuryAddressUpdated represents a TreasuryAddressUpdated event raised by the WorkerHub contract.
type WorkerHubTreasuryAddressUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTreasuryAddressUpdated is a free log retrieval operation binding the contract event 0x430359a6d97ced2b6f93c77a91e7ce9dfd43252eb91e916adba170485cd8a6a4.
//
// Solidity: event TreasuryAddressUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) FilterTreasuryAddressUpdated(opts *bind.FilterOpts) (*WorkerHubTreasuryAddressUpdatedIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "TreasuryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &WorkerHubTreasuryAddressUpdatedIterator{contract: _WorkerHub.contract, event: "TreasuryAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryAddressUpdated is a free log subscription operation binding the contract event 0x430359a6d97ced2b6f93c77a91e7ce9dfd43252eb91e916adba170485cd8a6a4.
//
// Solidity: event TreasuryAddressUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) WatchTreasuryAddressUpdated(opts *bind.WatchOpts, sink chan<- *WorkerHubTreasuryAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "TreasuryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubTreasuryAddressUpdated)
				if err := _WorkerHub.contract.UnpackLog(event, "TreasuryAddressUpdated", log); err != nil {
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

// ParseTreasuryAddressUpdated is a log parse operation binding the contract event 0x430359a6d97ced2b6f93c77a91e7ce9dfd43252eb91e916adba170485cd8a6a4.
//
// Solidity: event TreasuryAddressUpdated(address oldAddress, address newAddress)
func (_WorkerHub *WorkerHubFilterer) ParseTreasuryAddressUpdated(log types.Log) (*WorkerHubTreasuryAddressUpdated, error) {
	event := new(WorkerHubTreasuryAddressUpdated)
	if err := _WorkerHub.contract.UnpackLog(event, "TreasuryAddressUpdated", log); err != nil {
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

// WorkerHubUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the WorkerHub contract.
type WorkerHubUnstakeDelayTimeIterator struct {
	Event *WorkerHubUnstakeDelayTime // Event containing the contract specifics and raw log

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
func (it *WorkerHubUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkerHubUnstakeDelayTime)
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
		it.Event = new(WorkerHubUnstakeDelayTime)
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
func (it *WorkerHubUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkerHubUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkerHubUnstakeDelayTime represents a UnstakeDelayTime event raised by the WorkerHub contract.
type WorkerHubUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHub *WorkerHubFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*WorkerHubUnstakeDelayTimeIterator, error) {

	logs, sub, err := _WorkerHub.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &WorkerHubUnstakeDelayTimeIterator{contract: _WorkerHub.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_WorkerHub *WorkerHubFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *WorkerHubUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _WorkerHub.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkerHubUnstakeDelayTime)
				if err := _WorkerHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
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
func (_WorkerHub *WorkerHubFilterer) ParseUnstakeDelayTime(log types.Log) (*WorkerHubUnstakeDelayTime, error) {
	event := new(WorkerHubUnstakeDelayTime)
	if err := _WorkerHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
