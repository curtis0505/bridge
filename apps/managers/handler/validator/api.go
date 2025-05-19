package validator

import (
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/apps/managers/util"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

func (p *ValidatorHandler) ValidatorInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, p.validatorList)
}

func (p *ValidatorHandler) PendingTxRequest(ctx *gin.Context) {
	var req validatortypes.PendingTxRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
		return
	}

	p.pendingTxMapLock.Lock()
	defer p.pendingTxMapLock.Unlock()

	pendingTx := &validatortypes.PendingTx{
		ValidatorInfo:    p.GetValidatorInfo(req.Chain, req.From),
		TxHash:           req.TxHash,
		Chain:            req.Chain,
		Time:             time.Now(),
		ValidatorSummary: p.validatorSummary(req.Chain, req.From),

		//SubmitTransaction:     &inputSubmitTransaction,
		//SubmitTransactionData: &inputSubmitTransactionData,
	}

	if pendingTx.ValidatorInfo == nil || pendingTx.ValidatorSummary == nil {
		p.logger.Error("event", "PendingTx", "chain", pendingTx.Chain, "validator", "unknown", "txHash", pendingTx.TxHash)
		return
	}

	p.pendingTxMap[pendingTx.TxHash] = pendingTx

	p.logger.Info(
		"event", "PendingTx", "chain", pendingTx.Chain, "validator", pendingTx.ValidatorInfo.Name, "txHash", pendingTx.TxHash,
		"pending", len(p.pendingTxMap),
	)
}

func (p *ValidatorHandler) validatorSummary(chain, address string) *validatortypes.ValidatorSummary {
	for _, v := range p.validatorList {
		if strings.ToLower(address) == strings.ToLower(v.AddressInfo[chain].Address) {
			return &validatortypes.ValidatorSummary{
				Address: address,
				Name:    v.Name,
			}
		}
	}
	return nil
}

func (p *ValidatorHandler) ValidatorStatistics(ctx *gin.Context) {
	var addresses []string
	for _, v := range p.validatorList {
		for _, address := range v.AddressInfo {
			addresses = append(addresses, address.Address)
		}
	}

	statisticsList, err := p.historyService.AggregateValidatorStatistics(ctx, addresses...)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
		return
	}

	ctx.JSON(http.StatusOK, statisticsList)
}

// Tx TODO: Move to another module
func (p *ValidatorHandler) Tx(ctx *gin.Context) {
	var txResponse validatortypes.TxResponse
	var status validatortypes.BridgeStatus

	txHash := ctx.Param("txHash")
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
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
		return
	}

	transferList, err := p.historyService.FindBridgeTransferHistory(ctx, bson.M{
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
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
		return
	}

	for _, history := range historyList {
		switch history.Cate {
		case protocol.TxHistoryCateBridgeIn:
			if status < validatortypes.BridgeRequest {
				status = validatortypes.BridgeRequest
			}

			txResponse.From = history.From
			txResponse.FromChain = history.FromChain
			txResponse.ToChain = history.ToChain

			txResponse.AddBridgeHistory(nil, history)
			continue
		case protocol.TxHistoryCateBridgeSubmission:
			if status < validatortypes.BridgeSubmission {
				status = validatortypes.BridgeSubmission
			}

		case protocol.TxHistoryCateBridgeConfirm:
			if status < validatortypes.BridgeConfirm {
				status = validatortypes.BridgeConfirm
			}

			txResponse.Confirmation += 1
		case protocol.TxHistoryCateBridgeExecution:
			if status < validatortypes.BridgeExecute {
				status = validatortypes.BridgeExecute
			}
			txResponse.Execute = true

		case protocol.TxHistoryCateBridgeExecutionFailure:
			if status < validatortypes.BridgeExecuteFailure {
				status = validatortypes.BridgeExecuteFailure
			}
			txResponse.Execute = false
		}

		txResponse.AddBridgeHistory(p.validatorSummary(history.Chain, history.From), history)
	}

	for _, history := range transferList {
		switch history.Cate {
		case protocol.TxHistoryCateBridgeIn:
			txResponse.Amount = util.ToEther(history.Amount).String()
			txResponse.FromTxHash = history.TxHash
			txResponse.FromCurrencyId = history.CurrencyID
			txResponse.GroupId = history.GroupID

		case protocol.TxHistoryCateBridgeTax:
			txResponse.Tax = util.ToEther(history.Amount).String()

		case protocol.TxHistoryCateBridgeOut:
			txResponse.To = history.To
			txResponse.ToTxHash = history.TxHash
			txResponse.ToCurrencyId = history.CurrencyID
		}
		txResponse.AddBridgeTransferHistory(history)
	}

	txResponse.Status = status.String()

	ctx.JSON(http.StatusOK, txResponse)
}

func (p *ValidatorHandler) ValidatorPendingTxs(ctx *gin.Context) {
	p.pendingTxMapLock.Lock()
	defer p.pendingTxMapLock.Unlock()

	ctx.JSON(http.StatusOK, p.pendingTxMap)
}

func (p *ValidatorHandler) MultiSigAddress(ctx *gin.Context) {
	multiSigPubKey := p.MultiSigPubKey(ctx.Param("chain"), validatortypes.Threshold)
	ctx.JSON(http.StatusOK, gin.H{
		"address":   cosmoscommon.FromAddress(ctx.Param("chain"), multiSigPubKey.Address()).String(),
		"pubKey":    multiSigPubKey.GetPubKeys(),
		"threshold": validatortypes.Threshold,
	})
}
