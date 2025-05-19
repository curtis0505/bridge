package verify

import (
	"context"
	"encoding/json"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/apps/managers/util"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/entity"
	mongoentity "github.com/curtis0505/bridge/libs/entity/mongo"
	"github.com/curtis0505/bridge/libs/types"
	basetypes "github.com/curtis0505/bridge/libs/types/base"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/bridge/abi"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"math/big"
)

func (p *VerifyHandler) VerifyDeposit(ctx context.Context, tx *types.Transaction, log types.Log, event bridgetypes.EventDeposit) error {
	var tokenInfo *mongoentity.TokenInfo
	var err error
	var balanceOf *big.Int

	if event.TokenAddr.String() == bridgetypes.ZeroAddress || event.TokenAddr.String() == bridgetypes.CoinAddress {
		currencyID := tokentypes.CurrencyIdByChain(log.Chain())
		tokenInfo, err = cache.TokenCache().GetTokenInfoByCurrencyID(currencyID)
		if err != nil {
			return eventtypes.WrapError("VerifyDeposit", "GetTokenInfoByCurrencyId", err)
		}

		balanceOf, err = p.client.BalanceAt(ctx, log.Chain(), log.Address(), nil)
		if err != nil {
			return eventtypes.WrapError("VerifyDeposit", "GetBalanceAt", err)
		}
	} else {
		tokenInfo, err = cache.TokenCache().GetTokenInfoByAddress(log.Chain(), event.TokenAddr.String())
		if err != nil {
			return eventtypes.WrapError("VerifyDeposit", "GetTokenInfoByAddress", err)
		}

		var balance basetypes.OutputBigInt
		err := p.client.CallMsgUnmarshal(ctx, log.Chain(), "", tokenInfo.Address, "balanceOf", abi.GetAbiToMap(abi.ERC20Abi), &balance, common.HexToAddress(log.Chain(), log.Address()))
		if err != nil {
			return err
		}

		balanceOf = balance.Value
	}

	p.logger.Info(
		"event", "VerifyDeposit", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(), "address", log.Address(),
		"currencyId", tokenInfo.CurrencyID, "balanceOf", util.ToEtherWithDecimal(balanceOf, tokenInfo.Decimal).String(),
	)

	if balance, ok := p.vaultBalance[log.Chain()+tokenInfo.CurrencyID]; !ok {
		p.logger.Warn("event", "VerifyDeposit", "chain", log.Chain(), "msg", "initialize")

		p.vaultBalance[log.Chain()+tokenInfo.CurrencyID] = balanceOf
	} else {
		vaultBalance := new(big.Int).Add(balance, event.Amount)
		if vaultBalance.Cmp(balanceOf) != 0 {
			p.logger.Warn(
				"event", "VerifyDeposit", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "mismatch", "name", "Vault",
				"current", util.ToEtherWithDecimal(balanceOf, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(vaultBalance, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Amount, tokenInfo.Decimal).String(),
			)
		} else {
			p.logger.Info(
				"event", "VerifyDeposit", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "verified", "name", "Vault",
				"current", util.ToEtherWithDecimal(balanceOf, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(vaultBalance, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Amount, tokenInfo.Decimal).String(),
			)
		}
		p.vaultBalance[log.Chain()+tokenInfo.CurrencyID] = vaultBalance
	}

	balanceHistory := entity.NewBridgeBalanceHistory().
		SetChain(log.Chain()).
		SetAddress(log.Address()).
		SetTokenInfo(tokenInfo).
		SetBalance(balanceOf).
		SetEvent(log.EventName).
		SetName("Vault").
		SetBlockNumber(tx.Receipt().BlockNumber().Int64()).
		ConvertMongoEntity()

	err = p.historyService.UpsertBridgeBalanceHistory(ctx, balanceHistory, nil)
	if err != nil {
		p.logger.Error("UpsertBridgeBalanceHistory", err)
	}

	return nil
}

func (p *VerifyHandler) VerifyWithdraw(ctx context.Context, tx *types.Transaction, log types.Log, event bridgetypes.EventWithdraw) error {
	var tokenInfo *mongoentity.TokenInfo
	var err error
	var balanceOf *big.Int

	if event.Token.String() == bridgetypes.ZeroAddress || event.Token.String() == bridgetypes.CoinAddress {
		currencyID := tokentypes.CurrencyIdByChain(log.Chain())
		tokenInfo, err = cache.TokenCache().GetTokenInfoByCurrencyID(currencyID)
		if err != nil {
			return eventtypes.WrapError("VerifyDeposit", "GetTokenInfoByCurrencyId", err)
		}

		balanceOf, err = p.client.BalanceAt(ctx, log.Chain(), log.Address(), nil)
		if err != nil {
			return eventtypes.WrapError("VerifyDeposit", "GetBalanceAt", err)
		}
	} else {
		tokenInfo, err = cache.TokenCache().GetTokenInfoByAddress(log.Chain(), event.Token.String())
		if err != nil {
			return eventtypes.WrapError("VerifyDeposit", "GetTokenInfoByAddress", err)
		}

		var erc20ABI []map[string]interface{}
		err = json.Unmarshal([]byte(abi.MultiSigWalletAbi), &erc20ABI)
		if err != nil {
			return err
		}

		var balance basetypes.OutputBigInt
		err := p.client.CallMsgUnmarshal(ctx, log.Chain(), "", tokenInfo.Address, "balanceOf", abi.GetAbiToMap(abi.ERC20Abi), &balance, common.HexToAddress(log.Chain(), log.Address()))
		if err != nil {
			return err
		}

		balanceOf = balance.Value
	}

	p.logger.Info(
		"event", "VerifyWithdraw", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(), "address", log.Address(),
		"currencyId", tokenInfo.CurrencyID, "balanceOf", util.ToEtherWithDecimal(balanceOf, tokenInfo.Decimal).String(),
	)

	if balance, ok := p.vaultBalance[log.Chain()+tokenInfo.CurrencyID]; !ok {
		p.logger.Warn("event", "VerifyWithdraw", "chain", log.Chain(), "msg", "initialize")

		p.vaultBalance[log.Chain()+tokenInfo.CurrencyID] = balanceOf
	} else {
		vaultBalance := new(big.Int).Sub(balance, event.Uints[0])
		if vaultBalance.Cmp(balanceOf) != 0 {
			p.logger.Warn(
				"event", "VerifyWithdraw", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "mismatch", "name", "Vault",
				"current", util.ToEtherWithDecimal(balanceOf, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(vaultBalance, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Uints[0], tokenInfo.Decimal).String(),
			)
		} else {
			p.logger.Info(
				"event", "VerifyWithdraw", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "verified", "name", "Vault",
				"current", util.ToEtherWithDecimal(balanceOf, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(vaultBalance, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Uints[0], tokenInfo.Decimal).String(),
			)
		}

		p.vaultBalance[log.Chain()+tokenInfo.CurrencyID] = vaultBalance
	}

	balanceHistory := entity.NewBridgeBalanceHistory().
		SetChain(log.Chain()).
		SetAddress(log.Address()).
		SetTokenInfo(tokenInfo).
		SetBalance(balanceOf).
		SetEvent(log.EventName).
		SetName("Vault").
		SetBlockNumber(tx.Receipt().BlockNumber().Int64()).
		ConvertMongoEntity()

	err = p.historyService.UpsertBridgeBalanceHistory(ctx, balanceHistory, nil)
	if err != nil {
		p.logger.Error("UpsertBridgeBalanceHistory", err)
	}

	return nil
}

func (p *VerifyHandler) VerifyBurn(ctx context.Context, tx *types.Transaction, log types.Log, event bridgetypes.EventBurn) error {
	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), event.TokenAddr.String())
	if err != nil {
		return err
	}

	var totalSupply basetypes.OutputBigInt
	err = p.client.CallMsgUnmarshal(ctx, log.Chain(), "", tokenInfo.Address, "totalSupply", abi.GetAbiToMap(abi.ERC20Abi), &totalSupply)
	if err != nil {
		return err
	}

	p.logger.Info(
		"event", "VerifyBurn", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(), "address", log.Address(),
		"currencyId", tokenInfo.CurrencyID, "totalSupply", util.ToEtherWithDecimal(totalSupply, tokenInfo.Decimal).String(),
	)

	if supply, ok := p.minterSupply[log.Chain()+tokenInfo.CurrencyID]; !ok {
		p.logger.Warn("event", "VerifyBurn", "chain", log.Chain(), "msg", "initialize")

		p.minterSupply[log.Chain()+tokenInfo.CurrencyID] = totalSupply.Value
	} else {
		minterSupply := new(big.Int).Sub(supply, event.Amount)
		if minterSupply.Cmp(totalSupply.Value) != 0 {
			p.logger.Warn(
				"event", "VerifyBurn", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "mismatch", "name", "Minter",
				"current", util.ToEtherWithDecimal(totalSupply, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(minterSupply, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Amount, tokenInfo.Decimal).String(),
			)
		} else {
			p.logger.Info(
				"event", "VerifyBurn", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "verified", "name", "Minter",
				"current", util.ToEtherWithDecimal(totalSupply, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(minterSupply, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Amount, tokenInfo.Decimal).String(),
			)
		}

		p.minterSupply[log.Chain()+tokenInfo.CurrencyID] = minterSupply
	}

	balanceHistory := entity.NewBridgeBalanceHistory().
		SetChain(log.Chain()).
		SetAddress(log.Address()).
		SetTokenInfo(tokenInfo).
		SetTotalSupply(totalSupply.Value).
		SetEvent(log.EventName).
		SetName("Minter").
		SetBlockNumber(tx.Receipt().BlockNumber().Int64()).
		ConvertMongoEntity()

	err = p.historyService.UpsertBridgeBalanceHistory(ctx, balanceHistory, nil)
	if err != nil {
		p.logger.Error("UpsertBridgeBalanceHistory", err)
	}

	return nil
}

func (p *VerifyHandler) VerifyMint(ctx context.Context, tx *types.Transaction, log types.Log, event bridgetypes.EventMint) error {
	tokenInfo, err := cache.TokenCache().GetTokenInfoByAddress(log.Chain(), event.TokenAddr.String())
	if err != nil {
		return err
	}

	var totalSupply basetypes.OutputBigInt
	err = p.client.CallMsgUnmarshal(ctx, log.Chain(), "", tokenInfo.Address, "totalSupply", abi.GetAbiToMap(abi.ERC20Abi), &totalSupply)
	if err != nil {
		return err
	}

	p.logger.Info(
		"event", "VerifyMint", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(), "address", log.Address(),
		"currencyId", tokenInfo.CurrencyID, "totalSupply", util.ToEtherWithDecimal(totalSupply, tokenInfo.Decimal).String(),
	)

	if supply, ok := p.minterSupply[log.Chain()+tokenInfo.CurrencyID]; !ok {
		p.logger.Warn("event", "VerifyMint", "chain", log.Chain(), "msg", "initialize")

		p.minterSupply[log.Chain()+tokenInfo.CurrencyID] = totalSupply.Value
	} else {
		minterSupply := new(big.Int).Add(supply, event.Uints[0])
		if minterSupply.Cmp(totalSupply.Value) != 0 {
			p.logger.Warn(
				"event", "VerifyMint", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "mismatch", "name", "Minter",
				"current", util.ToEtherWithDecimal(totalSupply, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(minterSupply, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Uints[0], tokenInfo.Decimal).String(),
			)
		} else {
			p.logger.Info(
				"event", "VerifyMint", "chain", log.Chain(), "blockNumber", tx.Receipt().BlockNumber(),
				"msg", "verified", "name", "Minter",
				"current", util.ToEtherWithDecimal(totalSupply, tokenInfo.Decimal).String(),
				"before", util.ToEtherWithDecimal(minterSupply, tokenInfo.Decimal).String(),
				"transfer", util.ToEtherWithDecimal(event.Uints[0], tokenInfo.Decimal).String(),
			)
		}

		p.minterSupply[log.Chain()+tokenInfo.CurrencyID] = minterSupply
	}

	balanceHistory := entity.NewBridgeBalanceHistory().
		SetChain(log.Chain()).
		SetAddress(log.Address()).
		SetTokenInfo(tokenInfo).
		SetTotalSupply(totalSupply.Value).
		SetEvent(log.EventName).
		SetName("Minter").
		SetBlockNumber(tx.Receipt().BlockNumber().Int64()).
		ConvertMongoEntity()

	err = p.historyService.UpsertBridgeBalanceHistory(ctx, balanceHistory, nil)
	if err != nil {
		p.logger.Error("UpsertBridgeBalanceHistory", err)
	}

	return nil
}
