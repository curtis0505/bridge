package types

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	ethawskmssigner "github.com/welthee/go-ethereum-aws-kms-tx-signer/v2"
	"io/ioutil"
	"math/big"
)

type Account struct {
	TransactOpts *bind.TransactOpts
	KMS          *KMSInstance
	PrivateKey   *ecdsa.PrivateKey

	Address string
	Type    AccountType
}

const (
	EOANPTReward   = "EOA-NPT-REWARD"
	EOAKLAYReward  = "EOA-KLAY-REWARD"
	EOAEventReward = "EOA-EVENT-REWARD"
)

func NewAccount(accountInfo AccountConfig, chainId *big.Int) (*Account, error) {
	account := new(Account)
	var err error

	switch accountInfo.Type {
	case PrivateKey:
		if accountInfo.PrivateKeyInfo.PrivateKey == "" {
			return nil, fmt.Errorf("need to enter the accountInfo.PrivateKeyInfo data in the config file")
		}

		account, err = NewAccountFromPK(accountInfo.PrivateKeyInfo.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("NewAccountFromPK: %w", err)
		}

	case KeyStore:
		if accountInfo.KeystoreInfo.FilePath == "" {
			return nil, fmt.Errorf("account.KeystoreInfo filepath not specified. the service account will be used")
		}
		password := accountInfo.KeystoreInfo.GetPassword()
		if password == "" {
			return nil, errors.New("failed to get password. check the config file")
		}
		account, err = NewAccountFromJsonFile(accountInfo.KeystoreInfo.FilePath, password)
		if err != nil {
			return nil, fmt.Errorf("NewAccountFromJsonFile: %w", err)
		}
	case KMS:
		account, err = NewAccountFromKms(accountInfo, chainId)
		if err != nil {
			return nil, err
		}

	case KeyStoreECS:
		password := accountInfo.KeystoreECS.GetSecretValue()
		if password == "" {
			return nil, fmt.Errorf("failed to get password. check the config file")
		}

		account, err = NewAccountFromJsonFile(accountInfo.KeystoreECS.FilePath, password)
		if err != nil {
			return nil, fmt.Errorf("NewAccountFromJsonFile: %w", err)
		}
	default:
		panic("invalid config data - accountInfo.type in (PRIVATE_KEY, KEYSTORE, KMS), config - " + accountInfo.Type)
	}

	return account, nil
}

func NewAccountFromPK(pk string) (*Account, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	account := Account{
		PrivateKey: privateKey,
		Address:    address.String(),
		Type:       PrivateKey,
	}
	return &account, nil
}

func NewAccountFromJsonFile(jsonFile, password string) (*Account, error) {
	jsonBytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	return newAccountFromJson(jsonBytes, password)
}

func NewAccountFromJsonString(jsonString, password string) (*Account, error) {
	return newAccountFromJson([]byte(jsonString), password)
}

func newAccountFromJson(json []byte, password string) (*Account, error) {
	if key, err := keystore.DecryptKey(json, password); err != nil {
		return nil, err
	} else {
		privateKey := key.PrivateKey
		address := crypto.PubkeyToAddress(privateKey.PublicKey)

		account := Account{
			PrivateKey: privateKey,
			Address:    address.String(),
			Type:       KeyStore,
		}
		return &account, nil
	}
}

func NewAccountFromKms(accountInfo AccountConfig, chainId *big.Int) (*Account, error) {
	kms := GetKMS(accountInfo)
	if kms == nil {
		return nil, errors.New("get kms error")
	}
	transactOpts, err := ethawskmssigner.NewAwsKmsTransactorWithChainIDCtx(context.Background(), kms.Client, kms.KeyId, chainId)
	if err != nil {
		return nil, err
	}

	account := Account{
		KMS:          kms,
		TransactOpts: transactOpts,
		Address:      transactOpts.From.String(),
		Type:         KMS,
	}

	return &account, nil
}

func (a *Account) Sign(tx *Transaction, chainId *big.Int) (*Transaction, error) {
	if a.Type == KMS {
		signedTx, err := tx.SignTxKms(a.TransactOpts, chainId)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	} else {
		signedTx, err := tx.SignTx(a, chainId)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}
}

func (a *Account) ECDSA() *ecdsa.PrivateKey {
	return a.PrivateKey
}

func (a *Account) Secp256k1() *secp256k1.PrivKey {
	return &secp256k1.PrivKey{
		Key: a.PrivateKey.D.Bytes(),
	}
}

var (
	secp256k1N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
)

// Deprecated: Secp256k1Legacy
// 잘못된 키 생성으로 인해 사용하면 안됩니다..
// Must <<< Remove
func (a *Account) Secp256k1Legacy() *secp256k1.PrivKey {
	priv := new(big.Int).Add(a.PrivateKey.D, big.NewInt(1))
	priv.Mod(priv, secp256k1N)
	return &secp256k1.PrivKey{
		Key: priv.Bytes(),
	}
}

func createKeystoreV3(password string) ([]byte, error) {
	d, err := ioutil.TempDir("", "ks")
	if err != nil {
		return []byte{}, err
	}

	ks := keystore.NewKeyStore(d, 65536, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return []byte{}, err
	}

	keyJson, err := ks.Export(account, password, password)
	if err != nil {
		return []byte{}, err
	}

	return keyJson, nil
}

func CreateKeystore(password string) (string, *Account, error) {
	ks, err := createKeystoreV3(password)
	if err != nil {
		return "", nil, err
	}

	account, err := newAccountFromJson(ks, password)
	if err != nil {
		return "", nil, err
	}

	return string(ks), account, nil
}

func GetCosmosPubKeyFromString(pubKeyString string) *secp256k1.PubKey {
	pubKeyBytes, err := hex.DecodeString(pubKeyString)
	if err != nil {
		return nil
	}

	return &secp256k1.PubKey{Key: pubKeyBytes}
}
