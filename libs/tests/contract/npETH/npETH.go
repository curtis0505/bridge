// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package npETH

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

// NpETHMetaData contains all meta data concerning the NpETH contract.
var NpETHMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"BurnFeeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawableEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawReqEth\",\"type\":\"uint256\"}],\"name\":\"ClaimStakeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELExitReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELIncludeExitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionLayerRewardsVaultAddress\",\"type\":\"address\"}],\"name\":\"ELRewardsVaultAddressSetEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bufferedEth\",\"type\":\"uint256\"}],\"name\":\"HandleCompensateBufferEthEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"HandleDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountToMint\",\"type\":\"uint256\"}],\"name\":\"HandleElRewardsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawReqAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalExitingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_exitingAmount\",\"type\":\"uint256\"}],\"name\":\"HandleExitValidatorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawReqEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawableEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bufferedEth\",\"type\":\"uint256\"}],\"name\":\"HandleExitedValidatorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawReqEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawableEth\",\"type\":\"uint256\"}],\"name\":\"HandleWithdrawReqEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_cnDepositAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_protocolFeeBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_npElRewardsVaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_verification\",\"type\":\"address\"}],\"name\":\"InitializeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawableEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawReqEth\",\"type\":\"uint256\"}],\"name\":\"RequestClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountNpToBurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawReqEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawReqIndex\",\"type\":\"uint256\"}],\"name\":\"RequestWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldCnDeposit\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCnDeposit\",\"type\":\"address\"}],\"name\":\"SetCnDepositAddressEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newFeeTo\",\"type\":\"address\"}],\"name\":\"SetFeeToEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldLimitDepositLoop\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newLimitDepositLoop\",\"type\":\"uint32\"}],\"name\":\"SetLimitDepositLoopEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProtocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFee\",\"type\":\"uint256\"}],\"name\":\"SetProtocolFeeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVerificatoin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newVerification\",\"type\":\"address\"}],\"name\":\"SetVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountToMint\",\"type\":\"uint256\"}],\"name\":\"StakeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountToMint\",\"type\":\"uint256\"}],\"name\":\"SubmitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountNpToBurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawReqEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawReqIndex\",\"type\":\"uint256\"}],\"name\":\"UnstakeEvent\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_CREDENTIALS_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accBeaconValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accElExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accElRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bufferedEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_npEthAmount\",\"type\":\"uint256\"}],\"name\":\"burnFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountEth\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"claimStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cnDeposit\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInEth\",\"type\":\"uint256\"}],\"name\":\"convertEthToNpEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInNpEth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInNpEth\",\"type\":\"uint256\"}],\"name\":\"convertNpEthToEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInEth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"depositValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumNpEthStorage.STATUS\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"exitValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumNpEthStorage.STATUS\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"handleCompensateBufferEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_withdrawal_credentials\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_deposit_data_roots\",\"type\":\"bytes32[]\"}],\"name\":\"handleDepositValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accBeaconValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beaconBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"handleElRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_exitPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_exitingAmount\",\"type\":\"uint256\"}],\"name\":\"handleExitValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accBeaconValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_beaconBalance\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_exitedPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_exitedAmount\",\"type\":\"uint256\"}],\"name\":\"handleExitedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"handleWithdrawReq\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_cnDepositAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFeeBP\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_npElRewardsVaultAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verification\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitDepositLoop\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"npElRewardsVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveELExit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountEth\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"requestClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newCnDepositAddress\",\"type\":\"address\"}],\"name\":\"setCnDepositAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executionLayerRewardsVaultAddress\",\"type\":\"address\"}],\"name\":\"setELRewardsVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_newLimitDepositLoop\",\"type\":\"uint32\"}],\"name\":\"setLimitDepositLoop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newProtocolFeeBP\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerification\",\"type\":\"address\"}],\"name\":\"setVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalExitingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWithdrawReqAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWithdrawableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountNp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumNpEthStorage.STAKE_TYPE\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawReqAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verification\",\"outputs\":[{\"internalType\":\"contractIVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawReqLog\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumNpEthStorage.STAKE_TYPE\",\"name\":\"stype\",\"type\":\"uint8\"},{\"internalType\":\"enumNpEthStorage.STATUS\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawReqLogLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawReqLogNextStartIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NpETHABI is the input ABI used to generate the binding from.
// Deprecated: Use NpETHMetaData.ABI instead.
var NpETHABI = NpETHMetaData.ABI

// NpETH is an auto generated Go binding around an Ethereum contract.
type NpETH struct {
	NpETHCaller     // Read-only binding to the contract
	NpETHTransactor // Write-only binding to the contract
	NpETHFilterer   // Log filterer for contract events
}

// NpETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type NpETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NpETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NpETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NpETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NpETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NpETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NpETHSession struct {
	Contract     *NpETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NpETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NpETHCallerSession struct {
	Contract *NpETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NpETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NpETHTransactorSession struct {
	Contract     *NpETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NpETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type NpETHRaw struct {
	Contract *NpETH // Generic contract binding to access the raw methods on
}

// NpETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NpETHCallerRaw struct {
	Contract *NpETHCaller // Generic read-only contract binding to access the raw methods on
}

// NpETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NpETHTransactorRaw struct {
	Contract *NpETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNpETH creates a new instance of NpETH, bound to a specific deployed contract.
func NewNpETH(address common.Address, backend bind.ContractBackend) (*NpETH, error) {
	contract, err := bindNpETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NpETH{NpETHCaller: NpETHCaller{contract: contract}, NpETHTransactor: NpETHTransactor{contract: contract}, NpETHFilterer: NpETHFilterer{contract: contract}}, nil
}

// NewNpETHCaller creates a new read-only instance of NpETH, bound to a specific deployed contract.
func NewNpETHCaller(address common.Address, caller bind.ContractCaller) (*NpETHCaller, error) {
	contract, err := bindNpETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NpETHCaller{contract: contract}, nil
}

// NewNpETHTransactor creates a new write-only instance of NpETH, bound to a specific deployed contract.
func NewNpETHTransactor(address common.Address, transactor bind.ContractTransactor) (*NpETHTransactor, error) {
	contract, err := bindNpETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NpETHTransactor{contract: contract}, nil
}

// NewNpETHFilterer creates a new log filterer instance of NpETH, bound to a specific deployed contract.
func NewNpETHFilterer(address common.Address, filterer bind.ContractFilterer) (*NpETHFilterer, error) {
	contract, err := bindNpETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NpETHFilterer{contract: contract}, nil
}

// bindNpETH binds a generic wrapper to an already deployed contract.
func bindNpETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NpETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NpETH *NpETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NpETH.Contract.NpETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NpETH *NpETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.Contract.NpETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NpETH *NpETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NpETH.Contract.NpETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NpETH *NpETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NpETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NpETH *NpETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NpETH *NpETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NpETH.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NpETH *NpETHCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NpETH *NpETHSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NpETH.Contract.DEFAULTADMINROLE(&_NpETH.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NpETH *NpETHCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NpETH.Contract.DEFAULTADMINROLE(&_NpETH.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_NpETH *NpETHCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_NpETH *NpETHSession) DEPOSITSIZE() (*big.Int, error) {
	return _NpETH.Contract.DEPOSITSIZE(&_NpETH.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_NpETH *NpETHCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _NpETH.Contract.DEPOSITSIZE(&_NpETH.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_NpETH *NpETHCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_NpETH *NpETHSession) GOVERNANCEROLE() ([32]byte, error) {
	return _NpETH.Contract.GOVERNANCEROLE(&_NpETH.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_NpETH *NpETHCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _NpETH.Contract.GOVERNANCEROLE(&_NpETH.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_NpETH *NpETHCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_NpETH *NpETHSession) OPERATORROLE() ([32]byte, error) {
	return _NpETH.Contract.OPERATORROLE(&_NpETH.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_NpETH *NpETHCallerSession) OPERATORROLE() ([32]byte, error) {
	return _NpETH.Contract.OPERATORROLE(&_NpETH.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_NpETH *NpETHCaller) PUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_NpETH *NpETHSession) PUBKEYLENGTH() (*big.Int, error) {
	return _NpETH.Contract.PUBKEYLENGTH(&_NpETH.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_NpETH *NpETHCallerSession) PUBKEYLENGTH() (*big.Int, error) {
	return _NpETH.Contract.PUBKEYLENGTH(&_NpETH.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_NpETH *NpETHCaller) SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_NpETH *NpETHSession) SIGNATURELENGTH() (*big.Int, error) {
	return _NpETH.Contract.SIGNATURELENGTH(&_NpETH.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_NpETH *NpETHCallerSession) SIGNATURELENGTH() (*big.Int, error) {
	return _NpETH.Contract.SIGNATURELENGTH(&_NpETH.CallOpts)
}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_NpETH *NpETHCaller) WITHDRAWALCREDENTIALSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "WITHDRAWAL_CREDENTIALS_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_NpETH *NpETHSession) WITHDRAWALCREDENTIALSLENGTH() (*big.Int, error) {
	return _NpETH.Contract.WITHDRAWALCREDENTIALSLENGTH(&_NpETH.CallOpts)
}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_NpETH *NpETHCallerSession) WITHDRAWALCREDENTIALSLENGTH() (*big.Int, error) {
	return _NpETH.Contract.WITHDRAWALCREDENTIALSLENGTH(&_NpETH.CallOpts)
}

// AccBeaconValidators is a free data retrieval call binding the contract method 0x88bb2a8d.
//
// Solidity: function accBeaconValidators() view returns(uint256)
func (_NpETH *NpETHCaller) AccBeaconValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "accBeaconValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccBeaconValidators is a free data retrieval call binding the contract method 0x88bb2a8d.
//
// Solidity: function accBeaconValidators() view returns(uint256)
func (_NpETH *NpETHSession) AccBeaconValidators() (*big.Int, error) {
	return _NpETH.Contract.AccBeaconValidators(&_NpETH.CallOpts)
}

// AccBeaconValidators is a free data retrieval call binding the contract method 0x88bb2a8d.
//
// Solidity: function accBeaconValidators() view returns(uint256)
func (_NpETH *NpETHCallerSession) AccBeaconValidators() (*big.Int, error) {
	return _NpETH.Contract.AccBeaconValidators(&_NpETH.CallOpts)
}

// AccElExit is a free data retrieval call binding the contract method 0xf3516089.
//
// Solidity: function accElExit() view returns(uint256)
func (_NpETH *NpETHCaller) AccElExit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "accElExit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccElExit is a free data retrieval call binding the contract method 0xf3516089.
//
// Solidity: function accElExit() view returns(uint256)
func (_NpETH *NpETHSession) AccElExit() (*big.Int, error) {
	return _NpETH.Contract.AccElExit(&_NpETH.CallOpts)
}

// AccElExit is a free data retrieval call binding the contract method 0xf3516089.
//
// Solidity: function accElExit() view returns(uint256)
func (_NpETH *NpETHCallerSession) AccElExit() (*big.Int, error) {
	return _NpETH.Contract.AccElExit(&_NpETH.CallOpts)
}

// AccElRewards is a free data retrieval call binding the contract method 0x815f34df.
//
// Solidity: function accElRewards() view returns(uint256)
func (_NpETH *NpETHCaller) AccElRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "accElRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccElRewards is a free data retrieval call binding the contract method 0x815f34df.
//
// Solidity: function accElRewards() view returns(uint256)
func (_NpETH *NpETHSession) AccElRewards() (*big.Int, error) {
	return _NpETH.Contract.AccElRewards(&_NpETH.CallOpts)
}

// AccElRewards is a free data retrieval call binding the contract method 0x815f34df.
//
// Solidity: function accElRewards() view returns(uint256)
func (_NpETH *NpETHCallerSession) AccElRewards() (*big.Int, error) {
	return _NpETH.Contract.AccElRewards(&_NpETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NpETH *NpETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NpETH *NpETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NpETH.Contract.Allowance(&_NpETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NpETH *NpETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NpETH.Contract.Allowance(&_NpETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NpETH *NpETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NpETH *NpETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NpETH.Contract.BalanceOf(&_NpETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NpETH *NpETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NpETH.Contract.BalanceOf(&_NpETH.CallOpts, account)
}

// BeaconBalance is a free data retrieval call binding the contract method 0xa8d1f822.
//
// Solidity: function beaconBalance() view returns(uint256)
func (_NpETH *NpETHCaller) BeaconBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "beaconBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeaconBalance is a free data retrieval call binding the contract method 0xa8d1f822.
//
// Solidity: function beaconBalance() view returns(uint256)
func (_NpETH *NpETHSession) BeaconBalance() (*big.Int, error) {
	return _NpETH.Contract.BeaconBalance(&_NpETH.CallOpts)
}

// BeaconBalance is a free data retrieval call binding the contract method 0xa8d1f822.
//
// Solidity: function beaconBalance() view returns(uint256)
func (_NpETH *NpETHCallerSession) BeaconBalance() (*big.Int, error) {
	return _NpETH.Contract.BeaconBalance(&_NpETH.CallOpts)
}

// BufferedEth is a free data retrieval call binding the contract method 0xb8989d1e.
//
// Solidity: function bufferedEth() view returns(uint256)
func (_NpETH *NpETHCaller) BufferedEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "bufferedEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferedEth is a free data retrieval call binding the contract method 0xb8989d1e.
//
// Solidity: function bufferedEth() view returns(uint256)
func (_NpETH *NpETHSession) BufferedEth() (*big.Int, error) {
	return _NpETH.Contract.BufferedEth(&_NpETH.CallOpts)
}

// BufferedEth is a free data retrieval call binding the contract method 0xb8989d1e.
//
// Solidity: function bufferedEth() view returns(uint256)
func (_NpETH *NpETHCallerSession) BufferedEth() (*big.Int, error) {
	return _NpETH.Contract.BufferedEth(&_NpETH.CallOpts)
}

// CnDeposit is a free data retrieval call binding the contract method 0x9297af97.
//
// Solidity: function cnDeposit() view returns(address)
func (_NpETH *NpETHCaller) CnDeposit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "cnDeposit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CnDeposit is a free data retrieval call binding the contract method 0x9297af97.
//
// Solidity: function cnDeposit() view returns(address)
func (_NpETH *NpETHSession) CnDeposit() (common.Address, error) {
	return _NpETH.Contract.CnDeposit(&_NpETH.CallOpts)
}

// CnDeposit is a free data retrieval call binding the contract method 0x9297af97.
//
// Solidity: function cnDeposit() view returns(address)
func (_NpETH *NpETHCallerSession) CnDeposit() (common.Address, error) {
	return _NpETH.Contract.CnDeposit(&_NpETH.CallOpts)
}

// ConvertEthToNpEth is a free data retrieval call binding the contract method 0x91da18c0.
//
// Solidity: function convertEthToNpEth(uint256 _amountInEth) view returns(uint256 amountInNpEth)
func (_NpETH *NpETHCaller) ConvertEthToNpEth(opts *bind.CallOpts, _amountInEth *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "convertEthToNpEth", _amountInEth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertEthToNpEth is a free data retrieval call binding the contract method 0x91da18c0.
//
// Solidity: function convertEthToNpEth(uint256 _amountInEth) view returns(uint256 amountInNpEth)
func (_NpETH *NpETHSession) ConvertEthToNpEth(_amountInEth *big.Int) (*big.Int, error) {
	return _NpETH.Contract.ConvertEthToNpEth(&_NpETH.CallOpts, _amountInEth)
}

// ConvertEthToNpEth is a free data retrieval call binding the contract method 0x91da18c0.
//
// Solidity: function convertEthToNpEth(uint256 _amountInEth) view returns(uint256 amountInNpEth)
func (_NpETH *NpETHCallerSession) ConvertEthToNpEth(_amountInEth *big.Int) (*big.Int, error) {
	return _NpETH.Contract.ConvertEthToNpEth(&_NpETH.CallOpts, _amountInEth)
}

// ConvertNpEthToEth is a free data retrieval call binding the contract method 0x20065f52.
//
// Solidity: function convertNpEthToEth(uint256 _amountInNpEth) view returns(uint256 amountInEth)
func (_NpETH *NpETHCaller) ConvertNpEthToEth(opts *bind.CallOpts, _amountInNpEth *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "convertNpEthToEth", _amountInNpEth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertNpEthToEth is a free data retrieval call binding the contract method 0x20065f52.
//
// Solidity: function convertNpEthToEth(uint256 _amountInNpEth) view returns(uint256 amountInEth)
func (_NpETH *NpETHSession) ConvertNpEthToEth(_amountInNpEth *big.Int) (*big.Int, error) {
	return _NpETH.Contract.ConvertNpEthToEth(&_NpETH.CallOpts, _amountInNpEth)
}

// ConvertNpEthToEth is a free data retrieval call binding the contract method 0x20065f52.
//
// Solidity: function convertNpEthToEth(uint256 _amountInNpEth) view returns(uint256 amountInEth)
func (_NpETH *NpETHCallerSession) ConvertNpEthToEth(_amountInNpEth *big.Int) (*big.Int, error) {
	return _NpETH.Contract.ConvertNpEthToEth(&_NpETH.CallOpts, _amountInNpEth)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NpETH *NpETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NpETH *NpETHSession) Decimals() (uint8, error) {
	return _NpETH.Contract.Decimals(&_NpETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NpETH *NpETHCallerSession) Decimals() (uint8, error) {
	return _NpETH.Contract.Decimals(&_NpETH.CallOpts)
}

// DepositValidators is a free data retrieval call binding the contract method 0xcd7161f2.
//
// Solidity: function depositValidators(bytes ) view returns(uint256 amount, uint8 status)
func (_NpETH *NpETHCaller) DepositValidators(opts *bind.CallOpts, arg0 []byte) (struct {
	Amount *big.Int
	Status uint8
}, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "depositValidators", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Status uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// DepositValidators is a free data retrieval call binding the contract method 0xcd7161f2.
//
// Solidity: function depositValidators(bytes ) view returns(uint256 amount, uint8 status)
func (_NpETH *NpETHSession) DepositValidators(arg0 []byte) (struct {
	Amount *big.Int
	Status uint8
}, error) {
	return _NpETH.Contract.DepositValidators(&_NpETH.CallOpts, arg0)
}

// DepositValidators is a free data retrieval call binding the contract method 0xcd7161f2.
//
// Solidity: function depositValidators(bytes ) view returns(uint256 amount, uint8 status)
func (_NpETH *NpETHCallerSession) DepositValidators(arg0 []byte) (struct {
	Amount *big.Int
	Status uint8
}, error) {
	return _NpETH.Contract.DepositValidators(&_NpETH.CallOpts, arg0)
}

// DepositedValidators is a free data retrieval call binding the contract method 0xbcd8391b.
//
// Solidity: function depositedValidators() view returns(uint256)
func (_NpETH *NpETHCaller) DepositedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "depositedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositedValidators is a free data retrieval call binding the contract method 0xbcd8391b.
//
// Solidity: function depositedValidators() view returns(uint256)
func (_NpETH *NpETHSession) DepositedValidators() (*big.Int, error) {
	return _NpETH.Contract.DepositedValidators(&_NpETH.CallOpts)
}

// DepositedValidators is a free data retrieval call binding the contract method 0xbcd8391b.
//
// Solidity: function depositedValidators() view returns(uint256)
func (_NpETH *NpETHCallerSession) DepositedValidators() (*big.Int, error) {
	return _NpETH.Contract.DepositedValidators(&_NpETH.CallOpts)
}

// ExitValidators is a free data retrieval call binding the contract method 0x8977079f.
//
// Solidity: function exitValidators(bytes ) view returns(uint256 amount, uint8 status)
func (_NpETH *NpETHCaller) ExitValidators(opts *bind.CallOpts, arg0 []byte) (struct {
	Amount *big.Int
	Status uint8
}, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "exitValidators", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Status uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// ExitValidators is a free data retrieval call binding the contract method 0x8977079f.
//
// Solidity: function exitValidators(bytes ) view returns(uint256 amount, uint8 status)
func (_NpETH *NpETHSession) ExitValidators(arg0 []byte) (struct {
	Amount *big.Int
	Status uint8
}, error) {
	return _NpETH.Contract.ExitValidators(&_NpETH.CallOpts, arg0)
}

// ExitValidators is a free data retrieval call binding the contract method 0x8977079f.
//
// Solidity: function exitValidators(bytes ) view returns(uint256 amount, uint8 status)
func (_NpETH *NpETHCallerSession) ExitValidators(arg0 []byte) (struct {
	Amount *big.Int
	Status uint8
}, error) {
	return _NpETH.Contract.ExitValidators(&_NpETH.CallOpts, arg0)
}

// ExitedValidators is a free data retrieval call binding the contract method 0x3882b1c9.
//
// Solidity: function exitedValidators() view returns(uint256)
func (_NpETH *NpETHCaller) ExitedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "exitedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitedValidators is a free data retrieval call binding the contract method 0x3882b1c9.
//
// Solidity: function exitedValidators() view returns(uint256)
func (_NpETH *NpETHSession) ExitedValidators() (*big.Int, error) {
	return _NpETH.Contract.ExitedValidators(&_NpETH.CallOpts)
}

// ExitedValidators is a free data retrieval call binding the contract method 0x3882b1c9.
//
// Solidity: function exitedValidators() view returns(uint256)
func (_NpETH *NpETHCallerSession) ExitedValidators() (*big.Int, error) {
	return _NpETH.Contract.ExitedValidators(&_NpETH.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_NpETH *NpETHCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_NpETH *NpETHSession) FeeTo() (common.Address, error) {
	return _NpETH.Contract.FeeTo(&_NpETH.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_NpETH *NpETHCallerSession) FeeTo() (common.Address, error) {
	return _NpETH.Contract.FeeTo(&_NpETH.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NpETH *NpETHCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NpETH *NpETHSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NpETH.Contract.GetRoleAdmin(&_NpETH.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NpETH *NpETHCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NpETH.Contract.GetRoleAdmin(&_NpETH.CallOpts, role)
}

// GetTotalPooledEth is a free data retrieval call binding the contract method 0x4995d148.
//
// Solidity: function getTotalPooledEth() view returns(uint256)
func (_NpETH *NpETHCaller) GetTotalPooledEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "getTotalPooledEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEth is a free data retrieval call binding the contract method 0x4995d148.
//
// Solidity: function getTotalPooledEth() view returns(uint256)
func (_NpETH *NpETHSession) GetTotalPooledEth() (*big.Int, error) {
	return _NpETH.Contract.GetTotalPooledEth(&_NpETH.CallOpts)
}

// GetTotalPooledEth is a free data retrieval call binding the contract method 0x4995d148.
//
// Solidity: function getTotalPooledEth() view returns(uint256)
func (_NpETH *NpETHCallerSession) GetTotalPooledEth() (*big.Int, error) {
	return _NpETH.Contract.GetTotalPooledEth(&_NpETH.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NpETH *NpETHCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NpETH *NpETHSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NpETH.Contract.HasRole(&_NpETH.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NpETH *NpETHCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NpETH.Contract.HasRole(&_NpETH.CallOpts, role, account)
}

// LimitDepositLoop is a free data retrieval call binding the contract method 0xb8f93995.
//
// Solidity: function limitDepositLoop() view returns(uint32)
func (_NpETH *NpETHCaller) LimitDepositLoop(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "limitDepositLoop")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LimitDepositLoop is a free data retrieval call binding the contract method 0xb8f93995.
//
// Solidity: function limitDepositLoop() view returns(uint32)
func (_NpETH *NpETHSession) LimitDepositLoop() (uint32, error) {
	return _NpETH.Contract.LimitDepositLoop(&_NpETH.CallOpts)
}

// LimitDepositLoop is a free data retrieval call binding the contract method 0xb8f93995.
//
// Solidity: function limitDepositLoop() view returns(uint32)
func (_NpETH *NpETHCallerSession) LimitDepositLoop() (uint32, error) {
	return _NpETH.Contract.LimitDepositLoop(&_NpETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NpETH *NpETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NpETH *NpETHSession) Name() (string, error) {
	return _NpETH.Contract.Name(&_NpETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NpETH *NpETHCallerSession) Name() (string, error) {
	return _NpETH.Contract.Name(&_NpETH.CallOpts)
}

// NpElRewardsVaultAddress is a free data retrieval call binding the contract method 0xc846626d.
//
// Solidity: function npElRewardsVaultAddress() view returns(address)
func (_NpETH *NpETHCaller) NpElRewardsVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "npElRewardsVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NpElRewardsVaultAddress is a free data retrieval call binding the contract method 0xc846626d.
//
// Solidity: function npElRewardsVaultAddress() view returns(address)
func (_NpETH *NpETHSession) NpElRewardsVaultAddress() (common.Address, error) {
	return _NpETH.Contract.NpElRewardsVaultAddress(&_NpETH.CallOpts)
}

// NpElRewardsVaultAddress is a free data retrieval call binding the contract method 0xc846626d.
//
// Solidity: function npElRewardsVaultAddress() view returns(address)
func (_NpETH *NpETHCallerSession) NpElRewardsVaultAddress() (common.Address, error) {
	return _NpETH.Contract.NpElRewardsVaultAddress(&_NpETH.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NpETH *NpETHCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NpETH *NpETHSession) Paused() (bool, error) {
	return _NpETH.Contract.Paused(&_NpETH.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NpETH *NpETHCallerSession) Paused() (bool, error) {
	return _NpETH.Contract.Paused(&_NpETH.CallOpts)
}

// ProtocolFeeBP is a free data retrieval call binding the contract method 0xc5fba0b8.
//
// Solidity: function protocolFeeBP() view returns(uint256)
func (_NpETH *NpETHCaller) ProtocolFeeBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "protocolFeeBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeBP is a free data retrieval call binding the contract method 0xc5fba0b8.
//
// Solidity: function protocolFeeBP() view returns(uint256)
func (_NpETH *NpETHSession) ProtocolFeeBP() (*big.Int, error) {
	return _NpETH.Contract.ProtocolFeeBP(&_NpETH.CallOpts)
}

// ProtocolFeeBP is a free data retrieval call binding the contract method 0xc5fba0b8.
//
// Solidity: function protocolFeeBP() view returns(uint256)
func (_NpETH *NpETHCallerSession) ProtocolFeeBP() (*big.Int, error) {
	return _NpETH.Contract.ProtocolFeeBP(&_NpETH.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_NpETH *NpETHCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "stakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_NpETH *NpETHSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _NpETH.Contract.Stakes(&_NpETH.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_NpETH *NpETHCallerSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _NpETH.Contract.Stakes(&_NpETH.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NpETH *NpETHCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NpETH *NpETHSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NpETH.Contract.SupportsInterface(&_NpETH.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NpETH *NpETHCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NpETH.Contract.SupportsInterface(&_NpETH.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NpETH *NpETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NpETH *NpETHSession) Symbol() (string, error) {
	return _NpETH.Contract.Symbol(&_NpETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NpETH *NpETHCallerSession) Symbol() (string, error) {
	return _NpETH.Contract.Symbol(&_NpETH.CallOpts)
}

// TotalExitingAmount is a free data retrieval call binding the contract method 0x2e899401.
//
// Solidity: function totalExitingAmount() view returns(uint256)
func (_NpETH *NpETHCaller) TotalExitingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "totalExitingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalExitingAmount is a free data retrieval call binding the contract method 0x2e899401.
//
// Solidity: function totalExitingAmount() view returns(uint256)
func (_NpETH *NpETHSession) TotalExitingAmount() (*big.Int, error) {
	return _NpETH.Contract.TotalExitingAmount(&_NpETH.CallOpts)
}

// TotalExitingAmount is a free data retrieval call binding the contract method 0x2e899401.
//
// Solidity: function totalExitingAmount() view returns(uint256)
func (_NpETH *NpETHCallerSession) TotalExitingAmount() (*big.Int, error) {
	return _NpETH.Contract.TotalExitingAmount(&_NpETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NpETH *NpETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NpETH *NpETHSession) TotalSupply() (*big.Int, error) {
	return _NpETH.Contract.TotalSupply(&_NpETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NpETH *NpETHCallerSession) TotalSupply() (*big.Int, error) {
	return _NpETH.Contract.TotalSupply(&_NpETH.CallOpts)
}

// TotalWithdrawReqAmount is a free data retrieval call binding the contract method 0xd2060a5e.
//
// Solidity: function totalWithdrawReqAmount() view returns(uint256)
func (_NpETH *NpETHCaller) TotalWithdrawReqAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "totalWithdrawReqAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWithdrawReqAmount is a free data retrieval call binding the contract method 0xd2060a5e.
//
// Solidity: function totalWithdrawReqAmount() view returns(uint256)
func (_NpETH *NpETHSession) TotalWithdrawReqAmount() (*big.Int, error) {
	return _NpETH.Contract.TotalWithdrawReqAmount(&_NpETH.CallOpts)
}

// TotalWithdrawReqAmount is a free data retrieval call binding the contract method 0xd2060a5e.
//
// Solidity: function totalWithdrawReqAmount() view returns(uint256)
func (_NpETH *NpETHCallerSession) TotalWithdrawReqAmount() (*big.Int, error) {
	return _NpETH.Contract.TotalWithdrawReqAmount(&_NpETH.CallOpts)
}

// TotalWithdrawableAmount is a free data retrieval call binding the contract method 0x40491cd2.
//
// Solidity: function totalWithdrawableAmount() view returns(uint256)
func (_NpETH *NpETHCaller) TotalWithdrawableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "totalWithdrawableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWithdrawableAmount is a free data retrieval call binding the contract method 0x40491cd2.
//
// Solidity: function totalWithdrawableAmount() view returns(uint256)
func (_NpETH *NpETHSession) TotalWithdrawableAmount() (*big.Int, error) {
	return _NpETH.Contract.TotalWithdrawableAmount(&_NpETH.CallOpts)
}

// TotalWithdrawableAmount is a free data retrieval call binding the contract method 0x40491cd2.
//
// Solidity: function totalWithdrawableAmount() view returns(uint256)
func (_NpETH *NpETHCallerSession) TotalWithdrawableAmount() (*big.Int, error) {
	return _NpETH.Contract.TotalWithdrawableAmount(&_NpETH.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0x674a5423.
//
// Solidity: function users(uint8 , address ) view returns(uint256 withdrawReqAmount, uint256 withdrawableAmount)
func (_NpETH *NpETHCaller) Users(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (struct {
	WithdrawReqAmount  *big.Int
	WithdrawableAmount *big.Int
}, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "users", arg0, arg1)

	outstruct := new(struct {
		WithdrawReqAmount  *big.Int
		WithdrawableAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawReqAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawableAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x674a5423.
//
// Solidity: function users(uint8 , address ) view returns(uint256 withdrawReqAmount, uint256 withdrawableAmount)
func (_NpETH *NpETHSession) Users(arg0 uint8, arg1 common.Address) (struct {
	WithdrawReqAmount  *big.Int
	WithdrawableAmount *big.Int
}, error) {
	return _NpETH.Contract.Users(&_NpETH.CallOpts, arg0, arg1)
}

// Users is a free data retrieval call binding the contract method 0x674a5423.
//
// Solidity: function users(uint8 , address ) view returns(uint256 withdrawReqAmount, uint256 withdrawableAmount)
func (_NpETH *NpETHCallerSession) Users(arg0 uint8, arg1 common.Address) (struct {
	WithdrawReqAmount  *big.Int
	WithdrawableAmount *big.Int
}, error) {
	return _NpETH.Contract.Users(&_NpETH.CallOpts, arg0, arg1)
}

// Verification is a free data retrieval call binding the contract method 0x4ffe2a8b.
//
// Solidity: function verification() view returns(address)
func (_NpETH *NpETHCaller) Verification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "verification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verification is a free data retrieval call binding the contract method 0x4ffe2a8b.
//
// Solidity: function verification() view returns(address)
func (_NpETH *NpETHSession) Verification() (common.Address, error) {
	return _NpETH.Contract.Verification(&_NpETH.CallOpts)
}

// Verification is a free data retrieval call binding the contract method 0x4ffe2a8b.
//
// Solidity: function verification() view returns(address)
func (_NpETH *NpETHCallerSession) Verification() (common.Address, error) {
	return _NpETH.Contract.Verification(&_NpETH.CallOpts)
}

// WithdrawReqLog is a free data retrieval call binding the contract method 0x58eff737.
//
// Solidity: function withdrawReqLog(uint256 ) view returns(uint256 amount, address user, uint8 stype, uint8 status)
func (_NpETH *NpETHCaller) WithdrawReqLog(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount *big.Int
	User   common.Address
	Stype  uint8
	Status uint8
}, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "withdrawReqLog", arg0)

	outstruct := new(struct {
		Amount *big.Int
		User   common.Address
		Stype  uint8
		Status uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Stype = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// WithdrawReqLog is a free data retrieval call binding the contract method 0x58eff737.
//
// Solidity: function withdrawReqLog(uint256 ) view returns(uint256 amount, address user, uint8 stype, uint8 status)
func (_NpETH *NpETHSession) WithdrawReqLog(arg0 *big.Int) (struct {
	Amount *big.Int
	User   common.Address
	Stype  uint8
	Status uint8
}, error) {
	return _NpETH.Contract.WithdrawReqLog(&_NpETH.CallOpts, arg0)
}

// WithdrawReqLog is a free data retrieval call binding the contract method 0x58eff737.
//
// Solidity: function withdrawReqLog(uint256 ) view returns(uint256 amount, address user, uint8 stype, uint8 status)
func (_NpETH *NpETHCallerSession) WithdrawReqLog(arg0 *big.Int) (struct {
	Amount *big.Int
	User   common.Address
	Stype  uint8
	Status uint8
}, error) {
	return _NpETH.Contract.WithdrawReqLog(&_NpETH.CallOpts, arg0)
}

// WithdrawReqLogLength is a free data retrieval call binding the contract method 0x9fdeb253.
//
// Solidity: function withdrawReqLogLength() view returns(uint256)
func (_NpETH *NpETHCaller) WithdrawReqLogLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "withdrawReqLogLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawReqLogLength is a free data retrieval call binding the contract method 0x9fdeb253.
//
// Solidity: function withdrawReqLogLength() view returns(uint256)
func (_NpETH *NpETHSession) WithdrawReqLogLength() (*big.Int, error) {
	return _NpETH.Contract.WithdrawReqLogLength(&_NpETH.CallOpts)
}

// WithdrawReqLogLength is a free data retrieval call binding the contract method 0x9fdeb253.
//
// Solidity: function withdrawReqLogLength() view returns(uint256)
func (_NpETH *NpETHCallerSession) WithdrawReqLogLength() (*big.Int, error) {
	return _NpETH.Contract.WithdrawReqLogLength(&_NpETH.CallOpts)
}

// WithdrawReqLogNextStartIndex is a free data retrieval call binding the contract method 0xaa4bb4da.
//
// Solidity: function withdrawReqLogNextStartIndex() view returns(uint256)
func (_NpETH *NpETHCaller) WithdrawReqLogNextStartIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NpETH.contract.Call(opts, &out, "withdrawReqLogNextStartIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawReqLogNextStartIndex is a free data retrieval call binding the contract method 0xaa4bb4da.
//
// Solidity: function withdrawReqLogNextStartIndex() view returns(uint256)
func (_NpETH *NpETHSession) WithdrawReqLogNextStartIndex() (*big.Int, error) {
	return _NpETH.Contract.WithdrawReqLogNextStartIndex(&_NpETH.CallOpts)
}

// WithdrawReqLogNextStartIndex is a free data retrieval call binding the contract method 0xaa4bb4da.
//
// Solidity: function withdrawReqLogNextStartIndex() view returns(uint256)
func (_NpETH *NpETHCallerSession) WithdrawReqLogNextStartIndex() (*big.Int, error) {
	return _NpETH.Contract.WithdrawReqLogNextStartIndex(&_NpETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NpETH *NpETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NpETH *NpETHSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.Approve(&_NpETH.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NpETH *NpETHTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.Approve(&_NpETH.TransactOpts, spender, amount)
}

// BurnFee is a paid mutator transaction binding the contract method 0x49ae028a.
//
// Solidity: function burnFee(uint256 _npEthAmount) returns()
func (_NpETH *NpETHTransactor) BurnFee(opts *bind.TransactOpts, _npEthAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "burnFee", _npEthAmount)
}

// BurnFee is a paid mutator transaction binding the contract method 0x49ae028a.
//
// Solidity: function burnFee(uint256 _npEthAmount) returns()
func (_NpETH *NpETHSession) BurnFee(_npEthAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.BurnFee(&_NpETH.TransactOpts, _npEthAmount)
}

// BurnFee is a paid mutator transaction binding the contract method 0x49ae028a.
//
// Solidity: function burnFee(uint256 _npEthAmount) returns()
func (_NpETH *NpETHTransactorSession) BurnFee(_npEthAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.BurnFee(&_NpETH.TransactOpts, _npEthAmount)
}

// ClaimStake is a paid mutator transaction binding the contract method 0xdca0c8aa.
//
// Solidity: function claimStake(uint256 _amountEth, bytes _proof) returns()
func (_NpETH *NpETHTransactor) ClaimStake(opts *bind.TransactOpts, _amountEth *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "claimStake", _amountEth, _proof)
}

// ClaimStake is a paid mutator transaction binding the contract method 0xdca0c8aa.
//
// Solidity: function claimStake(uint256 _amountEth, bytes _proof) returns()
func (_NpETH *NpETHSession) ClaimStake(_amountEth *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.ClaimStake(&_NpETH.TransactOpts, _amountEth, _proof)
}

// ClaimStake is a paid mutator transaction binding the contract method 0xdca0c8aa.
//
// Solidity: function claimStake(uint256 _amountEth, bytes _proof) returns()
func (_NpETH *NpETHTransactorSession) ClaimStake(_amountEth *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.ClaimStake(&_NpETH.TransactOpts, _amountEth, _proof)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NpETH *NpETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NpETH *NpETHSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.DecreaseAllowance(&_NpETH.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NpETH *NpETHTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.DecreaseAllowance(&_NpETH.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NpETH *NpETHTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NpETH *NpETHSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.GrantRole(&_NpETH.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NpETH *NpETHTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.GrantRole(&_NpETH.TransactOpts, role, account)
}

// HandleCompensateBufferEth is a paid mutator transaction binding the contract method 0x0e090b06.
//
// Solidity: function handleCompensateBufferEth() payable returns()
func (_NpETH *NpETHTransactor) HandleCompensateBufferEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "handleCompensateBufferEth")
}

// HandleCompensateBufferEth is a paid mutator transaction binding the contract method 0x0e090b06.
//
// Solidity: function handleCompensateBufferEth() payable returns()
func (_NpETH *NpETHSession) HandleCompensateBufferEth() (*types.Transaction, error) {
	return _NpETH.Contract.HandleCompensateBufferEth(&_NpETH.TransactOpts)
}

// HandleCompensateBufferEth is a paid mutator transaction binding the contract method 0x0e090b06.
//
// Solidity: function handleCompensateBufferEth() payable returns()
func (_NpETH *NpETHTransactorSession) HandleCompensateBufferEth() (*types.Transaction, error) {
	return _NpETH.Contract.HandleCompensateBufferEth(&_NpETH.TransactOpts)
}

// HandleDepositValidators is a paid mutator transaction binding the contract method 0x39656a7f.
//
// Solidity: function handleDepositValidators(bytes[] _pubkeys, bytes[] _withdrawal_credentials, bytes[] _signatures, bytes32[] _deposit_data_roots) returns()
func (_NpETH *NpETHTransactor) HandleDepositValidators(opts *bind.TransactOpts, _pubkeys [][]byte, _withdrawal_credentials [][]byte, _signatures [][]byte, _deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "handleDepositValidators", _pubkeys, _withdrawal_credentials, _signatures, _deposit_data_roots)
}

// HandleDepositValidators is a paid mutator transaction binding the contract method 0x39656a7f.
//
// Solidity: function handleDepositValidators(bytes[] _pubkeys, bytes[] _withdrawal_credentials, bytes[] _signatures, bytes32[] _deposit_data_roots) returns()
func (_NpETH *NpETHSession) HandleDepositValidators(_pubkeys [][]byte, _withdrawal_credentials [][]byte, _signatures [][]byte, _deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _NpETH.Contract.HandleDepositValidators(&_NpETH.TransactOpts, _pubkeys, _withdrawal_credentials, _signatures, _deposit_data_roots)
}

// HandleDepositValidators is a paid mutator transaction binding the contract method 0x39656a7f.
//
// Solidity: function handleDepositValidators(bytes[] _pubkeys, bytes[] _withdrawal_credentials, bytes[] _signatures, bytes32[] _deposit_data_roots) returns()
func (_NpETH *NpETHTransactorSession) HandleDepositValidators(_pubkeys [][]byte, _withdrawal_credentials [][]byte, _signatures [][]byte, _deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _NpETH.Contract.HandleDepositValidators(&_NpETH.TransactOpts, _pubkeys, _withdrawal_credentials, _signatures, _deposit_data_roots)
}

// HandleElRewards is a paid mutator transaction binding the contract method 0xf770362d.
//
// Solidity: function handleElRewards(uint256 _accBeaconValidators, uint256 _beaconBalance, bool force) returns()
func (_NpETH *NpETHTransactor) HandleElRewards(opts *bind.TransactOpts, _accBeaconValidators *big.Int, _beaconBalance *big.Int, force bool) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "handleElRewards", _accBeaconValidators, _beaconBalance, force)
}

// HandleElRewards is a paid mutator transaction binding the contract method 0xf770362d.
//
// Solidity: function handleElRewards(uint256 _accBeaconValidators, uint256 _beaconBalance, bool force) returns()
func (_NpETH *NpETHSession) HandleElRewards(_accBeaconValidators *big.Int, _beaconBalance *big.Int, force bool) (*types.Transaction, error) {
	return _NpETH.Contract.HandleElRewards(&_NpETH.TransactOpts, _accBeaconValidators, _beaconBalance, force)
}

// HandleElRewards is a paid mutator transaction binding the contract method 0xf770362d.
//
// Solidity: function handleElRewards(uint256 _accBeaconValidators, uint256 _beaconBalance, bool force) returns()
func (_NpETH *NpETHTransactorSession) HandleElRewards(_accBeaconValidators *big.Int, _beaconBalance *big.Int, force bool) (*types.Transaction, error) {
	return _NpETH.Contract.HandleElRewards(&_NpETH.TransactOpts, _accBeaconValidators, _beaconBalance, force)
}

// HandleExitValidator is a paid mutator transaction binding the contract method 0xf5bd455a.
//
// Solidity: function handleExitValidator(bytes _exitPubkey, uint256 _exitingAmount) returns()
func (_NpETH *NpETHTransactor) HandleExitValidator(opts *bind.TransactOpts, _exitPubkey []byte, _exitingAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "handleExitValidator", _exitPubkey, _exitingAmount)
}

// HandleExitValidator is a paid mutator transaction binding the contract method 0xf5bd455a.
//
// Solidity: function handleExitValidator(bytes _exitPubkey, uint256 _exitingAmount) returns()
func (_NpETH *NpETHSession) HandleExitValidator(_exitPubkey []byte, _exitingAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.HandleExitValidator(&_NpETH.TransactOpts, _exitPubkey, _exitingAmount)
}

// HandleExitValidator is a paid mutator transaction binding the contract method 0xf5bd455a.
//
// Solidity: function handleExitValidator(bytes _exitPubkey, uint256 _exitingAmount) returns()
func (_NpETH *NpETHTransactorSession) HandleExitValidator(_exitPubkey []byte, _exitingAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.HandleExitValidator(&_NpETH.TransactOpts, _exitPubkey, _exitingAmount)
}

// HandleExitedValidator is a paid mutator transaction binding the contract method 0x23d622cb.
//
// Solidity: function handleExitedValidator(uint256 _accBeaconValidators, uint256 _beaconBalance, bytes _exitedPubkey, uint256 _exitedAmount) returns()
func (_NpETH *NpETHTransactor) HandleExitedValidator(opts *bind.TransactOpts, _accBeaconValidators *big.Int, _beaconBalance *big.Int, _exitedPubkey []byte, _exitedAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "handleExitedValidator", _accBeaconValidators, _beaconBalance, _exitedPubkey, _exitedAmount)
}

// HandleExitedValidator is a paid mutator transaction binding the contract method 0x23d622cb.
//
// Solidity: function handleExitedValidator(uint256 _accBeaconValidators, uint256 _beaconBalance, bytes _exitedPubkey, uint256 _exitedAmount) returns()
func (_NpETH *NpETHSession) HandleExitedValidator(_accBeaconValidators *big.Int, _beaconBalance *big.Int, _exitedPubkey []byte, _exitedAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.HandleExitedValidator(&_NpETH.TransactOpts, _accBeaconValidators, _beaconBalance, _exitedPubkey, _exitedAmount)
}

// HandleExitedValidator is a paid mutator transaction binding the contract method 0x23d622cb.
//
// Solidity: function handleExitedValidator(uint256 _accBeaconValidators, uint256 _beaconBalance, bytes _exitedPubkey, uint256 _exitedAmount) returns()
func (_NpETH *NpETHTransactorSession) HandleExitedValidator(_accBeaconValidators *big.Int, _beaconBalance *big.Int, _exitedPubkey []byte, _exitedAmount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.HandleExitedValidator(&_NpETH.TransactOpts, _accBeaconValidators, _beaconBalance, _exitedPubkey, _exitedAmount)
}

// HandleWithdrawReq is a paid mutator transaction binding the contract method 0x13d01dcf.
//
// Solidity: function handleWithdrawReq() returns()
func (_NpETH *NpETHTransactor) HandleWithdrawReq(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "handleWithdrawReq")
}

// HandleWithdrawReq is a paid mutator transaction binding the contract method 0x13d01dcf.
//
// Solidity: function handleWithdrawReq() returns()
func (_NpETH *NpETHSession) HandleWithdrawReq() (*types.Transaction, error) {
	return _NpETH.Contract.HandleWithdrawReq(&_NpETH.TransactOpts)
}

// HandleWithdrawReq is a paid mutator transaction binding the contract method 0x13d01dcf.
//
// Solidity: function handleWithdrawReq() returns()
func (_NpETH *NpETHTransactorSession) HandleWithdrawReq() (*types.Transaction, error) {
	return _NpETH.Contract.HandleWithdrawReq(&_NpETH.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NpETH *NpETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NpETH *NpETHSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.IncreaseAllowance(&_NpETH.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NpETH *NpETHTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.IncreaseAllowance(&_NpETH.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x33e1a223.
//
// Solidity: function initialize(address _cnDepositAddress, address _feeTo, uint256 _protocolFeeBP, address _npElRewardsVaultAddress, address _verification) returns()
func (_NpETH *NpETHTransactor) Initialize(opts *bind.TransactOpts, _cnDepositAddress common.Address, _feeTo common.Address, _protocolFeeBP *big.Int, _npElRewardsVaultAddress common.Address, _verification common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "initialize", _cnDepositAddress, _feeTo, _protocolFeeBP, _npElRewardsVaultAddress, _verification)
}

// Initialize is a paid mutator transaction binding the contract method 0x33e1a223.
//
// Solidity: function initialize(address _cnDepositAddress, address _feeTo, uint256 _protocolFeeBP, address _npElRewardsVaultAddress, address _verification) returns()
func (_NpETH *NpETHSession) Initialize(_cnDepositAddress common.Address, _feeTo common.Address, _protocolFeeBP *big.Int, _npElRewardsVaultAddress common.Address, _verification common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.Initialize(&_NpETH.TransactOpts, _cnDepositAddress, _feeTo, _protocolFeeBP, _npElRewardsVaultAddress, _verification)
}

// Initialize is a paid mutator transaction binding the contract method 0x33e1a223.
//
// Solidity: function initialize(address _cnDepositAddress, address _feeTo, uint256 _protocolFeeBP, address _npElRewardsVaultAddress, address _verification) returns()
func (_NpETH *NpETHTransactorSession) Initialize(_cnDepositAddress common.Address, _feeTo common.Address, _protocolFeeBP *big.Int, _npElRewardsVaultAddress common.Address, _verification common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.Initialize(&_NpETH.TransactOpts, _cnDepositAddress, _feeTo, _protocolFeeBP, _npElRewardsVaultAddress, _verification)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NpETH *NpETHTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NpETH *NpETHSession) Pause() (*types.Transaction, error) {
	return _NpETH.Contract.Pause(&_NpETH.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NpETH *NpETHTransactorSession) Pause() (*types.Transaction, error) {
	return _NpETH.Contract.Pause(&_NpETH.TransactOpts)
}

// ReceiveELExit is a paid mutator transaction binding the contract method 0x18229a21.
//
// Solidity: function receiveELExit() payable returns()
func (_NpETH *NpETHTransactor) ReceiveELExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "receiveELExit")
}

// ReceiveELExit is a paid mutator transaction binding the contract method 0x18229a21.
//
// Solidity: function receiveELExit() payable returns()
func (_NpETH *NpETHSession) ReceiveELExit() (*types.Transaction, error) {
	return _NpETH.Contract.ReceiveELExit(&_NpETH.TransactOpts)
}

// ReceiveELExit is a paid mutator transaction binding the contract method 0x18229a21.
//
// Solidity: function receiveELExit() payable returns()
func (_NpETH *NpETHTransactorSession) ReceiveELExit() (*types.Transaction, error) {
	return _NpETH.Contract.ReceiveELExit(&_NpETH.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_NpETH *NpETHTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_NpETH *NpETHSession) ReceiveELRewards() (*types.Transaction, error) {
	return _NpETH.Contract.ReceiveELRewards(&_NpETH.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_NpETH *NpETHTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _NpETH.Contract.ReceiveELRewards(&_NpETH.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NpETH *NpETHTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NpETH *NpETHSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.RenounceRole(&_NpETH.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NpETH *NpETHTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.RenounceRole(&_NpETH.TransactOpts, role, account)
}

// RequestClaim is a paid mutator transaction binding the contract method 0x372133d7.
//
// Solidity: function requestClaim(uint256 _amountEth, bytes _proof) returns()
func (_NpETH *NpETHTransactor) RequestClaim(opts *bind.TransactOpts, _amountEth *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "requestClaim", _amountEth, _proof)
}

// RequestClaim is a paid mutator transaction binding the contract method 0x372133d7.
//
// Solidity: function requestClaim(uint256 _amountEth, bytes _proof) returns()
func (_NpETH *NpETHSession) RequestClaim(_amountEth *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.RequestClaim(&_NpETH.TransactOpts, _amountEth, _proof)
}

// RequestClaim is a paid mutator transaction binding the contract method 0x372133d7.
//
// Solidity: function requestClaim(uint256 _amountEth, bytes _proof) returns()
func (_NpETH *NpETHTransactorSession) RequestClaim(_amountEth *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.RequestClaim(&_NpETH.TransactOpts, _amountEth, _proof)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x9087c1f2.
//
// Solidity: function requestWithdraw(uint256 _amountNp, bytes _proof) returns()
func (_NpETH *NpETHTransactor) RequestWithdraw(opts *bind.TransactOpts, _amountNp *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "requestWithdraw", _amountNp, _proof)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x9087c1f2.
//
// Solidity: function requestWithdraw(uint256 _amountNp, bytes _proof) returns()
func (_NpETH *NpETHSession) RequestWithdraw(_amountNp *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.RequestWithdraw(&_NpETH.TransactOpts, _amountNp, _proof)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x9087c1f2.
//
// Solidity: function requestWithdraw(uint256 _amountNp, bytes _proof) returns()
func (_NpETH *NpETHTransactorSession) RequestWithdraw(_amountNp *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.RequestWithdraw(&_NpETH.TransactOpts, _amountNp, _proof)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NpETH *NpETHTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NpETH *NpETHSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.RevokeRole(&_NpETH.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NpETH *NpETHTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.RevokeRole(&_NpETH.TransactOpts, role, account)
}

// SetCnDepositAddress is a paid mutator transaction binding the contract method 0x5311e699.
//
// Solidity: function setCnDepositAddress(address _newCnDepositAddress) returns()
func (_NpETH *NpETHTransactor) SetCnDepositAddress(opts *bind.TransactOpts, _newCnDepositAddress common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "setCnDepositAddress", _newCnDepositAddress)
}

// SetCnDepositAddress is a paid mutator transaction binding the contract method 0x5311e699.
//
// Solidity: function setCnDepositAddress(address _newCnDepositAddress) returns()
func (_NpETH *NpETHSession) SetCnDepositAddress(_newCnDepositAddress common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetCnDepositAddress(&_NpETH.TransactOpts, _newCnDepositAddress)
}

// SetCnDepositAddress is a paid mutator transaction binding the contract method 0x5311e699.
//
// Solidity: function setCnDepositAddress(address _newCnDepositAddress) returns()
func (_NpETH *NpETHTransactorSession) SetCnDepositAddress(_newCnDepositAddress common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetCnDepositAddress(&_NpETH.TransactOpts, _newCnDepositAddress)
}

// SetELRewardsVaultAddress is a paid mutator transaction binding the contract method 0x230c6212.
//
// Solidity: function setELRewardsVaultAddress(address _executionLayerRewardsVaultAddress) returns()
func (_NpETH *NpETHTransactor) SetELRewardsVaultAddress(opts *bind.TransactOpts, _executionLayerRewardsVaultAddress common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "setELRewardsVaultAddress", _executionLayerRewardsVaultAddress)
}

// SetELRewardsVaultAddress is a paid mutator transaction binding the contract method 0x230c6212.
//
// Solidity: function setELRewardsVaultAddress(address _executionLayerRewardsVaultAddress) returns()
func (_NpETH *NpETHSession) SetELRewardsVaultAddress(_executionLayerRewardsVaultAddress common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetELRewardsVaultAddress(&_NpETH.TransactOpts, _executionLayerRewardsVaultAddress)
}

// SetELRewardsVaultAddress is a paid mutator transaction binding the contract method 0x230c6212.
//
// Solidity: function setELRewardsVaultAddress(address _executionLayerRewardsVaultAddress) returns()
func (_NpETH *NpETHTransactorSession) SetELRewardsVaultAddress(_executionLayerRewardsVaultAddress common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetELRewardsVaultAddress(&_NpETH.TransactOpts, _executionLayerRewardsVaultAddress)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _newFeeTo) returns()
func (_NpETH *NpETHTransactor) SetFeeTo(opts *bind.TransactOpts, _newFeeTo common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "setFeeTo", _newFeeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _newFeeTo) returns()
func (_NpETH *NpETHSession) SetFeeTo(_newFeeTo common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetFeeTo(&_NpETH.TransactOpts, _newFeeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _newFeeTo) returns()
func (_NpETH *NpETHTransactorSession) SetFeeTo(_newFeeTo common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetFeeTo(&_NpETH.TransactOpts, _newFeeTo)
}

// SetLimitDepositLoop is a paid mutator transaction binding the contract method 0xa12dec33.
//
// Solidity: function setLimitDepositLoop(uint32 _newLimitDepositLoop) returns()
func (_NpETH *NpETHTransactor) SetLimitDepositLoop(opts *bind.TransactOpts, _newLimitDepositLoop uint32) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "setLimitDepositLoop", _newLimitDepositLoop)
}

// SetLimitDepositLoop is a paid mutator transaction binding the contract method 0xa12dec33.
//
// Solidity: function setLimitDepositLoop(uint32 _newLimitDepositLoop) returns()
func (_NpETH *NpETHSession) SetLimitDepositLoop(_newLimitDepositLoop uint32) (*types.Transaction, error) {
	return _NpETH.Contract.SetLimitDepositLoop(&_NpETH.TransactOpts, _newLimitDepositLoop)
}

// SetLimitDepositLoop is a paid mutator transaction binding the contract method 0xa12dec33.
//
// Solidity: function setLimitDepositLoop(uint32 _newLimitDepositLoop) returns()
func (_NpETH *NpETHTransactorSession) SetLimitDepositLoop(_newLimitDepositLoop uint32) (*types.Transaction, error) {
	return _NpETH.Contract.SetLimitDepositLoop(&_NpETH.TransactOpts, _newLimitDepositLoop)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newProtocolFeeBP) returns()
func (_NpETH *NpETHTransactor) SetProtocolFee(opts *bind.TransactOpts, _newProtocolFeeBP *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "setProtocolFee", _newProtocolFeeBP)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newProtocolFeeBP) returns()
func (_NpETH *NpETHSession) SetProtocolFee(_newProtocolFeeBP *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.SetProtocolFee(&_NpETH.TransactOpts, _newProtocolFeeBP)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _newProtocolFeeBP) returns()
func (_NpETH *NpETHTransactorSession) SetProtocolFee(_newProtocolFeeBP *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.SetProtocolFee(&_NpETH.TransactOpts, _newProtocolFeeBP)
}

// SetVerification is a paid mutator transaction binding the contract method 0x971f8bb1.
//
// Solidity: function setVerification(address _newVerification) returns()
func (_NpETH *NpETHTransactor) SetVerification(opts *bind.TransactOpts, _newVerification common.Address) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "setVerification", _newVerification)
}

// SetVerification is a paid mutator transaction binding the contract method 0x971f8bb1.
//
// Solidity: function setVerification(address _newVerification) returns()
func (_NpETH *NpETHSession) SetVerification(_newVerification common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetVerification(&_NpETH.TransactOpts, _newVerification)
}

// SetVerification is a paid mutator transaction binding the contract method 0x971f8bb1.
//
// Solidity: function setVerification(address _newVerification) returns()
func (_NpETH *NpETHTransactorSession) SetVerification(_newVerification common.Address) (*types.Transaction, error) {
	return _NpETH.Contract.SetVerification(&_NpETH.TransactOpts, _newVerification)
}

// Stake is a paid mutator transaction binding the contract method 0x2d1e0c02.
//
// Solidity: function stake(bytes _proof) payable returns(uint256)
func (_NpETH *NpETHTransactor) Stake(opts *bind.TransactOpts, _proof []byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "stake", _proof)
}

// Stake is a paid mutator transaction binding the contract method 0x2d1e0c02.
//
// Solidity: function stake(bytes _proof) payable returns(uint256)
func (_NpETH *NpETHSession) Stake(_proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Stake(&_NpETH.TransactOpts, _proof)
}

// Stake is a paid mutator transaction binding the contract method 0x2d1e0c02.
//
// Solidity: function stake(bytes _proof) payable returns(uint256)
func (_NpETH *NpETHTransactorSession) Stake(_proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Stake(&_NpETH.TransactOpts, _proof)
}

// Submit is a paid mutator transaction binding the contract method 0xef7fa71b.
//
// Solidity: function submit(bytes _proof) payable returns(uint256)
func (_NpETH *NpETHTransactor) Submit(opts *bind.TransactOpts, _proof []byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "submit", _proof)
}

// Submit is a paid mutator transaction binding the contract method 0xef7fa71b.
//
// Solidity: function submit(bytes _proof) payable returns(uint256)
func (_NpETH *NpETHSession) Submit(_proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Submit(&_NpETH.TransactOpts, _proof)
}

// Submit is a paid mutator transaction binding the contract method 0xef7fa71b.
//
// Solidity: function submit(bytes _proof) payable returns(uint256)
func (_NpETH *NpETHTransactorSession) Submit(_proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Submit(&_NpETH.TransactOpts, _proof)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NpETH *NpETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NpETH *NpETHSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.Transfer(&_NpETH.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NpETH *NpETHTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.Transfer(&_NpETH.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NpETH *NpETHTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NpETH *NpETHSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.TransferFrom(&_NpETH.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NpETH *NpETHTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NpETH.Contract.TransferFrom(&_NpETH.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NpETH *NpETHTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NpETH *NpETHSession) Unpause() (*types.Transaction, error) {
	return _NpETH.Contract.Unpause(&_NpETH.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NpETH *NpETHTransactorSession) Unpause() (*types.Transaction, error) {
	return _NpETH.Contract.Unpause(&_NpETH.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xc8fd6ed0.
//
// Solidity: function unstake(uint256 _amountNp, bytes _proof) returns()
func (_NpETH *NpETHTransactor) Unstake(opts *bind.TransactOpts, _amountNp *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.contract.Transact(opts, "unstake", _amountNp, _proof)
}

// Unstake is a paid mutator transaction binding the contract method 0xc8fd6ed0.
//
// Solidity: function unstake(uint256 _amountNp, bytes _proof) returns()
func (_NpETH *NpETHSession) Unstake(_amountNp *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Unstake(&_NpETH.TransactOpts, _amountNp, _proof)
}

// Unstake is a paid mutator transaction binding the contract method 0xc8fd6ed0.
//
// Solidity: function unstake(uint256 _amountNp, bytes _proof) returns()
func (_NpETH *NpETHTransactorSession) Unstake(_amountNp *big.Int, _proof []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Unstake(&_NpETH.TransactOpts, _amountNp, _proof)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NpETH *NpETHTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _NpETH.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NpETH *NpETHSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Fallback(&_NpETH.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_NpETH *NpETHTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _NpETH.Contract.Fallback(&_NpETH.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NpETH *NpETHTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NpETH.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NpETH *NpETHSession) Receive() (*types.Transaction, error) {
	return _NpETH.Contract.Receive(&_NpETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NpETH *NpETHTransactorSession) Receive() (*types.Transaction, error) {
	return _NpETH.Contract.Receive(&_NpETH.TransactOpts)
}

// NpETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NpETH contract.
type NpETHApprovalIterator struct {
	Event *NpETHApproval // Event containing the contract specifics and raw log

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
func (it *NpETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHApproval)
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
		it.Event = new(NpETHApproval)
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
func (it *NpETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHApproval represents a Approval event raised by the NpETH contract.
type NpETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NpETH *NpETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NpETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NpETHApprovalIterator{contract: _NpETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NpETH *NpETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NpETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHApproval)
				if err := _NpETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NpETH *NpETHFilterer) ParseApproval(log types.Log) (*NpETHApproval, error) {
	event := new(NpETHApproval)
	if err := _NpETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHBurnFeeEventIterator is returned from FilterBurnFeeEvent and is used to iterate over the raw logs and unpacked data for BurnFeeEvent events raised by the NpETH contract.
type NpETHBurnFeeEventIterator struct {
	Event *NpETHBurnFeeEvent // Event containing the contract specifics and raw log

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
func (it *NpETHBurnFeeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHBurnFeeEvent)
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
		it.Event = new(NpETHBurnFeeEvent)
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
func (it *NpETHBurnFeeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHBurnFeeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHBurnFeeEvent represents a BurnFeeEvent event raised by the NpETH contract.
type NpETHBurnFeeEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnFeeEvent is a free log retrieval operation binding the contract event 0x2c219f774186be924343ec8667f7c3b740a56a1d0acca561780c099ee3d9e36f.
//
// Solidity: event BurnFeeEvent(address indexed _from, uint256 _amount)
func (_NpETH *NpETHFilterer) FilterBurnFeeEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHBurnFeeEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "BurnFeeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHBurnFeeEventIterator{contract: _NpETH.contract, event: "BurnFeeEvent", logs: logs, sub: sub}, nil
}

// WatchBurnFeeEvent is a free log subscription operation binding the contract event 0x2c219f774186be924343ec8667f7c3b740a56a1d0acca561780c099ee3d9e36f.
//
// Solidity: event BurnFeeEvent(address indexed _from, uint256 _amount)
func (_NpETH *NpETHFilterer) WatchBurnFeeEvent(opts *bind.WatchOpts, sink chan<- *NpETHBurnFeeEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "BurnFeeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHBurnFeeEvent)
				if err := _NpETH.contract.UnpackLog(event, "BurnFeeEvent", log); err != nil {
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

// ParseBurnFeeEvent is a log parse operation binding the contract event 0x2c219f774186be924343ec8667f7c3b740a56a1d0acca561780c099ee3d9e36f.
//
// Solidity: event BurnFeeEvent(address indexed _from, uint256 _amount)
func (_NpETH *NpETHFilterer) ParseBurnFeeEvent(log types.Log) (*NpETHBurnFeeEvent, error) {
	event := new(NpETHBurnFeeEvent)
	if err := _NpETH.contract.UnpackLog(event, "BurnFeeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHClaimStakeEventIterator is returned from FilterClaimStakeEvent and is used to iterate over the raw logs and unpacked data for ClaimStakeEvent events raised by the NpETH contract.
type NpETHClaimStakeEventIterator struct {
	Event *NpETHClaimStakeEvent // Event containing the contract specifics and raw log

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
func (it *NpETHClaimStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHClaimStakeEvent)
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
		it.Event = new(NpETHClaimStakeEvent)
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
func (it *NpETHClaimStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHClaimStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHClaimStakeEvent represents a ClaimStakeEvent event raised by the NpETH contract.
type NpETHClaimStakeEvent struct {
	From            common.Address
	AmountEth       *big.Int
	WithdrawableEth *big.Int
	WithdrawReqEth  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimStakeEvent is a free log retrieval operation binding the contract event 0x0eb0d69542ad7ca875f6d25e5dbafb3b7e8aa411d894a69a16381b3f6e8859ae.
//
// Solidity: event ClaimStakeEvent(address indexed _from, uint256 _amountEth, uint256 withdrawableEth, uint256 withdrawReqEth)
func (_NpETH *NpETHFilterer) FilterClaimStakeEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHClaimStakeEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "ClaimStakeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHClaimStakeEventIterator{contract: _NpETH.contract, event: "ClaimStakeEvent", logs: logs, sub: sub}, nil
}

// WatchClaimStakeEvent is a free log subscription operation binding the contract event 0x0eb0d69542ad7ca875f6d25e5dbafb3b7e8aa411d894a69a16381b3f6e8859ae.
//
// Solidity: event ClaimStakeEvent(address indexed _from, uint256 _amountEth, uint256 withdrawableEth, uint256 withdrawReqEth)
func (_NpETH *NpETHFilterer) WatchClaimStakeEvent(opts *bind.WatchOpts, sink chan<- *NpETHClaimStakeEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "ClaimStakeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHClaimStakeEvent)
				if err := _NpETH.contract.UnpackLog(event, "ClaimStakeEvent", log); err != nil {
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

// ParseClaimStakeEvent is a log parse operation binding the contract event 0x0eb0d69542ad7ca875f6d25e5dbafb3b7e8aa411d894a69a16381b3f6e8859ae.
//
// Solidity: event ClaimStakeEvent(address indexed _from, uint256 _amountEth, uint256 withdrawableEth, uint256 withdrawReqEth)
func (_NpETH *NpETHFilterer) ParseClaimStakeEvent(log types.Log) (*NpETHClaimStakeEvent, error) {
	event := new(NpETHClaimStakeEvent)
	if err := _NpETH.contract.UnpackLog(event, "ClaimStakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHELExitReceivedEventIterator is returned from FilterELExitReceivedEvent and is used to iterate over the raw logs and unpacked data for ELExitReceivedEvent events raised by the NpETH contract.
type NpETHELExitReceivedEventIterator struct {
	Event *NpETHELExitReceivedEvent // Event containing the contract specifics and raw log

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
func (it *NpETHELExitReceivedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHELExitReceivedEvent)
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
		it.Event = new(NpETHELExitReceivedEvent)
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
func (it *NpETHELExitReceivedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHELExitReceivedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHELExitReceivedEvent represents a ELExitReceivedEvent event raised by the NpETH contract.
type NpETHELExitReceivedEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELExitReceivedEvent is a free log retrieval operation binding the contract event 0x7f5267c96417b9540b271987cadb6fcac4fe7b69b7a5a865872b0e4c5dfe4b5e.
//
// Solidity: event ELExitReceivedEvent(uint256 amount)
func (_NpETH *NpETHFilterer) FilterELExitReceivedEvent(opts *bind.FilterOpts) (*NpETHELExitReceivedEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "ELExitReceivedEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHELExitReceivedEventIterator{contract: _NpETH.contract, event: "ELExitReceivedEvent", logs: logs, sub: sub}, nil
}

// WatchELExitReceivedEvent is a free log subscription operation binding the contract event 0x7f5267c96417b9540b271987cadb6fcac4fe7b69b7a5a865872b0e4c5dfe4b5e.
//
// Solidity: event ELExitReceivedEvent(uint256 amount)
func (_NpETH *NpETHFilterer) WatchELExitReceivedEvent(opts *bind.WatchOpts, sink chan<- *NpETHELExitReceivedEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "ELExitReceivedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHELExitReceivedEvent)
				if err := _NpETH.contract.UnpackLog(event, "ELExitReceivedEvent", log); err != nil {
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

// ParseELExitReceivedEvent is a log parse operation binding the contract event 0x7f5267c96417b9540b271987cadb6fcac4fe7b69b7a5a865872b0e4c5dfe4b5e.
//
// Solidity: event ELExitReceivedEvent(uint256 amount)
func (_NpETH *NpETHFilterer) ParseELExitReceivedEvent(log types.Log) (*NpETHELExitReceivedEvent, error) {
	event := new(NpETHELExitReceivedEvent)
	if err := _NpETH.contract.UnpackLog(event, "ELExitReceivedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHELIncludeExitEventIterator is returned from FilterELIncludeExitEvent and is used to iterate over the raw logs and unpacked data for ELIncludeExitEvent events raised by the NpETH contract.
type NpETHELIncludeExitEventIterator struct {
	Event *NpETHELIncludeExitEvent // Event containing the contract specifics and raw log

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
func (it *NpETHELIncludeExitEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHELIncludeExitEvent)
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
		it.Event = new(NpETHELIncludeExitEvent)
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
func (it *NpETHELIncludeExitEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHELIncludeExitEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHELIncludeExitEvent represents a ELIncludeExitEvent event raised by the NpETH contract.
type NpETHELIncludeExitEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELIncludeExitEvent is a free log retrieval operation binding the contract event 0xe7ca0e69814f70fb6de6803952b205327b34351582f4da2ea7a141f6d17bb845.
//
// Solidity: event ELIncludeExitEvent(uint256 amount)
func (_NpETH *NpETHFilterer) FilterELIncludeExitEvent(opts *bind.FilterOpts) (*NpETHELIncludeExitEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "ELIncludeExitEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHELIncludeExitEventIterator{contract: _NpETH.contract, event: "ELIncludeExitEvent", logs: logs, sub: sub}, nil
}

// WatchELIncludeExitEvent is a free log subscription operation binding the contract event 0xe7ca0e69814f70fb6de6803952b205327b34351582f4da2ea7a141f6d17bb845.
//
// Solidity: event ELIncludeExitEvent(uint256 amount)
func (_NpETH *NpETHFilterer) WatchELIncludeExitEvent(opts *bind.WatchOpts, sink chan<- *NpETHELIncludeExitEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "ELIncludeExitEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHELIncludeExitEvent)
				if err := _NpETH.contract.UnpackLog(event, "ELIncludeExitEvent", log); err != nil {
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

// ParseELIncludeExitEvent is a log parse operation binding the contract event 0xe7ca0e69814f70fb6de6803952b205327b34351582f4da2ea7a141f6d17bb845.
//
// Solidity: event ELIncludeExitEvent(uint256 amount)
func (_NpETH *NpETHFilterer) ParseELIncludeExitEvent(log types.Log) (*NpETHELIncludeExitEvent, error) {
	event := new(NpETHELIncludeExitEvent)
	if err := _NpETH.contract.UnpackLog(event, "ELIncludeExitEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHELRewardsReceivedEventIterator is returned from FilterELRewardsReceivedEvent and is used to iterate over the raw logs and unpacked data for ELRewardsReceivedEvent events raised by the NpETH contract.
type NpETHELRewardsReceivedEventIterator struct {
	Event *NpETHELRewardsReceivedEvent // Event containing the contract specifics and raw log

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
func (it *NpETHELRewardsReceivedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHELRewardsReceivedEvent)
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
		it.Event = new(NpETHELRewardsReceivedEvent)
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
func (it *NpETHELRewardsReceivedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHELRewardsReceivedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHELRewardsReceivedEvent represents a ELRewardsReceivedEvent event raised by the NpETH contract.
type NpETHELRewardsReceivedEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceivedEvent is a free log retrieval operation binding the contract event 0x7017e96924b5a46dcb9dab161a4dd7c82918812e6e66f60a263241be6d27f285.
//
// Solidity: event ELRewardsReceivedEvent(uint256 amount)
func (_NpETH *NpETHFilterer) FilterELRewardsReceivedEvent(opts *bind.FilterOpts) (*NpETHELRewardsReceivedEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "ELRewardsReceivedEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHELRewardsReceivedEventIterator{contract: _NpETH.contract, event: "ELRewardsReceivedEvent", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceivedEvent is a free log subscription operation binding the contract event 0x7017e96924b5a46dcb9dab161a4dd7c82918812e6e66f60a263241be6d27f285.
//
// Solidity: event ELRewardsReceivedEvent(uint256 amount)
func (_NpETH *NpETHFilterer) WatchELRewardsReceivedEvent(opts *bind.WatchOpts, sink chan<- *NpETHELRewardsReceivedEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "ELRewardsReceivedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHELRewardsReceivedEvent)
				if err := _NpETH.contract.UnpackLog(event, "ELRewardsReceivedEvent", log); err != nil {
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

// ParseELRewardsReceivedEvent is a log parse operation binding the contract event 0x7017e96924b5a46dcb9dab161a4dd7c82918812e6e66f60a263241be6d27f285.
//
// Solidity: event ELRewardsReceivedEvent(uint256 amount)
func (_NpETH *NpETHFilterer) ParseELRewardsReceivedEvent(log types.Log) (*NpETHELRewardsReceivedEvent, error) {
	event := new(NpETHELRewardsReceivedEvent)
	if err := _NpETH.contract.UnpackLog(event, "ELRewardsReceivedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHELRewardsVaultAddressSetEventIterator is returned from FilterELRewardsVaultAddressSetEvent and is used to iterate over the raw logs and unpacked data for ELRewardsVaultAddressSetEvent events raised by the NpETH contract.
type NpETHELRewardsVaultAddressSetEventIterator struct {
	Event *NpETHELRewardsVaultAddressSetEvent // Event containing the contract specifics and raw log

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
func (it *NpETHELRewardsVaultAddressSetEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHELRewardsVaultAddressSetEvent)
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
		it.Event = new(NpETHELRewardsVaultAddressSetEvent)
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
func (it *NpETHELRewardsVaultAddressSetEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHELRewardsVaultAddressSetEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHELRewardsVaultAddressSetEvent represents a ELRewardsVaultAddressSetEvent event raised by the NpETH contract.
type NpETHELRewardsVaultAddressSetEvent struct {
	ExecutionLayerRewardsVaultAddress common.Address
	Raw                               types.Log // Blockchain specific contextual infos
}

// FilterELRewardsVaultAddressSetEvent is a free log retrieval operation binding the contract event 0xa19c96fb75557c188ed86adf36136b75d287050c99ca8f91cae6f14b80a0aac5.
//
// Solidity: event ELRewardsVaultAddressSetEvent(address executionLayerRewardsVaultAddress)
func (_NpETH *NpETHFilterer) FilterELRewardsVaultAddressSetEvent(opts *bind.FilterOpts) (*NpETHELRewardsVaultAddressSetEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "ELRewardsVaultAddressSetEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHELRewardsVaultAddressSetEventIterator{contract: _NpETH.contract, event: "ELRewardsVaultAddressSetEvent", logs: logs, sub: sub}, nil
}

// WatchELRewardsVaultAddressSetEvent is a free log subscription operation binding the contract event 0xa19c96fb75557c188ed86adf36136b75d287050c99ca8f91cae6f14b80a0aac5.
//
// Solidity: event ELRewardsVaultAddressSetEvent(address executionLayerRewardsVaultAddress)
func (_NpETH *NpETHFilterer) WatchELRewardsVaultAddressSetEvent(opts *bind.WatchOpts, sink chan<- *NpETHELRewardsVaultAddressSetEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "ELRewardsVaultAddressSetEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHELRewardsVaultAddressSetEvent)
				if err := _NpETH.contract.UnpackLog(event, "ELRewardsVaultAddressSetEvent", log); err != nil {
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

// ParseELRewardsVaultAddressSetEvent is a log parse operation binding the contract event 0xa19c96fb75557c188ed86adf36136b75d287050c99ca8f91cae6f14b80a0aac5.
//
// Solidity: event ELRewardsVaultAddressSetEvent(address executionLayerRewardsVaultAddress)
func (_NpETH *NpETHFilterer) ParseELRewardsVaultAddressSetEvent(log types.Log) (*NpETHELRewardsVaultAddressSetEvent, error) {
	event := new(NpETHELRewardsVaultAddressSetEvent)
	if err := _NpETH.contract.UnpackLog(event, "ELRewardsVaultAddressSetEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHHandleCompensateBufferEthEventIterator is returned from FilterHandleCompensateBufferEthEvent and is used to iterate over the raw logs and unpacked data for HandleCompensateBufferEthEvent events raised by the NpETH contract.
type NpETHHandleCompensateBufferEthEventIterator struct {
	Event *NpETHHandleCompensateBufferEthEvent // Event containing the contract specifics and raw log

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
func (it *NpETHHandleCompensateBufferEthEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHHandleCompensateBufferEthEvent)
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
		it.Event = new(NpETHHandleCompensateBufferEthEvent)
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
func (it *NpETHHandleCompensateBufferEthEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHHandleCompensateBufferEthEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHHandleCompensateBufferEthEvent represents a HandleCompensateBufferEthEvent event raised by the NpETH contract.
type NpETHHandleCompensateBufferEthEvent struct {
	From        common.Address
	Amount      *big.Int
	BufferedEth *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHandleCompensateBufferEthEvent is a free log retrieval operation binding the contract event 0xad3e5565aa8e4b6ccd19e4ac6385c286d88eee4ee220db681662eafba612b201.
//
// Solidity: event HandleCompensateBufferEthEvent(address indexed _from, uint256 amount, uint256 bufferedEth)
func (_NpETH *NpETHFilterer) FilterHandleCompensateBufferEthEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHHandleCompensateBufferEthEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "HandleCompensateBufferEthEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHHandleCompensateBufferEthEventIterator{contract: _NpETH.contract, event: "HandleCompensateBufferEthEvent", logs: logs, sub: sub}, nil
}

// WatchHandleCompensateBufferEthEvent is a free log subscription operation binding the contract event 0xad3e5565aa8e4b6ccd19e4ac6385c286d88eee4ee220db681662eafba612b201.
//
// Solidity: event HandleCompensateBufferEthEvent(address indexed _from, uint256 amount, uint256 bufferedEth)
func (_NpETH *NpETHFilterer) WatchHandleCompensateBufferEthEvent(opts *bind.WatchOpts, sink chan<- *NpETHHandleCompensateBufferEthEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "HandleCompensateBufferEthEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHHandleCompensateBufferEthEvent)
				if err := _NpETH.contract.UnpackLog(event, "HandleCompensateBufferEthEvent", log); err != nil {
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

// ParseHandleCompensateBufferEthEvent is a log parse operation binding the contract event 0xad3e5565aa8e4b6ccd19e4ac6385c286d88eee4ee220db681662eafba612b201.
//
// Solidity: event HandleCompensateBufferEthEvent(address indexed _from, uint256 amount, uint256 bufferedEth)
func (_NpETH *NpETHFilterer) ParseHandleCompensateBufferEthEvent(log types.Log) (*NpETHHandleCompensateBufferEthEvent, error) {
	event := new(NpETHHandleCompensateBufferEthEvent)
	if err := _NpETH.contract.UnpackLog(event, "HandleCompensateBufferEthEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHHandleDepositEventIterator is returned from FilterHandleDepositEvent and is used to iterate over the raw logs and unpacked data for HandleDepositEvent events raised by the NpETH contract.
type NpETHHandleDepositEventIterator struct {
	Event *NpETHHandleDepositEvent // Event containing the contract specifics and raw log

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
func (it *NpETHHandleDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHHandleDepositEvent)
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
		it.Event = new(NpETHHandleDepositEvent)
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
func (it *NpETHHandleDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHHandleDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHHandleDepositEvent represents a HandleDepositEvent event raised by the NpETH contract.
type NpETHHandleDepositEvent struct {
	Pubkey []byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHandleDepositEvent is a free log retrieval operation binding the contract event 0x90f9d4b28bac0f290220eb701029cb071d9e30224f496fab6c864d9c9e776c19.
//
// Solidity: event HandleDepositEvent(bytes pubkey, uint256 _amount)
func (_NpETH *NpETHFilterer) FilterHandleDepositEvent(opts *bind.FilterOpts) (*NpETHHandleDepositEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "HandleDepositEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHHandleDepositEventIterator{contract: _NpETH.contract, event: "HandleDepositEvent", logs: logs, sub: sub}, nil
}

// WatchHandleDepositEvent is a free log subscription operation binding the contract event 0x90f9d4b28bac0f290220eb701029cb071d9e30224f496fab6c864d9c9e776c19.
//
// Solidity: event HandleDepositEvent(bytes pubkey, uint256 _amount)
func (_NpETH *NpETHFilterer) WatchHandleDepositEvent(opts *bind.WatchOpts, sink chan<- *NpETHHandleDepositEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "HandleDepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHHandleDepositEvent)
				if err := _NpETH.contract.UnpackLog(event, "HandleDepositEvent", log); err != nil {
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

// ParseHandleDepositEvent is a log parse operation binding the contract event 0x90f9d4b28bac0f290220eb701029cb071d9e30224f496fab6c864d9c9e776c19.
//
// Solidity: event HandleDepositEvent(bytes pubkey, uint256 _amount)
func (_NpETH *NpETHFilterer) ParseHandleDepositEvent(log types.Log) (*NpETHHandleDepositEvent, error) {
	event := new(NpETHHandleDepositEvent)
	if err := _NpETH.contract.UnpackLog(event, "HandleDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHHandleElRewardsEventIterator is returned from FilterHandleElRewardsEvent and is used to iterate over the raw logs and unpacked data for HandleElRewardsEvent events raised by the NpETH contract.
type NpETHHandleElRewardsEventIterator struct {
	Event *NpETHHandleElRewardsEvent // Event containing the contract specifics and raw log

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
func (it *NpETHHandleElRewardsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHHandleElRewardsEvent)
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
		it.Event = new(NpETHHandleElRewardsEvent)
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
func (it *NpETHHandleElRewardsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHHandleElRewardsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHHandleElRewardsEvent represents a HandleElRewardsEvent event raised by the NpETH contract.
type NpETHHandleElRewardsEvent struct {
	Amount       *big.Int
	AmountToMint *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHandleElRewardsEvent is a free log retrieval operation binding the contract event 0x909757eebe357724feb9f38a29c4c578c994d7eb9fa75ad1e7ab38f686f0bee3.
//
// Solidity: event HandleElRewardsEvent(uint256 indexed _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) FilterHandleElRewardsEvent(opts *bind.FilterOpts, _amount []*big.Int) (*NpETHHandleElRewardsEventIterator, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "HandleElRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return &NpETHHandleElRewardsEventIterator{contract: _NpETH.contract, event: "HandleElRewardsEvent", logs: logs, sub: sub}, nil
}

// WatchHandleElRewardsEvent is a free log subscription operation binding the contract event 0x909757eebe357724feb9f38a29c4c578c994d7eb9fa75ad1e7ab38f686f0bee3.
//
// Solidity: event HandleElRewardsEvent(uint256 indexed _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) WatchHandleElRewardsEvent(opts *bind.WatchOpts, sink chan<- *NpETHHandleElRewardsEvent, _amount []*big.Int) (event.Subscription, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "HandleElRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHHandleElRewardsEvent)
				if err := _NpETH.contract.UnpackLog(event, "HandleElRewardsEvent", log); err != nil {
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

// ParseHandleElRewardsEvent is a log parse operation binding the contract event 0x909757eebe357724feb9f38a29c4c578c994d7eb9fa75ad1e7ab38f686f0bee3.
//
// Solidity: event HandleElRewardsEvent(uint256 indexed _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) ParseHandleElRewardsEvent(log types.Log) (*NpETHHandleElRewardsEvent, error) {
	event := new(NpETHHandleElRewardsEvent)
	if err := _NpETH.contract.UnpackLog(event, "HandleElRewardsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHHandleExitValidatorEventIterator is returned from FilterHandleExitValidatorEvent and is used to iterate over the raw logs and unpacked data for HandleExitValidatorEvent events raised by the NpETH contract.
type NpETHHandleExitValidatorEventIterator struct {
	Event *NpETHHandleExitValidatorEvent // Event containing the contract specifics and raw log

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
func (it *NpETHHandleExitValidatorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHHandleExitValidatorEvent)
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
		it.Event = new(NpETHHandleExitValidatorEvent)
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
func (it *NpETHHandleExitValidatorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHHandleExitValidatorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHHandleExitValidatorEvent represents a HandleExitValidatorEvent event raised by the NpETH contract.
type NpETHHandleExitValidatorEvent struct {
	Pubkey                 []byte
	TotalWithdrawReqAmount *big.Int
	TotalExitingAmount     *big.Int
	ExitingAmount          *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterHandleExitValidatorEvent is a free log retrieval operation binding the contract event 0x2e872574cbf56bd8587cd1f2f9a9747e56bd116917473e24f2a5408b694d769e.
//
// Solidity: event HandleExitValidatorEvent(bytes _pubkey, uint256 totalWithdrawReqAmount, uint256 totalExitingAmount, uint256 _exitingAmount)
func (_NpETH *NpETHFilterer) FilterHandleExitValidatorEvent(opts *bind.FilterOpts) (*NpETHHandleExitValidatorEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "HandleExitValidatorEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHHandleExitValidatorEventIterator{contract: _NpETH.contract, event: "HandleExitValidatorEvent", logs: logs, sub: sub}, nil
}

// WatchHandleExitValidatorEvent is a free log subscription operation binding the contract event 0x2e872574cbf56bd8587cd1f2f9a9747e56bd116917473e24f2a5408b694d769e.
//
// Solidity: event HandleExitValidatorEvent(bytes _pubkey, uint256 totalWithdrawReqAmount, uint256 totalExitingAmount, uint256 _exitingAmount)
func (_NpETH *NpETHFilterer) WatchHandleExitValidatorEvent(opts *bind.WatchOpts, sink chan<- *NpETHHandleExitValidatorEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "HandleExitValidatorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHHandleExitValidatorEvent)
				if err := _NpETH.contract.UnpackLog(event, "HandleExitValidatorEvent", log); err != nil {
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

// ParseHandleExitValidatorEvent is a log parse operation binding the contract event 0x2e872574cbf56bd8587cd1f2f9a9747e56bd116917473e24f2a5408b694d769e.
//
// Solidity: event HandleExitValidatorEvent(bytes _pubkey, uint256 totalWithdrawReqAmount, uint256 totalExitingAmount, uint256 _exitingAmount)
func (_NpETH *NpETHFilterer) ParseHandleExitValidatorEvent(log types.Log) (*NpETHHandleExitValidatorEvent, error) {
	event := new(NpETHHandleExitValidatorEvent)
	if err := _NpETH.contract.UnpackLog(event, "HandleExitValidatorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHHandleExitedValidatorEventIterator is returned from FilterHandleExitedValidatorEvent and is used to iterate over the raw logs and unpacked data for HandleExitedValidatorEvent events raised by the NpETH contract.
type NpETHHandleExitedValidatorEventIterator struct {
	Event *NpETHHandleExitedValidatorEvent // Event containing the contract specifics and raw log

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
func (it *NpETHHandleExitedValidatorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHHandleExitedValidatorEvent)
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
		it.Event = new(NpETHHandleExitedValidatorEvent)
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
func (it *NpETHHandleExitedValidatorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHHandleExitedValidatorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHHandleExitedValidatorEvent represents a HandleExitedValidatorEvent event raised by the NpETH contract.
type NpETHHandleExitedValidatorEvent struct {
	Pubkey              []byte
	TotalWithdrawReqEth *big.Int
	WithdrawableEth     *big.Int
	BufferedEth         *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHandleExitedValidatorEvent is a free log retrieval operation binding the contract event 0x4ab14b2d6d234e61c5aa42da7dad9270560b99624ffc1efa3f6e7b0af8f52282.
//
// Solidity: event HandleExitedValidatorEvent(bytes _pubkey, uint256 totalWithdrawReqEth, uint256 withdrawableEth, uint256 bufferedEth)
func (_NpETH *NpETHFilterer) FilterHandleExitedValidatorEvent(opts *bind.FilterOpts) (*NpETHHandleExitedValidatorEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "HandleExitedValidatorEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHHandleExitedValidatorEventIterator{contract: _NpETH.contract, event: "HandleExitedValidatorEvent", logs: logs, sub: sub}, nil
}

// WatchHandleExitedValidatorEvent is a free log subscription operation binding the contract event 0x4ab14b2d6d234e61c5aa42da7dad9270560b99624ffc1efa3f6e7b0af8f52282.
//
// Solidity: event HandleExitedValidatorEvent(bytes _pubkey, uint256 totalWithdrawReqEth, uint256 withdrawableEth, uint256 bufferedEth)
func (_NpETH *NpETHFilterer) WatchHandleExitedValidatorEvent(opts *bind.WatchOpts, sink chan<- *NpETHHandleExitedValidatorEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "HandleExitedValidatorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHHandleExitedValidatorEvent)
				if err := _NpETH.contract.UnpackLog(event, "HandleExitedValidatorEvent", log); err != nil {
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

// ParseHandleExitedValidatorEvent is a log parse operation binding the contract event 0x4ab14b2d6d234e61c5aa42da7dad9270560b99624ffc1efa3f6e7b0af8f52282.
//
// Solidity: event HandleExitedValidatorEvent(bytes _pubkey, uint256 totalWithdrawReqEth, uint256 withdrawableEth, uint256 bufferedEth)
func (_NpETH *NpETHFilterer) ParseHandleExitedValidatorEvent(log types.Log) (*NpETHHandleExitedValidatorEvent, error) {
	event := new(NpETHHandleExitedValidatorEvent)
	if err := _NpETH.contract.UnpackLog(event, "HandleExitedValidatorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHHandleWithdrawReqEventIterator is returned from FilterHandleWithdrawReqEvent and is used to iterate over the raw logs and unpacked data for HandleWithdrawReqEvent events raised by the NpETH contract.
type NpETHHandleWithdrawReqEventIterator struct {
	Event *NpETHHandleWithdrawReqEvent // Event containing the contract specifics and raw log

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
func (it *NpETHHandleWithdrawReqEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHHandleWithdrawReqEvent)
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
		it.Event = new(NpETHHandleWithdrawReqEvent)
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
func (it *NpETHHandleWithdrawReqEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHHandleWithdrawReqEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHHandleWithdrawReqEvent represents a HandleWithdrawReqEvent event raised by the NpETH contract.
type NpETHHandleWithdrawReqEvent struct {
	TotalWithdrawReqEth *big.Int
	WithdrawableEth     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHandleWithdrawReqEvent is a free log retrieval operation binding the contract event 0xb03b005340d0f7599ce4baadad4d2db70fef7e9eac3abb5f310badd93f4ec3bf.
//
// Solidity: event HandleWithdrawReqEvent(uint256 totalWithdrawReqEth, uint256 withdrawableEth)
func (_NpETH *NpETHFilterer) FilterHandleWithdrawReqEvent(opts *bind.FilterOpts) (*NpETHHandleWithdrawReqEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "HandleWithdrawReqEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHHandleWithdrawReqEventIterator{contract: _NpETH.contract, event: "HandleWithdrawReqEvent", logs: logs, sub: sub}, nil
}

// WatchHandleWithdrawReqEvent is a free log subscription operation binding the contract event 0xb03b005340d0f7599ce4baadad4d2db70fef7e9eac3abb5f310badd93f4ec3bf.
//
// Solidity: event HandleWithdrawReqEvent(uint256 totalWithdrawReqEth, uint256 withdrawableEth)
func (_NpETH *NpETHFilterer) WatchHandleWithdrawReqEvent(opts *bind.WatchOpts, sink chan<- *NpETHHandleWithdrawReqEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "HandleWithdrawReqEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHHandleWithdrawReqEvent)
				if err := _NpETH.contract.UnpackLog(event, "HandleWithdrawReqEvent", log); err != nil {
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

// ParseHandleWithdrawReqEvent is a log parse operation binding the contract event 0xb03b005340d0f7599ce4baadad4d2db70fef7e9eac3abb5f310badd93f4ec3bf.
//
// Solidity: event HandleWithdrawReqEvent(uint256 totalWithdrawReqEth, uint256 withdrawableEth)
func (_NpETH *NpETHFilterer) ParseHandleWithdrawReqEvent(log types.Log) (*NpETHHandleWithdrawReqEvent, error) {
	event := new(NpETHHandleWithdrawReqEvent)
	if err := _NpETH.contract.UnpackLog(event, "HandleWithdrawReqEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHInitializeEventIterator is returned from FilterInitializeEvent and is used to iterate over the raw logs and unpacked data for InitializeEvent events raised by the NpETH contract.
type NpETHInitializeEventIterator struct {
	Event *NpETHInitializeEvent // Event containing the contract specifics and raw log

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
func (it *NpETHInitializeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHInitializeEvent)
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
		it.Event = new(NpETHInitializeEvent)
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
func (it *NpETHInitializeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHInitializeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHInitializeEvent represents a InitializeEvent event raised by the NpETH contract.
type NpETHInitializeEvent struct {
	From                    common.Address
	CnDepositAddress        common.Address
	FeeTo                   common.Address
	ProtocolFeeBP           *big.Int
	NpElRewardsVaultAddress common.Address
	Verification            common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterInitializeEvent is a free log retrieval operation binding the contract event 0xc457695b69b34cde33748dd58ce23728096642cac37a81bbf007af81834c094b.
//
// Solidity: event InitializeEvent(address indexed _from, address _cnDepositAddress, address _feeTo, uint256 _protocolFeeBP, address _npElRewardsVaultAddress, address _verification)
func (_NpETH *NpETHFilterer) FilterInitializeEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHInitializeEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "InitializeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHInitializeEventIterator{contract: _NpETH.contract, event: "InitializeEvent", logs: logs, sub: sub}, nil
}

// WatchInitializeEvent is a free log subscription operation binding the contract event 0xc457695b69b34cde33748dd58ce23728096642cac37a81bbf007af81834c094b.
//
// Solidity: event InitializeEvent(address indexed _from, address _cnDepositAddress, address _feeTo, uint256 _protocolFeeBP, address _npElRewardsVaultAddress, address _verification)
func (_NpETH *NpETHFilterer) WatchInitializeEvent(opts *bind.WatchOpts, sink chan<- *NpETHInitializeEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "InitializeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHInitializeEvent)
				if err := _NpETH.contract.UnpackLog(event, "InitializeEvent", log); err != nil {
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

// ParseInitializeEvent is a log parse operation binding the contract event 0xc457695b69b34cde33748dd58ce23728096642cac37a81bbf007af81834c094b.
//
// Solidity: event InitializeEvent(address indexed _from, address _cnDepositAddress, address _feeTo, uint256 _protocolFeeBP, address _npElRewardsVaultAddress, address _verification)
func (_NpETH *NpETHFilterer) ParseInitializeEvent(log types.Log) (*NpETHInitializeEvent, error) {
	event := new(NpETHInitializeEvent)
	if err := _NpETH.contract.UnpackLog(event, "InitializeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NpETH contract.
type NpETHInitializedIterator struct {
	Event *NpETHInitialized // Event containing the contract specifics and raw log

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
func (it *NpETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHInitialized)
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
		it.Event = new(NpETHInitialized)
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
func (it *NpETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHInitialized represents a Initialized event raised by the NpETH contract.
type NpETHInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NpETH *NpETHFilterer) FilterInitialized(opts *bind.FilterOpts) (*NpETHInitializedIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NpETHInitializedIterator{contract: _NpETH.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NpETH *NpETHFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NpETHInitialized) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHInitialized)
				if err := _NpETH.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NpETH *NpETHFilterer) ParseInitialized(log types.Log) (*NpETHInitialized, error) {
	event := new(NpETHInitialized)
	if err := _NpETH.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NpETH contract.
type NpETHPausedIterator struct {
	Event *NpETHPaused // Event containing the contract specifics and raw log

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
func (it *NpETHPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHPaused)
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
		it.Event = new(NpETHPaused)
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
func (it *NpETHPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHPaused represents a Paused event raised by the NpETH contract.
type NpETHPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NpETH *NpETHFilterer) FilterPaused(opts *bind.FilterOpts) (*NpETHPausedIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NpETHPausedIterator{contract: _NpETH.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NpETH *NpETHFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NpETHPaused) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHPaused)
				if err := _NpETH.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NpETH *NpETHFilterer) ParsePaused(log types.Log) (*NpETHPaused, error) {
	event := new(NpETHPaused)
	if err := _NpETH.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHRequestClaimEventIterator is returned from FilterRequestClaimEvent and is used to iterate over the raw logs and unpacked data for RequestClaimEvent events raised by the NpETH contract.
type NpETHRequestClaimEventIterator struct {
	Event *NpETHRequestClaimEvent // Event containing the contract specifics and raw log

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
func (it *NpETHRequestClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHRequestClaimEvent)
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
		it.Event = new(NpETHRequestClaimEvent)
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
func (it *NpETHRequestClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHRequestClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHRequestClaimEvent represents a RequestClaimEvent event raised by the NpETH contract.
type NpETHRequestClaimEvent struct {
	From            common.Address
	AmountEth       *big.Int
	WithdrawableEth *big.Int
	WithdrawReqEth  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRequestClaimEvent is a free log retrieval operation binding the contract event 0xb45d30e450715074e15d544f940947774641146a8d107ce16f4f2014a5b7bfda.
//
// Solidity: event RequestClaimEvent(address indexed _from, uint256 _amountEth, uint256 withdrawableEth, uint256 withdrawReqEth)
func (_NpETH *NpETHFilterer) FilterRequestClaimEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHRequestClaimEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "RequestClaimEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHRequestClaimEventIterator{contract: _NpETH.contract, event: "RequestClaimEvent", logs: logs, sub: sub}, nil
}

// WatchRequestClaimEvent is a free log subscription operation binding the contract event 0xb45d30e450715074e15d544f940947774641146a8d107ce16f4f2014a5b7bfda.
//
// Solidity: event RequestClaimEvent(address indexed _from, uint256 _amountEth, uint256 withdrawableEth, uint256 withdrawReqEth)
func (_NpETH *NpETHFilterer) WatchRequestClaimEvent(opts *bind.WatchOpts, sink chan<- *NpETHRequestClaimEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "RequestClaimEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHRequestClaimEvent)
				if err := _NpETH.contract.UnpackLog(event, "RequestClaimEvent", log); err != nil {
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

// ParseRequestClaimEvent is a log parse operation binding the contract event 0xb45d30e450715074e15d544f940947774641146a8d107ce16f4f2014a5b7bfda.
//
// Solidity: event RequestClaimEvent(address indexed _from, uint256 _amountEth, uint256 withdrawableEth, uint256 withdrawReqEth)
func (_NpETH *NpETHFilterer) ParseRequestClaimEvent(log types.Log) (*NpETHRequestClaimEvent, error) {
	event := new(NpETHRequestClaimEvent)
	if err := _NpETH.contract.UnpackLog(event, "RequestClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHRequestWithdrawEventIterator is returned from FilterRequestWithdrawEvent and is used to iterate over the raw logs and unpacked data for RequestWithdrawEvent events raised by the NpETH contract.
type NpETHRequestWithdrawEventIterator struct {
	Event *NpETHRequestWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *NpETHRequestWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHRequestWithdrawEvent)
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
		it.Event = new(NpETHRequestWithdrawEvent)
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
func (it *NpETHRequestWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHRequestWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHRequestWithdrawEvent represents a RequestWithdrawEvent event raised by the NpETH contract.
type NpETHRequestWithdrawEvent struct {
	From                common.Address
	AmountNpToBurn      *big.Int
	AmountEth           *big.Int
	TotalWithdrawReqEth *big.Int
	WithdrawReqIndex    *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawEvent is a free log retrieval operation binding the contract event 0x23a64f28f76e0f3c39bcc0688be732db46d715eee68104b5acc971d7498f164e.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 _amountNpToBurn, uint256 _amountEth, uint256 totalWithdrawReqEth, uint256 withdrawReqIndex)
func (_NpETH *NpETHFilterer) FilterRequestWithdrawEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHRequestWithdrawEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "RequestWithdrawEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHRequestWithdrawEventIterator{contract: _NpETH.contract, event: "RequestWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawEvent is a free log subscription operation binding the contract event 0x23a64f28f76e0f3c39bcc0688be732db46d715eee68104b5acc971d7498f164e.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 _amountNpToBurn, uint256 _amountEth, uint256 totalWithdrawReqEth, uint256 withdrawReqIndex)
func (_NpETH *NpETHFilterer) WatchRequestWithdrawEvent(opts *bind.WatchOpts, sink chan<- *NpETHRequestWithdrawEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "RequestWithdrawEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHRequestWithdrawEvent)
				if err := _NpETH.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
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

// ParseRequestWithdrawEvent is a log parse operation binding the contract event 0x23a64f28f76e0f3c39bcc0688be732db46d715eee68104b5acc971d7498f164e.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 _amountNpToBurn, uint256 _amountEth, uint256 totalWithdrawReqEth, uint256 withdrawReqIndex)
func (_NpETH *NpETHFilterer) ParseRequestWithdrawEvent(log types.Log) (*NpETHRequestWithdrawEvent, error) {
	event := new(NpETHRequestWithdrawEvent)
	if err := _NpETH.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NpETH contract.
type NpETHRoleAdminChangedIterator struct {
	Event *NpETHRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NpETHRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHRoleAdminChanged)
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
		it.Event = new(NpETHRoleAdminChanged)
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
func (it *NpETHRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHRoleAdminChanged represents a RoleAdminChanged event raised by the NpETH contract.
type NpETHRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NpETH *NpETHFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NpETHRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NpETHRoleAdminChangedIterator{contract: _NpETH.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NpETH *NpETHFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NpETHRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHRoleAdminChanged)
				if err := _NpETH.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NpETH *NpETHFilterer) ParseRoleAdminChanged(log types.Log) (*NpETHRoleAdminChanged, error) {
	event := new(NpETHRoleAdminChanged)
	if err := _NpETH.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NpETH contract.
type NpETHRoleGrantedIterator struct {
	Event *NpETHRoleGranted // Event containing the contract specifics and raw log

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
func (it *NpETHRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHRoleGranted)
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
		it.Event = new(NpETHRoleGranted)
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
func (it *NpETHRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHRoleGranted represents a RoleGranted event raised by the NpETH contract.
type NpETHRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NpETH *NpETHFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NpETHRoleGrantedIterator, error) {

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

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NpETHRoleGrantedIterator{contract: _NpETH.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NpETH *NpETHFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NpETHRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHRoleGranted)
				if err := _NpETH.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NpETH *NpETHFilterer) ParseRoleGranted(log types.Log) (*NpETHRoleGranted, error) {
	event := new(NpETHRoleGranted)
	if err := _NpETH.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NpETH contract.
type NpETHRoleRevokedIterator struct {
	Event *NpETHRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NpETHRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHRoleRevoked)
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
		it.Event = new(NpETHRoleRevoked)
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
func (it *NpETHRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHRoleRevoked represents a RoleRevoked event raised by the NpETH contract.
type NpETHRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NpETH *NpETHFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NpETHRoleRevokedIterator, error) {

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

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NpETHRoleRevokedIterator{contract: _NpETH.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NpETH *NpETHFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NpETHRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHRoleRevoked)
				if err := _NpETH.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NpETH *NpETHFilterer) ParseRoleRevoked(log types.Log) (*NpETHRoleRevoked, error) {
	event := new(NpETHRoleRevoked)
	if err := _NpETH.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHSetCnDepositAddressEventIterator is returned from FilterSetCnDepositAddressEvent and is used to iterate over the raw logs and unpacked data for SetCnDepositAddressEvent events raised by the NpETH contract.
type NpETHSetCnDepositAddressEventIterator struct {
	Event *NpETHSetCnDepositAddressEvent // Event containing the contract specifics and raw log

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
func (it *NpETHSetCnDepositAddressEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHSetCnDepositAddressEvent)
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
		it.Event = new(NpETHSetCnDepositAddressEvent)
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
func (it *NpETHSetCnDepositAddressEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHSetCnDepositAddressEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHSetCnDepositAddressEvent represents a SetCnDepositAddressEvent event raised by the NpETH contract.
type NpETHSetCnDepositAddressEvent struct {
	OldCnDeposit common.Address
	NewCnDeposit common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetCnDepositAddressEvent is a free log retrieval operation binding the contract event 0xbb7d216de7ba77c7c89619b09a6a1f8b55155e2a0efbebd1aba33c202fc5707c.
//
// Solidity: event SetCnDepositAddressEvent(address oldCnDeposit, address newCnDeposit)
func (_NpETH *NpETHFilterer) FilterSetCnDepositAddressEvent(opts *bind.FilterOpts) (*NpETHSetCnDepositAddressEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "SetCnDepositAddressEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHSetCnDepositAddressEventIterator{contract: _NpETH.contract, event: "SetCnDepositAddressEvent", logs: logs, sub: sub}, nil
}

// WatchSetCnDepositAddressEvent is a free log subscription operation binding the contract event 0xbb7d216de7ba77c7c89619b09a6a1f8b55155e2a0efbebd1aba33c202fc5707c.
//
// Solidity: event SetCnDepositAddressEvent(address oldCnDeposit, address newCnDeposit)
func (_NpETH *NpETHFilterer) WatchSetCnDepositAddressEvent(opts *bind.WatchOpts, sink chan<- *NpETHSetCnDepositAddressEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "SetCnDepositAddressEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHSetCnDepositAddressEvent)
				if err := _NpETH.contract.UnpackLog(event, "SetCnDepositAddressEvent", log); err != nil {
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

// ParseSetCnDepositAddressEvent is a log parse operation binding the contract event 0xbb7d216de7ba77c7c89619b09a6a1f8b55155e2a0efbebd1aba33c202fc5707c.
//
// Solidity: event SetCnDepositAddressEvent(address oldCnDeposit, address newCnDeposit)
func (_NpETH *NpETHFilterer) ParseSetCnDepositAddressEvent(log types.Log) (*NpETHSetCnDepositAddressEvent, error) {
	event := new(NpETHSetCnDepositAddressEvent)
	if err := _NpETH.contract.UnpackLog(event, "SetCnDepositAddressEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHSetFeeToEventIterator is returned from FilterSetFeeToEvent and is used to iterate over the raw logs and unpacked data for SetFeeToEvent events raised by the NpETH contract.
type NpETHSetFeeToEventIterator struct {
	Event *NpETHSetFeeToEvent // Event containing the contract specifics and raw log

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
func (it *NpETHSetFeeToEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHSetFeeToEvent)
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
		it.Event = new(NpETHSetFeeToEvent)
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
func (it *NpETHSetFeeToEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHSetFeeToEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHSetFeeToEvent represents a SetFeeToEvent event raised by the NpETH contract.
type NpETHSetFeeToEvent struct {
	OldFeeTo common.Address
	NewFeeTo common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFeeToEvent is a free log retrieval operation binding the contract event 0x33e1eeb6c5b29228be2616c47a3a094c7a3c4b81a8960fe62c1d5ade831dbe39.
//
// Solidity: event SetFeeToEvent(address oldFeeTo, address _newFeeTo)
func (_NpETH *NpETHFilterer) FilterSetFeeToEvent(opts *bind.FilterOpts) (*NpETHSetFeeToEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "SetFeeToEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHSetFeeToEventIterator{contract: _NpETH.contract, event: "SetFeeToEvent", logs: logs, sub: sub}, nil
}

// WatchSetFeeToEvent is a free log subscription operation binding the contract event 0x33e1eeb6c5b29228be2616c47a3a094c7a3c4b81a8960fe62c1d5ade831dbe39.
//
// Solidity: event SetFeeToEvent(address oldFeeTo, address _newFeeTo)
func (_NpETH *NpETHFilterer) WatchSetFeeToEvent(opts *bind.WatchOpts, sink chan<- *NpETHSetFeeToEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "SetFeeToEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHSetFeeToEvent)
				if err := _NpETH.contract.UnpackLog(event, "SetFeeToEvent", log); err != nil {
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

// ParseSetFeeToEvent is a log parse operation binding the contract event 0x33e1eeb6c5b29228be2616c47a3a094c7a3c4b81a8960fe62c1d5ade831dbe39.
//
// Solidity: event SetFeeToEvent(address oldFeeTo, address _newFeeTo)
func (_NpETH *NpETHFilterer) ParseSetFeeToEvent(log types.Log) (*NpETHSetFeeToEvent, error) {
	event := new(NpETHSetFeeToEvent)
	if err := _NpETH.contract.UnpackLog(event, "SetFeeToEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHSetLimitDepositLoopEventIterator is returned from FilterSetLimitDepositLoopEvent and is used to iterate over the raw logs and unpacked data for SetLimitDepositLoopEvent events raised by the NpETH contract.
type NpETHSetLimitDepositLoopEventIterator struct {
	Event *NpETHSetLimitDepositLoopEvent // Event containing the contract specifics and raw log

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
func (it *NpETHSetLimitDepositLoopEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHSetLimitDepositLoopEvent)
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
		it.Event = new(NpETHSetLimitDepositLoopEvent)
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
func (it *NpETHSetLimitDepositLoopEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHSetLimitDepositLoopEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHSetLimitDepositLoopEvent represents a SetLimitDepositLoopEvent event raised by the NpETH contract.
type NpETHSetLimitDepositLoopEvent struct {
	OldLimitDepositLoop uint32
	NewLimitDepositLoop uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetLimitDepositLoopEvent is a free log retrieval operation binding the contract event 0xee81468b2d09bb51681a19b4255734301b7fa838947a1d332ca98170dad3b631.
//
// Solidity: event SetLimitDepositLoopEvent(uint32 oldLimitDepositLoop, uint32 newLimitDepositLoop)
func (_NpETH *NpETHFilterer) FilterSetLimitDepositLoopEvent(opts *bind.FilterOpts) (*NpETHSetLimitDepositLoopEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "SetLimitDepositLoopEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHSetLimitDepositLoopEventIterator{contract: _NpETH.contract, event: "SetLimitDepositLoopEvent", logs: logs, sub: sub}, nil
}

// WatchSetLimitDepositLoopEvent is a free log subscription operation binding the contract event 0xee81468b2d09bb51681a19b4255734301b7fa838947a1d332ca98170dad3b631.
//
// Solidity: event SetLimitDepositLoopEvent(uint32 oldLimitDepositLoop, uint32 newLimitDepositLoop)
func (_NpETH *NpETHFilterer) WatchSetLimitDepositLoopEvent(opts *bind.WatchOpts, sink chan<- *NpETHSetLimitDepositLoopEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "SetLimitDepositLoopEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHSetLimitDepositLoopEvent)
				if err := _NpETH.contract.UnpackLog(event, "SetLimitDepositLoopEvent", log); err != nil {
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

// ParseSetLimitDepositLoopEvent is a log parse operation binding the contract event 0xee81468b2d09bb51681a19b4255734301b7fa838947a1d332ca98170dad3b631.
//
// Solidity: event SetLimitDepositLoopEvent(uint32 oldLimitDepositLoop, uint32 newLimitDepositLoop)
func (_NpETH *NpETHFilterer) ParseSetLimitDepositLoopEvent(log types.Log) (*NpETHSetLimitDepositLoopEvent, error) {
	event := new(NpETHSetLimitDepositLoopEvent)
	if err := _NpETH.contract.UnpackLog(event, "SetLimitDepositLoopEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHSetProtocolFeeEventIterator is returned from FilterSetProtocolFeeEvent and is used to iterate over the raw logs and unpacked data for SetProtocolFeeEvent events raised by the NpETH contract.
type NpETHSetProtocolFeeEventIterator struct {
	Event *NpETHSetProtocolFeeEvent // Event containing the contract specifics and raw log

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
func (it *NpETHSetProtocolFeeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHSetProtocolFeeEvent)
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
		it.Event = new(NpETHSetProtocolFeeEvent)
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
func (it *NpETHSetProtocolFeeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHSetProtocolFeeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHSetProtocolFeeEvent represents a SetProtocolFeeEvent event raised by the NpETH contract.
type NpETHSetProtocolFeeEvent struct {
	OldProtocolFee *big.Int
	NewProtocolFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetProtocolFeeEvent is a free log retrieval operation binding the contract event 0xc9578f161ee28120e271f491b542884e6298a2d0857eacb2a5af29099827b15b.
//
// Solidity: event SetProtocolFeeEvent(uint256 oldProtocolFee, uint256 newProtocolFee)
func (_NpETH *NpETHFilterer) FilterSetProtocolFeeEvent(opts *bind.FilterOpts) (*NpETHSetProtocolFeeEventIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "SetProtocolFeeEvent")
	if err != nil {
		return nil, err
	}
	return &NpETHSetProtocolFeeEventIterator{contract: _NpETH.contract, event: "SetProtocolFeeEvent", logs: logs, sub: sub}, nil
}

// WatchSetProtocolFeeEvent is a free log subscription operation binding the contract event 0xc9578f161ee28120e271f491b542884e6298a2d0857eacb2a5af29099827b15b.
//
// Solidity: event SetProtocolFeeEvent(uint256 oldProtocolFee, uint256 newProtocolFee)
func (_NpETH *NpETHFilterer) WatchSetProtocolFeeEvent(opts *bind.WatchOpts, sink chan<- *NpETHSetProtocolFeeEvent) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "SetProtocolFeeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHSetProtocolFeeEvent)
				if err := _NpETH.contract.UnpackLog(event, "SetProtocolFeeEvent", log); err != nil {
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

// ParseSetProtocolFeeEvent is a log parse operation binding the contract event 0xc9578f161ee28120e271f491b542884e6298a2d0857eacb2a5af29099827b15b.
//
// Solidity: event SetProtocolFeeEvent(uint256 oldProtocolFee, uint256 newProtocolFee)
func (_NpETH *NpETHFilterer) ParseSetProtocolFeeEvent(log types.Log) (*NpETHSetProtocolFeeEvent, error) {
	event := new(NpETHSetProtocolFeeEvent)
	if err := _NpETH.contract.UnpackLog(event, "SetProtocolFeeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHSetVerificationIterator is returned from FilterSetVerification and is used to iterate over the raw logs and unpacked data for SetVerification events raised by the NpETH contract.
type NpETHSetVerificationIterator struct {
	Event *NpETHSetVerification // Event containing the contract specifics and raw log

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
func (it *NpETHSetVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHSetVerification)
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
		it.Event = new(NpETHSetVerification)
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
func (it *NpETHSetVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHSetVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHSetVerification represents a SetVerification event raised by the NpETH contract.
type NpETHSetVerification struct {
	OldVerificatoin common.Address
	NewVerification common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetVerification is a free log retrieval operation binding the contract event 0x59968603049b5a22be25b140053bf34a3d104f2350cec1c64fd2a7aee52ad81a.
//
// Solidity: event SetVerification(address oldVerificatoin, address _newVerification)
func (_NpETH *NpETHFilterer) FilterSetVerification(opts *bind.FilterOpts) (*NpETHSetVerificationIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "SetVerification")
	if err != nil {
		return nil, err
	}
	return &NpETHSetVerificationIterator{contract: _NpETH.contract, event: "SetVerification", logs: logs, sub: sub}, nil
}

// WatchSetVerification is a free log subscription operation binding the contract event 0x59968603049b5a22be25b140053bf34a3d104f2350cec1c64fd2a7aee52ad81a.
//
// Solidity: event SetVerification(address oldVerificatoin, address _newVerification)
func (_NpETH *NpETHFilterer) WatchSetVerification(opts *bind.WatchOpts, sink chan<- *NpETHSetVerification) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "SetVerification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHSetVerification)
				if err := _NpETH.contract.UnpackLog(event, "SetVerification", log); err != nil {
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

// ParseSetVerification is a log parse operation binding the contract event 0x59968603049b5a22be25b140053bf34a3d104f2350cec1c64fd2a7aee52ad81a.
//
// Solidity: event SetVerification(address oldVerificatoin, address _newVerification)
func (_NpETH *NpETHFilterer) ParseSetVerification(log types.Log) (*NpETHSetVerification, error) {
	event := new(NpETHSetVerification)
	if err := _NpETH.contract.UnpackLog(event, "SetVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHStakeEventIterator is returned from FilterStakeEvent and is used to iterate over the raw logs and unpacked data for StakeEvent events raised by the NpETH contract.
type NpETHStakeEventIterator struct {
	Event *NpETHStakeEvent // Event containing the contract specifics and raw log

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
func (it *NpETHStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHStakeEvent)
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
		it.Event = new(NpETHStakeEvent)
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
func (it *NpETHStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHStakeEvent represents a StakeEvent event raised by the NpETH contract.
type NpETHStakeEvent struct {
	From         common.Address
	Amount       *big.Int
	AmountToMint *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeEvent is a free log retrieval operation binding the contract event 0x9dbaf9c586508abc91d6ee4e67d3c7a82ccb09bca5d9fe2c3b690f27b7e0a256.
//
// Solidity: event StakeEvent(address indexed _from, uint256 _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) FilterStakeEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHStakeEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "StakeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHStakeEventIterator{contract: _NpETH.contract, event: "StakeEvent", logs: logs, sub: sub}, nil
}

// WatchStakeEvent is a free log subscription operation binding the contract event 0x9dbaf9c586508abc91d6ee4e67d3c7a82ccb09bca5d9fe2c3b690f27b7e0a256.
//
// Solidity: event StakeEvent(address indexed _from, uint256 _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) WatchStakeEvent(opts *bind.WatchOpts, sink chan<- *NpETHStakeEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "StakeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHStakeEvent)
				if err := _NpETH.contract.UnpackLog(event, "StakeEvent", log); err != nil {
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

// ParseStakeEvent is a log parse operation binding the contract event 0x9dbaf9c586508abc91d6ee4e67d3c7a82ccb09bca5d9fe2c3b690f27b7e0a256.
//
// Solidity: event StakeEvent(address indexed _from, uint256 _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) ParseStakeEvent(log types.Log) (*NpETHStakeEvent, error) {
	event := new(NpETHStakeEvent)
	if err := _NpETH.contract.UnpackLog(event, "StakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHSubmitEventIterator is returned from FilterSubmitEvent and is used to iterate over the raw logs and unpacked data for SubmitEvent events raised by the NpETH contract.
type NpETHSubmitEventIterator struct {
	Event *NpETHSubmitEvent // Event containing the contract specifics and raw log

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
func (it *NpETHSubmitEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHSubmitEvent)
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
		it.Event = new(NpETHSubmitEvent)
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
func (it *NpETHSubmitEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHSubmitEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHSubmitEvent represents a SubmitEvent event raised by the NpETH contract.
type NpETHSubmitEvent struct {
	From         common.Address
	Amount       *big.Int
	AmountToMint *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubmitEvent is a free log retrieval operation binding the contract event 0x8794c1d5cd4c9d7f866ffcc0b772ecacde50cbbbda5084d5d82cd853a42cc96e.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) FilterSubmitEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHSubmitEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "SubmitEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHSubmitEventIterator{contract: _NpETH.contract, event: "SubmitEvent", logs: logs, sub: sub}, nil
}

// WatchSubmitEvent is a free log subscription operation binding the contract event 0x8794c1d5cd4c9d7f866ffcc0b772ecacde50cbbbda5084d5d82cd853a42cc96e.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) WatchSubmitEvent(opts *bind.WatchOpts, sink chan<- *NpETHSubmitEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "SubmitEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHSubmitEvent)
				if err := _NpETH.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
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

// ParseSubmitEvent is a log parse operation binding the contract event 0x8794c1d5cd4c9d7f866ffcc0b772ecacde50cbbbda5084d5d82cd853a42cc96e.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 _amount, uint256 _amountToMint)
func (_NpETH *NpETHFilterer) ParseSubmitEvent(log types.Log) (*NpETHSubmitEvent, error) {
	event := new(NpETHSubmitEvent)
	if err := _NpETH.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NpETH contract.
type NpETHTransferIterator struct {
	Event *NpETHTransfer // Event containing the contract specifics and raw log

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
func (it *NpETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHTransfer)
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
		it.Event = new(NpETHTransfer)
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
func (it *NpETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHTransfer represents a Transfer event raised by the NpETH contract.
type NpETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NpETH *NpETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NpETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NpETHTransferIterator{contract: _NpETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NpETH *NpETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NpETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHTransfer)
				if err := _NpETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NpETH *NpETHFilterer) ParseTransfer(log types.Log) (*NpETHTransfer, error) {
	event := new(NpETHTransfer)
	if err := _NpETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NpETH contract.
type NpETHUnpausedIterator struct {
	Event *NpETHUnpaused // Event containing the contract specifics and raw log

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
func (it *NpETHUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHUnpaused)
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
		it.Event = new(NpETHUnpaused)
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
func (it *NpETHUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHUnpaused represents a Unpaused event raised by the NpETH contract.
type NpETHUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NpETH *NpETHFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NpETHUnpausedIterator, error) {

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NpETHUnpausedIterator{contract: _NpETH.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NpETH *NpETHFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NpETHUnpaused) (event.Subscription, error) {

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHUnpaused)
				if err := _NpETH.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NpETH *NpETHFilterer) ParseUnpaused(log types.Log) (*NpETHUnpaused, error) {
	event := new(NpETHUnpaused)
	if err := _NpETH.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NpETHUnstakeEventIterator is returned from FilterUnstakeEvent and is used to iterate over the raw logs and unpacked data for UnstakeEvent events raised by the NpETH contract.
type NpETHUnstakeEventIterator struct {
	Event *NpETHUnstakeEvent // Event containing the contract specifics and raw log

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
func (it *NpETHUnstakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NpETHUnstakeEvent)
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
		it.Event = new(NpETHUnstakeEvent)
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
func (it *NpETHUnstakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NpETHUnstakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NpETHUnstakeEvent represents a UnstakeEvent event raised by the NpETH contract.
type NpETHUnstakeEvent struct {
	From                common.Address
	AmountNpToBurn      *big.Int
	AmountEth           *big.Int
	TotalWithdrawReqEth *big.Int
	WithdrawReqIndex    *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnstakeEvent is a free log retrieval operation binding the contract event 0x4c9d4b85f38031af67ea3a40072b077bd130e1eb6489016053e8c36b412f8609.
//
// Solidity: event UnstakeEvent(address indexed _from, uint256 _amountNpToBurn, uint256 _amountEth, uint256 totalWithdrawReqEth, uint256 withdrawReqIndex)
func (_NpETH *NpETHFilterer) FilterUnstakeEvent(opts *bind.FilterOpts, _from []common.Address) (*NpETHUnstakeEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.FilterLogs(opts, "UnstakeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return &NpETHUnstakeEventIterator{contract: _NpETH.contract, event: "UnstakeEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeEvent is a free log subscription operation binding the contract event 0x4c9d4b85f38031af67ea3a40072b077bd130e1eb6489016053e8c36b412f8609.
//
// Solidity: event UnstakeEvent(address indexed _from, uint256 _amountNpToBurn, uint256 _amountEth, uint256 totalWithdrawReqEth, uint256 withdrawReqIndex)
func (_NpETH *NpETHFilterer) WatchUnstakeEvent(opts *bind.WatchOpts, sink chan<- *NpETHUnstakeEvent, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _NpETH.contract.WatchLogs(opts, "UnstakeEvent", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NpETHUnstakeEvent)
				if err := _NpETH.contract.UnpackLog(event, "UnstakeEvent", log); err != nil {
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

// ParseUnstakeEvent is a log parse operation binding the contract event 0x4c9d4b85f38031af67ea3a40072b077bd130e1eb6489016053e8c36b412f8609.
//
// Solidity: event UnstakeEvent(address indexed _from, uint256 _amountNpToBurn, uint256 _amountEth, uint256 totalWithdrawReqEth, uint256 withdrawReqIndex)
func (_NpETH *NpETHFilterer) ParseUnstakeEvent(log types.Log) (*NpETHUnstakeEvent, error) {
	event := new(NpETHUnstakeEvent)
	if err := _NpETH.contract.UnpackLog(event, "UnstakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
