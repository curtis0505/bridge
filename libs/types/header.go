package types

import (
	basetypes "github.com/curtis0505/base/core/types"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	klaytypes "github.com/kaiachain/kaia/blockchain/types"
	"math/big"
)

type Header struct {
	inner interface{}
}

func NewHeader(header interface{}) *Header {
	return &Header{
		inner: header,
	}
}

func (r *Header) BlockNumber() *big.Int {
	switch v := r.inner.(type) {
	case *klaytypes.Header:
		return v.Number
	case *ethertypes.Header:
		return v.Number
	case *basetypes.Header:
		return v.Number
	}
	return big.NewInt(0)
}

func (r *Header) BaseFee() *big.Int {
	switch v := r.inner.(type) {
	case *klaytypes.Header:
		return v.BaseFee
	case *ethertypes.Header:
		return v.BaseFee
	case *basetypes.Header:
		return v.BaseFee
	}
	return big.NewInt(0)
}

func (r *Header) Time() uint64 {
	switch v := r.inner.(type) {
	case *klaytypes.Header:
		return v.Time.Uint64()
	case *ethertypes.Header:
		return v.Time
	case *basetypes.Header:
		return v.Time
	}
	return 0
}

func (r *Header) GasUsed() uint64 {
	switch v := r.inner.(type) {
	case *klaytypes.Header:
		return v.GasUsed
	case *ethertypes.Header:
		return v.GasUsed
	case *basetypes.Header:
		return v.GasUsed
	}
	return 0
}
