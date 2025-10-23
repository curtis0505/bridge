package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	arbcommon "github.com/curtis0505/arbitrum/common"
	basecommon "github.com/curtis0505/base/common"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/types"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	klaycommon "github.com/kaiachain/kaia/common"
	"strings"
)

const EmptyAddressString string = "0x0000000000000000000000000000000000000000"

type Address interface {
	Bytes() []byte
	String() string
}

func HexToAddress(chain, address string) Address {
	switch strings.ToUpper(chain) {
	// TODO: 체인 추가시 체크 필요
	case types.ChainETH:
		return ethercommon.HexToAddress(address)
	case types.ChainKLAY:
		return klaycommon.HexToAddress(address)
	case types.ChainARB:
		return arbcommon.HexToAddress(address)
	case types.ChainBASE:
		return basecommon.HexToAddress(address)
	case types.ChainATOM, types.ChainKAVA, types.ChainOSMO, types.ChainFNSA, types.ChainTFNSA:
		return cosmoscommon.FromBech32UnSafe(chain, address)
	default:
		return ethercommon.HexToAddress(address)
	}
}

func HexToBytes(address string) [20]byte {
	return ethercommon.HexToAddress(address)
}

func EmptyAddress(address Address) bool {
	switch v := address.(type) {
	case klaycommon.Address:
		return v == klaycommon.Address{}
	case ethercommon.Address:
		return v == ethercommon.Address{}
	case fmt.Stringer:
		return v.String() == "0x0000000000000000000000000000000000000000" || v.String() == ""
	case nil:
		return true
	default:
		return v.String() == "0x0000000000000000000000000000000000000000" || v.String() == ""
	}
}

func BytesToAddress(chain string, address []byte) string {
	switch types.GetChainType(chain) {
	// TODO: 체인 추가시 체크 필요
	case types.ChainTypeCOSMOS:
		a, _ := bech32.ConvertAndEncode(cosmoscommon.GetAddressPrefixByChain(chain), address)
		return a
	case types.ChainTypeEVM:
		if chain == types.ChainKLAY {
			return klaycommon.BytesToAddress(address).String()
		} else if chain == types.ChainARB {
			return arbcommon.BytesToAddress(address).String()
		}
		return ethercommon.BytesToAddress(address).String()
	}

	return ""
}

func PublicKeyToAddress(chain string, publicKey any) (string, error) {
	switch v := publicKey.(type) {
	case ecdsa.PublicKey:
		switch types.GetChainType(chain) {
		case types.ChainTypeEVM:
			return crypto.PubkeyToAddress(v).String(), nil
		case types.ChainTypeCOSMOS:
			return cosmoscommon.FromPublicKeyUnSafe(chain, elliptic.Marshal(v.Curve, v.X, v.Y)).String(), nil
		}
	case cosmoscryptotypes.PubKey:
		switch types.GetChainType(chain) {
		case types.ChainTypeCOSMOS:
			return cosmoscommon.FromPublicKeyUnSafe(chain, v.Bytes()).String(), nil
		default:
			return "", fmt.Errorf("cannot be converted publickey type, %s: %v", chain, v)
		}
	case []byte:
		return BytesToAddress(chain, v), nil
	case nil:
		return "", fmt.Errorf("empty publicKey")
	}
	return "", fmt.Errorf("cannot be converted publickey")
}
