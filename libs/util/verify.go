package util

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	kaiahexutil "github.com/kaiachain/kaia/common/hexutil"
	kaiacrypto "github.com/kaiachain/kaia/crypto"
	"strconv"
	"strings"
	"time"
)

func RecoverSigner(message, signature string) (string, error) {
	// Decode the signature
	sig, err := hexutil.Decode(signature)
	if err != nil {
		return "", err
	}
	if sig[64] != 27 && sig[64] != 28 {
		return "", fmt.Errorf("invalid signature 'v' value : %v", sig[64])
	}
	sig[64] -= 27 // Transform the 'v' value to 0 or 1

	// Hash the message
	hash := HashMessage(message)

	// Recover the public key
	pubKey, err := ethercrypto.SigToPub(hash, sig)
	if err != nil {
		return "", err
	}

	// Get the address from the public key
	address := ethercrypto.PubkeyToAddress(*pubKey).Hex()
	return address, nil
}

func HashMessage(message string) []byte {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	prefixedMessage := []byte(prefix + message)
	hash := ethercrypto.Keccak256Hash(prefixedMessage)
	return hash[:]
}

func RecoverKaiaSigner(message, signature string) (string, error) {
	// Decode the signature
	sig, err := kaiahexutil.Decode(signature)
	if err != nil {
		return "", err
	}
	if sig[64] != 27 && sig[64] != 28 {
		return "", fmt.Errorf("invalid signature 'v' value : %v", sig[64])
	}
	sig[64] -= 27 // Transform the 'v' value to 0 or 1

	// Hash the message
	hash := hashKaiaMessage(message)
	pubKey, err := kaiacrypto.SigToPub(hash, sig)
	if err != nil {
		return "", err
	}

	// Get the address from the public key
	address := kaiacrypto.PubkeyToAddress(*pubKey).Hex()
	return address, nil
}

func hashKaiaMessage(message string) []byte {
	prefix := fmt.Sprintf("\x19Klaytn Signed Message:\n%d", len(message))
	prefixedMessage := []byte(prefix + message)
	hash := kaiacrypto.Keccak256(prefixedMessage)
	return hash[:]
}

func SignMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := HashMessage(message)
	signature, err := ethercrypto.Sign(hash, privateKey)
	if err != nil {
		return "", err
	}
	if signature[ethercrypto.RecoveryIDOffset] <= 1 {
		signature[ethercrypto.RecoveryIDOffset] += 27 // for ethereum
	}
	return hexutil.Encode(signature), nil
}

func SignKaiaMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := hashKaiaMessage(message)
	signature, err := kaiacrypto.Sign(hash, privateKey)
	if err != nil {
		return "", err
	}
	if signature[kaiacrypto.RecoveryIDOffset] <= 1 {
		signature[kaiacrypto.RecoveryIDOffset] += 27
	}
	return kaiahexutil.Encode(signature), nil
}

func ValidateAddress(message, signature string) (string, error) {
	messages := strings.Split(message, "|")
	if len(messages) != 2 {
		return "", errors.New("len(messages) != 2")
	}
	address := messages[0]
	timestamp, err := strconv.ParseInt(messages[1], 10, 64)
	if err != nil {
		return "", fmt.Errorf("strconv.ParseInt: %w", err)
	}
	inputTime := time.Unix(timestamp, 0)
	if time.Now().Sub(inputTime).Abs() > time.Minute*5 {
		return "", fmt.Errorf("exceed 5 minutes, inputTime: %d", timestamp)
	}

	signer, err := RecoverSigner(message, signature)
	if err != nil {
		logger.Warn("ValidateAddress", logger.BuildLogInput().WithError(err).WithData("message", message, "signature", signature))
	}

	logger.Trace("ValidateAddress", logger.BuildLogInput().WithData("signer", signer))

	if !strings.EqualFold(address, signer) {
		signer, err = RecoverKaiaSigner(message, signature)
		if err != nil {
			return "", fmt.Errorf("util.RecoverSigner: %w", err)
		}

		logger.Trace("ValidateAddress", logger.BuildLogInput().WithData("signer", signer))

		if !strings.EqualFold(address, signer) {
			return "", fmt.Errorf("not matched address(input: %s, signer: %s)", address, signer)
		}
	}

	return signer, nil
}
