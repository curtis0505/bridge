package types

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/kaiachain/kaia/common"
)

type SingleAccountFilter struct {
	filter *bloom.BloomFilter
}

func NewSingleAccountFilter(estimatedItems uint, falsePositiveRate float64) *SingleAccountFilter {
	return &SingleAccountFilter{
		filter: bloom.NewWithEstimates(estimatedItems, falsePositiveRate),
	}
}

func (f *SingleAccountFilter) AddString(str string) {
	f.filter.Add([]byte(str))
}

func (f *SingleAccountFilter) AddBytes(b []byte) {
	f.filter.Add(b)
}

func (f *SingleAccountFilter) IsMemberString(str string) bool {
	if f.filter.Test([]byte(str)) {
		return true
	} else {
		return false
	}
}

func (f *SingleAccountFilter) IsMember(b []byte) bool {
	if f.filter.Test(b) {
		return true
	} else {
		return false
	}
}

type AccountFilter struct {
	filter            map[string]*SingleAccountFilter
	estimatedItems    uint
	falsePositiveRate float64
}

func NewAccountFilter(chains []string, estimatedItems uint, falsePositiveRate float64) *AccountFilter {
	accFilter := &AccountFilter{
		filter:            make(map[string]*SingleAccountFilter),
		estimatedItems:    estimatedItems,
		falsePositiveRate: falsePositiveRate,
	}

	for _, chain := range chains {
		accFilter.filter[chain] = NewSingleAccountFilter(estimatedItems, falsePositiveRate)
	}
	return accFilter
}

func (f *AccountFilter) AddString(chain, str string) {
	f.filter[chain].AddString(str)
}

func (f *AccountFilter) AddHexString(chain, addr string) {
	bytes := common.FromHex(addr)
	f.filter[chain].AddBytes(bytes)
}

func (f *AccountFilter) IsMemberString(chain, str string) bool {
	return f.filter[chain].IsMemberString(str)
}

func (f *AccountFilter) IsMemberHexString(chain, hexStr string) bool {
	bytes := common.FromHex(hexStr)
	return f.filter[chain].IsMember(bytes)
}
