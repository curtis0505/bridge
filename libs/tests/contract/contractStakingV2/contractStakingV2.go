// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractstakingV2

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

// ContractstakingV2MetaData contains all meta data concerning the ContractstakingV2 contract.
var ContractstakingV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockupDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyRedeemReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"UpdateBonusChef\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextRewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextBlockNumber\",\"type\":\"uint256\"}],\"name\":\"UpdateNextRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"UpdateRewardBar\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PASS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIStBonusChef\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_originUser\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"depositFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRedeemReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"enterStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_originUser\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"enterStakingFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getStakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStRewardBar\",\"name\":\"_rewardBar\",\"type\":\"address\"},{\"internalType\":\"contractIVerification\",\"name\":\"_verification\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"leaveStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextRewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIStBonusChef\",\"name\":\"bonusChef\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bpid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardBar\",\"outputs\":[{\"internalType\":\"contractIStRewardBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"contractIStBonusChef\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"updateBonusChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextRewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextBlockNumber\",\"type\":\"uint256\"}],\"name\":\"updateNextRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStRewardBar\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"updateRewardBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVerification\",\"name\":\"_verification\",\"type\":\"address\"}],\"name\":\"updateVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractstakingV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractstakingV2MetaData.ABI instead.
var ContractstakingV2ABI = ContractstakingV2MetaData.ABI

// ContractstakingV2 is an auto generated Go binding around an Ethereum contract.
type ContractstakingV2 struct {
	ContractstakingV2Caller     // Read-only binding to the contract
	ContractstakingV2Transactor // Write-only binding to the contract
	ContractstakingV2Filterer   // Log filterer for contract events
}

// ContractstakingV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ContractstakingV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractstakingV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractstakingV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractstakingV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractstakingV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractstakingV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractstakingV2Session struct {
	Contract     *ContractstakingV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractstakingV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractstakingV2CallerSession struct {
	Contract *ContractstakingV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractstakingV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractstakingV2TransactorSession struct {
	Contract     *ContractstakingV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractstakingV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ContractstakingV2Raw struct {
	Contract *ContractstakingV2 // Generic contract binding to access the raw methods on
}

// ContractstakingV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractstakingV2CallerRaw struct {
	Contract *ContractstakingV2Caller // Generic read-only contract binding to access the raw methods on
}

// ContractstakingV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractstakingV2TransactorRaw struct {
	Contract *ContractstakingV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContractstakingV2 creates a new instance of ContractstakingV2, bound to a specific deployed contract.
func NewContractstakingV2(address common.Address, backend bind.ContractBackend) (*ContractstakingV2, error) {
	contract, err := bindContractstakingV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2{ContractstakingV2Caller: ContractstakingV2Caller{contract: contract}, ContractstakingV2Transactor: ContractstakingV2Transactor{contract: contract}, ContractstakingV2Filterer: ContractstakingV2Filterer{contract: contract}}, nil
}

// NewContractstakingV2Caller creates a new read-only instance of ContractstakingV2, bound to a specific deployed contract.
func NewContractstakingV2Caller(address common.Address, caller bind.ContractCaller) (*ContractstakingV2Caller, error) {
	contract, err := bindContractstakingV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2Caller{contract: contract}, nil
}

// NewContractstakingV2Transactor creates a new write-only instance of ContractstakingV2, bound to a specific deployed contract.
func NewContractstakingV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ContractstakingV2Transactor, error) {
	contract, err := bindContractstakingV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2Transactor{contract: contract}, nil
}

// NewContractstakingV2Filterer creates a new log filterer instance of ContractstakingV2, bound to a specific deployed contract.
func NewContractstakingV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ContractstakingV2Filterer, error) {
	contract, err := bindContractstakingV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2Filterer{contract: contract}, nil
}

// bindContractstakingV2 binds a generic wrapper to an already deployed contract.
func bindContractstakingV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractstakingV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractstakingV2 *ContractstakingV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractstakingV2.Contract.ContractstakingV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractstakingV2 *ContractstakingV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.ContractstakingV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractstakingV2 *ContractstakingV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.ContractstakingV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractstakingV2 *ContractstakingV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractstakingV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractstakingV2 *ContractstakingV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractstakingV2 *ContractstakingV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractstakingV2.Contract.DEFAULTADMINROLE(&_ContractstakingV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractstakingV2.Contract.DEFAULTADMINROLE(&_ContractstakingV2.CallOpts)
}

// PASSROLE is a free data retrieval call binding the contract method 0xc8804c55.
//
// Solidity: function PASS_ROLE() view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2Caller) PASSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "PASS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PASSROLE is a free data retrieval call binding the contract method 0xc8804c55.
//
// Solidity: function PASS_ROLE() view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2Session) PASSROLE() ([32]byte, error) {
	return _ContractstakingV2.Contract.PASSROLE(&_ContractstakingV2.CallOpts)
}

// PASSROLE is a free data retrieval call binding the contract method 0xc8804c55.
//
// Solidity: function PASS_ROLE() view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2CallerSession) PASSROLE() ([32]byte, error) {
	return _ContractstakingV2.Contract.PASSROLE(&_ContractstakingV2.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) pure returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Caller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) pure returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Session) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _ContractstakingV2.Contract.GetMultiplier(&_ContractstakingV2.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) pure returns(uint256)
func (_ContractstakingV2 *ContractstakingV2CallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _ContractstakingV2.Contract.GetMultiplier(&_ContractstakingV2.CallOpts, _from, _to)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractstakingV2.Contract.GetRoleAdmin(&_ContractstakingV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractstakingV2 *ContractstakingV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractstakingV2.Contract.GetRoleAdmin(&_ContractstakingV2.CallOpts, role)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x5e534428.
//
// Solidity: function getStakeToken(uint256 _pid) view returns(address)
func (_ContractstakingV2 *ContractstakingV2Caller) GetStakeToken(opts *bind.CallOpts, _pid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "getStakeToken", _pid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakeToken is a free data retrieval call binding the contract method 0x5e534428.
//
// Solidity: function getStakeToken(uint256 _pid) view returns(address)
func (_ContractstakingV2 *ContractstakingV2Session) GetStakeToken(_pid *big.Int) (common.Address, error) {
	return _ContractstakingV2.Contract.GetStakeToken(&_ContractstakingV2.CallOpts, _pid)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x5e534428.
//
// Solidity: function getStakeToken(uint256 _pid) view returns(address)
func (_ContractstakingV2 *ContractstakingV2CallerSession) GetStakeToken(_pid *big.Int) (common.Address, error) {
	return _ContractstakingV2.Contract.GetStakeToken(&_ContractstakingV2.CallOpts, _pid)
}

// GetUserAmount is a free data retrieval call binding the contract method 0x3437586b.
//
// Solidity: function getUserAmount(uint256 _pid, address _user) view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Caller) GetUserAmount(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "getUserAmount", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAmount is a free data retrieval call binding the contract method 0x3437586b.
//
// Solidity: function getUserAmount(uint256 _pid, address _user) view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Session) GetUserAmount(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _ContractstakingV2.Contract.GetUserAmount(&_ContractstakingV2.CallOpts, _pid, _user)
}

// GetUserAmount is a free data retrieval call binding the contract method 0x3437586b.
//
// Solidity: function getUserAmount(uint256 _pid, address _user) view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2CallerSession) GetUserAmount(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _ContractstakingV2.Contract.GetUserAmount(&_ContractstakingV2.CallOpts, _pid, _user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractstakingV2.Contract.HasRole(&_ContractstakingV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractstakingV2 *ContractstakingV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractstakingV2.Contract.HasRole(&_ContractstakingV2.CallOpts, role, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Caller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Session) Initialized() (bool, error) {
	return _ContractstakingV2.Contract.Initialized(&_ContractstakingV2.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ContractstakingV2 *ContractstakingV2CallerSession) Initialized() (bool, error) {
	return _ContractstakingV2.Contract.Initialized(&_ContractstakingV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractstakingV2 *ContractstakingV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractstakingV2 *ContractstakingV2Session) Owner() (common.Address, error) {
	return _ContractstakingV2.Contract.Owner(&_ContractstakingV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractstakingV2 *ContractstakingV2CallerSession) Owner() (common.Address, error) {
	return _ContractstakingV2.Contract.Owner(&_ContractstakingV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Session) Paused() (bool, error) {
	return _ContractstakingV2.Contract.Paused(&_ContractstakingV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractstakingV2 *ContractstakingV2CallerSession) Paused() (bool, error) {
	return _ContractstakingV2.Contract.Paused(&_ContractstakingV2.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Caller) PendingReward(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "pendingReward", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Session) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _ContractstakingV2.Contract.PendingReward(&_ContractstakingV2.CallOpts, _pid, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2CallerSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _ContractstakingV2.Contract.PendingReward(&_ContractstakingV2.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, address rewardToken, uint256 lastRewardBlock, uint256 rewardPerBlock, uint256 accRewardPerShare, uint256 nextRewardPerBlock, uint256 nextBlockNumber, uint256 lockupDuration, address bonusChef, uint256 bpid)
func (_ContractstakingV2 *ContractstakingV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakeToken         common.Address
	RewardToken        common.Address
	LastRewardBlock    *big.Int
	RewardPerBlock     *big.Int
	AccRewardPerShare  *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	LockupDuration     *big.Int
	BonusChef          common.Address
	Bpid               *big.Int
}, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		StakeToken         common.Address
		RewardToken        common.Address
		LastRewardBlock    *big.Int
		RewardPerBlock     *big.Int
		AccRewardPerShare  *big.Int
		NextRewardPerBlock *big.Int
		NextBlockNumber    *big.Int
		LockupDuration     *big.Int
		BonusChef          common.Address
		Bpid               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardPerBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccRewardPerShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NextRewardPerBlock = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.NextBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LockupDuration = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.BonusChef = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.Bpid = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, address rewardToken, uint256 lastRewardBlock, uint256 rewardPerBlock, uint256 accRewardPerShare, uint256 nextRewardPerBlock, uint256 nextBlockNumber, uint256 lockupDuration, address bonusChef, uint256 bpid)
func (_ContractstakingV2 *ContractstakingV2Session) PoolInfo(arg0 *big.Int) (struct {
	StakeToken         common.Address
	RewardToken        common.Address
	LastRewardBlock    *big.Int
	RewardPerBlock     *big.Int
	AccRewardPerShare  *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	LockupDuration     *big.Int
	BonusChef          common.Address
	Bpid               *big.Int
}, error) {
	return _ContractstakingV2.Contract.PoolInfo(&_ContractstakingV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, address rewardToken, uint256 lastRewardBlock, uint256 rewardPerBlock, uint256 accRewardPerShare, uint256 nextRewardPerBlock, uint256 nextBlockNumber, uint256 lockupDuration, address bonusChef, uint256 bpid)
func (_ContractstakingV2 *ContractstakingV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	StakeToken         common.Address
	RewardToken        common.Address
	LastRewardBlock    *big.Int
	RewardPerBlock     *big.Int
	AccRewardPerShare  *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	LockupDuration     *big.Int
	BonusChef          common.Address
	Bpid               *big.Int
}, error) {
	return _ContractstakingV2.Contract.PoolInfo(&_ContractstakingV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2Session) PoolLength() (*big.Int, error) {
	return _ContractstakingV2.Contract.PoolLength(&_ContractstakingV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ContractstakingV2 *ContractstakingV2CallerSession) PoolLength() (*big.Int, error) {
	return _ContractstakingV2.Contract.PoolLength(&_ContractstakingV2.CallOpts)
}

// RewardBar is a free data retrieval call binding the contract method 0x3fbe9c5d.
//
// Solidity: function rewardBar() view returns(address)
func (_ContractstakingV2 *ContractstakingV2Caller) RewardBar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "rewardBar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardBar is a free data retrieval call binding the contract method 0x3fbe9c5d.
//
// Solidity: function rewardBar() view returns(address)
func (_ContractstakingV2 *ContractstakingV2Session) RewardBar() (common.Address, error) {
	return _ContractstakingV2.Contract.RewardBar(&_ContractstakingV2.CallOpts)
}

// RewardBar is a free data retrieval call binding the contract method 0x3fbe9c5d.
//
// Solidity: function rewardBar() view returns(address)
func (_ContractstakingV2 *ContractstakingV2CallerSession) RewardBar() (common.Address, error) {
	return _ContractstakingV2.Contract.RewardBar(&_ContractstakingV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractstakingV2 *ContractstakingV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractstakingV2.Contract.SupportsInterface(&_ContractstakingV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractstakingV2 *ContractstakingV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractstakingV2.Contract.SupportsInterface(&_ContractstakingV2.CallOpts, interfaceId)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 claimedReward, uint256 unlockBlockNumber)
func (_ContractstakingV2 *ContractstakingV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimedReward     *big.Int
	UnlockBlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _ContractstakingV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount            *big.Int
		RewardDebt        *big.Int
		ClaimedReward     *big.Int
		UnlockBlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnlockBlockNumber = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 claimedReward, uint256 unlockBlockNumber)
func (_ContractstakingV2 *ContractstakingV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimedReward     *big.Int
	UnlockBlockNumber *big.Int
}, error) {
	return _ContractstakingV2.Contract.UserInfo(&_ContractstakingV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 claimedReward, uint256 unlockBlockNumber)
func (_ContractstakingV2 *ContractstakingV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimedReward     *big.Int
	UnlockBlockNumber *big.Int
}, error) {
	return _ContractstakingV2.Contract.UserInfo(&_ContractstakingV2.CallOpts, arg0, arg1)
}

// AddPool is a paid mutator transaction binding the contract method 0xccd861fd.
//
// Solidity: function addPool(address _stakeToken, address _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) AddPool(opts *bind.TransactOpts, _stakeToken common.Address, _rewardToken common.Address, _startBlock *big.Int, _rewardPerBlock *big.Int, _lockupDuration *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "addPool", _stakeToken, _rewardToken, _startBlock, _rewardPerBlock, _lockupDuration, _bonusChef, _bpid)
}

// AddPool is a paid mutator transaction binding the contract method 0xccd861fd.
//
// Solidity: function addPool(address _stakeToken, address _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid) returns()
func (_ContractstakingV2 *ContractstakingV2Session) AddPool(_stakeToken common.Address, _rewardToken common.Address, _startBlock *big.Int, _rewardPerBlock *big.Int, _lockupDuration *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.AddPool(&_ContractstakingV2.TransactOpts, _stakeToken, _rewardToken, _startBlock, _rewardPerBlock, _lockupDuration, _bonusChef, _bpid)
}

// AddPool is a paid mutator transaction binding the contract method 0xccd861fd.
//
// Solidity: function addPool(address _stakeToken, address _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) AddPool(_stakeToken common.Address, _rewardToken common.Address, _startBlock *big.Int, _rewardPerBlock *big.Int, _lockupDuration *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.AddPool(&_ContractstakingV2.TransactOpts, _stakeToken, _rewardToken, _startBlock, _rewardPerBlock, _lockupDuration, _bonusChef, _bpid)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "deposit", _pid, _amount, _proof)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) Deposit(_pid *big.Int, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Deposit(&_ContractstakingV2.TransactOpts, _pid, _amount, _proof)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Deposit(&_ContractstakingV2.TransactOpts, _pid, _amount, _proof)
}

// DepositFrom is a paid mutator transaction binding the contract method 0x73adabb3.
//
// Solidity: function depositFrom(uint256 _pid, uint256 _amount, address _originUser, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) DepositFrom(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _originUser common.Address, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "depositFrom", _pid, _amount, _originUser, _proof)
}

// DepositFrom is a paid mutator transaction binding the contract method 0x73adabb3.
//
// Solidity: function depositFrom(uint256 _pid, uint256 _amount, address _originUser, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) DepositFrom(_pid *big.Int, _amount *big.Int, _originUser common.Address, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.DepositFrom(&_ContractstakingV2.TransactOpts, _pid, _amount, _originUser, _proof)
}

// DepositFrom is a paid mutator transaction binding the contract method 0x73adabb3.
//
// Solidity: function depositFrom(uint256 _pid, uint256 _amount, address _originUser, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) DepositFrom(_pid *big.Int, _amount *big.Int, _originUser common.Address, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.DepositFrom(&_ContractstakingV2.TransactOpts, _pid, _amount, _originUser, _proof)
}

// EmergencyRedeemReward is a paid mutator transaction binding the contract method 0x5e174f14.
//
// Solidity: function emergencyRedeemReward(uint256 _pid, uint256 _amount) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) EmergencyRedeemReward(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "emergencyRedeemReward", _pid, _amount)
}

// EmergencyRedeemReward is a paid mutator transaction binding the contract method 0x5e174f14.
//
// Solidity: function emergencyRedeemReward(uint256 _pid, uint256 _amount) returns()
func (_ContractstakingV2 *ContractstakingV2Session) EmergencyRedeemReward(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EmergencyRedeemReward(&_ContractstakingV2.TransactOpts, _pid, _amount)
}

// EmergencyRedeemReward is a paid mutator transaction binding the contract method 0x5e174f14.
//
// Solidity: function emergencyRedeemReward(uint256 _pid, uint256 _amount) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) EmergencyRedeemReward(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EmergencyRedeemReward(&_ContractstakingV2.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xb355bf62.
//
// Solidity: function emergencyWithdraw(uint256 _pid, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "emergencyWithdraw", _pid, _proof)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xb355bf62.
//
// Solidity: function emergencyWithdraw(uint256 _pid, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) EmergencyWithdraw(_pid *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EmergencyWithdraw(&_ContractstakingV2.TransactOpts, _pid, _proof)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xb355bf62.
//
// Solidity: function emergencyWithdraw(uint256 _pid, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) EmergencyWithdraw(_pid *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EmergencyWithdraw(&_ContractstakingV2.TransactOpts, _pid, _proof)
}

// EnterStaking is a paid mutator transaction binding the contract method 0xc1f9e58d.
//
// Solidity: function enterStaking(uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) EnterStaking(opts *bind.TransactOpts, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "enterStaking", _amount, _proof)
}

// EnterStaking is a paid mutator transaction binding the contract method 0xc1f9e58d.
//
// Solidity: function enterStaking(uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) EnterStaking(_amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EnterStaking(&_ContractstakingV2.TransactOpts, _amount, _proof)
}

// EnterStaking is a paid mutator transaction binding the contract method 0xc1f9e58d.
//
// Solidity: function enterStaking(uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) EnterStaking(_amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EnterStaking(&_ContractstakingV2.TransactOpts, _amount, _proof)
}

// EnterStakingFrom is a paid mutator transaction binding the contract method 0x86ea1974.
//
// Solidity: function enterStakingFrom(uint256 _amount, address _originUser, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) EnterStakingFrom(opts *bind.TransactOpts, _amount *big.Int, _originUser common.Address, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "enterStakingFrom", _amount, _originUser, _proof)
}

// EnterStakingFrom is a paid mutator transaction binding the contract method 0x86ea1974.
//
// Solidity: function enterStakingFrom(uint256 _amount, address _originUser, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) EnterStakingFrom(_amount *big.Int, _originUser common.Address, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EnterStakingFrom(&_ContractstakingV2.TransactOpts, _amount, _originUser, _proof)
}

// EnterStakingFrom is a paid mutator transaction binding the contract method 0x86ea1974.
//
// Solidity: function enterStakingFrom(uint256 _amount, address _originUser, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) EnterStakingFrom(_amount *big.Int, _originUser common.Address, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.EnterStakingFrom(&_ContractstakingV2.TransactOpts, _amount, _originUser, _proof)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.GrantRole(&_ContractstakingV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.GrantRole(&_ContractstakingV2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rewardBar, address _verification) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) Initialize(opts *bind.TransactOpts, _rewardBar common.Address, _verification common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "initialize", _rewardBar, _verification)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rewardBar, address _verification) returns()
func (_ContractstakingV2 *ContractstakingV2Session) Initialize(_rewardBar common.Address, _verification common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Initialize(&_ContractstakingV2.TransactOpts, _rewardBar, _verification)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _rewardBar, address _verification) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) Initialize(_rewardBar common.Address, _verification common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Initialize(&_ContractstakingV2.TransactOpts, _rewardBar, _verification)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x77eaa191.
//
// Solidity: function leaveStaking(uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) LeaveStaking(opts *bind.TransactOpts, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "leaveStaking", _amount, _proof)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x77eaa191.
//
// Solidity: function leaveStaking(uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) LeaveStaking(_amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.LeaveStaking(&_ContractstakingV2.TransactOpts, _amount, _proof)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x77eaa191.
//
// Solidity: function leaveStaking(uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) LeaveStaking(_amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.LeaveStaking(&_ContractstakingV2.TransactOpts, _amount, _proof)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractstakingV2 *ContractstakingV2Session) Pause() (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Pause(&_ContractstakingV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) Pause() (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Pause(&_ContractstakingV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractstakingV2 *ContractstakingV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _ContractstakingV2.Contract.RenounceOwnership(&_ContractstakingV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractstakingV2.Contract.RenounceOwnership(&_ContractstakingV2.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.RenounceRole(&_ContractstakingV2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.RenounceRole(&_ContractstakingV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.RevokeRole(&_ContractstakingV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.RevokeRole(&_ContractstakingV2.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractstakingV2 *ContractstakingV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.TransferOwnership(&_ContractstakingV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.TransferOwnership(&_ContractstakingV2.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractstakingV2 *ContractstakingV2Session) Unpause() (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Unpause(&_ContractstakingV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Unpause(&_ContractstakingV2.TransactOpts)
}

// UpdateBonusChef is a paid mutator transaction binding the contract method 0xc2a287ce.
//
// Solidity: function updateBonusChef(uint256 _pid, address _bonusChef, uint256 _bpid) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) UpdateBonusChef(opts *bind.TransactOpts, _pid *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "updateBonusChef", _pid, _bonusChef, _bpid)
}

// UpdateBonusChef is a paid mutator transaction binding the contract method 0xc2a287ce.
//
// Solidity: function updateBonusChef(uint256 _pid, address _bonusChef, uint256 _bpid) returns()
func (_ContractstakingV2 *ContractstakingV2Session) UpdateBonusChef(_pid *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateBonusChef(&_ContractstakingV2.TransactOpts, _pid, _bonusChef, _bpid)
}

// UpdateBonusChef is a paid mutator transaction binding the contract method 0xc2a287ce.
//
// Solidity: function updateBonusChef(uint256 _pid, address _bonusChef, uint256 _bpid) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) UpdateBonusChef(_pid *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateBonusChef(&_ContractstakingV2.TransactOpts, _pid, _bonusChef, _bpid)
}

// UpdateNextRewardPerBlock is a paid mutator transaction binding the contract method 0xb74b998e.
//
// Solidity: function updateNextRewardPerBlock(uint256 _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) UpdateNextRewardPerBlock(opts *bind.TransactOpts, _pid *big.Int, _nextRewardPerBlock *big.Int, _nextBlockNumber *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "updateNextRewardPerBlock", _pid, _nextRewardPerBlock, _nextBlockNumber)
}

// UpdateNextRewardPerBlock is a paid mutator transaction binding the contract method 0xb74b998e.
//
// Solidity: function updateNextRewardPerBlock(uint256 _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber) returns()
func (_ContractstakingV2 *ContractstakingV2Session) UpdateNextRewardPerBlock(_pid *big.Int, _nextRewardPerBlock *big.Int, _nextBlockNumber *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateNextRewardPerBlock(&_ContractstakingV2.TransactOpts, _pid, _nextRewardPerBlock, _nextBlockNumber)
}

// UpdateNextRewardPerBlock is a paid mutator transaction binding the contract method 0xb74b998e.
//
// Solidity: function updateNextRewardPerBlock(uint256 _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) UpdateNextRewardPerBlock(_pid *big.Int, _nextRewardPerBlock *big.Int, _nextBlockNumber *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateNextRewardPerBlock(&_ContractstakingV2.TransactOpts, _pid, _nextRewardPerBlock, _nextBlockNumber)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_ContractstakingV2 *ContractstakingV2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdatePool(&_ContractstakingV2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdatePool(&_ContractstakingV2.TransactOpts, _pid)
}

// UpdateRewardBar is a paid mutator transaction binding the contract method 0x99b4ca9b.
//
// Solidity: function updateRewardBar(address _rewardBar) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) UpdateRewardBar(opts *bind.TransactOpts, _rewardBar common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "updateRewardBar", _rewardBar)
}

// UpdateRewardBar is a paid mutator transaction binding the contract method 0x99b4ca9b.
//
// Solidity: function updateRewardBar(address _rewardBar) returns()
func (_ContractstakingV2 *ContractstakingV2Session) UpdateRewardBar(_rewardBar common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateRewardBar(&_ContractstakingV2.TransactOpts, _rewardBar)
}

// UpdateRewardBar is a paid mutator transaction binding the contract method 0x99b4ca9b.
//
// Solidity: function updateRewardBar(address _rewardBar) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) UpdateRewardBar(_rewardBar common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateRewardBar(&_ContractstakingV2.TransactOpts, _rewardBar)
}

// UpdateVerification is a paid mutator transaction binding the contract method 0x4714a411.
//
// Solidity: function updateVerification(address _verification) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) UpdateVerification(opts *bind.TransactOpts, _verification common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "updateVerification", _verification)
}

// UpdateVerification is a paid mutator transaction binding the contract method 0x4714a411.
//
// Solidity: function updateVerification(address _verification) returns()
func (_ContractstakingV2 *ContractstakingV2Session) UpdateVerification(_verification common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateVerification(&_ContractstakingV2.TransactOpts, _verification)
}

// UpdateVerification is a paid mutator transaction binding the contract method 0x4714a411.
//
// Solidity: function updateVerification(address _verification) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) UpdateVerification(_verification common.Address) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.UpdateVerification(&_ContractstakingV2.TransactOpts, _verification)
}

// Withdraw is a paid mutator transaction binding the contract method 0x744fb6ca.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.contract.Transact(opts, "withdraw", _pid, _amount, _proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x744fb6ca.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2Session) Withdraw(_pid *big.Int, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Withdraw(&_ContractstakingV2.TransactOpts, _pid, _amount, _proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x744fb6ca.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount, bytes _proof) returns()
func (_ContractstakingV2 *ContractstakingV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int, _proof []byte) (*types.Transaction, error) {
	return _ContractstakingV2.Contract.Withdraw(&_ContractstakingV2.TransactOpts, _pid, _amount, _proof)
}

// ContractstakingV2AddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the ContractstakingV2 contract.
type ContractstakingV2AddPoolIterator struct {
	Event *ContractstakingV2AddPool // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2AddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2AddPool)
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
		it.Event = new(ContractstakingV2AddPool)
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
func (it *ContractstakingV2AddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2AddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2AddPool represents a AddPool event raised by the ContractstakingV2 contract.
type ContractstakingV2AddPool struct {
	StakeToken     common.Address
	RewardToken    common.Address
	StartBlock     *big.Int
	RewardPerBlock *big.Int
	LockupDuration *big.Int
	BonusChef      common.Address
	Bpid           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x586e58511a6de8dc332e5c7748a03c93fe91c153bcfc475f610aea6a1d07a41d.
//
// Solidity: event AddPool(address indexed _stakeToken, address indexed _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterAddPool(opts *bind.FilterOpts, _stakeToken []common.Address, _rewardToken []common.Address) (*ContractstakingV2AddPoolIterator, error) {

	var _stakeTokenRule []interface{}
	for _, _stakeTokenItem := range _stakeToken {
		_stakeTokenRule = append(_stakeTokenRule, _stakeTokenItem)
	}
	var _rewardTokenRule []interface{}
	for _, _rewardTokenItem := range _rewardToken {
		_rewardTokenRule = append(_rewardTokenRule, _rewardTokenItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "AddPool", _stakeTokenRule, _rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2AddPoolIterator{contract: _ContractstakingV2.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x586e58511a6de8dc332e5c7748a03c93fe91c153bcfc475f610aea6a1d07a41d.
//
// Solidity: event AddPool(address indexed _stakeToken, address indexed _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *ContractstakingV2AddPool, _stakeToken []common.Address, _rewardToken []common.Address) (event.Subscription, error) {

	var _stakeTokenRule []interface{}
	for _, _stakeTokenItem := range _stakeToken {
		_stakeTokenRule = append(_stakeTokenRule, _stakeTokenItem)
	}
	var _rewardTokenRule []interface{}
	for _, _rewardTokenItem := range _rewardToken {
		_rewardTokenRule = append(_rewardTokenRule, _rewardTokenItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "AddPool", _stakeTokenRule, _rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2AddPool)
				if err := _ContractstakingV2.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0x586e58511a6de8dc332e5c7748a03c93fe91c153bcfc475f610aea6a1d07a41d.
//
// Solidity: event AddPool(address indexed _stakeToken, address indexed _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseAddPool(log types.Log) (*ContractstakingV2AddPool, error) {
	event := new(ContractstakingV2AddPool)
	if err := _ContractstakingV2.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2ClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the ContractstakingV2 contract.
type ContractstakingV2ClaimRewardIterator struct {
	Event *ContractstakingV2ClaimReward // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2ClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2ClaimReward)
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
		it.Event = new(ContractstakingV2ClaimReward)
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
func (it *ContractstakingV2ClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2ClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2ClaimReward represents a ClaimReward event raised by the ContractstakingV2 contract.
type ContractstakingV2ClaimReward struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterClaimReward(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingV2ClaimRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "ClaimReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2ClaimRewardIterator{contract: _ContractstakingV2.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *ContractstakingV2ClaimReward, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "ClaimReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2ClaimReward)
				if err := _ContractstakingV2.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseClaimReward(log types.Log) (*ContractstakingV2ClaimReward, error) {
	event := new(ContractstakingV2ClaimReward)
	if err := _ContractstakingV2.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ContractstakingV2 contract.
type ContractstakingV2DepositIterator struct {
	Event *ContractstakingV2Deposit // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2Deposit)
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
		it.Event = new(ContractstakingV2Deposit)
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
func (it *ContractstakingV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2Deposit represents a Deposit event raised by the ContractstakingV2 contract.
type ContractstakingV2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingV2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2DepositIterator{contract: _ContractstakingV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractstakingV2Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2Deposit)
				if err := _ContractstakingV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseDeposit(log types.Log) (*ContractstakingV2Deposit, error) {
	event := new(ContractstakingV2Deposit)
	if err := _ContractstakingV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2EmergencyRedeemRewardIterator is returned from FilterEmergencyRedeemReward and is used to iterate over the raw logs and unpacked data for EmergencyRedeemReward events raised by the ContractstakingV2 contract.
type ContractstakingV2EmergencyRedeemRewardIterator struct {
	Event *ContractstakingV2EmergencyRedeemReward // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2EmergencyRedeemRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2EmergencyRedeemReward)
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
		it.Event = new(ContractstakingV2EmergencyRedeemReward)
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
func (it *ContractstakingV2EmergencyRedeemRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2EmergencyRedeemRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2EmergencyRedeemReward represents a EmergencyRedeemReward event raised by the ContractstakingV2 contract.
type ContractstakingV2EmergencyRedeemReward struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyRedeemReward is a free log retrieval operation binding the contract event 0x81b07abdd46d57aedef714770bd3b6999f6815998451356717c9f06041eb5ca3.
//
// Solidity: event EmergencyRedeemReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterEmergencyRedeemReward(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingV2EmergencyRedeemRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "EmergencyRedeemReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2EmergencyRedeemRewardIterator{contract: _ContractstakingV2.contract, event: "EmergencyRedeemReward", logs: logs, sub: sub}, nil
}

// WatchEmergencyRedeemReward is a free log subscription operation binding the contract event 0x81b07abdd46d57aedef714770bd3b6999f6815998451356717c9f06041eb5ca3.
//
// Solidity: event EmergencyRedeemReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchEmergencyRedeemReward(opts *bind.WatchOpts, sink chan<- *ContractstakingV2EmergencyRedeemReward, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "EmergencyRedeemReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2EmergencyRedeemReward)
				if err := _ContractstakingV2.contract.UnpackLog(event, "EmergencyRedeemReward", log); err != nil {
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

// ParseEmergencyRedeemReward is a log parse operation binding the contract event 0x81b07abdd46d57aedef714770bd3b6999f6815998451356717c9f06041eb5ca3.
//
// Solidity: event EmergencyRedeemReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseEmergencyRedeemReward(log types.Log) (*ContractstakingV2EmergencyRedeemReward, error) {
	event := new(ContractstakingV2EmergencyRedeemReward)
	if err := _ContractstakingV2.contract.UnpackLog(event, "EmergencyRedeemReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the ContractstakingV2 contract.
type ContractstakingV2EmergencyWithdrawIterator struct {
	Event *ContractstakingV2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2EmergencyWithdraw)
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
		it.Event = new(ContractstakingV2EmergencyWithdraw)
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
func (it *ContractstakingV2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2EmergencyWithdraw represents a EmergencyWithdraw event raised by the ContractstakingV2 contract.
type ContractstakingV2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingV2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2EmergencyWithdrawIterator{contract: _ContractstakingV2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *ContractstakingV2EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2EmergencyWithdraw)
				if err := _ContractstakingV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseEmergencyWithdraw(log types.Log) (*ContractstakingV2EmergencyWithdraw, error) {
	event := new(ContractstakingV2EmergencyWithdraw)
	if err := _ContractstakingV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2InitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the ContractstakingV2 contract.
type ContractstakingV2InitializeIterator struct {
	Event *ContractstakingV2Initialize // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2InitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2Initialize)
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
		it.Event = new(ContractstakingV2Initialize)
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
func (it *ContractstakingV2InitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2InitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2Initialize represents a Initialize event raised by the ContractstakingV2 contract.
type ContractstakingV2Initialize struct {
	RewardBar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed _rewardBar)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterInitialize(opts *bind.FilterOpts, _rewardBar []common.Address) (*ContractstakingV2InitializeIterator, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "Initialize", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2InitializeIterator{contract: _ContractstakingV2.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed _rewardBar)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *ContractstakingV2Initialize, _rewardBar []common.Address) (event.Subscription, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "Initialize", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2Initialize)
				if err := _ContractstakingV2.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed _rewardBar)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseInitialize(log types.Log) (*ContractstakingV2Initialize, error) {
	event := new(ContractstakingV2Initialize)
	if err := _ContractstakingV2.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractstakingV2 contract.
type ContractstakingV2OwnershipTransferredIterator struct {
	Event *ContractstakingV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2OwnershipTransferred)
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
		it.Event = new(ContractstakingV2OwnershipTransferred)
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
func (it *ContractstakingV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2OwnershipTransferred represents a OwnershipTransferred event raised by the ContractstakingV2 contract.
type ContractstakingV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractstakingV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2OwnershipTransferredIterator{contract: _ContractstakingV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractstakingV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2OwnershipTransferred)
				if err := _ContractstakingV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseOwnershipTransferred(log types.Log) (*ContractstakingV2OwnershipTransferred, error) {
	event := new(ContractstakingV2OwnershipTransferred)
	if err := _ContractstakingV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractstakingV2 contract.
type ContractstakingV2PausedIterator struct {
	Event *ContractstakingV2Paused // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2Paused)
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
		it.Event = new(ContractstakingV2Paused)
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
func (it *ContractstakingV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2Paused represents a Paused event raised by the ContractstakingV2 contract.
type ContractstakingV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterPaused(opts *bind.FilterOpts) (*ContractstakingV2PausedIterator, error) {

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2PausedIterator{contract: _ContractstakingV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractstakingV2Paused) (event.Subscription, error) {

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2Paused)
				if err := _ContractstakingV2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractstakingV2 *ContractstakingV2Filterer) ParsePaused(log types.Log) (*ContractstakingV2Paused, error) {
	event := new(ContractstakingV2Paused)
	if err := _ContractstakingV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractstakingV2 contract.
type ContractstakingV2RoleAdminChangedIterator struct {
	Event *ContractstakingV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2RoleAdminChanged)
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
		it.Event = new(ContractstakingV2RoleAdminChanged)
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
func (it *ContractstakingV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2RoleAdminChanged represents a RoleAdminChanged event raised by the ContractstakingV2 contract.
type ContractstakingV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractstakingV2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2RoleAdminChangedIterator{contract: _ContractstakingV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractstakingV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2RoleAdminChanged)
				if err := _ContractstakingV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseRoleAdminChanged(log types.Log) (*ContractstakingV2RoleAdminChanged, error) {
	event := new(ContractstakingV2RoleAdminChanged)
	if err := _ContractstakingV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractstakingV2 contract.
type ContractstakingV2RoleGrantedIterator struct {
	Event *ContractstakingV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2RoleGranted)
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
		it.Event = new(ContractstakingV2RoleGranted)
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
func (it *ContractstakingV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2RoleGranted represents a RoleGranted event raised by the ContractstakingV2 contract.
type ContractstakingV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractstakingV2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2RoleGrantedIterator{contract: _ContractstakingV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractstakingV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2RoleGranted)
				if err := _ContractstakingV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseRoleGranted(log types.Log) (*ContractstakingV2RoleGranted, error) {
	event := new(ContractstakingV2RoleGranted)
	if err := _ContractstakingV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractstakingV2 contract.
type ContractstakingV2RoleRevokedIterator struct {
	Event *ContractstakingV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2RoleRevoked)
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
		it.Event = new(ContractstakingV2RoleRevoked)
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
func (it *ContractstakingV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2RoleRevoked represents a RoleRevoked event raised by the ContractstakingV2 contract.
type ContractstakingV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractstakingV2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2RoleRevokedIterator{contract: _ContractstakingV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractstakingV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2RoleRevoked)
				if err := _ContractstakingV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseRoleRevoked(log types.Log) (*ContractstakingV2RoleRevoked, error) {
	event := new(ContractstakingV2RoleRevoked)
	if err := _ContractstakingV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractstakingV2 contract.
type ContractstakingV2UnpausedIterator struct {
	Event *ContractstakingV2Unpaused // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2Unpaused)
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
		it.Event = new(ContractstakingV2Unpaused)
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
func (it *ContractstakingV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2Unpaused represents a Unpaused event raised by the ContractstakingV2 contract.
type ContractstakingV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractstakingV2UnpausedIterator, error) {

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2UnpausedIterator{contract: _ContractstakingV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractstakingV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2Unpaused)
				if err := _ContractstakingV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseUnpaused(log types.Log) (*ContractstakingV2Unpaused, error) {
	event := new(ContractstakingV2Unpaused)
	if err := _ContractstakingV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2UpdateBonusChefIterator is returned from FilterUpdateBonusChef and is used to iterate over the raw logs and unpacked data for UpdateBonusChef events raised by the ContractstakingV2 contract.
type ContractstakingV2UpdateBonusChefIterator struct {
	Event *ContractstakingV2UpdateBonusChef // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2UpdateBonusChefIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2UpdateBonusChef)
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
		it.Event = new(ContractstakingV2UpdateBonusChef)
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
func (it *ContractstakingV2UpdateBonusChefIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2UpdateBonusChefIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2UpdateBonusChef represents a UpdateBonusChef event raised by the ContractstakingV2 contract.
type ContractstakingV2UpdateBonusChef struct {
	Pid       *big.Int
	BonusChef common.Address
	Bpid      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateBonusChef is a free log retrieval operation binding the contract event 0x86414693d0e345d8ceaa07867d3849691a7d0f2faa82c3a2e468d087dffcfb4f.
//
// Solidity: event UpdateBonusChef(uint256 indexed _pid, address _bonusChef, uint256 _bpid)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterUpdateBonusChef(opts *bind.FilterOpts, _pid []*big.Int) (*ContractstakingV2UpdateBonusChefIterator, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "UpdateBonusChef", _pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2UpdateBonusChefIterator{contract: _ContractstakingV2.contract, event: "UpdateBonusChef", logs: logs, sub: sub}, nil
}

// WatchUpdateBonusChef is a free log subscription operation binding the contract event 0x86414693d0e345d8ceaa07867d3849691a7d0f2faa82c3a2e468d087dffcfb4f.
//
// Solidity: event UpdateBonusChef(uint256 indexed _pid, address _bonusChef, uint256 _bpid)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchUpdateBonusChef(opts *bind.WatchOpts, sink chan<- *ContractstakingV2UpdateBonusChef, _pid []*big.Int) (event.Subscription, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "UpdateBonusChef", _pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2UpdateBonusChef)
				if err := _ContractstakingV2.contract.UnpackLog(event, "UpdateBonusChef", log); err != nil {
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

// ParseUpdateBonusChef is a log parse operation binding the contract event 0x86414693d0e345d8ceaa07867d3849691a7d0f2faa82c3a2e468d087dffcfb4f.
//
// Solidity: event UpdateBonusChef(uint256 indexed _pid, address _bonusChef, uint256 _bpid)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseUpdateBonusChef(log types.Log) (*ContractstakingV2UpdateBonusChef, error) {
	event := new(ContractstakingV2UpdateBonusChef)
	if err := _ContractstakingV2.contract.UnpackLog(event, "UpdateBonusChef", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2UpdateNextRewardPerBlockIterator is returned from FilterUpdateNextRewardPerBlock and is used to iterate over the raw logs and unpacked data for UpdateNextRewardPerBlock events raised by the ContractstakingV2 contract.
type ContractstakingV2UpdateNextRewardPerBlockIterator struct {
	Event *ContractstakingV2UpdateNextRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2UpdateNextRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2UpdateNextRewardPerBlock)
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
		it.Event = new(ContractstakingV2UpdateNextRewardPerBlock)
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
func (it *ContractstakingV2UpdateNextRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2UpdateNextRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2UpdateNextRewardPerBlock represents a UpdateNextRewardPerBlock event raised by the ContractstakingV2 contract.
type ContractstakingV2UpdateNextRewardPerBlock struct {
	Pid                *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateNextRewardPerBlock is a free log retrieval operation binding the contract event 0xf003ec729117091675a31e6d8b50921adf145cc483c7a7243eef40c781d4ecd6.
//
// Solidity: event UpdateNextRewardPerBlock(uint256 indexed _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterUpdateNextRewardPerBlock(opts *bind.FilterOpts, _pid []*big.Int) (*ContractstakingV2UpdateNextRewardPerBlockIterator, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "UpdateNextRewardPerBlock", _pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2UpdateNextRewardPerBlockIterator{contract: _ContractstakingV2.contract, event: "UpdateNextRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchUpdateNextRewardPerBlock is a free log subscription operation binding the contract event 0xf003ec729117091675a31e6d8b50921adf145cc483c7a7243eef40c781d4ecd6.
//
// Solidity: event UpdateNextRewardPerBlock(uint256 indexed _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchUpdateNextRewardPerBlock(opts *bind.WatchOpts, sink chan<- *ContractstakingV2UpdateNextRewardPerBlock, _pid []*big.Int) (event.Subscription, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "UpdateNextRewardPerBlock", _pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2UpdateNextRewardPerBlock)
				if err := _ContractstakingV2.contract.UnpackLog(event, "UpdateNextRewardPerBlock", log); err != nil {
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

// ParseUpdateNextRewardPerBlock is a log parse operation binding the contract event 0xf003ec729117091675a31e6d8b50921adf145cc483c7a7243eef40c781d4ecd6.
//
// Solidity: event UpdateNextRewardPerBlock(uint256 indexed _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseUpdateNextRewardPerBlock(log types.Log) (*ContractstakingV2UpdateNextRewardPerBlock, error) {
	event := new(ContractstakingV2UpdateNextRewardPerBlock)
	if err := _ContractstakingV2.contract.UnpackLog(event, "UpdateNextRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2UpdateRewardBarIterator is returned from FilterUpdateRewardBar and is used to iterate over the raw logs and unpacked data for UpdateRewardBar events raised by the ContractstakingV2 contract.
type ContractstakingV2UpdateRewardBarIterator struct {
	Event *ContractstakingV2UpdateRewardBar // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2UpdateRewardBarIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2UpdateRewardBar)
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
		it.Event = new(ContractstakingV2UpdateRewardBar)
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
func (it *ContractstakingV2UpdateRewardBarIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2UpdateRewardBarIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2UpdateRewardBar represents a UpdateRewardBar event raised by the ContractstakingV2 contract.
type ContractstakingV2UpdateRewardBar struct {
	RewardBar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateRewardBar is a free log retrieval operation binding the contract event 0x47901f23c1e60500b7e858760fa90a5e4abd09779c564b37494cbcc269e8a76f.
//
// Solidity: event UpdateRewardBar(address indexed _rewardBar)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterUpdateRewardBar(opts *bind.FilterOpts, _rewardBar []common.Address) (*ContractstakingV2UpdateRewardBarIterator, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "UpdateRewardBar", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2UpdateRewardBarIterator{contract: _ContractstakingV2.contract, event: "UpdateRewardBar", logs: logs, sub: sub}, nil
}

// WatchUpdateRewardBar is a free log subscription operation binding the contract event 0x47901f23c1e60500b7e858760fa90a5e4abd09779c564b37494cbcc269e8a76f.
//
// Solidity: event UpdateRewardBar(address indexed _rewardBar)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchUpdateRewardBar(opts *bind.WatchOpts, sink chan<- *ContractstakingV2UpdateRewardBar, _rewardBar []common.Address) (event.Subscription, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "UpdateRewardBar", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2UpdateRewardBar)
				if err := _ContractstakingV2.contract.UnpackLog(event, "UpdateRewardBar", log); err != nil {
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

// ParseUpdateRewardBar is a log parse operation binding the contract event 0x47901f23c1e60500b7e858760fa90a5e4abd09779c564b37494cbcc269e8a76f.
//
// Solidity: event UpdateRewardBar(address indexed _rewardBar)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseUpdateRewardBar(log types.Log) (*ContractstakingV2UpdateRewardBar, error) {
	event := new(ContractstakingV2UpdateRewardBar)
	if err := _ContractstakingV2.contract.UnpackLog(event, "UpdateRewardBar", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ContractstakingV2 contract.
type ContractstakingV2WithdrawIterator struct {
	Event *ContractstakingV2Withdraw // Event containing the contract specifics and raw log

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
func (it *ContractstakingV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingV2Withdraw)
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
		it.Event = new(ContractstakingV2Withdraw)
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
func (it *ContractstakingV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingV2Withdraw represents a Withdraw event raised by the ContractstakingV2 contract.
type ContractstakingV2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingV2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingV2WithdrawIterator{contract: _ContractstakingV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractstakingV2Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ContractstakingV2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingV2Withdraw)
				if err := _ContractstakingV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ContractstakingV2 *ContractstakingV2Filterer) ParseWithdraw(log types.Log) (*ContractstakingV2Withdraw, error) {
	event := new(ContractstakingV2Withdraw)
	if err := _ContractstakingV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
