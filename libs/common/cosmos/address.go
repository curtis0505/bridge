package cosmos

import (
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"golang.org/x/crypto/ripemd160"
)

type Address struct {
	chain   string
	bytes   []byte
	address []byte
}

func (a *Address) String() string {
	prefix := GetAddressPrefixByChain(a.chain)
	if prefix == "" {
		return ""
	}

	if a.address == nil {
		sha := sha256.Sum256(a.bytes)
		ripemd := ripemd160.New()
		ripemd.Write(sha[:])
		a.address = ripemd.Sum(nil)
	}

	addr, err := bech32.ConvertAndEncode(prefix, a.address)
	if err != nil {
		return ""
	}

	return addr
}

func (a *Address) Bytes() []byte {
	return a.bytes
}

func (a *Address) Address() []byte {
	return a.address
}

func FromAddress(chain string, address []byte) *Address {
	return &Address{chain: chain, address: address}
}

func FromBech32UnSafe(chain, address string) *Address {
	addr, _ := FromBech32(chain, address)
	return addr
}

func FromBech32(chain, address string) (*Address, error) {
	prefix := GetAddressPrefixByChain(chain)
	if prefix == "" {
		return nil, fmt.Errorf("not supported chain: %s", chain)
	}

	hrp, bz, err := bech32.DecodeAndConvert(address)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid Bech32 prefix; expected %s, got %s", prefix, hrp)
	}

	addr := &Address{chain: chain, address: bz}
	return addr, nil
}

const (
	PubKeySize             = 33
	PubKeyAnySize          = 35
	pubKeyUncompressedSize = 65
)

func FromPublicKeyUnSafe(chain string, pubkey []byte) *Address {
	acc, _ := FromPublicKey(chain, pubkey)
	return acc
}

func FromPublicKey(chain string, pubkey []byte) (*Address, error) {
	pubKey, err := NewPubKey(pubkey)
	if err != nil {
		return nil, err
	}
	return &Address{chain: chain, bytes: pubKey.Bytes(), address: pubKey.Address().Bytes()}, nil
}

// NewPubKey
// https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-028-public-key-addresses.md
func NewPubKey(pubkey []byte) (cryptotypes.PubKey, error) {
	var key []byte
	switch len(pubkey) {
	case PubKeySize:
		// PubKeySize
		// public key
		// [public key...]

		key = pubkey
	case PubKeyAnySize:
		// PubKeyAnySize
		// protobuf any public key
		// [10] [33] [public key...]

		if pubkey[0] != byte(10) {
			return nil, fmt.Errorf("first byte is incorrect got: (%v) expected: (%v)", pubkey[0], byte(10))
		}

		if pubkey[1] != byte(PubKeySize) {
			return nil, fmt.Errorf("second byte is incorrect got: (%v) expected: (%v)", pubkey[1], byte(PubKeySize))
		}

		key = pubkey[2:]
	case pubKeyUncompressedSize:
		// pubKeyUncompressedSize
		// public key
		// compress to 33 bytes

		pubKeyParsed, err := btcec.ParsePubKey(pubkey)
		if err != nil {
			return nil, err
		}
		key = pubKeyParsed.SerializeCompressed()
	default:
		return nil, fmt.Errorf("length of pubkey is incorrect got: (%d) expected: native (%d) or protobuf any (:%d)", len(pubkey), PubKeySize, PubKeyAnySize)
	}

	return &secp256k1.PubKey{
		Key: key,
	}, nil
}
