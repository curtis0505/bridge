package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	kaiacrypto "github.com/kaiachain/kaia/crypto"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func Test_RecoverSigner(t *testing.T) {
	t.Run("Ethereum", func(t *testing.T) {
		privateKeyString := "83436c52cc96073eb27af24e245161a179ef93dfcff7fc7a40520cd360d0242a"

		// Generate a new private key
		privateKey, err := crypto.HexToECDSA(privateKeyString)
		assert.NoError(t, err)

		// Get the public address
		publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
		fmt.Printf("Public Address: %s\n", publicAddress)

		// Example message
		message := fmt.Sprintf("%s|%d", publicAddress, time.Now().Unix())
		fmt.Println("message:", message)

		// Sign the message
		signature, err := SignMessage(message, privateKey)
		assert.NoError(t, err)
		fmt.Printf("Signature: %s\n", signature)

		// Recover the signer address
		signer, err := RecoverSigner(message, signature)
		assert.NoError(t, err)

		assert.Equal(t, signer, publicAddress)
	})

	t.Run("KaiaWallet", func(t *testing.T) {
		privateKeyString := "1fbc1a8cc35bf09ac652579c18bad487b26e507405c1807db8dec7919601c7f8"

		// Generate a new private key
		privateKey, err := kaiacrypto.HexToECDSA(privateKeyString)
		if err != nil {
			t.Fatalf("Failed to generate private key: %v", err)
		}

		// Get the public address
		publicAddress := kaiacrypto.PubkeyToAddress(privateKey.PublicKey).Hex()
		if !strings.EqualFold(publicAddress, "0x2f2bb80C298bD35d4Ab30F79Ba11260174915580") {
			t.Fatalf("Public address is not equal to expected address: %s", publicAddress)
		}

		// Example message
		//message := fmt.Sprintf("%s|%d", publicAddress, time.Now().Unix())
		message := "0x2f2bb80C298bD35d4Ab30F79Ba11260174915580|1726725296"

		// Sign the message
		signature, err := SignKaiaMessage(message, privateKey)
		assert.NoError(t, err)
		assert.Equal(t, signature, "0xae535c96be4ea10d831de5b77c938235c8ca6d59ff4a7f3160ccc0d24b09dd660cd3445656a917e1bf5eff4b4b953c7b991d6ff2bb71489349b0a3993bf064d51b")

		// Recover the signer address
		signer, err := RecoverKaiaSigner(message, signature)
		assert.NoError(t, err)

		assert.Equal(t, signer, publicAddress)
	})
}
