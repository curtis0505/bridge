package util

import (
	"github.com/codahale/sss"
	"github.com/kaiachain/kaia/common/hexutil"
)

func ShamirEnc(secret string, create int, threshold int) (map[byte][]byte, error) {
	n := byte(create)
	k := byte(threshold)

	shares, err := sss.Split(n, k, []byte(secret))
	if err != nil {
		return nil, err
	}

	return shares, nil
}

func ShamirDec(shares map[byte][]byte, threshold int) string {

	subset := make(map[byte][]byte, threshold)
	for x, y := range shares {
		hexData := hexutil.Encode(y)
		byteData, _ := hexutil.Decode(hexData)
		subset[x] = byteData
		if len(subset) == int(threshold) {
			break
		}
	}

	return string(sss.Combine(subset))
}
