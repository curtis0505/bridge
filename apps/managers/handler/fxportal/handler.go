package fxportal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/apps/managers/conf"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/apps/managers/types/fxportal"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/entity"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/util"

	//"github.com/curtis0505/bridge/apps/managers/util"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/common"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/elog"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	"github.com/curtis0505/bridge/libs/service"
	commontypes "github.com/curtis0505/bridge/libs/types"
	basetypes "github.com/curtis0505/bridge/libs/types/base"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	_ types.Handler = &FxPortal{}
)

type FxPortal struct {
	cfg     *conf.Config
	client  *chain.Client
	logger  *elog.Logger
	account *commontypes.Account

	PendingWithdrawTxs  map[string]fxportal.WithdrawPendingTx
	pendingWithdrawLock *sync.RWMutex

	itemService        *service.ItemService
	historyService     *service.HistoryService
	transactionManager *service.TransactionManager
}

func New(cfg conf.Config, client *chain.Client) *FxPortal {
	chainID, err := client.GetChainID(context.Background(), commontypes.ChainETH)
	if err != nil {
		panic(err)
	}

	account, err := commontypes.NewAccount(cfg.Account, chainID)
	if err != nil {
		panic(err)
	}

	fxPortalClient := FxPortal{
		cfg:                 &cfg,
		client:              client,
		logger:              elog.NewLogger("FxPortal"),
		account:             account,
		PendingWithdrawTxs:  make(map[string]fxportal.WithdrawPendingTx),
		pendingWithdrawLock: &sync.RWMutex{},

		itemService:        service.GetRegistry().ItemService(),
		historyService:     service.GetRegistry().HistoryService(),
		transactionManager: service.GetRegistry().TransactionManager(),
	}

	go fxPortalClient.iterate()

	return &fxPortalClient
}

func (f *FxPortal) Name() string { return "FxPortal" }

func (f *FxPortal) ApiHandler(e *gin.Engine) {
	g := e.Group("/manager/fxportal")

	g.GET("/pending", f.GetPendingWithdrawTxMap)
	g.POST("/pending/check", f.CheckPendingTx)
	g.POST("/exit/:txHash", f.Exit)
	g.GET("/exit/sync", f.SyncReceiveMessage)
}

func (f *FxPortal) LogHandler(log commontypes.Log) error {
	return nil
}

func (f *FxPortal) RegisterPendingWithdrawTx(withdrawPendingTx fxportal.WithdrawPendingTx) {
	f.pendingWithdrawLock.Lock()
	defer f.pendingWithdrawLock.Unlock()
	f.PendingWithdrawTxs[withdrawPendingTx.TxHash] = withdrawPendingTx
	f.logger.Debug("registerPendingWithdrawTx", "Pending Tx Registered", withdrawPendingTx.TxHash)
}

func (f *FxPortal) iterate() {
	checkpointProofWaitTicker := time.NewTicker(15 * time.Minute)
	balanceTicker := time.NewTicker(30 * time.Minute)

	for {
		select {
		case <-checkpointProofWaitTicker.C:
			f.ReceiveMessage()
		case <-balanceTicker.C:
			f.checkBalance()
		}
	}
}

func (f *FxPortal) checkBalance() {
	ctx := context.Background()
	const minimumBalance = 0.1

	// Check Eth-NPT balance of FxERC20RootTunnel
	tokenContract, err := cache.ContractCache().GetContractByAddress(commontypes.ChainETH, f.cfg.FxPortal.RootTokenAddress)
	if err != nil {
		f.logger.Error("checkBalance", "GetContractByAddress", err)
		return
	}

	contract, err := cache.ContractCache().GetContractByContractID(commontypes.ChainETH, bridgetypes.FxERC20RootTunnelCustomID)
	if err != nil {
		f.logger.Error("checkBalance", "GetContractByContractID", err)
		return
	}

	var fxERC20RootTunnelClientNPTBalance basetypes.OutputBigInt
	err = f.client.CallMsgUnmarshalContract2(
		ctx,
		tokenContract,
		"balanceOf",
		&fxERC20RootTunnelClientNPTBalance,
		common.HexToAddress(contract.Address, contract.Address),
	)
	if err != nil {
		f.logger.Error("checkBalance", "BalanceOf", err)
	}
	f.logger.Info("checkBalance",
		"fxERC20RootTunnelClientNPTBalance", util.ToEther(fxERC20RootTunnelClientNPTBalance.Value).String())

	// Check Matic-NPT balance of FxERC20ChildTunnel
	tokenContract, err = cache.ContractCache().GetContractByAddress(commontypes.ChainETH, f.cfg.FxPortal.RootTokenAddress)
	if err != nil {
		f.logger.Error("checkBalance", "GetContractByAddress", err)
		return
	}

	var fxERC20ChildTunnelClientNPTBalance basetypes.OutputBigInt
	contract, err = cache.ContractCache().GetContractByContractID(commontypes.ChainMATIC, bridgetypes.FxERC20ChildTunnelCustomID)
	if err != nil {
		f.logger.Error("checkBalance", "GetContractByContractID", err)
		return
	}

	err = f.client.CallMsgUnmarshalContract2(
		ctx,
		tokenContract,
		"balanceOf",
		&fxERC20ChildTunnelClientNPTBalance,
		common.HexToAddress(contract.Address, contract.Address),
	)
	f.logger.Info("checkBalance",
		"fxERC20ChldTunnelClientNPTBalance", util.ToEther(fxERC20ChildTunnelClientNPTBalance.Value).String())

	// Check Eth balance of bridge-manager for exit transaction
	bridgeManagerEthBalance, err := f.client.BalanceAt(context.Background(), commontypes.ChainETH, f.cfg.Account.Address, nil)
	if err != nil {
		f.logger.Error("checkBalance", "BalanceOf", err)
	}
	f.logger.Info("checkBalance",
		"bridgeManagerEthBalance", util.ToEther(bridgeManagerEthBalance).String())
}

// ReceiveMessage
// Chain:ETH (MATIC->ETH)
func (f *FxPortal) ReceiveMessage() {
	const maxRetryCount = 30

	contract, err := cache.ContractCache().GetContractByContractID(commontypes.ChainETH, bridgetypes.FxERC20RootTunnelCustomID)
	if err != nil {
		f.logger.Error("ReceiveMessage", "GetContractByContractID", err)
		return
	}

	fxERC20RootTunnelClient, err := f.client.NewFxERC20RootTunnel(contract.Address, nil)
	f.pendingWithdrawLock.Lock()
	defer func() {
		f.pendingWithdrawLock.Unlock()
	}()

	for _, pendingWithdrawTx := range f.PendingWithdrawTxs {
		txHash := pendingWithdrawTx.TxHash
		proof, err := f.GetProofFromMaticAPI(txHash)
		f.logger.Info("NewHeaderBlock.getProofFromMaticAPI", "txHash", txHash, "proof", proof)
		if err != nil {
			continue
		}
		proofBytes, err := hexutil.Decode(proof)

		txResult, err := fxERC20RootTunnelClient.ReceiveMessage(proofBytes, f.account)
		if err != nil {
			f.logger.Error("NewHeaderBlock.ReceiveMessage", "ReceiveMessage", err)
			if strings.Contains(err.Error(), "FxRootTunnel: EXIT_ALREADY_PROCESSED") {
				f.logger.Warn("NewHeaderBlock.ReceiveMessage",
					"msg", "ReceiveMessage already processed",
					"txHash", txHash)
				delete(f.PendingWithdrawTxs, txHash)
			}
			if isReceiveMessageSkippableError(err.Error()) {
				continue
			}
			if pendingWithdrawTx.RetryCount > maxRetryCount {
				f.logger.Error("NewHeaderBlock.ReceiveMessage",
					"msg", "ReceiveMessage max retryCount exceeded",
					"RetryCount", pendingWithdrawTx.RetryCount,
					"txHash", txHash)
				delete(f.PendingWithdrawTxs, txHash)
			}
			pendingWithdrawTx.RetryCount += 1
			continue
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

		ctx := context.Background()
		_, err = f.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
			// trader에서 in_tx_hash 값을 가져오기 위해 result_tx_hash 값을 저장
			txHistory, err := f.historyService.FindOneTxHistory(sessCtx, bson.M{
				"cate":       mongoentity.TxHistoryCateBridgeOut,
				"in_tx_hash": txHash,
			})
			if err != nil {
				return nil, fmt.Errorf("FindOneTxHistory : %w", err)
			}

			err = f.historyService.UpdateOneTxHistory(ctx,
				bson.M{
					"_id": txHistory.ID,
				},
				bson.M{
					"$set": bson.M{
						"tx_hash": txResult.TxHash(),
					},
				},
			)
			if err != nil {
				return nil, err
			}

			tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(commontypes.ChainETH, f.cfg.FxPortal.RootTokenAddress)
			if err != nil {
				return nil, fmt.Errorf("GetTokenInfoByAddress: %w", err)
			}

			if tokenInfo == nil {
				return nil, fmt.Errorf("GetTokenInfoByAddress: tokenInfo is nil")
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
				SetFrom(pendingWithdrawTx.From).
				SetFromChain(commontypes.ChainMATIC).
				SetFromTxHash(txHash).
				SetToChain(commontypes.ChainETH).
				SetToTxHash(txResult.TxHash()).
				ConvertMongoEntity()

			err = f.historyService.UpsertBridgeTxHistory(ctx, bridgeTxHistory, nil)
			if err != nil {
				return nil, fmt.Errorf("UpsertBridgeTxHistory: %w", err)
			}

			transferInfo := entity.NewBridgeTransferHistory().
				SetTxHash(txResult.TxHash()).
				SetFrom(pendingWithdrawTx.From).
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

			err = f.historyService.UpsertBridgeTransferHistory(ctx, transferInfo, nil)
			if err != nil {
				return nil, fmt.Errorf("UpsertBridgeTransferHistory: %w", err)
			}

			//activeHistory := mongoentity.NewBridgeActiveHistory().
			//	SetSourceChain(commontypes.ChainMATIC).
			//	SetSourceTxHash(txHash).
			//	SetTargetChain(commontypes.ChainETH).
			//	SetTargetTxHash(txResult.TxHash())
			//
			//err = f.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
			//if err != nil {
			//	return nil, fmt.Errorf("UpsetBridgeActiveHistory: %w", err)
			//}

			delete(f.PendingWithdrawTxs, txHash)

			return nil, nil
		})
		if err != nil {
			f.logger.Debug("ReceiveMessage", "err", err)
			delete(f.PendingWithdrawTxs, txHash)
			continue
		}
	}
}

func isReceiveMessageSkippableError(errorMsg string) bool {
	// i.e) tx fee (21.48 ether) exceeds the configured cap (1.00 ether)
	txFeeMatch, _ := regexp.MatchString("tx fee(.*)exceeds the configured cap(.*)", errorMsg)
	if txFeeMatch ||
		strings.Contains(errorMsg, "Burn transaction has not been checkpointed yet") ||
		strings.Contains(errorMsg, "Insufficient funds for gas*price+value") {
		return true
	}
	return false
}

func (f *FxPortal) GetProofFromMaticAPI(txHash string) (string, error) {
	const ExitPayloadPath = "exit-payload"
	const messageSentEventSig = "0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036"
	maticApiBaseUrl := f.cfg.FxPortal.MaticApiUrl

	requestUrl := fmt.Sprintf("%s/%s/%s", maticApiBaseUrl, ExitPayloadPath, txHash)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		f.logger.Error("getProofFromMaticAPI", err)
		return "", err
	}
	query := req.URL.Query()
	query.Add("tx_hash", txHash)
	query.Add("eventSignature", messageSentEventSig)
	req.URL.RawQuery = query.Encode()
	f.logger.Debug("getProofFromMaticAPI", "requestUrl", req.URL, "txHash", txHash)

	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		f.logger.Error("getProofFromMaticAPI", "error", err)
		return "", err
	}

	var result protocol.PolygonBurnProofResp

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		f.logger.Error("getProofFromMaticAPI", err)
		return "", err
	}

	if result.Message != "Payload generation success" {
		f.logger.Error("getProofFromMaticAPI", "message", result.Message, "result", result.Result)
		return "", errors.New(result.Message)
	}

	return result.Result, nil
}
