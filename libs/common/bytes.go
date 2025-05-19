package common

import "github.com/ethereum/go-ethereum/common/hexutil"

type Bytes []byte

func (bytes Bytes) String() string { return hexutil.Encode(bytes) }
