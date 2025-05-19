package types

import (
	"bytes"
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestKMS(t *testing.T) {
	// aws kms get-public-key response
	publicKeyResponse := `
{
    "KeyId": "arn:aws:kms:ap-northeast-2:718039325254:key/fb022605-f5c4-456d-9166-bcfa34b7c9da",
    "PublicKey": "MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAE89qu2layeXZQFTkBxwFs6jD7ft8PLnOKvzVgLsh6bzIkTq1iv+xiD/rVb12N2KQr55MsRfazfTlJNKkAsd9urg==",
    "CustomerMasterKeySpec": "ECC_SECG_P256K1",
    "KeySpec": "ECC_SECG_P256K1",
    "KeyUsage": "SIGN_VERIFY",
    "SigningAlgorithms": [
        "ECDSA_SHA_256"
    ]
}`

	var publicKeyOutput kms.GetPublicKeyOutput
	assert.NoError(t, json.Unmarshal([]byte(publicKeyResponse), &publicKeyOutput))

	var publicKey ANS1PublicKey
	_, err := asn1.Unmarshal(publicKeyOutput.PublicKey, &publicKey)
	assert.NoError(t, err)
	assert.Equal(t, len(publicKey.PublicKey.Bytes), 65)

	// https://git.dev.neopin.io/neopindeploys/deploys/-/blob/dq/bridge-validator/conf/config-4.toml
	evmAddress := common.BytesToAddress(crypto.Keccak256(publicKey.PublicKey.Bytes[1:])[12:])
	assert.Equal(t, evmAddress.String(), "0xc1b5d1393f23Ec47277D554ce6D202842eD2Cb2c")

	cosmosAddress, err := cosmoscommon.FromPublicKey(ChainFNSA, publicKey.PublicKey.Bytes)
	assert.NoError(t, err)

	t.Log("address", evmAddress, cosmosAddress)

	// aws kms sign response
	signResponse := `
{
    "KeyId": "arn:aws:kms:ap-northeast-2:718039325254:key/fb022605-f5c4-456d-9166-bcfa34b7c9da",
    "Signature": "MEYCIQC6nYTZQveJ2remK3KGt5KPuNkb8cDR00tirhQHT0dSqgIhAN6ZLho8iYVN8BdikuTyRGbeniM0FEC+0omq8Hygg28v",
    "SigningAlgorithm": "ECDSA_SHA_256"
}
`
	var signOutput kms.SignOutput
	assert.NoError(t, json.Unmarshal([]byte(signResponse), &signOutput))

	var signature ANS1Signature
	_, err = asn1.Unmarshal(signOutput.Signature, &signature)

	sBytes := signature.S.Bytes
	rBytes := signature.R.Bytes
	adjustSignatureLength := func(buffer []byte) []byte {
		buffer = bytes.TrimLeft(buffer, "\x00")
		for len(buffer) < 32 {
			zeroBuf := []byte{0}
			buffer = append(zeroBuf, buffer...)
		}
		return buffer
	}

	sBigInt := new(big.Int).SetBytes(sBytes)
	if sBigInt.Cmp(secp256k1HalfN) > 0 {
		sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
	}

	rsSignature := append(adjustSignatureLength(rBytes), adjustSignatureLength(sBytes)...)

	// sign to bytes from cosmos builder
	hexSignToBytes := `7b226163636f756e745f6e756d626572223a2231313730222c22636861696e5f6964223a2266696e7363686961222c22666565223a7b22616d6f756e74223a5b7b22616d6f756e74223a2235303030222c2264656e6f6d223a22636f6e79227d5d2c22676173223a2232303030303030227d2c226d656d6f223a22222c226d736773223a5b7b2274797065223a22636f736d6f732d73646b2f4d736753656e64222c2276616c7565223a7b22616d6f756e74223a5b7b22616d6f756e74223a2231303030222c2264656e6f6d223a22636f6e79227d5d2c2266726f6d5f61646472657373223a226c696e6b316465783370323438657a32657a6e6c73617872376565793939767277367272613039716d357a222c22746f5f61646472657373223a226c696e6b31383837747939776d79777a7673706c346439716b766a687a617230657a68636c76713030686a227d7d5d2c2273657175656e6365223a223134227d`
	rawTx, _ := hex.DecodeString(hexSignToBytes)
	assert.NoError(t, err)

	pubKey := secp256k1.PubKey{
		Key: cosmoscommon.FromPublicKeyUnSafe(ChainFNSA, publicKey.PublicKey.Bytes).Bytes(),
	}
	assert.NoError(t, err)
	assert.Equal(t, pubKey.VerifySignature(rawTx, rsSignature), true)
}
