package token

import (
	"github.com/curtis0505/bridge/libs/types"
)

const (
	ChainNameKAIA = "Kaia"
)

const (
	SymbolETH   = "ETH"
	SymbolKLAY  = "KLAY"
	SymbolKAIA  = "KAIA"
	SymbolTRX   = "TRX"
	SymbolMATIC = "MATIC"
	SymbolPOL   = "POL"
	SymbolNPT   = "NPT"
	SymbolS2T   = "S2T"
	SymbolBVT   = "BVT"
	SymbolARB   = "ARB"
	SymbolDAI   = "DAI"
	SymbolUSDe  = "USDe"
	SymbolUSDC  = "USDC"
	SymbolSRUSD = "srUSD"
)

const (
	CurrencyETH    = "ETH001"
	CurrencyKLAY   = "KLAY001"
	CurrencyTRX    = "TRX001"
	CurrencyMATIC  = "MATIC001"
	CurrencyWETH   = "WETH001"
	CurrencyWETH2  = "WETH002"
	CurrencyWKLAY  = "WKLAY001"
	CurrencyWMATIC = "WMATIC001"
	CurrencyARBETH = "ETH002"

	// CurrencyNPT Klaytn
	CurrencyNPT = "NPT001"
	// CurrencyNPT2 Ethereum
	CurrencyNPT2 = "NPT002"
	// CurrencyNPT3 Polygon
	CurrencyNPT3 = "NPT003"

	CurrencyS2T    = "S2T001"
	CurrencyBVT    = "BVT001"
	CurrencyNPKLAY = "npKLAY001"
	CurrencyNPETH  = "npETH001"

	CurrencyNFNSA = "nFNSA001"

	// CurrencyDAI2 Ethereum
	CurrencyDAI2  = "DAI002"
	CurrencySDAI  = "sDAI001"
	CurrencyUSDe  = "USDe001"
	CurrencySUSDe = "sUSDe001"

	CurrencySTETH  = "stETH001"
	CurrencyWSTETH = "wstETH001"

	CurrencyFNSA  = "FNSA001"
	CurrencyTFNSA = "TFNSA001"
	CurrencyATOM  = "ATOM001"
	CurrencyOSMO  = "OSMO001"
	CurrencyKAVA  = "KAVA001"

	CurrencySRUSD     = "srUSD001"     //RWABoostVault LP
	CurrencyRUSUY     = "rUSDY001"     //RWABoostVault LP
	CurrencyUsualUSDC = "USUALUSDC001" //RWABoostVault LP
)

var WrappedToken = map[string]string{
	CurrencyKLAY:  CurrencyWKLAY,
	CurrencyETH:   CurrencyWETH,
	CurrencyMATIC: CurrencyWMATIC,
}

const (
	// Lp Tokens
	LPStrategyTurboStETH = "LP-StrategyTurboStETH001"
)

const (
	EventNameApproval = "Approval"
	EventNameTransfer = "Transfer"
)

const (
	GroupTRX   = "G-TRX1"
	GroupKLAY  = "G-KLAY1"
	GroupNPT   = "G-NPT1"
	GroupMATIC = "G-MATIC1"
	GroupETH   = "G-ETH1"
	GroupARB   = "G-ARB1"
)

const (
	MethodTransfer    = "transfer"
	MethodApprove     = "approve"
	MethodBalanceOf   = "balanceOf"
	MethodSymbol      = "symbol"
	MethodDecimals    = "decimals"
	MethodName        = "name"
	MethodTotalSupply = "totalSupply"

	IERC20Id = "ierc20"
	IERC20   = "IERC20"
)

const (
	Maintenance = "transfer"
)

const (
	DenomATOM  = "uatom"
	DenomKAVA  = "ukava"
	DenomOSMO  = "uosmo"
	DenomFNSA  = "cony"
	DenomTFNSA = "tcony"
)

// CurrencyIdByChain - 해당 체인의 코인을 조회
func CurrencyIdByChain(symbol string) string {
	switch symbol {
	case types.ChainKLAY:
		return CurrencyKLAY
	case types.ChainETH:
		return CurrencyETH
	case types.ChainMATIC:
		return CurrencyMATIC
	case types.ChainARB:
		return CurrencyARBETH
	case types.ChainTRX:
		return CurrencyTRX

	case types.ChainATOM:
		return CurrencyATOM
	case types.ChainOSMO:
		return CurrencyOSMO
	case types.ChainKAVA:
		return CurrencyKAVA

	case types.ChainFNSA:
		return CurrencyFNSA
	case types.ChainTFNSA:
		return CurrencyTFNSA
	}

	return ""
}

func CheckWrappedToken(currencyID string) bool {
	_, ok := WrappedToken[currencyID]
	return ok
}

func WrappedCurrencyIdByChain(chain string) string {
	switch chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		return CurrencyWKLAY
	case types.ChainETH:
		return CurrencyWETH
	case types.ChainMATIC:
		return CurrencyWMATIC
	}

	return ""
}

func LiquidCurrencyIdByChain(chain string) string {
	switch chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainKLAY:
		return CurrencyNPKLAY
	case types.ChainETH:
		return CurrencyNPETH
	}

	return ""
}

func DenomByChain(chain string) string {
	switch chain {
	// TODO: 체인 추가시 체크 필요
	case types.ChainATOM:
		return DenomATOM
	case types.ChainKAVA:
		return DenomKAVA
	case types.ChainOSMO:
		return DenomOSMO

	case types.ChainFNSA:
		return DenomFNSA
	case types.ChainTFNSA:
		return DenomTFNSA
	}

	return ""
}
