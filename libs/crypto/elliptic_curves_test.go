package crypto

import (
	"crypto/elliptic"
	"crypto/sha256"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/types"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEllipticCurves(t *testing.T) {
	account, _ := types.NewAccountFromPK("83c13018679900c69be911758784782f693215f754c995fc17e304c9fbf90afb")

	signMessage := []byte("test")
	secp256k1Sig, _ := account.Secp256k1().Sign(signMessage)

	hashed := sha256.Sum256(signMessage)
	ecdsaSig, _ := ethercrypto.Sign(hashed[:], account.ECDSA())

	t.Log(cosmoscommon.FromPublicKey("FNSA", account.Secp256k1().PubKey().Bytes()))
	t.Log(cosmoscommon.FromPublicKey("FNSA", elliptic.Marshal(account.ECDSA().PublicKey, account.ECDSA().PublicKey.X, account.ECDSA().PublicKey.Y)))
	// COSMOS PubicKey: 0x13 + PublicKey
	// ETH PublicKey: PublicKey
	assert.Equal(t, account.Secp256k1().PubKey().Bytes()[1:], account.ECDSA().PublicKey.X.Bytes()) // slices 0x13
	// COSMOS secp256k1 returns [ R || S ]
	// ETHEREUM secp256k1 returns  [ R || S || V (0 or 1) ]
	assert.Equal(t, secp256k1Sig, ecdsaSig[:len(ecdsaSig)-1]) // slices V
}
