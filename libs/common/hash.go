package common

import (
	ethercommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"reflect"
)

var hashT = reflect.TypeOf(ethercommon.Hash{})

type Hash [32]byte

func (hash Hash) String() string { return hexutil.Encode(hash[:]) }

func (hash *Hash) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(hashT, input, hash[:])
}

func HexToHash(hash string) Hash {
	var h Hash
	b := ethercommon.HexToHash(hash).Bytes()
	copy(h[:], b)
	return h
}
