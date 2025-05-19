package pair

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/samber/lo"
	"math/big"
	"strings"
)

const (
	MethodCreatePairByUser    = "createPairByUser"
	MethodCreatePairETHByUser = "createPairETHByUser"

	EasyRouter        = "EasyRouter"
	UniswapV2Router02 = "UniswapV2Router02"
)

const (
	EventNameCreatePairByUser    = "CreatePairByUserEvent"
	EventNameCreatePairETHByUser = "CreatePairETHByUserEvent"
	EventNameTransfer            = "Transfer"
)

type CreatePairByUserTokenInfo struct {
	TokenA   common.Address
	TokenB   common.Address
	LpAmount *big.Int
	AmountA  *big.Int
	AmountB  *big.Int
	HaveCoin bool
}

func GetPairKey(tokenA, tokenB string) string {
	return fmt.Sprintf("%s-%s", tokenA, tokenB)
}

func GetPairSymbol(tokenA, tokenB string) string {
	return fmt.Sprintf("%s-%s-LP", tokenA, tokenB)
}

func GetPairGroupID(tokenA, tokenB string) string {
	return fmt.Sprintf("G-%s-%s-LP", tokenA, tokenB)
}

func GetCurrencyID(chain, symbol string) string {
	chainList := []string{
		"BLANK",
		types.ChainKLAY,
		types.ChainMATIC,
		types.ChainETH,
	}

	// when case of coin
	if strings.Compare(chain, symbol) == 0 {
		return fmt.Sprintf("%s001", symbol)
	}

	index := lo.IndexOf(chainList, chain)
	if index == -1 {
		return fmt.Sprintf("%s001", symbol)
	}

	return fmt.Sprintf("%s%03d", symbol, index)
}
