package validator

import (
	"context"
	"errors"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/cache"
	protocol "github.com/curtis0505/bridge/libs/dto"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

// Recover  API 용 컨트롤러
func (p *ValidatorHandler) Recover(ctx *gin.Context) {
	switch commontypes.GetChainType(ctx.Param("chain")) {
	case commontypes.ChainTypeEVM:
		p.RecoverEVM(ctx.Param("chain"), ctx.Param("txHash"))
		ctx.JSON(http.StatusOK, types.NewResponseSuccess())

	default:
		ctx.JSON(http.StatusOK, types.NewResponseSuccess())
	}

}

func (p *ValidatorHandler) RecoverMultiSig(ctx *gin.Context) {
	p.logger.Info("event", "RecoverMultiSig", "chain", ctx.Param("chain"), "txHash", ctx.Param("txHash"))

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, ctx.Param("chain"), ctx.Param("txHash"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(http.StatusBadRequest, errors.New("invalid txHash")))
		return
	}

	contract, err := cache.ContractCache().GetContractByAddress(ctx.Param("chain"), tx.To())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(http.StatusBadRequest, errors.New("invalid to address")))
		return
	}

	if contract.ContractID != bridgetypes.VaultContractID && contract.ContractID != bridgetypes.MinterContractID {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(http.StatusBadRequest, errors.New("invalid to address")))
		return
	}

	for _, log := range tx.Receipt().Logs() {
		switch log.EventName {
		case bridgetypes.EventNameDeposit:
			eventDeposit := bridgetypes.EventDeposit{}
			if err = log.Unmarshal(&eventDeposit); err != nil {
				p.logger.Error("err", err)
				continue
			}
			// Cosmos 계열 체인으로 가는 경우 MultiSig Transaction 을 Manager 에서 전송
			if commontypes.GetChainType(eventDeposit.ToChainName) == commontypes.ChainTypeCOSMOS {
				commonLog, err := bridgetypes.GetDepositCommonLog(log, bridgetypes.ContractAddresses{})
				if err != nil {
					p.logger.Error("err", err)
					continue
				}
				p.AddCosmosMultiSigTx(validatortypes.NewCosmosMultiSigTx(log.Chain(), log.TxHash(), eventtypes.EventDeposit, commonLog))
			}
		case bridgetypes.EventNameBurn:
			eventBurn := bridgetypes.EventBurn{}
			if err = log.Unmarshal(&eventBurn); err != nil {
				p.logger.Error("err", err)
				continue
			}
			// Cosmos 계열 체인으로 가는 경우 MultiSig Transaction 을 Manager 에서 전송
			if commontypes.GetChainType(eventBurn.ToChainName) == commontypes.ChainTypeCOSMOS {
				commonLog, err := bridgetypes.GetBurnCommonLog(log, bridgetypes.ContractAddresses{})
				if err != nil {
					p.logger.Error("err", err)
					continue
				}
				p.AddCosmosMultiSigTx(validatortypes.NewCosmosMultiSigTx(log.Chain(), log.TxHash(), eventtypes.EventBurn, commonLog))
			}
		}
	}

	ctx.JSON(http.StatusOK, types.NewResponseHeader(http.StatusOK, errors.New("invalid txHash")))
}

func (p *ValidatorHandler) RecoverEVM(chain string, txHash string) {
	for _, v := range p.validatorList {
		go func(validator *validatortypes.ValidatorInfo) {
			if err := validator.RecoverTx(chain, txHash); err != nil {
				p.logger.Error("event", "RecoverEVM", "chain", chain, "txHash", txHash, "validator", validator.Name)
			}
		}(v)

		p.logger.Info("event", "RecoverEVM", "chain", chain, "txHash", txHash, "validator", v.Name)
	}

	return
}

// CheckRecoverTx 스케줄러, 주기적으로 브릿지 펜딩중인건들 recover
func (p *ValidatorHandler) CheckRecoverTx() {
	ctx := context.Background()

	bridgeOutPendingHistoryList, err := p.historyService.FindTxHistories(ctx, bson.M{
		"stat":      mongoentity.TxHistoryStatPending,
		"cate":      mongoentity.TxHistoryCateBridgeOut,
		"create_at": bson.M{"$lte": time.Now().Add(-time.Hour)},
	})
	if err != nil {
		p.logger.Error("event", "CheckRecoverTx", "GetBridgeOutPendingHistoryList", err)
		return
	}
	p.logger.Info("CheckRecoverTx", "start", "count", len(bridgeOutPendingHistoryList))

	for _, bridgeOutPendingHistory := range bridgeOutPendingHistoryList {
		bridgeStatus := p.GetBridgeStatus(bridgeOutPendingHistory.InTxHash)

		bridgePath, err := p.bridgeService.GetUniqueBridgePath(
			ctx,
			bridgeOutPendingHistory.BaseCurrencyID,
			bridgeOutPendingHistory.BaseChain,
			bridgeOutPendingHistory.TargetChain,
		)
		if err != nil {
			p.logger.Error("event", "CheckRecoverTx", "GetUniqueBridgePath", err)
			continue
		}

		// target evm 브릿지만 재실행
		if commontypes.GetChainType(bridgePath.TargetChain) != commontypes.ChainTypeEVM {
			continue
		}

		// 네오핀에서 구현한 브릿지만 재실행
		if bridgePath.BridgeType != mongoentity.BridgeTypeNeopin {
			continue
		}

		// bridge_tx_history 에 기록이 되지 않았거나, 이미 실행되지 않은 내역만 처리
		if (validatortypes.BridgeUnknown < bridgeStatus) && (bridgeStatus < validatortypes.BridgeExecute) {
			p.logger.Info("event", "CheckRecoverTx", "tx_hash", bridgeOutPendingHistory.InTxHash, "status", bridgeStatus)
			p.RecoverEVM(bridgeOutPendingHistory.BaseChain, bridgeOutPendingHistory.InTxHash)

			// 동시에 여러번 호출시 gateway 에서 proof 데이터 조회시, gateway 에서 죽음
			time.Sleep(time.Second * 10)
		}
	}
}

func (p *ValidatorHandler) GetBridgeStatus(txHash string) validatortypes.BridgeStatus {
	ctx := context.Background()
	var status validatortypes.BridgeStatus
	historyList, err := p.historyService.FindBridgeTxHistory(ctx, bson.M{
		"$or": bson.A{
			bson.M{
				"tx_hash": txHash,
			},
			bson.M{
				"from_tx_hash": txHash,
			},
			bson.M{
				"to_tx_hash": txHash,
			},
		},
	})
	if err != nil {
		return status
	}

	for _, history := range historyList {
		switch history.Cate {
		case protocol.TxHistoryCateBridgeIn, protocol.TxHistoryCateBridgeBurn:
			if status < validatortypes.BridgeRequest {
				status = validatortypes.BridgeRequest
			}
		case protocol.TxHistoryCateBridgeSubmission:
			if status < validatortypes.BridgeSubmission {
				status = validatortypes.BridgeSubmission
			}
		case protocol.TxHistoryCateBridgeConfirm:
			if status < validatortypes.BridgeConfirm {
				status = validatortypes.BridgeConfirm
			}
		case protocol.TxHistoryCateBridgeExecution:
			if status < validatortypes.BridgeExecute {
				status = validatortypes.BridgeExecute
			}
		case protocol.TxHistoryCateBridgeExecutionFailure:
			if status < validatortypes.BridgeExecuteFailure {
				status = validatortypes.BridgeExecuteFailure
			}
		}
	}

	return status
}
