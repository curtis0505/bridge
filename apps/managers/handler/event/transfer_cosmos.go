package event

import (
	"context"
	"fmt"
	cosmostxtypes "github.com/cosmos/cosmos-sdk/types/tx"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/libs/cache"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/entity"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

func (p *EventHandler) WasmDeposit(log types.Log) error {
	ctx := context.Background()

	err := p.CheckAlertAmount(ctx, log)
	if err != nil {
		p.logger.Warn("CheckAlertAmount", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(context.Background(), log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "GetTransactionWithReceipt", err)
	}

	var eventDeposit bridge.EventDeposit
	err = log.Unmarshal(&eventDeposit)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "Unmarshal", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventDeposit.TokenAddr)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "GetTokenInfoByAddress", err)
	}

	txHash := "0x" + strings.ToLower(log.TxHash())
	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Deposit Event
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(txHash).
			SetCate(protocol.TxHistoryCateBridgeIn).
			SetEvent(bridgetypes.EventNameWasmDeposit).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(log.Chain()).
			SetFromTxHash(txHash).
			SetToChain(eventDeposit.ChildChainName).
			SetTxFee(tx.Inner().(*cosmostxtypes.Tx).AuthInfo.Fee.Amount.AmountOf(tokentypes.DenomByChain(log.Chain())).BigInt()).
			SetBlockNumber(util.ToBigInt(log.BlockNumber())).
			ConvertMongoEntity()

		p.logger.Info("event", bridgetypes.EventNameWasmDeposit, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash())

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                            // Sender
			SetTo(tx.To()).                                // Vault
			SetToken(*tokenInfo).                          // Token
			SetAmount(util.ToBigInt(eventDeposit.Amount)). // Transfer Amount
			SetChain(log.Chain()).                         // Chain
			SetTxHash(txHash).                             // TxHash
			SetCate(protocol.TxHistoryCateBridgeIn).       // Category
			SetFromChain(log.Chain()).                     // From Tx
			SetFromTxHash(txHash).
			SetToChain(eventDeposit.ChildChainName).
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
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "UpsertBridgeTransferHistory", err)
		}

		// Tax Transfer Event
		taxTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                         // Sender
			SetTo(tx.To()).                             // Tax
			SetToken(*tokenInfo).                       // Token
			SetAmount(util.ToBigInt(eventDeposit.Tax)). // Tax
			SetChain(log.Chain()).                      // Chain
			SetTxHash(txHash).                          // TxHash
			SetCate(protocol.TxHistoryCateBridgeTax).   // Category
			SetFromChain(log.Chain()).                  // Source Tx
			SetFromTxHash(txHash).
			SetToChain(eventDeposit.ChildChainName).
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
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "UpsertBridgeTransferHistory", err)
		}

		//Chain Fee Transfer Event
		chainFeeTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                              // Sender
			SetTo(tx.To()).                                  // To
			SetToken(*tokenInfo).                            // Token
			SetAmount(util.ToBigInt(eventDeposit.ChainFee)). // Transfer Amount
			SetChain(log.Chain()).                           // Chain
			SetTxHash(txHash).                               // TxHash
			SetCate(eventtypes.TxHistoryBridgeChainFee).     // Category
			SetFromChain(log.Chain()).
			SetFromTxHash(txHash).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventFeeTransfer,
			"chain", log.Chain(),
			"cate", chainFeeTransferInfo.Cate,
			"from", chainFeeTransferInfo.From,
			"to", chainFeeTransferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, chainFeeTransferInfo, nil)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetFrom(tx.From()).
			SetTo(eventDeposit.ToAddr).
			SetSourceChain(log.Chain()).
			SetSourceTxHash(log.TxHash()).
			SetSourceToken(tokenInfo.Address).
			SetTargetChain(eventDeposit.ChildChainName).
			SetTargetTxHash("").
			SetTargetToken("").
			SetAmount(util.ToBigInt(eventDeposit.Amount)).
			SetChainFee(chainFeeTransferInfo.Amount).
			SetTax(taxTransferInfo.Amount)

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", log.Chain(), eventDeposit.ChildChainName)).
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

func (p *EventHandler) WasmDepositCoin(log types.Log) error {
	ctx := context.Background()

	err := p.CheckAlertAmount(ctx, log)
	if err != nil {
		p.logger.Warn("CheckAlertAmount", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(context.Background(), log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "GetTransactionWithReceipt", err)
	}

	var eventDeposit bridge.EventDepositCoin
	err = log.Unmarshal(&eventDeposit)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "Unmarshal", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventDeposit.NativeDenom)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "tokenInfo", fmt.Errorf("tokenInfo is nil"))
	}

	txHash := "0x" + strings.ToLower(log.TxHash()) //강제 convert를 할 것 인가?!

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Deposit Event
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(txHash).
			SetCate(protocol.TxHistoryCateBridgeIn).
			SetEvent(bridgetypes.EventNameWasmDepositCoin).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(log.Chain()).
			SetFromTxHash(txHash).
			SetToChain(eventDeposit.ChildChainName).
			SetTxFee(tx.Inner().(*cosmostxtypes.Tx).AuthInfo.Fee.Amount.AmountOf(tokentypes.DenomByChain(log.Chain())).BigInt()).
			SetBlockNumber(util.ToBigInt(log.BlockNumber())).
			ConvertMongoEntity()

		p.logger.Info("event", bridgetypes.EventNameWasmDepositCoin, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash())

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                            // Sender
			SetTo(tx.To()).                                // Vault
			SetToken(*tokenInfo).                          // Token
			SetAmount(util.ToBigInt(eventDeposit.Amount)). // Transfer Amount
			SetChain(log.Chain()).                         // Chain
			SetTxHash(txHash).                             // TxHash
			SetCate(protocol.TxHistoryCateBridgeIn).       // Category
			SetFromChain(log.Chain()).                     // From Tx
			SetFromTxHash(txHash).
			SetToChain(eventDeposit.ChildChainName).
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
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "UpsertBridgeTransferHistory", err)
		}

		// Tax Transfer Event
		taxTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                         // Sender
			SetTo(tx.To()).                             // Tax
			SetToken(*tokenInfo).                       // Token
			SetAmount(util.ToBigInt(eventDeposit.Tax)). // Tax
			SetChain(log.Chain()).                      // Chain
			SetTxHash(txHash).                          // TxHash
			SetCate(protocol.TxHistoryCateBridgeTax).   // Category
			SetFromChain(log.Chain()).                  // Source Tx
			SetFromTxHash(txHash).
			SetToChain(eventDeposit.ChildChainName).
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
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDepositCoin, "UpsertBridgeTransferHistory", err)
		}

		//Chain Fee Transfer Event
		chainFeeTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                              // Sender
			SetTo(tx.To()).                                  // To
			SetToken(*tokenInfo).                            // Token
			SetAmount(util.ToBigInt(eventDeposit.ChainFee)). // Transfer Amount
			SetChain(log.Chain()).                           // Chain
			SetTxHash(txHash).                               // TxHash
			SetCate(eventtypes.TxHistoryBridgeChainFee).     // Category
			SetFromChain(log.Chain()).
			SetFromTxHash(txHash).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventFeeTransfer,
			"chain", log.Chain(),
			"cate", chainFeeTransferInfo.Cate,
			"from", chainFeeTransferInfo.From,
			"to", chainFeeTransferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, chainFeeTransferInfo, nil)
		if err != nil {
			return nil, fmt.Errorf("UpsertBridgeTransferHistory: %w", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetFrom(tx.From()).
			SetTo(eventDeposit.ToAddr).
			SetSourceChain(log.Chain()).
			SetSourceTxHash(log.TxHash()).
			SetSourceToken(tokenInfo.Address).
			SetTargetChain(eventDeposit.ChildChainName).
			SetTargetTxHash("").
			SetTargetToken("").
			SetAmount(util.ToBigInt(eventDeposit.Amount)).
			SetChainFee(chainFeeTransferInfo.Amount).
			SetTax(taxTransferInfo.Amount)

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmDeposit, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", log.Chain(), eventDeposit.ChildChainName)).
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

func (p *EventHandler) WasmWithdraw(log types.Log) error {
	ctx := context.Background()

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "GetTransactionWithReceipt", err)
	}

	var eventWithdraw bridge.EventWithdraw
	err = log.Unmarshal(&eventWithdraw)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "Unmarshal", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventWithdraw.TokenAddr)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "GetTokenInfoByAddress", fmt.Errorf("tokenInfo is nil"))
	}

	txHash := "0x" + strings.ToLower(log.TxHash())

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(txHash).
			SetCate(protocol.TxHistoryCateBridgeOut).
			SetEvent(bridgetypes.EventNameWasmWithdraw).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(eventWithdraw.ChildChainName).
			SetFromTxHash(eventWithdraw.ChildTx).
			SetToChain(log.Chain()).
			SetToTxHash(txHash).
			SetTxFee(tx.Inner().(*cosmostxtypes.Tx).AuthInfo.Fee.Amount.AmountOf(tokentypes.DenomByChain(log.Chain())).BigInt()).
			SetBlockNumber(util.ToBigInt(log.BlockNumber())).
			ConvertMongoEntity()

		p.logger.Info("event", bridgetypes.EventNameWasmWithdraw, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash())

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(log.Address()).                         // Vault
			SetTo(eventWithdraw.ToAddr).                    // Recipient
			SetToken(*tokenInfo).                           // Token
			SetAmount(util.ToBigInt(eventWithdraw.Amount)). // Transfer Amount
			SetChain(log.Chain()).                          // Chain
			SetTxHash(txHash).                              // TxHash
			SetCate(protocol.TxHistoryCateBridgeOut).       // Category
			SetFromChain(eventWithdraw.ChildChainName).
			SetFromTxHash(eventWithdraw.ChildTx).
			SetToChain(log.Chain()).
			SetToTxHash(txHash).
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
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetSourceChain(eventWithdraw.ChildChainName).
			SetSourceTxHash(eventWithdraw.ChildTx).
			SetTargetChain(log.Chain()).
			SetTargetTxHash(txHash).
			SetTargetToken(tokenInfo.Address).
			SetUpdatedAt(time.Now())

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", eventWithdraw.ChildChainName, log.Chain())).
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

func (p *EventHandler) WasmWithdrawCoin(log types.Log) error {
	ctx := context.Background()

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdrawCoin, "GetTransactionWithReceipt", err)
	}

	var eventWithdraw bridge.EventWithdrawCoin
	err = log.Unmarshal(&eventWithdraw)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdrawCoin, "Unmarshal", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventWithdraw.NativeDenom)
	if err != nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdrawCoin, "GetTokenInfoByAddress", err)
	}
	if tokenInfo == nil {
		return eventtypes.WrapError(bridgetypes.EventNameWasmWithdrawCoin, "GetTokenInfoByAddress", fmt.Errorf("tokenInfo is nil"))
	}

	txHash := "0x" + strings.ToLower(log.TxHash())

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(txHash).
			SetCate(protocol.TxHistoryCateBridgeOut).
			SetEvent(bridgetypes.EventNameWasmWithdrawCoin).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(eventWithdraw.ChildChainName).
			SetFromTxHash(eventWithdraw.ChildTx).
			SetToChain(log.Chain()).
			SetToTxHash(txHash).
			SetTxFee(tx.Inner().(*cosmostxtypes.Tx).AuthInfo.Fee.Amount.AmountOf(tokentypes.DenomByChain(log.Chain())).BigInt()).
			SetBlockNumber(util.ToBigInt(log.BlockNumber())).
			ConvertMongoEntity()

		p.logger.Info("event", bridgetypes.EventNameWasmWithdrawCoin, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash())

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmWithdrawCoin, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(log.Address()).                         // Vault
			SetTo(eventWithdraw.ToAddr).                    // Recipient
			SetToken(*tokenInfo).                           // Token
			SetAmount(util.ToBigInt(eventWithdraw.Amount)). // Transfer Amount
			SetChain(log.Chain()).                          // Chain
			SetTxHash(txHash).                              // TxHash
			SetCate(protocol.TxHistoryCateBridgeOut).       // Category
			SetFromChain(eventWithdraw.ChildChainName).
			SetFromTxHash(eventWithdraw.ChildTx).
			SetToChain(log.Chain()).
			SetToTxHash(txHash).
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
			return nil, fmt.Errorf("UpsertBridgeTransferHistory: %w", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetSourceChain(eventWithdraw.ChildChainName).
			SetSourceTxHash(eventWithdraw.ChildTx).
			SetTargetChain(log.Chain()).
			SetTargetTxHash(txHash).
			SetTargetToken(tokenInfo.Address).
			SetUpdatedAt(time.Now())

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(bridgetypes.EventNameWasmWithdraw, "UpsetBridgeActiveHistory", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", eventWithdraw.ChildChainName, log.Chain())).
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

func (p *EventHandler) WasmMint(log types.Log) error {
	return nil
}

func (p *EventHandler) WasmBurn(log types.Log) error {
	ctx := context.Background()
	err := p.CheckAlertAmount(ctx, log)
	if err != nil {
		p.logger.Warn("CheckAlertAmount", err)
	}

	return nil
}
