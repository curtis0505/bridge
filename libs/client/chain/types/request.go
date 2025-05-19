package types

import (
	"math/big"
)

type ProxyRequest struct {
	// CheckFee 수수료 10회 체크 여부
	CheckFee bool
	// PushRedis Redis 에 tx_hash 전송 여부
	PushRedis bool
	// InsertTxHistory tx_history insert 여부
	InsertTxHistory bool
	// 가속화 혹은 취소 트랜잭션 여부, normal - 0, speedup - 1, cancel - 2
	ProxyRequestType int

	TxHistoryId      string   `json:"tx_history_id"`
	Uuid             string   `json:"uuid"`
	Cate             string   `json:"cate"`
	Chain            string   `json:"chain"`
	BaseChain        string   `json:"base_chain"`
	TargetChain      string   `json:"target_chain"`
	TaxRate          float64  `json:"taxRate"`      // 삭제 예정: 20230317
	EstimateRate     float64  `json:"estimateRate"` // 스왑 통계용
	Tax              *big.Int `json:"tax"`
	ChainFee         *big.Int `json:"chainFee"`
	Symbol           string   `json:"symbol"`
	CurrencyId       string   `json:"currencyId"`
	BaseCurrencyId   string   `json:"base_currency_id"`
	TargetCurrencyId string   `json:"target_currency_id"`
	BaseSymbol       string   `json:"base_symbol"`
	TargetSymbol     string   `json:"target_symbol"`
	GroupId          string   `json:"groupId"`
	Product          string   `json:"product"`
	Amount           *big.Int `json:"amount"`
	BaseAmount       *big.Int `json:"baseAmount"`
	TargetAmount     *big.Int `json:"targetAmount"`
	DeviceName       string   `json:"deviceName"`    // 디바이스 이름 : ios, android, neopin_extension
	DeviceVersion    string   `json:"deviceVersion"` // 디바이스의 버전 : x.x.x
}
