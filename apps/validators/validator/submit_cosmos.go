package validator

import (
	"context"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/curtis0505/bridge/libs/client/chain/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	"strings"
)

func (p *Validator) MintToCOSMOS(commonLog bridgetypes.CommonEventLog, multiSigAddress string) (signing.SignatureV2, error) {
	contractAddress, err := p.Contract.GetChain(commontypes.Chain(commonLog.ToChainName))
	if err != nil {
		return signing.SignatureV2{}, err
	}

	sign, _, err := p.SignMultiSig(
		commonLog,
		multiSigAddress,
		cosmostypes.NewMsgExecuteContract(
			multiSigAddress,
			contractAddress.Minter,
			bridge.MsgMint{
				ParentChainName: commonLog.FromChainName,
				ParentTokenAddr: strings.ToLower(commonLog.FromTokenAddr.String()),
				ParentTx:        commonLog.TxHash,
				FromAddr:        commonLog.From.String(),
				ToAddr:          cosmoscommon.FromAddress(commonLog.ToChainName, commonLog.ToAddr.Bytes()).String(),
				Amount:          commonLog.Amount.String(),
			},
			cosmossdk.NewCoins(),
		),
	)
	if err != nil {
		p.logger.Error("SignMultiSig(Mint)", err, "commonLog", commonLog)
		return sign, err
	}

	p.logger.Info(
		"success", "MintToCOSMOS",
		"contractAddress", contractAddress.Minter,
		"toChainName", commonLog.ToChainName,
		"eventMethod", "deposit",
		"txHash", commonLog.TxHash,
		"fromChainName", commonLog.FromChainName,
		"executeMethod", "mint",
	)

	return sign, nil
}

func (p *Validator) WithdrawToCOSMOS(commonLog bridgetypes.CommonEventLog, multiSigAddress string) (signing.SignatureV2, error) {
	contractAddress, err := p.Contract.GetChain(commontypes.Chain(commonLog.ToChainName))
	if err != nil {
		return signing.SignatureV2{}, err
	}

	response := bridge.QueryCoinResponse{}
	p.client.CallWasm(context.Background(), commonLog.ToChainName, contractAddress.Vault, bridge.QueryCoinRequest{
		ChildChainName: commonLog.FromChainName,
		ChildTokenAddr: strings.ToLower(commonLog.FromTokenAddr.String()),
	}, &response)

	var sign signing.SignatureV2
	if response.NativeDenom == "" {
		sign, _, err = p.SignMultiSig(
			commonLog,
			multiSigAddress,
			cosmostypes.NewMsgExecuteContract(
				multiSigAddress,
				contractAddress.Vault,
				bridge.MsgWithdraw{
					ChildChainName: commonLog.FromChainName,
					ChildTokenAddr: strings.ToLower(commonLog.FromTokenAddr.String()),
					ChildTx:        commonLog.TxHash,
					FromAddr:       commonLog.From.String(),
					ToAddr:         cosmoscommon.FromAddress(commonLog.ToChainName, commonLog.ToAddr.Bytes()).String(),
					Amount:         commonLog.Amount.String(),
				},
				cosmossdk.NewCoins(),
			),
		)
		if err != nil {
			p.logger.Error("SignMultiSig(Withdraw)", err, "commonLog", commonLog)
			return sign, err
		}
	} else {
		sign, _, err = p.SignMultiSig(
			commonLog,
			multiSigAddress,
			cosmostypes.NewMsgExecuteContract(
				multiSigAddress,
				contractAddress.Vault,
				bridge.MsgWithdrawCoin{
					ChildChainName: commonLog.FromChainName,
					ChildTokenAddr: strings.ToLower(commonLog.FromTokenAddr.String()),
					ChildTx:        commonLog.TxHash,
					FromAddr:       commonLog.From.String(),
					ToAddr:         cosmoscommon.FromAddress(commonLog.ToChainName, commonLog.ToAddr.Bytes()).String(),
					Amount:         commonLog.Amount.String(),
				},
				cosmossdk.NewCoins(),
			),
		)
		if err != nil {
			p.logger.Error("SignMultiSig(Withdraw)", err, "commonLog", commonLog)
			return sign, err
		}
	}

	p.logger.Info(
		"success", "WithdrawToCOSMOS",
		"contractAddress", contractAddress.Vault,
		"toChainName", commonLog.ToChainName,
		"eventMethod", "burn",
		"txHash", commonLog.TxHash,
		"fromChainName", commonLog.FromChainName,
		"executeMethod", "withdraw",
	)

	return sign, nil
}

func (p *Validator) SignMultiSig(commonLog bridgetypes.CommonEventLog, multiSigAddress string, msg *wasmtypes.MsgExecuteContract) (signing.SignatureV2, *types.SendTxAsyncResult, error) {
	signature, err := p.client.SignMultiSig(
		context.Background(),
		commonLog.ToChainName,
		p.Account[commonLog.ToChainName],
		multiSigAddress,
		msg,
	)
	if err != nil {
		return signing.SignatureV2{}, nil, commontypes.WrapError("GetSignedTransaction", err)
	}

	p.logger.Info("multiSigAddress", multiSigAddress)

	return signature, nil, nil
}
