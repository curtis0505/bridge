package cosmos

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func NewTxHash(txBytes []byte) Hash {
	return sha256.Sum256(txBytes)
}

type Hash [32]byte

func (h Hash) String() string {
	return strings.ToUpper(hex.EncodeToString(h[:]))
}

func (h Hash) Bytes() [32]byte {
	return h
}
