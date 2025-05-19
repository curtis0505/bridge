package validator

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/libs/common"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/bridge/abi"
	"math/big"
	"strings"
)

func (p *Validator) Deposit(eventLog commontypes.Log) error {
	commonLog, err := bridgetypes.GetDepositCommonLog(eventLog, p.Contract)
	if err != nil {
		p.logger.Error("GetDepositCommonLog", err, "eventLog", eventLog)
		return err
	}

	// validate transaction
	err = p.ValidationTransaction(
		context.Background(),
		commonLog,
		p.Account[commonLog.ToChainName].Address,
	)
	if err != nil {
		p.logger.Error("GetCheckTransaction", err, "commonLog", commonLog)
		return err
	}

	fromContractAddress, err := p.Contract.GetChain(commontypes.Chain(commonLog.FromChainName))
	if err != nil {
		p.logger.Error("p.Contract.GetChain", err, "chain", commonLog.FromChainName)
		return err
	}

	switch commontypes.GetChainType(commonLog.ToChainName) {
	case commontypes.ChainTypeEVM:
		if commonLog.ToChainName == commontypes.ChainBASE &&
			strings.EqualFold(commonLog.FromTokenAddr.String(), fromContractAddress.NptToken) {
			return p.NptFromKlayToBase(commonLog)
		} else {
			return p.MintToEVM(commonLog)
		}
	case commontypes.ChainTypeCOSMOS:
		return fmt.Errorf("not a supported chain") //return p.MintToCOSMOS(commonLog)
	default:
		return fmt.Errorf("not a supported chain")
	}
}

func (p *Validator) DepositRestakeToken(eventLog commontypes.Log) error {
	commonLog, err := bridgetypes.GetDepositCommonLog(eventLog, p.Contract)
	if err != nil {
		p.logger.Error("GetDepositCommonLog", err, "eventLog", eventLog)
		return err
	}

	// validate transaction
	err = p.ValidationTransaction(
		context.Background(),
		commonLog,
		p.Account[commonLog.ToChainName].Address,
	)
	if err != nil {
		p.logger.Error("GetCheckTransaction", err, "commonLog", commonLog)
		return err
	}

	switch commontypes.GetChainType(commonLog.ToChainName) {
	case commontypes.ChainTypeEVM:
		return p.MintToEVM(commonLog)
	case commontypes.ChainTypeCOSMOS:
		return fmt.Errorf("not a supported chain") //return p.MintToCOSMOS(commonLog)
	default:
		return fmt.Errorf("not a supported chain")
	}
}

func (p *Validator) Burn(eventLog commontypes.Log) error {
	commonLog, err := bridgetypes.GetBurnCommonLog(eventLog, p.Contract)
	if err != nil {
		p.logger.Error("GetBurnCommonLog", err, "eventLog", eventLog)
		return fmt.Errorf("GetBurnLogToCommonEventLog: %w", err)
	}

	// validate transaction
	err = p.ValidationTransaction(
		context.Background(),
		commonLog,
		p.Account[commonLog.ToChainName].Address,
	)
	if err != nil {
		p.logger.Error("GetCheckTransaction", err, "commonLog", commonLog)
		return fmt.Errorf("ValidationTransaction(address:%s): %w", p.Account[commonLog.ToChainName].Address, err)
	}

	switch commontypes.GetChainType(commonLog.ToChainName) {
	case commontypes.ChainTypeEVM:
		return p.WithdrawToEVM(commonLog)
	case commontypes.ChainTypeCOSMOS:
		return fmt.Errorf("not a supported chain") //p.WithdrawToCOSMOS(commonLog)
	default:
		return fmt.Errorf("not a supported chain")
	}
}

func (p *Validator) Mint(eventLog commontypes.Log) error {
	p.logger.Info(
		"eventName", eventLog.EventName,
		"chain", eventLog.Chain(),
		"txHash", eventLog.TxHash(),
	)
	return nil
}

func (p *Validator) Withdraw(eventLog commontypes.Log) error {
	p.logger.Info(
		"eventName", eventLog.EventName,
		"chain", eventLog.Chain(),
		"txHash", eventLog.TxHash(),
	)
	return nil
}

func (p *Validator) ValidationTransaction(ctx context.Context, commonLog bridgetypes.CommonEventLog, validatorAddress string) error {
	txReceipt, isPending, err := p.client.GetTransactionWithReceipt(ctx, commonLog.FromChainName, commonLog.TxHash)
	if err != nil {
		return err
	}

	if isPending {
		return fmt.Errorf("tx is pending")
	}

	if !txReceipt.Receipt().Success() {
		return fmt.Errorf("tx status is not success")
	}

	switch commontypes.GetChainType(commonLog.ToChainName) {
	case commontypes.ChainTypeEVM:
		ctx := context.Background()
		chain := commonLog.ToChainName

		contractAddress, err := p.Contract.GetChain(commontypes.Chain(chain))
		if err != nil {
			return err
		}

		multiSigWalletAbi := abi.GetAbiToMap(abi.MultiSigWalletAbi)

		msg, err := p.client.CallMsg(
			ctx,
			chain,
			"",
			contractAddress.MultiSigWallet,
			"txHashIdMap",
			multiSigWalletAbi,
			common.HexToHash(commonLog.TxHash),
		)
		if err != nil {
			return commontypes.WrapError("GetTxId", err)
		}

		txId, ok := msg[0].(*big.Int)
		if ok == false {
			return commontypes.WrapError("GetTxId Type", err)
		}

		// txId 가 0 이 아니면, 이미 해당 트랜잭션 Hash 로 txId 가 생성되었음
		if txId.Cmp(big.NewInt(0)) != 0 {
			msg, err := p.client.CallMsg(
				ctx,
				chain,
				"",
				contractAddress.MultiSigWallet,
				"transactions",
				multiSigWalletAbi,
				txId,
			)
			if err != nil {
				return commontypes.WrapError(fmt.Sprintf("GetTransaction : txId - %s", txId), err)
			}

			transaction := bridgetypes.TransactionResponse{
				Proof:    msg[0].([32]byte),
				Executed: msg[1].(bool),
			}

			if transaction.Executed {
				return commontypes.WrapError(fmt.Sprintf("transaction.Executed : txId - %s, txHash : %s", txId, commonLog.TxHash), fmt.Errorf("already confirmation"))
			}

			msg3, err := p.client.CallMsg(
				ctx,
				chain,
				"",
				contractAddress.MultiSigWallet,
				"confirmations",
				multiSigWalletAbi,
				txId,
				common.HexToAddress(chain, validatorAddress),
			)
			if err != nil {
				return commontypes.WrapError(fmt.Sprintf("GetConfirmation - txId: %s, address : %s", txId, validatorAddress), err)
			}

			confirmation, ok := msg3[0].(bool)
			if ok == false {
				return commontypes.WrapError(fmt.Sprintf("GetConfirmation Type - txId: %s, address : %s", txId, validatorAddress), err)
			}

			if confirmation {
				return commontypes.WrapError(fmt.Sprintf("transaction.Executed - txId: %s, txHash : %s", txId, commonLog.TxHash), fmt.Errorf("already confirmation"))
			}
		}
	case commontypes.ChainTypeCOSMOS:
		//validate
	}

	return nil
}
