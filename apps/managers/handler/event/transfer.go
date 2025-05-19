package event

import (
	"context"
	"fmt"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/common"
	protocol "github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/entity"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	"github.com/curtis0505/bridge/libs/service"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (p *EventHandler) Deposit(log types.Log) error {
	ctx := context.Background()
	err := p.CheckAlertAmount(ctx, log)
	if err != nil {
		p.logger.Warn("CheckAlertAmount", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "HeaderByNumber", err)
	}

	eventDeposit := bridgetypes.EventDeposit{}
	if err = log.Unmarshal(&eventDeposit); err != nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "EventUnmarshal", err)
	}

	contract, err := cache.ContractCache().GetContractByContractID(log.Chain(), bridgetypes.VaultContractID)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "GetContractByContractID", err)
	}

	inputDeposit := bridgetypes.InputDeposit{}
	if err = tx.UnmarshalABI(contract.ABI, &inputDeposit); err != nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "InputUnmarshal", err)
	}

	var tokenInfo *mongoentity.TokenInfo
	if eventDeposit.TokenAddr.String() == bridgetypes.ZeroAddress || eventDeposit.TokenAddr.String() == bridgetypes.CoinAddress {
		currencyID := tokentypes.CurrencyIdByChain(log.Chain())
		tokenInfo, err = cache.TokenCache().GetTokenInfoByCurrencyID(currencyID)
		if err != nil {
			return eventtypes.WrapError(eventtypes.EventDeposit, "GetTokenInfoByCurrencyId", err)
		}
	} else {
		tokenInfo, err = cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventDeposit.TokenAddr.String())
		if err != nil {
			return eventtypes.WrapError(eventtypes.EventDeposit, "GetTokenInfoByAddress", err)
		}
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "tokenInfo", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Deposit Event
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeIn).
			SetEvent(eventtypes.EventDeposit).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(log.Chain()).
			SetFromTxHash(log.TxHash()).
			SetToChain(eventDeposit.ToChainName).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.EventDeposit,
			"chain", log.Chain(),
			"from", tx.From(),
			"txHash", tx.TxHash(),
			"fee", tx.TxFee(header.BaseFee()),
		)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, fmt.Errorf("UpsertBridgeTxHistory: %w", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                      // Sender
			SetTo(tx.To()).                          // Vault
			SetToken(*tokenInfo).                    // Token
			SetAmount(eventDeposit.Amount).          // Transfer Amount
			SetChain(log.Chain()).                   // Chain
			SetTxHash(log.TxHash()).                 // TxHash
			SetCate(protocol.TxHistoryCateBridgeIn). // Category
			SetFromChain(log.Chain()).               // From Tx
			SetFromTxHash(log.TxHash()).
			SetToChain(eventDeposit.ToChainName).
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

		sourceAmount := big.NewInt(0)
		taxAmount := big.NewInt(0)
		chainFeeAmount := big.NewInt(0)
		if eventDeposit.TokenAddr.String() == bridgetypes.ZeroAddress || eventDeposit.TokenAddr.String() == bridgetypes.CoinAddress {
			bridgePath, err := p.bridgeService.GetBridgePathBySourceCurrencyID(ctx, tokenInfo.CurrencyID, eventDeposit.ToChainName)
			if err != nil {
				return nil, fmt.Errorf("GetBridgePathBySourceCurrencyID: %w", err)
			}

			sourceAmount = func() *big.Int {
				if bridgePath.TaxRate == 0.0 {
					return eventDeposit.Amount
				} else {
					taxRate := decimal.NewFromFloat(bridgePath.TaxRate)
					transferRate := decimal.NewFromInt(1).Sub(taxRate)                       // 1 - taxRate
					return util.ToDecimal(eventDeposit.Amount, 0).Div(transferRate).BigInt() // targetAmount / transferRate
				}
			}()
			taxAmount = big.NewInt(0).Sub(sourceAmount, eventDeposit.Amount)

			value, err := tx.Value()
			if err != nil {
				return nil, fmt.Errorf("value: %w", err)
			}
			chainFeeAmount = big.NewInt(0).Sub(value, sourceAmount)
		} else {
			sourceAmount = inputDeposit.Amount
			taxAmount = new(big.Int).Sub(inputDeposit.Amount, eventDeposit.Amount)
			chainFeeAmount, err = tx.Value()
			if err != nil {
				return nil, fmt.Errorf("value: %w", err)
			}
		}

		// Tax Transfer Event
		taxTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                       // Sender
			SetTo(tx.To()).                           // Tax TODO
			SetToken(*tokenInfo).                     // Token
			SetAmount(taxAmount).                     // Transfer Amount
			SetChain(log.Chain()).                    // Chain
			SetTxHash(log.TxHash()).                  // TxHash
			SetCate(protocol.TxHistoryCateBridgeTax). // Category
			SetFromChain(log.Chain()).                // Source Tx
			SetFromTxHash(log.TxHash()).
			SetToChain(eventDeposit.ToChainName).
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
			return nil, eventtypes.WrapError(eventtypes.EventDeposit, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetFrom(tx.From()).
			SetTo(tx.From()).
			SetSourceChain(log.Chain()).
			SetSourceTxHash(log.TxHash()).
			SetSourceToken(common.HexToAddress(log.Chain(), tokenInfo.Address).String()).
			SetTargetChain(eventDeposit.ToChainName).
			SetTargetTxHash("").
			SetTargetToken("").
			SetAmount(sourceAmount).
			SetChainFee(chainFeeAmount).
			SetTax(taxTransferInfo.Amount)

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventDeposit, "UpsetBridgeActiveHistory", err)
		}

		if err = p.CheckChainFee(sessCtx, tx, log, chainFeeAmount); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventDeposit, "CheckChainFee", err)
		}

		if err = p.verifyHandler.VerifyDeposit(sessCtx, tx, log, eventDeposit); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventDeposit, "VerifyDeposit", err)
		}

		// Cosmos 계열 체인으로 가는 경우 MultiSig Transaction 을 Manager 에서 전송
		if types.GetChainType(eventDeposit.ToChainName) == types.ChainTypeCOSMOS {
			commonLog, err := bridgetypes.GetDepositCommonLog(log, bridgetypes.ContractAddresses{})
			if err != nil {
				return nil, fmt.Errorf("GetDepositLogToCommon: %w", err)
			}
			p.validatorHandler.AddCosmosMultiSigTx(validatortypes.NewCosmosMultiSigTx(log.Chain(), log.TxHash(), eventtypes.EventDeposit, commonLog))
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", log.Chain(), eventDeposit.ToChainName)).
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

func (p *EventHandler) Withdraw(log types.Log) error {
	ctx := context.Background()

	eventWithdraw := bridgetypes.EventWithdraw{}
	if err := log.Unmarshal(&eventWithdraw); err != nil {
		return eventtypes.WrapError(eventtypes.EventWithdraw, "Unmarshal", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventWithdraw, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventWithdraw, "HeaderByNumber", err)
	}

	var tokenInfo *mongoentity.TokenInfo
	if eventWithdraw.Token.String() == bridgetypes.ZeroAddress || eventWithdraw.Token.String() == bridgetypes.CoinAddress {
		currencyID := tokentypes.CurrencyIdByChain(log.Chain())
		tokenInfo, err = cache.TokenCache().GetTokenInfoByCurrencyID(currencyID)
		if err != nil {
			return eventtypes.WrapError(eventtypes.EventWithdraw, "GetTokenInfoByCurrencyId", err)
		}
	} else {
		tokenInfo, err = cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventWithdraw.Token.String())
		if err != nil {
			return eventtypes.WrapError(eventtypes.EventWithdraw, "GetTokenInfoByAddress", err)
		}
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventWithdraw, "tokenInfo", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeOut).
			SetEvent(eventtypes.EventWithdraw).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(eventWithdraw.FromChainName).
			SetFromTxHash(eventWithdraw.Bytes32s[1].String()).
			SetToChain(log.Chain()).
			SetToTxHash(log.TxHash()).
			SetTxFee(tx.TxFee(header.BaseFee())).
			ConvertMongoEntity()

		p.logger.Info("event", eventtypes.EventWithdraw, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash(), "fee", tx.TxFee(header.BaseFee()))

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, fmt.Errorf("UpsertBridgeTxHistory: %w", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(log.Address()).                   // Vault
			SetTo(eventWithdraw.To.String()).         // Recipient
			SetToken(*tokenInfo).                     // Token
			SetAmount(eventWithdraw.Uints[0]).        // Transfer Amount
			SetChain(log.Chain()).                    // Chain
			SetTxHash(log.TxHash()).                  // TxHash
			SetCate(protocol.TxHistoryCateBridgeOut). // Category
			SetFromChain(eventWithdraw.FromChainName).
			SetFromTxHash(eventWithdraw.Bytes32s[1].String()).
			SetToChain(log.Chain()).
			SetToTxHash(log.TxHash()).
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
			return nil, eventtypes.WrapError(eventtypes.EventWithdraw, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetSourceChain(eventWithdraw.FromChainName).
			SetSourceTxHash(common.HexToHash(eventWithdraw.Bytes32s[1].String()).String()).
			SetTargetChain(log.Chain()).
			SetTargetTxHash(log.TxHash()).
			SetTargetToken(common.HexToAddress(log.Chain(), tokenInfo.Address).String()).
			SetUpdatedAt(time.Now())

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventSyncDeposit, "UpsetBridgeActiveHistory", err)
		}

		if err = p.verifyHandler.VerifyWithdraw(sessCtx, tx, log, eventWithdraw); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventWithdraw, "VerifyWithdraw", err)

		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", eventWithdraw.FromChainName, log.Chain())).
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

func (p *EventHandler) Mint(log types.Log) error {
	ctx := context.Background()

	eventMint := bridgetypes.EventMint{}
	if err := log.Unmarshal(&eventMint); err != nil {
		return eventtypes.WrapError(eventtypes.EventMint, "Unmarshal", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventMint, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventMint, "HeaderByNumber", err)
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventMint.TokenAddr.String())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventMint, "GetTokenInfoByAddress", err)
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventMint, "tokenInfo", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeMint).
			SetEvent(eventtypes.EventMint).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(eventMint.FromChainName).
			SetFromTxHash(eventMint.Bytes32s[1].String()).
			SetToChain(log.Chain()).
			SetToTxHash(log.TxHash()).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info("event", eventtypes.EventMint, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash(), "fee", tx.TxFee(header.BaseFee()))

		p.logger.Info("event", eventtypes.EventMint, "chain", log.Chain(), "from", tx.From(), "txHash", tx.TxHash())

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventMint, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(log.Address()).                        // Minter
			SetTo(eventMint.To.String()).                  // Recipient
			SetToken(*tokenInfo).                          // Token
			SetAmount(eventMint.Uints[0]).                 // Transfer Amount
			SetChain(log.Chain()).                         // Chain
			SetTxHash(log.TxHash()).                       // TxHash
			SetCate(protocol.TxHistoryCateBridgeMint).     // Category
			SetFromChain(eventMint.FromChainName).         // FromChain
			SetFromTxHash(eventMint.Bytes32s[1].String()). // Source Tx
			SetToChain(log.Chain()).
			SetToTxHash(log.TxHash()).
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
			return nil, eventtypes.WrapError(eventtypes.EventMint, "UpsertBridgeTransferHistory", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetSourceChain(eventMint.FromChainName).
			SetSourceTxHash(common.HexToHash(eventMint.Bytes32s[1].String()).String()).
			SetTargetChain(log.Chain()).
			SetTargetTxHash(log.TxHash()).
			SetTargetToken(common.HexToAddress(log.Chain(), tokenInfo.Address).String()).
			SetUpdatedAt(time.Now())

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventMint, "UpsetBridgeActiveHistory", err)
		}

		if err = p.verifyHandler.VerifyMint(sessCtx, tx, log, eventMint); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventMint, "VerifyMint", err)

		}
		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", eventMint.FromChainName, log.Chain())).
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

func (p *EventHandler) Burn(log types.Log) error {
	ctx := context.Background()

	err := p.CheckAlertAmount(ctx, log)
	if err != nil {
		p.logger.Warn("CheckAlertAmount", err)
	}

	tx, _, err := p.client.GetTransactionWithReceipt(ctx, log.Chain(), log.TxHash())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventBurn, "GetTransactionWithReceipt", err)
	}

	header, err := p.client.HeaderByNumber(ctx, log.Chain(), tx.Receipt().BlockNumber())
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventBurn, "HeaderByNumber", err)
	}

	eventBurn := bridgetypes.EventBurn{}
	if err := log.Unmarshal(&eventBurn); err != nil {
		return eventtypes.WrapError(eventtypes.EventBurn, "EventUnmarshal", err)
	}

	contract, err := cache.ContractCache().GetContractByContractID(log.Chain(), bridgetypes.MinterContractID)
	if err != nil {
		return eventtypes.WrapError(eventtypes.EventDeposit, "GetContractByContractID", err)
	}

	inputBurn := bridgetypes.InputBurn{}
	if err := tx.UnmarshalABI(contract.ABI, &inputBurn); err != nil {
		return eventtypes.WrapError(eventtypes.EventBurn, "InputUnmarshal", err)
	}

	var tokenInfo *mongoentity.TokenInfo
	if eventBurn.TokenAddr.String() == bridgetypes.ZeroAddress || eventBurn.TokenAddr.String() == bridgetypes.CoinAddress {
		currencyID := tokentypes.CurrencyIdByChain(log.Chain())
		tokenInfo, err = cache.TokenCache().GetTokenInfoByCurrencyID(currencyID)
		if err != nil {
			return eventtypes.WrapError(eventtypes.EventBurn, "GetTokenInfoByCurrencyId", err)
		}
	} else {
		tokenInfo, err = cache.TokenCache().GetTokenInfoByAddress(log.Chain(), eventBurn.TokenAddr.String())
		if err != nil {
			return eventtypes.WrapError(eventtypes.EventBurn, "GetTokenInfoByAddress", err)
		}
	}

	if tokenInfo == nil {
		return eventtypes.WrapError(eventtypes.EventBurn, "tokenInfo", fmt.Errorf("TokenInfo is nil"))
	}

	_, err = p.transactionManager.WithMongoTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		txHistory := entity.NewBridgeTxHistory().
			SetTxHash(log.TxHash()).
			SetCate(protocol.TxHistoryCateBridgeBurn).
			SetEvent(eventtypes.EventBurn).
			SetChain(log.Chain()).
			SetFrom(tx.From()).
			SetFromChain(log.Chain()).
			SetFromTxHash(log.TxHash()).
			SetToChain(eventBurn.ToChainName).
			SetTxFee(tx.TxFee(header.BaseFee())).
			SetBlockNumber(header.BlockNumber()).
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.EventBurn,
			"chain", log.Chain(),
			"from", tx.From(),
			"txHash", tx.TxHash(),
			"fee", tx.TxFee(header.BaseFee()),
		)

		err = p.historyService.UpsertBridgeTxHistory(sessCtx, txHistory, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "UpsertBridgeTxHistory", err)
		}

		// Transfer Event
		transferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                        // Sender
			SetTo(tx.To()).                            // Minter
			SetToken(*tokenInfo).                      // Token
			SetAmount(eventBurn.Amount).               // Transfer Amount
			SetChain(log.Chain()).                     // Chain
			SetTxHash(log.TxHash()).                   // TxHash
			SetCate(protocol.TxHistoryCateBridgeBurn). // Category
			SetFromChain(log.Chain()).
			SetFromTxHash(log.TxHash()).
			SetToChain(eventBurn.ToChainName). // Source Tx
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
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "UpsertBridgeTransferHistory", err)
		}

		// Tax Transfer Event
		taxTransferInfo := entity.NewBridgeTransferHistory().
			SetFrom(tx.From()).                                              // Sender
			SetTo(tx.To()).                                                  // Tax TODO
			SetToken(*tokenInfo).                                            // Token
			SetAmount(new(big.Int).Sub(inputBurn.Amount, eventBurn.Amount)). // Transfer Amount
			SetChain(log.Chain()).                                           // Chain
			SetTxHash(log.TxHash()).                                         // TxHash
			SetCate(protocol.TxHistoryCateBridgeTax).                        // Category
			SetFromChain(log.Chain()).
			SetFromTxHash(log.TxHash()).
			SetToChain(eventBurn.ToChainName). // To Chain
			ConvertMongoEntity()

		p.logger.Info(
			"event", eventtypes.LogEventTransfer,
			"chain", transferInfo.Chain,
			"cate", taxTransferInfo.Cate,
			"from", taxTransferInfo.From,
			"to", taxTransferInfo.To,
		)

		err = p.historyService.UpsertBridgeTransferHistory(sessCtx, taxTransferInfo, nil)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "UpsertBridgeTransferHistory", err)
		}

		chainFeeAmount, err := tx.Value()
		if err != nil {
			p.logger.Error("Value", err)
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "Value", err)
		}

		activeHistory := mongoentity.NewBridgeActiveHistory().
			SetFrom(tx.From()).
			SetTo(tx.From()).
			SetSourceChain(log.Chain()).
			SetSourceTxHash(log.TxHash()).
			SetSourceToken(common.HexToAddress(log.Chain(), tokenInfo.Address).String()).
			SetTargetChain(eventBurn.ToChainName).
			SetTargetTxHash("").
			SetTargetToken("").
			SetAmount(inputBurn.Amount).
			SetChainFee(chainFeeAmount).
			SetTax(taxTransferInfo.Amount)

		err = p.historyService.UpsetBridgeActiveHistory(sessCtx, activeHistory)
		if err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "UpsetBridgeActiveHistory", err)
		}

		if err = p.CheckChainFee(sessCtx, tx, log, chainFeeAmount); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "VerifyBurn", err)
		}

		if err = p.verifyHandler.VerifyBurn(sessCtx, tx, log, eventBurn); err != nil {
			return nil, eventtypes.WrapError(eventtypes.EventBurn, "VerifyBurn", err)
		}

		// Cosmos 계열 체인으로 가는 경우 MultiSig Transaction 을 Manager 에서 전송
		if types.GetChainType(eventBurn.ToChainName) == types.ChainTypeCOSMOS {
			commonLog, err := bridgetypes.GetBurnCommonLog(log, bridgetypes.ContractAddresses{})
			if err != nil {
				return nil, eventtypes.WrapError(eventtypes.EventBurn, "GetBurnCommonLog", err)
			}
			p.validatorHandler.AddCosmosMultiSigTx(validatortypes.NewCosmosMultiSigTx(log.Chain(), log.TxHash(), eventtypes.EventBurn, commonLog))
		}

		return nil, nil
	})
	if err != nil {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", log.Chain(), eventBurn.ToChainName)).
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

func (p *EventHandler) CheckChainFee(ctx context.Context, tx *types.Transaction, log types.Log, chainFeeAmount *big.Int) error {
	if chainFeeAmount.Cmp(big.NewInt(0)) == 0 {
		p.logger.Warn(
			"event", eventtypes.LogEventFeeTransfer, "chain", log.Chain(), "cate", "ChainFee",
		)
		return nil
	}

	tokenInfo, err := cache.TokenCache().GetTokenInfoByCurrencyID(tokentypes.CurrencyIdByChain(log.Chain()))
	if err != nil {
		p.logger.Warn(
			"event", eventtypes.LogEventFeeTransfer, "chain", log.Chain(), "cate", "ChainFee",
		)
		return err
	}

	if tokenInfo == nil {
		return fmt.Errorf("TokenInfo is nil")
	}

	//Chain Fee Transfer Event
	taxTransferInfo := entity.NewBridgeTransferHistory().
		SetFrom(tx.From()).                          // Sender
		SetTo(tx.To()).                              // To
		SetToken(*tokenInfo).                        // Token
		SetAmount(chainFeeAmount).                   // Transfer Amount
		SetChain(log.Chain()).                       // Chain
		SetTxHash(log.TxHash()).                     // TxHash
		SetCate(eventtypes.TxHistoryBridgeChainFee). // Category
		SetFromChain(log.Chain()).
		SetFromTxHash(log.TxHash()).
		ConvertMongoEntity()

	p.logger.Info(
		"event", eventtypes.LogEventFeeTransfer,
		"chain", log.Chain(),
		"cate", taxTransferInfo.Cate,
		"from", taxTransferInfo.From,
		"to", taxTransferInfo.To,
	)

	err = p.historyService.UpsertBridgeTransferHistory(ctx, taxTransferInfo, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *EventHandler) CheckAlertAmount(ctx context.Context, eventLog types.Log) error {
	var log bridgetypes.CommonEventLog
	var err error
	switch bridgetypes.GetEventType(eventLog.EventName) {
	case bridgetypes.EventNameDeposit:
		log, err = bridgetypes.GetDepositCommonLog(eventLog, p.contracts)
		if err != nil {
			p.logger.Error("GetDepositLogToCommon", err)
			return types.WrapError("GetDepositLogToCommon", err)
		}

	case bridgetypes.EventNameBurn:
		log, err = bridgetypes.GetBurnCommonLog(eventLog, p.contracts)
		if err != nil {
			p.logger.Error("GetBurnCommonLog", err)
			return types.WrapError("GetBurnCommonLog", err)
		}

	default:
		return types.WrapError("CheckAlertAmount", fmt.Errorf("not support event name"))
	}

	var sourceTokenInfo *mongoentity.TokenInfo
	if log.FromTokenAddr.String() == bridgetypes.CoinAddress {
		sourceTokenInfo, err = cache.TokenCache().GetTokenInfoByDenom(log.FromChainName, tokentypes.DenomByChain(log.FromChainName))
		if err != nil {
			return types.WrapError("GetTokenInfoByAddress", err)
		}
	} else {
		sourceTokenInfo, err = cache.TokenCache().GetTokenInfoByAddress(log.FromChainName, log.FromTokenAddr.String())
		if err != nil {
			return types.WrapError("GetTokenInfoByAddress", err)
		}
	}

	if sourceTokenInfo == nil {
		return types.WrapError("CheckAlertAmount", fmt.Errorf("TokenInfo is nil"))
	}

	config, err := service.GetRegistry().ConfigService().FindOneAppSettingByKey(ctx, "bridge_alert_balance")
	if err != nil {
		return types.WrapError("FindOneAppSettingByKey", err)
	}

	tokenValue := util.BalanceToDecimal(
		log.Amount,
		float32(cache.PriceCache().GetPriceBySymbolWithNoErr(strings.ToLower(sourceTokenInfo.Symbol))),
		sourceTokenInfo.Decimal,
	)

	if reflect.TypeOf(config.Value).Kind() != reflect.String {
		return types.WrapError("TypeOf", fmt.Errorf("invalid type"))
	}

	alertBalance, err := strconv.ParseFloat(config.Value.(string), 64)
	if err != nil {
		return types.WrapError("ParseFloat", err)
	}

	value, _ := tokenValue.Float64()
	if value > alertBalance {
		util.NewMessage().
			SetZone(p.cfg.Server.ServiceId).
			SetMessageType(util.MessageTypeAlert).
			SetTitle(fmt.Sprintf("🚨%s -> %s Bridge🚨", log.FromChainName, log.ToChainName)).
			AddKeyValueWidget("Event", eventLog.EventName).
			AddKeyValueWidget("Event Amount", fmt.Sprintf("%v %s", util.ToDecimal(log.Amount, sourceTokenInfo.Decimal).String(), sourceTokenInfo.Symbol)).
			AddKeyValueWidget("From", common.BytesToAddress(log.FromChainName, log.From.Bytes())).
			AddKeyValueWidget("To", common.BytesToAddress(log.ToChainName, log.ToAddr.Bytes())).
			AddKeyValueWidget("TxHash", log.TxHash).
			AddKeyValueWidget("Description", fmt.Sprintf("%s transaction was executed for more than %s dollars", eventLog.EventName, config.Value.(string))).
			SendMessage()
		return nil
	}

	return nil
}
