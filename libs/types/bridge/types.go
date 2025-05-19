package bridge

import (
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/common"
	commontypes "github.com/curtis0505/bridge/libs/types"
	cosmosbridgetypes "github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	"github.com/curtis0505/bridge/libs/types/restaking"
	"github.com/curtis0505/bridge/libs/util"
	"math/big"
	"slices"
	"strings"
)

const (
	TransactionAlreadyExecuted = "transaction already executed"
)

const (
	EventNameWithdraw            = "Withdraw"
	EventNameDeposit             = "Deposit"
	EventNameBurn                = "Burn"
	EventNameMint                = "Mint"
	EventNameDepositRestakeToken = "DepositRestakeToken"

	EventNameWasmWithdraw     = "withdraw"
	EventNameWasmDeposit      = "deposit"
	EventNameWasmDepositCoin  = "deposit_coin"
	EventNameWasmWithdrawCoin = "withdraw_coin"
	EventNameWasmMint         = "mint"
	EventNameWasmBurn         = "burn"

	EventNameConfirmation = "Confirmation"
	EventNameSubmission   = "Submission"
	EventNameExecution    = "Execution"

	MethodNameBridgeDeposit         = "deposit"
	MethodNameBridgeDepositToken    = "depositToken"
	MethodNameBridgeBurn            = "burn"
	MethodNameBridgeWithdraw        = "withdraw"
	MethodNameFxPortalDepositERC20  = "deposit"
	MethodNameFxPortalWithdrawERC20 = "withdraw"

	MethodNameFxPortalChainFee   = "chainFee"
	MethodNameFxPortalChainFeeTo = "chainFeeTo"
	MethodNameFxPortalTaxRateBP  = "taxRateBP"
	MethodNameFxPortalTaxTo      = "taxTo"

	VaultContractID            = "vault"
	MinterContractID           = "minter"
	MultisigWalletContractID   = "multisigwallet"
	FxERC20RootTunnelCustomID  = "fxerc20-root-tunnel-custom"
	FxERC20ChildTunnelCustomID = "fxerc20-child-tunnel-custom"
)

const (
	ZeroAddress = "0x0000000000000000000000000000000000000000"
	CoinAddress = "0x0000000000000000000000000000000000000001" // contract 에서 코인은 위 주소로 저장 되어 있음
)

const (
	Maintenance = "bridge"
)

type BridgeType int

const (
	BridgeTypeDefault BridgeType = iota
	BridgeTypeVaultMinter
	BridgeTypeFxportal
)

type BridgeRole int

const (
	BridgeRoleDefault BridgeRole = iota
	BridgeRoleVault
	BridgeRoleMinter
	BridgeRoleFxPortalRootTunnel
	BridgeRoleFxPortalChildTunnel
)

type Version string

const (
	VersionBridge        Version = "Bridge"
	VersionRestakeBridge         = "RestakeBridge"
)

type CommonEventLog struct {
	Destination   string
	FromChainName string
	ToChainName   string
	From          common.Address
	ToAddr        common.Address
	FromTokenAddr common.Address
	ToTokenAddr   common.Address
	TxHash        string
	Amount        *big.Int
	Decimal       *big.Int
	Version       Version
}

func GetEventType(eventName string) string {
	switch eventName {
	case EventNameDeposit, EventNameWasmDeposit, EventNameWasmDepositCoin:
		return EventNameDeposit
	case EventNameDepositRestakeToken:
		return EventNameDepositRestakeToken
	case EventNameWithdraw, EventNameWasmWithdraw, EventNameWasmWithdrawCoin:
		return EventNameWithdraw
	case EventNameBurn, EventNameWasmBurn:
		return EventNameBurn
	case EventNameMint, EventNameWasmMint:
		return EventNameMint
	default:
		return ""
	}
}

func GetInTxHash(chain string, txHash common.Hash) string {
	if commontypes.GetChainType(chain) == commontypes.ChainTypeCOSMOS {
		return strings.ToUpper(txHash.String())[2:]
	}
	return txHash.String()
}

func bridgeChains() map[string][]string {
	return map[string][]string{
		"EVM": {
			commontypes.ChainETH,
			commontypes.ChainMATIC,
			commontypes.ChainKLAY,
		},
		"COSMOS": {
			commontypes.ChainFNSA,
			commontypes.ChainTFNSA,
		},
	}
}

func GetBridgeChains() []string {
	chains := bridgeChains()
	return append(chains["EVM"], chains["COSMOS"]...)
}

func GetBridgeEVMChains() []string {
	return bridgeChains()["EVM"]
}

func GetBridgeCosmosChains() []string {
	return bridgeChains()["COSMOS"]
}

func GetBridgeContractIDs() []string {
	return []string{
		VaultContractID,
		MinterContractID,
		MultisigWalletContractID,
		FxERC20RootTunnelCustomID,
		FxERC20ChildTunnelCustomID,
		restaking.RestakeVaultContractID,
		restaking.RestakeMinterContractID,
	}
}

func GetDepositCommonLog(eventLog commontypes.Log, contracts ContractAddresses) (CommonEventLog, error) {
	switch commontypes.GetChainType(eventLog.Chain()) {
	case commontypes.ChainTypeEVM:
		eventDeposit := EventDeposit{}
		err := eventLog.Unmarshal(&eventDeposit)
		if err != nil {
			return CommonEventLog{}, commontypes.WrapError("Unmarshal", err)
		}

		fromContractAddress, err := contracts.GetChain(commontypes.Chain(eventLog.Chain()))
		if err != nil {
			return CommonEventLog{}, err
		}

		toContractAddress, err := contracts.GetChain(commontypes.Chain(eventDeposit.ToChainName))
		if err != nil {
			return CommonEventLog{}, err
		}

		tokenAddr, err := eventDeposit.SafeTokenAddr()
		if err != nil {
			return CommonEventLog{}, err
		}

		amount, err := eventDeposit.SafeAmount()
		if err != nil {
			return CommonEventLog{}, err
		}

		destination := ""
		switch eventDeposit.BridgeVersion() {
		case VersionBridge:
			if eventDeposit.ToChainName == commontypes.ChainBASE &&
				strings.EqualFold(eventDeposit.TokenAddr.String(), fromContractAddress.NptToken) {
				destination = toContractAddress.Vault
			} else {
				destination = toContractAddress.Minter
			}
		case VersionRestakeBridge:
			destination = toContractAddress.RestakeMinter
		default:
			return CommonEventLog{}, errors.New("invalid bridge version")
		}

		return CommonEventLog{
			Destination: destination,
			ToTokenAddr: func() common.Address {
				if eventDeposit.ToChainName == commontypes.ChainBASE &&
					strings.EqualFold(eventDeposit.TokenAddr.String(), fromContractAddress.NptToken) {
					return common.HexToAddress(commontypes.ChainBASE, toContractAddress.NptToken)
				} else {
					return nil
				}
			}(),
			FromChainName: eventLog.Chain(),
			From:          eventDeposit.FromAddr,
			ToChainName:   eventDeposit.ToChainName,
			ToAddr:        common.HexToAddress(eventLog.Chain(), eventDeposit.To.String()),
			FromTokenAddr: tokenAddr,
			TxHash:        eventLog.TxHash(),
			Amount:        amount,
			Decimal:       big.NewInt(int64(eventDeposit.Decimal)),
			Version:       eventDeposit.BridgeVersion(),
		}, nil
	case commontypes.ChainTypeCOSMOS:
		eventDeposit := cosmosbridgetypes.EventDeposit{}
		err := eventLog.Unmarshal(&eventDeposit)
		if err != nil {
			return CommonEventLog{}, commontypes.WrapError("Unmarshal", err)
		}

		var fromTokenAddr common.Address
		if eventLog.EventName == EventNameWasmDepositCoin {
			fromTokenAddr = common.HexToAddress(eventDeposit.ChildChainName, CoinAddress)
		} else {
			fromTokenAddr = common.HexToAddress(eventLog.Chain(), eventDeposit.TokenAddr)
		}

		contractAddress, err := contracts.GetChain(commontypes.Chain(eventDeposit.ChildChainName))
		if err != nil {
			return CommonEventLog{}, err
		}

		return CommonEventLog{
			Destination:   contractAddress.Minter,
			FromChainName: eventLog.Chain(),
			From:          common.HexToAddress(eventDeposit.ChildChainName, eventDeposit.FromAddr),
			ToChainName:   eventDeposit.ChildChainName,
			ToAddr:        common.HexToAddress(eventDeposit.ChildChainName, eventDeposit.ToAddr),
			ToTokenAddr:   common.HexToAddress(eventDeposit.ChildChainName, eventDeposit.ChildTokenAddr),
			FromTokenAddr: fromTokenAddr,
			TxHash:        eventLog.TxHash(),
			Amount:        util.ToBigInt(eventDeposit.Amount),
			Decimal:       util.ToBigInt(eventDeposit.Decimals),
			Version:       VersionBridge,
		}, nil
	}

	return CommonEventLog{}, fmt.Errorf("invalid chain")
}

func GetBurnCommonLog(eventLog commontypes.Log, contracts ContractAddresses) (CommonEventLog, error) {
	switch commontypes.GetChainType(eventLog.Chain()) {
	case commontypes.ChainTypeEVM:
		eventBurn := EventBurn{}
		err := eventLog.Unmarshal(&eventBurn)
		if err != nil {
			return CommonEventLog{}, commontypes.WrapError("Unmarshal", err)
		}

		contractAddress, err := contracts.GetChain(commontypes.Chain(eventBurn.ToChainName))
		if err != nil {
			return CommonEventLog{}, err
		}

		fromTokenAddr, err := eventBurn.SafeTokenAddr()
		if err != nil {
			return CommonEventLog{}, err
		}

		toToken, err := eventBurn.SafeToken()
		if err != nil {
			return CommonEventLog{}, err
		}

		return CommonEventLog{
			Destination: func() string {
				if eventBurn.TokenAddr != nil {
					return contractAddress.Vault
				} else {
					return contractAddress.RestakeVault
				}
			}(),
			FromChainName: eventLog.Chain(),
			From:          eventBurn.FromAddr,
			ToChainName:   eventBurn.ToChainName,
			ToAddr:        common.HexToAddress(eventLog.Chain(), eventBurn.To.String()),
			FromTokenAddr: fromTokenAddr,
			ToTokenAddr:   common.HexToAddress(eventBurn.ToChainName, toToken.String()),
			TxHash:        eventLog.TxHash(),
			Amount:        eventBurn.Amount,
			Decimal:       big.NewInt(int64(eventBurn.Decimal)),
			Version:       VersionBridge,
		}, nil
	case commontypes.ChainTypeCOSMOS:
		eventBurn := cosmosbridgetypes.EventBurn{}
		err := eventLog.Unmarshal(&eventBurn)
		if err != nil {
			return CommonEventLog{}, commontypes.WrapError("Unmarshal", err)
		}

		contractAddress, err := contracts.GetChain(commontypes.Chain(eventBurn.ParentChainName))
		if err != nil {
			return CommonEventLog{}, err
		}

		return CommonEventLog{
			Destination:   contractAddress.Vault,
			FromChainName: eventLog.Chain(),
			From:          common.HexToAddress(eventBurn.ParentChainName, eventBurn.FromAddr),
			ToChainName:   eventBurn.ParentChainName,
			ToAddr:        common.HexToAddress(eventBurn.ParentChainName, eventBurn.ToAddr),
			FromTokenAddr: common.HexToAddress(eventBurn.ParentChainName, eventBurn.ParentTokenAddr),
			TxHash:        eventLog.TxHash(),
			Amount:        util.ToBigInt(eventBurn.Amount),
			Decimal:       util.ToBigInt(eventBurn.Decimals),
			Version:       VersionBridge,
			//ToTokenAddr: "", //to token address query 조회
		}, nil
	}
	return CommonEventLog{}, fmt.Errorf("invalid chain")
}

type ContractAddress struct {
	NptToken              string
	Vault                 string
	Minter                string
	MultiSigWallet        string
	RestakeVault          string
	RestakeMinter         string
	RestakeMultiSigWallet string
}

type ContractAddresses map[commontypes.Chain]*ContractAddress

func (c ContractAddresses) GetChain(chain commontypes.Chain) (*ContractAddress, error) {
	if slices.Contains(commontypes.SupportedChains, chain) {
		contract, ok := c[chain]
		if ok == false {
			return nil, fmt.Errorf("chain %s not included", chain)
		} else {
			return contract, nil
		}
	} else {
		return nil, fmt.Errorf("chain %s not supported", chain)
	}
}
