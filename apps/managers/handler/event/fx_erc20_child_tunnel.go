package event

import (
	"context"
	"fmt"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/common"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/entity"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"time"
)

// FxChildWithdrawERC20
// Chain :MATIC(MATIC -> ETH)
func (p *EventHandler) FxChildWithdrawERC20(log types.Log) error {
	ctx := context.Background()
	eventFxChildWithdrawERC20 := bridgetypes.EventFxChildWithdrawERC20{}
	err := log.Unmarshal(&eventFxChildWithdrawERC20)
	if err != nil {
		logger.Error("FxChildWithdrawERC20", "Unmarshal", err)
		return err
	}
	tx, _, err := p.client.GetTransactionWithReceipt(context.Background(), log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "GetTransactionWithReceipt", err)
	}

	contract, err := cache.ContractCache().GetContractByContractID(log.Chain(), bridgetypes.FxERC20ChildTunnelCustomID)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "GetContractByContractId", err)
	}

	var inputWithdraw bridgetypes.InputWithdrawFxERC20ChildTunnel
	err = tx.UnmarshalABI(contract.ABI, &inputWithdraw)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "UnmarshalABI", err)
	}

	header, err := p.client.HeaderByNumber(context.Background(), log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "GetBlockHeader", err)
	}

	blockNumber := header.BlockNumber()

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventFxChildWithdrawERC20.ChildToken.String())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "error", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeIn). // fxchild withdraw
			SetEvent(eventtypes.EventFxChildWithdrawERC20).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(types.ChainMATIC).
			SetToChain(types.ChainETH).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info("FxChildWithdrawERC20",
			"txHash", log.TxHash(),
			"blockNumber", blockNumber,
			"amount", eventFxChildWithdrawERC20.Amount,
			"msgSender", eventFxChildWithdrawERC20.MsgSender,
			"receiver", eventFxChildWithdrawERC20.Receiver)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                          // Sender
			SetTo(tx.To()).                              // Vault
			SetToken(*tokenInfo).                        // Token
			SetAmount(eventFxChildWithdrawERC20.Amount). // Transfer Amount
			SetChain(log.Chain()).                       // Chain
			SetTxHash(log.TxHash()).                     // TxHash
			SetCate(protocol.TxHistoryCateBridgeIn).     // Category
			SetFromChain(log.Chain()).                   // From Tx
			SetFromTxHash(log.TxHash()).
			SetToChain(types.ChainETH).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventTransfer,
			"chain", transferInfo.Chain,
			"cate", transferInfo.Cate,
			"from", transferInfo.From,
			"to", transferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, transferInfo, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "UpsertBridgeTransferHistory", err)
		}

		// Tax Transfer Event
		taxTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                                                                  // Sender
			SetTo(tx.To()).                                                                      // Tax TODO
			SetToken(*tokenInfo).                                                                // Token
			SetAmount(new(big.Int).Sub(inputWithdraw.Amount, eventFxChildWithdrawERC20.Amount)). // Transfer Amount
			SetChain(log.Chain()).                                                               // Chain
			SetTxHash(log.TxHash()).                                                             // TxHash
			SetCate(protocol.TxHistoryCateBridgeTax).                                            // Category
			SetFromChain(log.Chain()).                                                           // Source Tx
			SetFromTxHash(log.TxHash()).
			SetToChain(types.ChainETH).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventTransfer,
			"chain", taxTransferInfo.Chain,
			"cate", taxTransferInfo.Cate,
			"from", taxTransferInfo.From,
			"to", taxTransferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, taxTransferInfo, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "UpsertBridgeTransferHistory", err)
		}

		chainFeeAmount, err := tx.Value()
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "Value", fmt.Errorf("value is nil"))
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetFrom(tx.From()).
			SetTo(tx.From()).
			SetSourceChain(types.ChainMATIC).
			SetSourceTxHash(log.TxHash()).
			SetSourceToken(common.HexToAddress(types.ChainMATIC, eventFxChildWithdrawERC20.ChildToken.String()).String()).
			SetTargetChain(types.ChainETH).
			SetTargetTxHash("").
			SetTargetToken("").
			SetAmount(inputWithdraw.Amount).
			SetChainFee(chainFeeAmount).
			SetTax(taxTransferInfo.Amount)

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", types.ChainMATIC, types.ChainETH)).
			AddKeyValueWidget("Event", log.EventName).
			AddKeyValueWidget("Chain", log.Chain()).
			AddKeyValueWidget("Address", log.Address()).
			AddKeyValueWidget("TxHash", log.TxHash()).
			AddKeyValueWidget("Error", fmt.Sprintf("%s %s", "WithMongoTransaction:", err.Error())).
			SendMessage()
		return err
	}

	return nil
}

// FxWithdrawERC20
// Chain: ETH(MATIC -> ETH)
func (p *EventHandler) FxWithdrawERC20(log types.Log) error {
	// transfer나 tax event 의 경우, handler/fxportal 참고
	eventSyncWithdraw := bridgetypes.EventFxWithdrawERC20{}
	err := log.Unmarshal(&eventSyncWithdraw)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxWithdrawERC20, "Unmarshal", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(context.Background(), log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxWithdrawERC20, "GetTransactionWithReceipt", err)
	}

	// handler/event에서 receiveMessage를 실행할 때 미리 BRIDGE-OUT cate의 txHash 업데이트를 해준다.
	txHistory, err := p.historyService.FindOneTxHistory(context.Background(), bson.M{
		"tx_hash": tx.TxHash(),
	})
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxWithdrawERC20, "txHistory", err)
	}
	activeHistory := mongoentity.NewBridgeActiveHistory().
		SetSourceChain(types.ChainMATIC).
		SetSourceTxHash(txHistory.InTxHash).
		SetTargetChain(types.ChainETH).
		SetTargetTxHash(log.TxHash()).
		SetTargetToken(common.HexToAddress(types.ChainETH, eventSyncWithdraw.RootToken.String()).String()).
		SetUpdatedAt(time.Now())

	err = p.historyService.UpsetBridgeActiveHistory(context.Background(), activeHistory)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxWithdrawERC20, "UpsetBridgeActiveHistory", err)
	}

	return nil
}

// FxDepositERC20
// Chain :ETH(ETH -> MATIC)
func (p *EventHandler) FxDepositERC20(log types.Log) error {
	ctx := context.Background()

	var eventDeposit bridgetypes.EventFxDepositERC20
	err := log.Unmarshal(&eventDeposit)
	if err != nil {
		p.logger.Error("FxPortalDepositERC20", "Unmarshal", err)
		return err
	}

	tx, _, err := p.client.GetTransactionWithReceipt(context.Background(), log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "GetTransactionWithReceipt", err)
	}

	contract, err := cache.ContractCache().GetContractByContractID(log.Chain(), bridgetypes.FxERC20RootTunnelCustomID)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "GetContractByContractId", err)
	}

	var inputDeposit bridgetypes.InputDepositFxERC20RootTunnel
	err = tx.UnmarshalABI(contract.ABI, &inputDeposit)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "UnmarshalABI", err)
	}

	header, err := p.client.HeaderByNumber(context.Background(), log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "GetBlockHeader", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventDeposit.RootToken.String())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "error", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Deposit Event
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeIn).
			SetEvent(eventtypes.EventFxDepositERC20).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(log.Chain()).
			SetFromTxHash(log.TxHash()).
			SetToChain(types.ChainMATIC).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			SetDepositNonce(eventDeposit.DepositNonce).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.EventFxDepositERC20,
			"chain", log.Chain(),
			"from", tx.From(),
			"txHash", tx.TxHash(),
			"fee", tx.TxFee(header.BaseFee()),
		)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxDepositERC20, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                      // Vault
			SetTo(eventDeposit.Depositor.String()).  // Recipient
			SetToken(*tokenInfo).                    // Token
			SetAmount(eventDeposit.Amount).          // Transfer Amount
			SetChain(log.Chain()).                   // Chain
			SetTxHash(log.TxHash()).                 // TxHash
			SetCate(protocol.TxHistoryCateBridgeIn). // Category
			SetFromChain(log.Chain()).
			SetFromTxHash(log.TxHash()).
			SetToChain(types.ChainMATIC).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventTransfer,
			"chain", transferInfo.Chain,
			"cate", transferInfo.Cate,
			"from", transferInfo.From,
			"to", transferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, transferInfo, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxDepositERC20, "UpsertBridgeTransferHistory", err)
		}

		// Tax Transfer Event
		taxTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                                                    // Sender
			SetTo(tx.To()).                                                        // Tax TODO
			SetToken(*tokenInfo).                                                  // Token
			SetAmount(new(big.Int).Sub(inputDeposit.Amount, eventDeposit.Amount)). // Transfer Amount
			SetChain(log.Chain()).                                                 // Chain
			SetTxHash(log.TxHash()).                                               // TxHash
			SetCate(protocol.TxHistoryCateBridgeTax).                              // Category
			SetFromChain(log.Chain()).                                             // Source Tx
			SetFromTxHash(log.TxHash()).
			SetToChain(types.ChainMATIC).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventTransfer,
			"chain", taxTransferInfo.Chain,
			"cate", taxTransferInfo.Cate,
			"from", taxTransferInfo.From,
			"to", taxTransferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, taxTransferInfo, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxDepositERC20, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetFrom(tx.From()).
			SetTo(tx.From()).
			SetSourceChain(types.ChainETH).
			SetSourceTxHash(log.TxHash()).
			SetSourceToken(eventDeposit.RootToken.String()).
			SetTargetChain(types.ChainMATIC).
			SetTargetTxHash("").
			SetTargetToken("").
			SetAmount(inputDeposit.Amount).
			SetTax(taxTransferInfo.Amount)

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventFxChildWithdrawERC20, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", types.ChainETH, types.ChainMATIC)).
			AddKeyValueWidget("Event", log.EventName).
			AddKeyValueWidget("Chain", log.Chain()).
			AddKeyValueWidget("Address", log.Address()).
			AddKeyValueWidget("TxHash", log.TxHash()).
			AddKeyValueWidget("Error", fmt.Sprintf("%s %s", "WithMongoTransaction:", err.Error())).
			SendMessage()
		return err
	}

	return nil
}

// SyncDeposit
// Chain :MATIC(ETH -> MATIC)
func (p *EventHandler) SyncDeposit(log types.Log) error {
	ctx := context.Background()

	eventSyncDeposit := bridgetypes.EventFxSyncDeposit{}
	err := log.Unmarshal(&eventSyncDeposit)
	if err != nil {
		logger.Error("EventSyncDeposit", "Unmarshal", err)
		return err
	}

	tx, _, err := p.client.GetTransactionWithReceipt(context.Background(), log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventSyncDeposit, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(context.Background(), log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventSyncDeposit, "GetBlockHeader", err)
	}

	blockNumber := header.BlockNumber()

	bridgeHistory, err := p.historyService.FindOneBridgeTxHistory(ctx, bson.M{
		"from_chain":    types.ChainETH,
		"to_chain":      types.ChainMATIC,
		"deposit_nonce": eventSyncDeposit.DepositNonce.Int64(),
	})
	if err != nil {
		return types.WrapError("GetBridgeHistoryByDepositNonce", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), p.cfg.FxPortal.ChildTokenAddress)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventFxDepositERC20, "error", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeOut).
			SetEvent(eventtypes.EventSyncDeposit).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(types.ChainETH).
			SetFromTxHash(bridgeHistory.TxHash).
			SetToChain(types.ChainMATIC).
			SetToTxHash(log.TxHash()).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info("EventSyncDeposit",
			"txHash", log.TxHash(),
			"blockNumber", blockNumber,
			"amount", eventSyncDeposit.Amount,
			"RootToken", eventSyncDeposit.RootToken,
			"msgSender", eventSyncDeposit.MsgSender,
			"To", eventSyncDeposit.To)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, bson.M{
			"cate":         txHistory.Cate,
			"tx_hash":      txHistory.TxHash,
			"from_tx_hash": txHistory.FromTxHash,
		})
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSyncDeposit, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(log.Address()).                   // Vault
			SetTo(eventSyncDeposit.To.String()).      // Recipient
			SetToken(*tokenInfo).                     // Token
			SetAmount(eventSyncDeposit.Amount).       // Transfer Amount
			SetChain(log.Chain()).                    // Chain
			SetTxHash(log.TxHash()).                  // TxHash
			SetCate(protocol.TxHistoryCateBridgeOut). // Category
			SetFromChain(types.ChainETH).
			SetFromTxHash(bridgeHistory.TxHash).
			SetToChain(types.ChainMATIC).
			SetToTxHash(log.TxHash()).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventTransfer,
			"chain", transferInfo.Chain,
			"cate", transferInfo.Cate,
			"from", transferInfo.From,
			"to", transferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, transferInfo, bson.M{
			"cate":         txHistory.Cate,
			"tx_hash":      txHistory.TxHash,
			"from_tx_hash": txHistory.FromTxHash,
		})
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSyncDeposit, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetSourceChain(types.ChainETH).
			SetSourceTxHash(bridgeHistory.TxHash).
			SetTargetChain(types.ChainMATIC).
			SetTargetTxHash(log.TxHash()).
			SetTargetToken(common.HexToAddress(types.ChainMATIC, tokenInfo.Address).String()).
			SetUpdatedAt(time.Now())

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSyncDeposit, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", types.ChainETH, types.ChainMATIC)).
			AddKeyValueWidget("Event", log.EventName).
			AddKeyValueWidget("Chain", log.Chain()).
			AddKeyValueWidget("Address", log.Address()).
			AddKeyValueWidget("TxHash", log.TxHash()).
			AddKeyValueWidget("Error", fmt.Sprintf("%s %s", "WithMongoTransaction:", err.Error())).
			SendMessage()
		return err
	}

	return nil
}
