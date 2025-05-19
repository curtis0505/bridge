// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractstaking

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

// ContractstakingMetaData contains all meta data concerning the Contractstaking contract.
var ContractstakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockupDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyRedeemReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"UpdateBonusChef\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextRewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextBlockNumber\",\"type\":\"uint256\"}],\"name\":\"UpdateNextRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"UpdateRewardBar\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIStBonusChef\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRedeemReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"enterStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getStakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStRewardBar\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"leaveStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextRewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIStBonusChef\",\"name\":\"bonusChef\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bpid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardBar\",\"outputs\":[{\"internalType\":\"contractIStRewardBar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"contractIStBonusChef\",\"name\":\"_bonusChef\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bpid\",\"type\":\"uint256\"}],\"name\":\"updateBonusChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextRewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextBlockNumber\",\"type\":\"uint256\"}],\"name\":\"updateNextRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStRewardBar\",\"name\":\"_rewardBar\",\"type\":\"address\"}],\"name\":\"updateRewardBar\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractstakingABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractstakingMetaData.ABI instead.
var ContractstakingABI = ContractstakingMetaData.ABI

// Contractstaking is an auto generated Go binding around an Ethereum contract.
type Contractstaking struct {
	ContractstakingCaller     // Read-only binding to the contract
	ContractstakingTransactor // Write-only binding to the contract
	ContractstakingFilterer   // Log filterer for contract events
}

// ContractstakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractstakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractstakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractstakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractstakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractstakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractstakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractstakingSession struct {
	Contract     *Contractstaking  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractstakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractstakingCallerSession struct {
	Contract *ContractstakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractstakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractstakingTransactorSession struct {
	Contract     *ContractstakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractstakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractstakingRaw struct {
	Contract *Contractstaking // Generic contract binding to access the raw methods on
}

// ContractstakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractstakingCallerRaw struct {
	Contract *ContractstakingCaller // Generic read-only contract binding to access the raw methods on
}

// ContractstakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractstakingTransactorRaw struct {
	Contract *ContractstakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractstaking creates a new instance of Contractstaking, bound to a specific deployed contract.
func NewContractstaking(address common.Address, backend bind.ContractBackend) (*Contractstaking, error) {
	contract, err := bindContractstaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractstaking{ContractstakingCaller: ContractstakingCaller{contract: contract}, ContractstakingTransactor: ContractstakingTransactor{contract: contract}, ContractstakingFilterer: ContractstakingFilterer{contract: contract}}, nil
}

// NewContractstakingCaller creates a new read-only instance of Contractstaking, bound to a specific deployed contract.
func NewContractstakingCaller(address common.Address, caller bind.ContractCaller) (*ContractstakingCaller, error) {
	contract, err := bindContractstaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractstakingCaller{contract: contract}, nil
}

// NewContractstakingTransactor creates a new write-only instance of Contractstaking, bound to a specific deployed contract.
func NewContractstakingTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractstakingTransactor, error) {
	contract, err := bindContractstaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractstakingTransactor{contract: contract}, nil
}

// NewContractstakingFilterer creates a new log filterer instance of Contractstaking, bound to a specific deployed contract.
func NewContractstakingFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractstakingFilterer, error) {
	contract, err := bindContractstaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractstakingFilterer{contract: contract}, nil
}

// bindContractstaking binds a generic wrapper to an already deployed contract.
func bindContractstaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractstakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractstaking *ContractstakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractstaking.Contract.ContractstakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractstaking *ContractstakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractstaking.Contract.ContractstakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractstaking *ContractstakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractstaking.Contract.ContractstakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractstaking *ContractstakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractstaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractstaking *ContractstakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractstaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractstaking *ContractstakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractstaking.Contract.contract.Transact(opts, method, params...)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) pure returns(uint256)
func (_Contractstaking *ContractstakingCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) pure returns(uint256)
func (_Contractstaking *ContractstakingSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Contractstaking.Contract.GetMultiplier(&_Contractstaking.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) pure returns(uint256)
func (_Contractstaking *ContractstakingCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Contractstaking.Contract.GetMultiplier(&_Contractstaking.CallOpts, _from, _to)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x5e534428.
//
// Solidity: function getStakeToken(uint256 _pid) view returns(address)
func (_Contractstaking *ContractstakingCaller) GetStakeToken(opts *bind.CallOpts, _pid *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "getStakeToken", _pid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakeToken is a free data retrieval call binding the contract method 0x5e534428.
//
// Solidity: function getStakeToken(uint256 _pid) view returns(address)
func (_Contractstaking *ContractstakingSession) GetStakeToken(_pid *big.Int) (common.Address, error) {
	return _Contractstaking.Contract.GetStakeToken(&_Contractstaking.CallOpts, _pid)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x5e534428.
//
// Solidity: function getStakeToken(uint256 _pid) view returns(address)
func (_Contractstaking *ContractstakingCallerSession) GetStakeToken(_pid *big.Int) (common.Address, error) {
	return _Contractstaking.Contract.GetStakeToken(&_Contractstaking.CallOpts, _pid)
}

// GetUserAmount is a free data retrieval call binding the contract method 0x3437586b.
//
// Solidity: function getUserAmount(uint256 _pid, address _user) view returns(uint256)
func (_Contractstaking *ContractstakingCaller) GetUserAmount(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "getUserAmount", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAmount is a free data retrieval call binding the contract method 0x3437586b.
//
// Solidity: function getUserAmount(uint256 _pid, address _user) view returns(uint256)
func (_Contractstaking *ContractstakingSession) GetUserAmount(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Contractstaking.Contract.GetUserAmount(&_Contractstaking.CallOpts, _pid, _user)
}

// GetUserAmount is a free data retrieval call binding the contract method 0x3437586b.
//
// Solidity: function getUserAmount(uint256 _pid, address _user) view returns(uint256)
func (_Contractstaking *ContractstakingCallerSession) GetUserAmount(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Contractstaking.Contract.GetUserAmount(&_Contractstaking.CallOpts, _pid, _user)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contractstaking *ContractstakingCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contractstaking *ContractstakingSession) Initialized() (bool, error) {
	return _Contractstaking.Contract.Initialized(&_Contractstaking.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contractstaking *ContractstakingCallerSession) Initialized() (bool, error) {
	return _Contractstaking.Contract.Initialized(&_Contractstaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractstaking *ContractstakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractstaking *ContractstakingSession) Owner() (common.Address, error) {
	return _Contractstaking.Contract.Owner(&_Contractstaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractstaking *ContractstakingCallerSession) Owner() (common.Address, error) {
	return _Contractstaking.Contract.Owner(&_Contractstaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contractstaking *ContractstakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contractstaking *ContractstakingSession) Paused() (bool, error) {
	return _Contractstaking.Contract.Paused(&_Contractstaking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contractstaking *ContractstakingCallerSession) Paused() (bool, error) {
	return _Contractstaking.Contract.Paused(&_Contractstaking.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_Contractstaking *ContractstakingCaller) PendingReward(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "pendingReward", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_Contractstaking *ContractstakingSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Contractstaking.Contract.PendingReward(&_Contractstaking.CallOpts, _pid, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_Contractstaking *ContractstakingCallerSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Contractstaking.Contract.PendingReward(&_Contractstaking.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, address rewardToken, uint256 lastRewardBlock, uint256 rewardPerBlock, uint256 accRewardPerShare, uint256 nextRewardPerBlock, uint256 nextBlockNumber, uint256 lockupDuration, address bonusChef, uint256 bpid)
func (_Contractstaking *ContractstakingCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Contractstaking.contract.Call(opts, &out, "poolInfo", arg0)

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
func (_Contractstaking *ContractstakingSession) PoolInfo(arg0 *big.Int) (struct {
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
	return _Contractstaking.Contract.PoolInfo(&_Contractstaking.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address stakeToken, address rewardToken, uint256 lastRewardBlock, uint256 rewardPerBlock, uint256 accRewardPerShare, uint256 nextRewardPerBlock, uint256 nextBlockNumber, uint256 lockupDuration, address bonusChef, uint256 bpid)
func (_Contractstaking *ContractstakingCallerSession) PoolInfo(arg0 *big.Int) (struct {
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
	return _Contractstaking.Contract.PoolInfo(&_Contractstaking.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Contractstaking *ContractstakingCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Contractstaking *ContractstakingSession) PoolLength() (*big.Int, error) {
	return _Contractstaking.Contract.PoolLength(&_Contractstaking.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Contractstaking *ContractstakingCallerSession) PoolLength() (*big.Int, error) {
	return _Contractstaking.Contract.PoolLength(&_Contractstaking.CallOpts)
}

// RewardBar is a free data retrieval call binding the contract method 0x3fbe9c5d.
//
// Solidity: function rewardBar() view returns(address)
func (_Contractstaking *ContractstakingCaller) RewardBar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "rewardBar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardBar is a free data retrieval call binding the contract method 0x3fbe9c5d.
//
// Solidity: function rewardBar() view returns(address)
func (_Contractstaking *ContractstakingSession) RewardBar() (common.Address, error) {
	return _Contractstaking.Contract.RewardBar(&_Contractstaking.CallOpts)
}

// RewardBar is a free data retrieval call binding the contract method 0x3fbe9c5d.
//
// Solidity: function rewardBar() view returns(address)
func (_Contractstaking *ContractstakingCallerSession) RewardBar() (common.Address, error) {
	return _Contractstaking.Contract.RewardBar(&_Contractstaking.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 claimedReward, uint256 unlockBlockNumber)
func (_Contractstaking *ContractstakingCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimedReward     *big.Int
	UnlockBlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _Contractstaking.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_Contractstaking *ContractstakingSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimedReward     *big.Int
	UnlockBlockNumber *big.Int
}, error) {
	return _Contractstaking.Contract.UserInfo(&_Contractstaking.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 claimedReward, uint256 unlockBlockNumber)
func (_Contractstaking *ContractstakingCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	ClaimedReward     *big.Int
	UnlockBlockNumber *big.Int
}, error) {
	return _Contractstaking.Contract.UserInfo(&_Contractstaking.CallOpts, arg0, arg1)
}

// AddPool is a paid mutator transaction binding the contract method 0xccd861fd.
//
// Solidity: function addPool(address _stakeToken, address _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid) returns()
func (_Contractstaking *ContractstakingTransactor) AddPool(opts *bind.TransactOpts, _stakeToken common.Address, _rewardToken common.Address, _startBlock *big.Int, _rewardPerBlock *big.Int, _lockupDuration *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "addPool", _stakeToken, _rewardToken, _startBlock, _rewardPerBlock, _lockupDuration, _bonusChef, _bpid)
}

// AddPool is a paid mutator transaction binding the contract method 0xccd861fd.
//
// Solidity: function addPool(address _stakeToken, address _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid) returns()
func (_Contractstaking *ContractstakingSession) AddPool(_stakeToken common.Address, _rewardToken common.Address, _startBlock *big.Int, _rewardPerBlock *big.Int, _lockupDuration *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.AddPool(&_Contractstaking.TransactOpts, _stakeToken, _rewardToken, _startBlock, _rewardPerBlock, _lockupDuration, _bonusChef, _bpid)
}

// AddPool is a paid mutator transaction binding the contract method 0xccd861fd.
//
// Solidity: function addPool(address _stakeToken, address _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid) returns()
func (_Contractstaking *ContractstakingTransactorSession) AddPool(_stakeToken common.Address, _rewardToken common.Address, _startBlock *big.Int, _rewardPerBlock *big.Int, _lockupDuration *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.AddPool(&_Contractstaking.TransactOpts, _stakeToken, _rewardToken, _startBlock, _rewardPerBlock, _lockupDuration, _bonusChef, _bpid)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.Deposit(&_Contractstaking.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.Deposit(&_Contractstaking.TransactOpts, _pid, _amount)
}

// EmergencyRedeemReward is a paid mutator transaction binding the contract method 0x5e174f14.
//
// Solidity: function emergencyRedeemReward(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactor) EmergencyRedeemReward(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "emergencyRedeemReward", _pid, _amount)
}

// EmergencyRedeemReward is a paid mutator transaction binding the contract method 0x5e174f14.
//
// Solidity: function emergencyRedeemReward(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingSession) EmergencyRedeemReward(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.EmergencyRedeemReward(&_Contractstaking.TransactOpts, _pid, _amount)
}

// EmergencyRedeemReward is a paid mutator transaction binding the contract method 0x5e174f14.
//
// Solidity: function emergencyRedeemReward(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactorSession) EmergencyRedeemReward(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.EmergencyRedeemReward(&_Contractstaking.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Contractstaking *ContractstakingTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Contractstaking *ContractstakingSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.EmergencyWithdraw(&_Contractstaking.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Contractstaking *ContractstakingTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.EmergencyWithdraw(&_Contractstaking.TransactOpts, _pid)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactor) EnterStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "enterStaking", _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Contractstaking *ContractstakingSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.EnterStaking(&_Contractstaking.TransactOpts, _amount)
}

// EnterStaking is a paid mutator transaction binding the contract method 0x41441d3b.
//
// Solidity: function enterStaking(uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactorSession) EnterStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.EnterStaking(&_Contractstaking.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rewardBar) returns()
func (_Contractstaking *ContractstakingTransactor) Initialize(opts *bind.TransactOpts, _rewardBar common.Address) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "initialize", _rewardBar)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rewardBar) returns()
func (_Contractstaking *ContractstakingSession) Initialize(_rewardBar common.Address) (*types.Transaction, error) {
	return _Contractstaking.Contract.Initialize(&_Contractstaking.TransactOpts, _rewardBar)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _rewardBar) returns()
func (_Contractstaking *ContractstakingTransactorSession) Initialize(_rewardBar common.Address) (*types.Transaction, error) {
	return _Contractstaking.Contract.Initialize(&_Contractstaking.TransactOpts, _rewardBar)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactor) LeaveStaking(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "leaveStaking", _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Contractstaking *ContractstakingSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.LeaveStaking(&_Contractstaking.TransactOpts, _amount)
}

// LeaveStaking is a paid mutator transaction binding the contract method 0x1058d281.
//
// Solidity: function leaveStaking(uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactorSession) LeaveStaking(_amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.LeaveStaking(&_Contractstaking.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractstaking *ContractstakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractstaking *ContractstakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contractstaking.Contract.RenounceOwnership(&_Contractstaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contractstaking *ContractstakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contractstaking.Contract.RenounceOwnership(&_Contractstaking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractstaking *ContractstakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractstaking *ContractstakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contractstaking.Contract.TransferOwnership(&_Contractstaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contractstaking *ContractstakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contractstaking.Contract.TransferOwnership(&_Contractstaking.TransactOpts, newOwner)
}

// UpdateBonusChef is a paid mutator transaction binding the contract method 0xc2a287ce.
//
// Solidity: function updateBonusChef(uint256 _pid, address _bonusChef, uint256 _bpid) returns()
func (_Contractstaking *ContractstakingTransactor) UpdateBonusChef(opts *bind.TransactOpts, _pid *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "updateBonusChef", _pid, _bonusChef, _bpid)
}

// UpdateBonusChef is a paid mutator transaction binding the contract method 0xc2a287ce.
//
// Solidity: function updateBonusChef(uint256 _pid, address _bonusChef, uint256 _bpid) returns()
func (_Contractstaking *ContractstakingSession) UpdateBonusChef(_pid *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdateBonusChef(&_Contractstaking.TransactOpts, _pid, _bonusChef, _bpid)
}

// UpdateBonusChef is a paid mutator transaction binding the contract method 0xc2a287ce.
//
// Solidity: function updateBonusChef(uint256 _pid, address _bonusChef, uint256 _bpid) returns()
func (_Contractstaking *ContractstakingTransactorSession) UpdateBonusChef(_pid *big.Int, _bonusChef common.Address, _bpid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdateBonusChef(&_Contractstaking.TransactOpts, _pid, _bonusChef, _bpid)
}

// UpdateNextRewardPerBlock is a paid mutator transaction binding the contract method 0xb74b998e.
//
// Solidity: function updateNextRewardPerBlock(uint256 _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber) returns()
func (_Contractstaking *ContractstakingTransactor) UpdateNextRewardPerBlock(opts *bind.TransactOpts, _pid *big.Int, _nextRewardPerBlock *big.Int, _nextBlockNumber *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "updateNextRewardPerBlock", _pid, _nextRewardPerBlock, _nextBlockNumber)
}

// UpdateNextRewardPerBlock is a paid mutator transaction binding the contract method 0xb74b998e.
//
// Solidity: function updateNextRewardPerBlock(uint256 _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber) returns()
func (_Contractstaking *ContractstakingSession) UpdateNextRewardPerBlock(_pid *big.Int, _nextRewardPerBlock *big.Int, _nextBlockNumber *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdateNextRewardPerBlock(&_Contractstaking.TransactOpts, _pid, _nextRewardPerBlock, _nextBlockNumber)
}

// UpdateNextRewardPerBlock is a paid mutator transaction binding the contract method 0xb74b998e.
//
// Solidity: function updateNextRewardPerBlock(uint256 _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber) returns()
func (_Contractstaking *ContractstakingTransactorSession) UpdateNextRewardPerBlock(_pid *big.Int, _nextRewardPerBlock *big.Int, _nextBlockNumber *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdateNextRewardPerBlock(&_Contractstaking.TransactOpts, _pid, _nextRewardPerBlock, _nextBlockNumber)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Contractstaking *ContractstakingTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Contractstaking *ContractstakingSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdatePool(&_Contractstaking.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Contractstaking *ContractstakingTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdatePool(&_Contractstaking.TransactOpts, _pid)
}

// UpdateRewardBar is a paid mutator transaction binding the contract method 0x99b4ca9b.
//
// Solidity: function updateRewardBar(address _rewardBar) returns()
func (_Contractstaking *ContractstakingTransactor) UpdateRewardBar(opts *bind.TransactOpts, _rewardBar common.Address) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "updateRewardBar", _rewardBar)
}

// UpdateRewardBar is a paid mutator transaction binding the contract method 0x99b4ca9b.
//
// Solidity: function updateRewardBar(address _rewardBar) returns()
func (_Contractstaking *ContractstakingSession) UpdateRewardBar(_rewardBar common.Address) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdateRewardBar(&_Contractstaking.TransactOpts, _rewardBar)
}

// UpdateRewardBar is a paid mutator transaction binding the contract method 0x99b4ca9b.
//
// Solidity: function updateRewardBar(address _rewardBar) returns()
func (_Contractstaking *ContractstakingTransactorSession) UpdateRewardBar(_rewardBar common.Address) (*types.Transaction, error) {
	return _Contractstaking.Contract.UpdateRewardBar(&_Contractstaking.TransactOpts, _rewardBar)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.Withdraw(&_Contractstaking.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Contractstaking *ContractstakingTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Contractstaking.Contract.Withdraw(&_Contractstaking.TransactOpts, _pid, _amount)
}

// ContractstakingAddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the Contractstaking contract.
type ContractstakingAddPoolIterator struct {
	Event *ContractstakingAddPool // Event containing the contract specifics and raw log

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
func (it *ContractstakingAddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingAddPool)
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
		it.Event = new(ContractstakingAddPool)
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
func (it *ContractstakingAddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingAddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingAddPool represents a AddPool event raised by the Contractstaking contract.
type ContractstakingAddPool struct {
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
func (_Contractstaking *ContractstakingFilterer) FilterAddPool(opts *bind.FilterOpts, _stakeToken []common.Address, _rewardToken []common.Address) (*ContractstakingAddPoolIterator, error) {

	var _stakeTokenRule []interface{}
	for _, _stakeTokenItem := range _stakeToken {
		_stakeTokenRule = append(_stakeTokenRule, _stakeTokenItem)
	}
	var _rewardTokenRule []interface{}
	for _, _rewardTokenItem := range _rewardToken {
		_rewardTokenRule = append(_rewardTokenRule, _rewardTokenItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "AddPool", _stakeTokenRule, _rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingAddPoolIterator{contract: _Contractstaking.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x586e58511a6de8dc332e5c7748a03c93fe91c153bcfc475f610aea6a1d07a41d.
//
// Solidity: event AddPool(address indexed _stakeToken, address indexed _rewardToken, uint256 _startBlock, uint256 _rewardPerBlock, uint256 _lockupDuration, address _bonusChef, uint256 _bpid)
func (_Contractstaking *ContractstakingFilterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *ContractstakingAddPool, _stakeToken []common.Address, _rewardToken []common.Address) (event.Subscription, error) {

	var _stakeTokenRule []interface{}
	for _, _stakeTokenItem := range _stakeToken {
		_stakeTokenRule = append(_stakeTokenRule, _stakeTokenItem)
	}
	var _rewardTokenRule []interface{}
	for _, _rewardTokenItem := range _rewardToken {
		_rewardTokenRule = append(_rewardTokenRule, _rewardTokenItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "AddPool", _stakeTokenRule, _rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingAddPool)
				if err := _Contractstaking.contract.UnpackLog(event, "AddPool", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseAddPool(log types.Log) (*ContractstakingAddPool, error) {
	event := new(ContractstakingAddPool)
	if err := _Contractstaking.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the Contractstaking contract.
type ContractstakingClaimRewardIterator struct {
	Event *ContractstakingClaimReward // Event containing the contract specifics and raw log

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
func (it *ContractstakingClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingClaimReward)
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
		it.Event = new(ContractstakingClaimReward)
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
func (it *ContractstakingClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingClaimReward represents a ClaimReward event raised by the Contractstaking contract.
type ContractstakingClaimReward struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) FilterClaimReward(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingClaimRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "ClaimReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingClaimRewardIterator{contract: _Contractstaking.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0xe74e5c9d4ac1fc33412485f18c159a0a391efe287ab3fd271123f30e6bacf4e3.
//
// Solidity: event ClaimReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *ContractstakingClaimReward, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "ClaimReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingClaimReward)
				if err := _Contractstaking.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseClaimReward(log types.Log) (*ContractstakingClaimReward, error) {
	event := new(ContractstakingClaimReward)
	if err := _Contractstaking.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Contractstaking contract.
type ContractstakingDepositIterator struct {
	Event *ContractstakingDeposit // Event containing the contract specifics and raw log

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
func (it *ContractstakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingDeposit)
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
		it.Event = new(ContractstakingDeposit)
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
func (it *ContractstakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingDeposit represents a Deposit event raised by the Contractstaking contract.
type ContractstakingDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingDepositIterator{contract: _Contractstaking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractstakingDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingDeposit)
				if err := _Contractstaking.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseDeposit(log types.Log) (*ContractstakingDeposit, error) {
	event := new(ContractstakingDeposit)
	if err := _Contractstaking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingEmergencyRedeemRewardIterator is returned from FilterEmergencyRedeemReward and is used to iterate over the raw logs and unpacked data for EmergencyRedeemReward events raised by the Contractstaking contract.
type ContractstakingEmergencyRedeemRewardIterator struct {
	Event *ContractstakingEmergencyRedeemReward // Event containing the contract specifics and raw log

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
func (it *ContractstakingEmergencyRedeemRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingEmergencyRedeemReward)
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
		it.Event = new(ContractstakingEmergencyRedeemReward)
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
func (it *ContractstakingEmergencyRedeemRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingEmergencyRedeemRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingEmergencyRedeemReward represents a EmergencyRedeemReward event raised by the Contractstaking contract.
type ContractstakingEmergencyRedeemReward struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyRedeemReward is a free log retrieval operation binding the contract event 0x81b07abdd46d57aedef714770bd3b6999f6815998451356717c9f06041eb5ca3.
//
// Solidity: event EmergencyRedeemReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) FilterEmergencyRedeemReward(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingEmergencyRedeemRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "EmergencyRedeemReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingEmergencyRedeemRewardIterator{contract: _Contractstaking.contract, event: "EmergencyRedeemReward", logs: logs, sub: sub}, nil
}

// WatchEmergencyRedeemReward is a free log subscription operation binding the contract event 0x81b07abdd46d57aedef714770bd3b6999f6815998451356717c9f06041eb5ca3.
//
// Solidity: event EmergencyRedeemReward(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) WatchEmergencyRedeemReward(opts *bind.WatchOpts, sink chan<- *ContractstakingEmergencyRedeemReward, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "EmergencyRedeemReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingEmergencyRedeemReward)
				if err := _Contractstaking.contract.UnpackLog(event, "EmergencyRedeemReward", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseEmergencyRedeemReward(log types.Log) (*ContractstakingEmergencyRedeemReward, error) {
	event := new(ContractstakingEmergencyRedeemReward)
	if err := _Contractstaking.contract.UnpackLog(event, "EmergencyRedeemReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Contractstaking contract.
type ContractstakingEmergencyWithdrawIterator struct {
	Event *ContractstakingEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractstakingEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingEmergencyWithdraw)
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
		it.Event = new(ContractstakingEmergencyWithdraw)
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
func (it *ContractstakingEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingEmergencyWithdraw represents a EmergencyWithdraw event raised by the Contractstaking contract.
type ContractstakingEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingEmergencyWithdrawIterator{contract: _Contractstaking.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *ContractstakingEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingEmergencyWithdraw)
				if err := _Contractstaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseEmergencyWithdraw(log types.Log) (*ContractstakingEmergencyWithdraw, error) {
	event := new(ContractstakingEmergencyWithdraw)
	if err := _Contractstaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Contractstaking contract.
type ContractstakingInitializeIterator struct {
	Event *ContractstakingInitialize // Event containing the contract specifics and raw log

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
func (it *ContractstakingInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingInitialize)
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
		it.Event = new(ContractstakingInitialize)
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
func (it *ContractstakingInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingInitialize represents a Initialize event raised by the Contractstaking contract.
type ContractstakingInitialize struct {
	RewardBar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed _rewardBar)
func (_Contractstaking *ContractstakingFilterer) FilterInitialize(opts *bind.FilterOpts, _rewardBar []common.Address) (*ContractstakingInitializeIterator, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "Initialize", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingInitializeIterator{contract: _Contractstaking.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x36b1453565f45af7b509b59d5e2eac8f1948ea9e3e193db6663b4101afb6382c.
//
// Solidity: event Initialize(address indexed _rewardBar)
func (_Contractstaking *ContractstakingFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *ContractstakingInitialize, _rewardBar []common.Address) (event.Subscription, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "Initialize", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingInitialize)
				if err := _Contractstaking.contract.UnpackLog(event, "Initialize", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseInitialize(log types.Log) (*ContractstakingInitialize, error) {
	event := new(ContractstakingInitialize)
	if err := _Contractstaking.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contractstaking contract.
type ContractstakingOwnershipTransferredIterator struct {
	Event *ContractstakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractstakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingOwnershipTransferred)
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
		it.Event = new(ContractstakingOwnershipTransferred)
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
func (it *ContractstakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingOwnershipTransferred represents a OwnershipTransferred event raised by the Contractstaking contract.
type ContractstakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractstaking *ContractstakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractstakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingOwnershipTransferredIterator{contract: _Contractstaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contractstaking *ContractstakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractstakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingOwnershipTransferred)
				if err := _Contractstaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseOwnershipTransferred(log types.Log) (*ContractstakingOwnershipTransferred, error) {
	event := new(ContractstakingOwnershipTransferred)
	if err := _Contractstaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contractstaking contract.
type ContractstakingPausedIterator struct {
	Event *ContractstakingPaused // Event containing the contract specifics and raw log

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
func (it *ContractstakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingPaused)
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
		it.Event = new(ContractstakingPaused)
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
func (it *ContractstakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingPaused represents a Paused event raised by the Contractstaking contract.
type ContractstakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contractstaking *ContractstakingFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractstakingPausedIterator, error) {

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractstakingPausedIterator{contract: _Contractstaking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contractstaking *ContractstakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractstakingPaused) (event.Subscription, error) {

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingPaused)
				if err := _Contractstaking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParsePaused(log types.Log) (*ContractstakingPaused, error) {
	event := new(ContractstakingPaused)
	if err := _Contractstaking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Contractstaking contract.
type ContractstakingUnpausedIterator struct {
	Event *ContractstakingUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractstakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingUnpaused)
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
		it.Event = new(ContractstakingUnpaused)
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
func (it *ContractstakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingUnpaused represents a Unpaused event raised by the Contractstaking contract.
type ContractstakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contractstaking *ContractstakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractstakingUnpausedIterator, error) {

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractstakingUnpausedIterator{contract: _Contractstaking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contractstaking *ContractstakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractstakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingUnpaused)
				if err := _Contractstaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseUnpaused(log types.Log) (*ContractstakingUnpaused, error) {
	event := new(ContractstakingUnpaused)
	if err := _Contractstaking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingUpdateBonusChefIterator is returned from FilterUpdateBonusChef and is used to iterate over the raw logs and unpacked data for UpdateBonusChef events raised by the Contractstaking contract.
type ContractstakingUpdateBonusChefIterator struct {
	Event *ContractstakingUpdateBonusChef // Event containing the contract specifics and raw log

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
func (it *ContractstakingUpdateBonusChefIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingUpdateBonusChef)
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
		it.Event = new(ContractstakingUpdateBonusChef)
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
func (it *ContractstakingUpdateBonusChefIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingUpdateBonusChefIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingUpdateBonusChef represents a UpdateBonusChef event raised by the Contractstaking contract.
type ContractstakingUpdateBonusChef struct {
	Pid       *big.Int
	BonusChef common.Address
	Bpid      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateBonusChef is a free log retrieval operation binding the contract event 0x86414693d0e345d8ceaa07867d3849691a7d0f2faa82c3a2e468d087dffcfb4f.
//
// Solidity: event UpdateBonusChef(uint256 indexed _pid, address _bonusChef, uint256 _bpid)
func (_Contractstaking *ContractstakingFilterer) FilterUpdateBonusChef(opts *bind.FilterOpts, _pid []*big.Int) (*ContractstakingUpdateBonusChefIterator, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "UpdateBonusChef", _pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingUpdateBonusChefIterator{contract: _Contractstaking.contract, event: "UpdateBonusChef", logs: logs, sub: sub}, nil
}

// WatchUpdateBonusChef is a free log subscription operation binding the contract event 0x86414693d0e345d8ceaa07867d3849691a7d0f2faa82c3a2e468d087dffcfb4f.
//
// Solidity: event UpdateBonusChef(uint256 indexed _pid, address _bonusChef, uint256 _bpid)
func (_Contractstaking *ContractstakingFilterer) WatchUpdateBonusChef(opts *bind.WatchOpts, sink chan<- *ContractstakingUpdateBonusChef, _pid []*big.Int) (event.Subscription, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "UpdateBonusChef", _pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingUpdateBonusChef)
				if err := _Contractstaking.contract.UnpackLog(event, "UpdateBonusChef", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseUpdateBonusChef(log types.Log) (*ContractstakingUpdateBonusChef, error) {
	event := new(ContractstakingUpdateBonusChef)
	if err := _Contractstaking.contract.UnpackLog(event, "UpdateBonusChef", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingUpdateNextRewardPerBlockIterator is returned from FilterUpdateNextRewardPerBlock and is used to iterate over the raw logs and unpacked data for UpdateNextRewardPerBlock events raised by the Contractstaking contract.
type ContractstakingUpdateNextRewardPerBlockIterator struct {
	Event *ContractstakingUpdateNextRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *ContractstakingUpdateNextRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingUpdateNextRewardPerBlock)
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
		it.Event = new(ContractstakingUpdateNextRewardPerBlock)
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
func (it *ContractstakingUpdateNextRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingUpdateNextRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingUpdateNextRewardPerBlock represents a UpdateNextRewardPerBlock event raised by the Contractstaking contract.
type ContractstakingUpdateNextRewardPerBlock struct {
	Pid                *big.Int
	NextRewardPerBlock *big.Int
	NextBlockNumber    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateNextRewardPerBlock is a free log retrieval operation binding the contract event 0xf003ec729117091675a31e6d8b50921adf145cc483c7a7243eef40c781d4ecd6.
//
// Solidity: event UpdateNextRewardPerBlock(uint256 indexed _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber)
func (_Contractstaking *ContractstakingFilterer) FilterUpdateNextRewardPerBlock(opts *bind.FilterOpts, _pid []*big.Int) (*ContractstakingUpdateNextRewardPerBlockIterator, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "UpdateNextRewardPerBlock", _pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingUpdateNextRewardPerBlockIterator{contract: _Contractstaking.contract, event: "UpdateNextRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchUpdateNextRewardPerBlock is a free log subscription operation binding the contract event 0xf003ec729117091675a31e6d8b50921adf145cc483c7a7243eef40c781d4ecd6.
//
// Solidity: event UpdateNextRewardPerBlock(uint256 indexed _pid, uint256 _nextRewardPerBlock, uint256 _nextBlockNumber)
func (_Contractstaking *ContractstakingFilterer) WatchUpdateNextRewardPerBlock(opts *bind.WatchOpts, sink chan<- *ContractstakingUpdateNextRewardPerBlock, _pid []*big.Int) (event.Subscription, error) {

	var _pidRule []interface{}
	for _, _pidItem := range _pid {
		_pidRule = append(_pidRule, _pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "UpdateNextRewardPerBlock", _pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingUpdateNextRewardPerBlock)
				if err := _Contractstaking.contract.UnpackLog(event, "UpdateNextRewardPerBlock", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseUpdateNextRewardPerBlock(log types.Log) (*ContractstakingUpdateNextRewardPerBlock, error) {
	event := new(ContractstakingUpdateNextRewardPerBlock)
	if err := _Contractstaking.contract.UnpackLog(event, "UpdateNextRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingUpdateRewardBarIterator is returned from FilterUpdateRewardBar and is used to iterate over the raw logs and unpacked data for UpdateRewardBar events raised by the Contractstaking contract.
type ContractstakingUpdateRewardBarIterator struct {
	Event *ContractstakingUpdateRewardBar // Event containing the contract specifics and raw log

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
func (it *ContractstakingUpdateRewardBarIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingUpdateRewardBar)
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
		it.Event = new(ContractstakingUpdateRewardBar)
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
func (it *ContractstakingUpdateRewardBarIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingUpdateRewardBarIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingUpdateRewardBar represents a UpdateRewardBar event raised by the Contractstaking contract.
type ContractstakingUpdateRewardBar struct {
	RewardBar common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateRewardBar is a free log retrieval operation binding the contract event 0x47901f23c1e60500b7e858760fa90a5e4abd09779c564b37494cbcc269e8a76f.
//
// Solidity: event UpdateRewardBar(address indexed _rewardBar)
func (_Contractstaking *ContractstakingFilterer) FilterUpdateRewardBar(opts *bind.FilterOpts, _rewardBar []common.Address) (*ContractstakingUpdateRewardBarIterator, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "UpdateRewardBar", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingUpdateRewardBarIterator{contract: _Contractstaking.contract, event: "UpdateRewardBar", logs: logs, sub: sub}, nil
}

// WatchUpdateRewardBar is a free log subscription operation binding the contract event 0x47901f23c1e60500b7e858760fa90a5e4abd09779c564b37494cbcc269e8a76f.
//
// Solidity: event UpdateRewardBar(address indexed _rewardBar)
func (_Contractstaking *ContractstakingFilterer) WatchUpdateRewardBar(opts *bind.WatchOpts, sink chan<- *ContractstakingUpdateRewardBar, _rewardBar []common.Address) (event.Subscription, error) {

	var _rewardBarRule []interface{}
	for _, _rewardBarItem := range _rewardBar {
		_rewardBarRule = append(_rewardBarRule, _rewardBarItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "UpdateRewardBar", _rewardBarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingUpdateRewardBar)
				if err := _Contractstaking.contract.UnpackLog(event, "UpdateRewardBar", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseUpdateRewardBar(log types.Log) (*ContractstakingUpdateRewardBar, error) {
	event := new(ContractstakingUpdateRewardBar)
	if err := _Contractstaking.contract.UnpackLog(event, "UpdateRewardBar", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractstakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contractstaking contract.
type ContractstakingWithdrawIterator struct {
	Event *ContractstakingWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractstakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractstakingWithdraw)
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
		it.Event = new(ContractstakingWithdraw)
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
func (it *ContractstakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractstakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractstakingWithdraw represents a Withdraw event raised by the Contractstaking contract.
type ContractstakingWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ContractstakingWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ContractstakingWithdrawIterator{contract: _Contractstaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Contractstaking *ContractstakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractstakingWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Contractstaking.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractstakingWithdraw)
				if err := _Contractstaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Contractstaking *ContractstakingFilterer) ParseWithdraw(log types.Log) (*ContractstakingWithdraw, error) {
	event := new(ContractstakingWithdraw)
	if err := _Contractstaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
