package event

import (
	"context"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/apps/managers/types/fxportal"
	"github.com/curtis0505/bridge/libs/dto"
	logger "github.com/curtis0505/bridge/libs/logger"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *EventHandler) NewHeaderBlock(log types.Log) error {
	eventNewHeaderBlock := bridgetypes.EventNewHeaderBlock{}
	err := log.Unmarshal(&eventNewHeaderBlock)
	if err != nil {
		logger.Error("NewHeaderBlock", "Unmarshal", err)
		return err
	}

	logger.Info("NewHeaderBlock",
		"proposer", eventNewHeaderBlock.Proposer.String(),
		"headerBlockId", eventNewHeaderBlock.HeaderBlockId.String(),
		"reward", eventNewHeaderBlock.Reward.String(),
		"start", eventNewHeaderBlock.Start.String(),
		"end", eventNewHeaderBlock.End.String(),
	)

	withdrawTxHistories, err := p.historyService.FindBridgeTxHistory(context.Background(), bson.M{
		"$and": bson.A{
			bson.M{
				"chain": types.ChainMATIC,
			},
			bson.M{
				"cate": dto.TxHistoryCateBridgeIn,
			},
			bson.M{
				"event": "FxChildWithdrawERC20",
			},
			bson.M{
				"block_number": bson.M{
					"$gte": eventNewHeaderBlock.Start.Int64(),
					"$lte": eventNewHeaderBlock.End.Int64(),
				},
			},
		},
	})
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventNewHeaderBlock, "GetBridgeWithdrawHistoryByBlockNumber", err)
	}
	logger.Info("NewHeaderBlock.GetBridgeWithdrawHistoryByBlockNumber", "start", eventNewHeaderBlock.Start.String(), "end", eventNewHeaderBlock.End.String(), "count", len(withdrawTxHistories))

	for _, history := range withdrawTxHistories {
		withdrawTxHistory := fxportal.WithdrawPendingTx{
			From:       history.From,
			TxHash:     history.TxHash,
			RetryCount: 0,
		}
		p.fxPortalHandler.RegisterPendingWithdrawTx(withdrawTxHistory)
		logger.Info("NewHeaderBlock.RegisterPendingWithdrawTx", "txHash", history.TxHash, "from", history.From, "chain", history.Chain, "cate", history.Cate)
	}
	return nil
}
