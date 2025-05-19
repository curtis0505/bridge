package fxportal

import (
	"errors"
	"fmt"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/apps/managers/types/fxportal"
	"github.com/curtis0505/bridge/libs/cache"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/entity"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// SyncReceiveMessage
// execute FxERC20RootTunnel ReceiveMessage contract call in pendingTxPool manually
func (f *FxPortal) SyncReceiveMessage(ctx *gin.Context) {
	f.ReceiveMessage()
	ctx.JSON(http.StatusOK, types.NewResponseSuccess())
}

// CheckPendingTx
// check pending FxPortalWithdrawal tx from DB which does not exist in pendingTxMap
func (f *FxPortal) CheckPendingTx(ctx *gin.Context) {
	pendingTxHistories, err := f.historyService.FindTxHistories(ctx, bson.M{
		"cate":         mongoentity.TxHistoryCateBridgeOut,
		"base_chain":   commontypes.ChainMATIC,
		"target_chain": commontypes.ChainETH,
		"stat":         mongoentity.TxHistoryStatPending,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewResponseHeader(types.Failed, err))
		return
	}
	for _, pendingTxHistory := range pendingTxHistories {
		withdrawTxHistory := fxportal.WithdrawPendingTx{
			From:       pendingTxHistory.From,
			TxHash:     pendingTxHistory.InTxHash,
			RetryCount: 0,
		}
		f.RegisterPendingWithdrawTx(withdrawTxHistory)
	}
	ctx.JSON(http.StatusOK, types.NewResponseSuccess())
}

// Exit
// execute FxERC20RootTunnel ReceiveMessage contract call by txHash manually
// Chain: ETH(MATIC -> ETH)
func (f *FxPortal) Exit(ctx *gin.Context) {
	txHash := ctx.Param("txHash")
	if txHash == "" {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, errors.New("txHash is empty")))
		return
	}

	withdrawTxHistory, err := f.historyService.FindOneBridgeTxHistory(ctx, bson.M{
		"tx_hash": txHash,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, fmt.Errorf("FindOneBridgeTxHistory: %w", err)))
		return
	} else if withdrawTxHistory.Event != eventtypes.EventFxChildWithdrawERC20 {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, errors.New("not FxChildWithdrawERC20 event")))
		return
	}

	proof, err := f.GetProofFromMaticAPI(txHash)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, fmt.Errorf("GetProofFromMaticAPI: %w", err)))
		return
	}

	fxERC20RootTunnelClient, err := f.client.NewFxERC20RootTunnel(f.cfg.FxPortal.RootTokenAddress, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewResponseHeader(types.Failed, fmt.Errorf("NewFxERC20RootTunnel: %w", err)))
		return
	}

	chainID, err := f.client.GetChainID(ctx, commontypes.ChainETH)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewResponseHeader(types.Failed, fmt.Errorf("GetChainID: %w", err)))
		return
	}

	account, err := commontypes.NewAccount(f.cfg.Account, chainID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewResponseHeader(types.Failed, fmt.Errorf("NewAccount: %w", err)))
		return
	}

	proofBytes, err := hexutil.Decode(proof)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, fmt.Errorf("decode: %w", err)))
		return
	}

	txResult, err := fxERC20RootTunnelClient.ReceiveMessage(proofBytes, account)
	if err != nil {
		f.logger.Error("NewHeaderBlock.ReceiveMessage", "ReceiveMessage", err)
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, fmt.Errorf("ReceiveMessage: %w", err)))
		return
	}

	txResultNonce, _ := txResult.Nonce()
	f.logger.Info("NewHeaderBlock.ReceiveMessage",
		"resultTxHash", txResult.TxHash(),
		"from", txResult.From(),
		"to", txResult.To(),
		"gasPrice", txResult.GasPrice().String(),
		"gasFeeCap", txResult.GasFeeCap().String(),
		"gasTipCap", txResult.GasTipCap().String(),
		"nonce", txResultNonce,
	)

	_, err = f.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		// trader에서 in_tx_hash 값을 가져오기 위해 result_tx_hash 값을 저장
		txHistory, err := f.historyService.FindOneTxHistory(sessCtx, bson.M{
			"cate":       mongoentity.TxHistoryCateBridgeOut,
			"in_tx_hash": txHash,
		})
		if err != nil {
			return nil, fmt.Errorf("FindOneTxHistory: %w", err)
		}

		err = f.historyService.UpdateOneTxHistory(ctx,
			bson.M{
				"_id": txHistory.ID,
			},
			bson.M{
				"$set": bson.M{
					"tx_hash": txResult.TxHash(),
				},
			})
		if err != nil {
			return nil, fmt.Errorf("UpdateOneTxHistory: %w", err)
		}

		tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(commontypes.ChainETH, f.cfg.FxPortal.RootTokenAddress)
		if err != nil {
			return nil, fmt.Errorf("GetTokenInfoByAddress: %w", err)
		}

		if tokenInfo == nil {
			return nil, fmt.Errorf("TokenInfo is nil")
		}

		amount, _, err := txHistory.TargetAmount.BigInt()
		if err != nil {
			return nil, fmt.Errorf("BigInt: %w", err)
		}

		bridgeTxHistory := entity.NewBridgeTxHistory().
			SetTxHash(txResult.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeOut).
			SetEvent(eventtypes.EventFxWithdrawERC20).
			SetChain(commontypes.ChainETH).
			SetFrom(withdrawTxHistory.From).
			SetFromChain(commontypes.ChainMATIC).
			SetFromTxHash(txHash).
			SetToChain(commontypes.ChainETH).
			SetToTxHash(txResult.TxHash()).
			ConvertMongoEntity()

		err = f.historyService.UpsertBridgeTxHistory(sessCtx, bridgeTxHistory, nil)
		if err != nil {
			return nil, fmt.Errorf("UpsertBridgeTxHistory: %w", err)
		}

		transferInfo := entity.NewBridgeTransferHistory().
			SetTxHash(txResult.TxHash()).
			SetFrom(withdrawTxHistory.From).
			SetTo(txResult.To()).
			SetChain(commontypes.ChainETH).
			SetToken(*tokenInfo).
			SetFromTxHash(txHash).
			SetAmount(amount).
			SetCate(protocol.TxHistoryCateBridgeOut).
			SetFromChain(commontypes.ChainMATIC).
			SetToChain(commontypes.ChainETH).
			SetToTxHash(txResult.TxHash()).
			ConvertMongoEntity()

		err = f.historyService.UpsertBridgeTransferHistory(sessCtx, transferInfo, nil)
		if err != nil {
			return nil, fmt.Errorf("UpsertBridgeTransferHistory: %w", err)
		}

		//activeHistory := mongoentity.NewBridgeActiveHistory().
		//	SetSourceChain(commontypes.ChainMATIC).
		//	SetTargetChain(commontypes.ChainETH).
		//	SetSourceTxHash(txHash).
		//	SetTargetTxHash(txResult.TxHash())
		//
		//err = f.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		//if err != nil {
		//	return nil, fmt.Errorf("UpsetBridgeActiveHistory: %w", err)
		//}

		return nil, nil
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.NewResponseHeader(types.Failed, err))
		return
	}

	ctx.JSON(http.StatusOK, types.NewResponseSuccess())
}

type PendingWithdrawTxResponse struct {
	*types.ResponseHeader
	PendingTx map[string]fxportal.WithdrawPendingTx `json:"pendingTx"`
}

// GetPendingWithdrawTxMap
// get pending withdraw tx map
func (f *FxPortal) GetPendingWithdrawTxMap(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, f.PendingWithdrawTxs)
}
