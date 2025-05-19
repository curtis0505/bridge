package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenType int

const (
	TokenTypeCoin TokenType = iota + 1
	TokenTypeToken
	TokenTypeLp
)

type Tokens struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Symbol            string             `bson:"symbol"`         // 코인, 토큰 심볼
	DisplaySymbol     string             `bson:"display_symbol"` // 표시용 심볼
	Name              string             `bson:"name"`           // 코인, 토큰 이름
	Chain             string             `bson:"chain"`          // 체인의 심볼 (NPT의 경우 KLAY)
	ContractID        primitive.ObjectID `bson:"contract_id"`    // 컨트랙트 ID(contractDB.contracts)
	Image             string             `bson:"image"`          // 토큰 심볼 이미지
	PriceSite         string             `bson:"price_site"`
	PriceKey          string             `bson:"price_key"`   // Price 구하는 Key
	Address           string             `bson:"address"`     // 0: 안보임, 1: 보임
	CurrencyID        string             `bson:"currency_id"` // 0: 안보임, 1: 보임
	GroupID           string             `bson:"group_id"`
	Denom             string             `bson:"denom"`   // cosmos denominator
	Type              TokenType          `bson:"type"`    // 1: 코인, 2: 체인에서 발행된 토큰
	Decimal           int                `bson:"decimal"` // 자릿수
	Stat              int                `bson:"stat"`    // 0: 사용 안함, 1: 사용
	Swap              int                `bson:"swap"`    // 0: 스왑 불가  1: 스왑 가능
	Bridge            int                `bson:"bridge"`  //
	Liquid            int                `bson:"liquid"`  //
	Digit             int                `bson:"digit"`
	NeopinBridgeToken int                `bson:"neopin_bridge_token"` // 0: 그냥 토큰 1: 네오핀 브릿지 토큰(e.g. nETH, nFNSA...)
	IsPartner         int                `bson:"is_partner"`          // 1: 제휴 토큰, 0: 제휴가 아닌 토큰
	IsView            int                `bson:"is_view"`             // 0: 안보임, 1: 보임
	Order             int                `bson:"order"`
	ChainOrder        int                `bson:"chain_order"`
	TraceHolder       int                `bson:"trace_holder"`
	PoolAllowed       int                `bson:"pool_allowed"` // pool 지원 여부
	IsBeta            bool               `bson:"is_beta"`      //beta 서비스 token 여부
	VisibleServices   []string           `bson:"visible_services"`
}

func (t *Tokens) Clone() *Tokens {
	if t == nil {
		return nil
	}

	return &Tokens{
		ID:                t.ID,
		Symbol:            t.Symbol,
		DisplaySymbol:     t.DisplaySymbol,
		Name:              t.Name,
		Chain:             t.Chain,
		ContractID:        t.ContractID,
		Image:             t.Image,
		PriceSite:         t.PriceSite,
		PriceKey:          t.PriceKey,
		Address:           t.Address,
		CurrencyID:        t.CurrencyID,
		GroupID:           t.GroupID,
		Denom:             t.Denom,
		Type:              t.Type,
		Decimal:           t.Decimal,
		Stat:              t.Stat,
		Swap:              t.Swap,
		Bridge:            t.Bridge,
		Liquid:            t.Liquid,
		Digit:             t.Digit,
		NeopinBridgeToken: t.NeopinBridgeToken,
		IsPartner:         t.IsPartner,
		IsView:            t.IsView,
		Order:             t.Order,
		ChainOrder:        t.ChainOrder,
		TraceHolder:       t.TraceHolder,
		PoolAllowed:       t.PoolAllowed,
		IsBeta:            t.IsBeta,
		VisibleServices:   t.VisibleServices,
	}
}
