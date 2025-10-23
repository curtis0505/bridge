package types

import (
	"fmt"
	arbhexutil "github.com/curtis0505/arbitrum/common/hexutil"
	basehexutil "github.com/curtis0505/base/common/hexutil"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	klayhexutil "github.com/kaiachain/kaia/common/hexutil"
	"strings"
)

func HexUtilDecode(chain string, input string) ([]byte, error) {
	// TODO: 체인 추가시 체크 필요
	switch strings.ToUpper(chain) {
	case ChainKLAY:
		return klayhexutil.Decode(input)
	case ChainETH, ChainMATIC:
		return etherhexutil.Decode(input)
	case ChainARB:
		return arbhexutil.Decode(input)
	case ChainBASE:
		return basehexutil.Decode(input)
	default:
		return nil, fmt.Errorf("invalid chain")
	}
}

func HexUtilEncode(chain string, b []byte) string {
	// TODO: 체인 추가시 체크 필요
	switch strings.ToUpper(chain) {
	case ChainKLAY:
		return klayhexutil.Encode(b)
	case ChainETH, ChainMATIC:
		return etherhexutil.Encode(b)
	case ChainARB:
		return arbhexutil.Encode(b)
	case ChainBASE:
		return basehexutil.Encode(b)
	default:
		return ""
	}
}
