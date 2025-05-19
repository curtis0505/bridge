package event

import (
	"context"
	"fmt"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/entity"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (p *EventHandler) Submission(log types.Log) error {
	ctx := context.Background()
	eventSubmission := bridgetypes.EventSubmission{}
	err := log.Unmarshal(&eventSubmission)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventSubmission, "Unmarshal", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventSubmission, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventSubmission, "HeaderByNumber", err)
	}

	history, err := p.historyService.FindOneBridgeTxHistory(ctx, bson.M{
		"tx_hash": eventSubmission.TxHash.String(),
	})
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventSubmission, "FindOneBridgeTxHistory", err)
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeSubmission).
			SetEvent(eventtypes.EventSubmission).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(history.Chain).
			SetFromTxHash(eventSubmission.TxHash.String()).
			SetFromTxId(eventSubmission.TransactionId).
			SetToChain(history.ToChain).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.EventSubmission,
			"chain", log.Chain(),
			"from", tx.From(),
			"txHash", tx.TxHash(),
			"fee", tx.TxFee(header.BaseFee()),
			"blockNumber", header.BlockNumber(),
		)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSubmission, "UpsertBridgeTxHistory", err)
		}

		if err := p.callBackSubmission(eventSubmission.TransactionId.Int64()); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSubmission, "callBackSubmission", err)
		}

		if err := p.verifyHandler.VerifySubmission(ctx, log, eventSubmission); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSubmission, "VerifySubmission", err)
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", history.Chain, history.ToChain)).
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

func (p *EventHandler) callBackSubmission(txId int64) error {
	for _, confirmationLog := range p.pendingConfirmation[txId] {
		if err := p.Confirmation(confirmationLog); err != nil {
			return err
		}
	}
	p.pendingConfirmation[txId] = p.pendingConfirmation[txId][:0]
	return nil
}

func (p *EventHandler) addPendingConfirmation(log types.Log, txId int64) {
	p.logger.Warn("event", eventtypes.EventConfirmation, "chain", log.Chain(), "msg", "wait for submission log", "txId", txId, "txHash", log.TxHash())
	p.pendingConfirmation[txId] = append(p.pendingConfirmation[txId], log)
}

func (p *EventHandler) Confirmation(log types.Log) error {
	ctx := context.Background()
	eventConfirmation := bridgetypes.EventConfirmation{}
	err := log.Unmarshal(&eventConfirmation)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventConfirmation, "Unmarshal", err)
	}

	history, err := p.historyService.FindOneBridgeTxHistory(ctx, bson.M{
		"chain": log.Chain(),
		"cate":  protocol.TxHistoryCateBridgeSubmission,
		"tx_id": eventConfirmation.TransactionId.Int64(),
	})
	if err != nil {
		p.addPendingConfirmation(log, eventConfirmation.TransactionId.Int64())
		return nil
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventConfirmation, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventConfirmation, "HeaderByNumber", err)
	}

	txHistory := entity.NewBridgeTxHistory().
		SetTxHash(log.TxHash()).
		SetCate(protocol.TxHistoryCateBridgeConfirm).
		SetEvent(eventtypes.EventConfirmation).
		SetChain(log.Chain()).
		SetFrom(tx.From()).
		SetFromChain(history.FromChain).
		SetFromTxHash(history.FromTxHash).
		SetFromTxId(eventConfirmation.TransactionId).
		SetToChain(history.ToChain).
		SetToTxHash(history.ToTxHash).
		SetTxFee(tx.TxFee(header.BaseFee())).
		SetBlockNumber(header.BlockNumber()).
		ConvertMongoEntity()

	p.logger.Info(
		"event", eventtypes.EventConfirmation,
		"chain", log.Chain(),
		"from", tx.From(),
		"txHash", tx.TxHash(),
		"fee", tx.TxFee(header.BaseFee()),
		"blockNumber", header.BlockNumber(),
	)

	err = p.historyService.UpsertBridgeTxHistory(ctx, txHistory, nil)
	if err != nil {
		p.logger.Error("UpsertBridgeTxHistory", err)
	}

	p.validatorHandler.ConfirmTx(tx)

	return nil
}

func (p *EventHandler) Execution(log types.Log) error {
	ctx := context.Background()

	eventExecution := bridgetypes.EventExecution{}
	err := log.Unmarshal(&eventExecution)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecution, "Unmarshal", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecution, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecution, "GetBlockHeader", err)
	}

	history, err := p.historyService.FindOneBridgeTxHistory(ctx, bson.M{
		"tx_hash": eventExecution.TxHash.String(),
	})
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecution, "FindOneBridgeTxHistory", err)
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeExecution).
			SetEvent(eventtypes.EventExecution).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(history.Chain).
			SetFromTxHash(eventExecution.TxHash.String()).
			SetFromTxId(eventExecution.TransactionId).
			SetToTxHash(history.ToTxHash).
			SetToChain(log.Chain()).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.EventExecution,
			"chain", log.Chain(),
			"from", tx.From(),
			"txHash", tx.TxHash(),
			"fee", tx.TxFee(header.BaseFee()),
			"blockNumber", header.BlockNumber(),
		)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventExecution, "UpsertBridgeTxHistory", err)
		}

		if err = p.verifyHandler.VerifyExecution(sessCtx, log, eventExecution); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventExecution, "VerifyExecution", err)
		}
		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", history.Chain, log.Chain())).
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

func (p *EventHandler) ExecutionFailure(log types.Log) error {
	ctx := context.Background()

	eventExecutionFailure := bridgetypes.EventExecutionFailure{}
	err := log.Unmarshal(&eventExecutionFailure)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecutionFailure, "Unmarshal", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecutionFailure, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecutionFailure, "GetBlockHeader", err)
	}

	history, err := p.historyService.FindOneBridgeTxHistory(ctx, bson.M{
		"tx_hash": eventExecutionFailure.TxHash.String(),
	})
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventExecutionFailure, "FindOneBridgeTxHistory", err)
	}

	txHistory := entity.NewBridgeTxHistory().
		SetTxHash(log.TxHash()).
		SetCate(protocol.TxHistoryCateBridgeExecutionFailure).
		SetEvent(eventtypes.EventExecutionFailure).
		SetChain(log.Chain()).
		SetFrom(tx.From()).
		SetFromChain(history.Chain).
		SetFromTxHash(eventExecutionFailure.TxHash.String()).
		SetFromTxId(eventExecutionFailure.TransactionId).
		SetToTxHash(history.ToTxHash).
		SetToChain(log.Chain()).
		SetTxFee(tx.TxFee(header.BaseFee())).
		ConvertMongoEntity()

	p.logger.Info(
		"event", eventtypes.EventExecutionFailure,
		"chain", log.Chain(),
		"from", tx.From(),
		"txHash", tx.TxHash(),
		"fee", tx.TxFee(header.BaseFee()),
	)

	err = p.historyService.UpsertBridgeTxHistory(ctx, txHistory, nil)
	if err != nil {
		p.logger.Error("UpsertBridgeTxHistory", err)
	}

	return nil
}
