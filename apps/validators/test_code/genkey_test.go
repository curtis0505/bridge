package test_code

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/libs/common"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type GenerateKey struct {
	Config *conf.Config
}

func NewGenerateKey(config *conf.Config) (*GenerateKey, error) {
	generateKey := GenerateKey{
		Config: config,
	}

	return &generateKey, nil
}

type GenerateKeyInfo struct {
	Keystore  string            `json:"keystore"`
	PublicKey map[string]string `json:"publicKey"`
}

func GeneKey(accountConf commontypes.AccountConfig) (GenerateKeyInfo, error) {
	var password string
	switch accountConf.Type {
	case commontypes.KeyStoreECS:
		password = accountConf.KeystoreECS.GetSecretValue()
	case commontypes.KeyStore:
		password = accountConf.KeystoreInfo.Password
	default:
		return GenerateKeyInfo{}, errors.New("")
	}

	if len(password) == 0 {
		return GenerateKeyInfo{}, errors.New("password")
	}

	//이미 게정 생성된 경우
	account, err := commontypes.NewAccount(accountConf, nil)
	if err == nil {
		return GenerateKeyInfo{
			PublicKey: map[string]string{
				"ecdsa":     fmt.Sprintf("%v", account.PrivateKey.PublicKey),
				"secp256k1": account.Secp256k1().PubKey().String(),
			},
		}, nil
	}

	// 계정 생성 및 pubkey를 추출
	ks, account, err := commontypes.CreateKeystore(password)
	if err != nil {
		return GenerateKeyInfo{}, err
	}
	return GenerateKeyInfo{
		Keystore: ks,
		PublicKey: map[string]string{
			"ecdsa":     fmt.Sprintf("%v", account.PrivateKey.PublicKey),
			"secp256k1": account.Secp256k1().PubKey().String(),
		},
	}, nil
}

func Test_GenKey(t *testing.T) {
	cfg, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	key, _ := GeneKey(cfg.Account)
	fmt.Println(key)
}

func TestMultiSigAddr(t *testing.T) {
	chain := commontypes.ChainFNSA
	pubKeys := []cryptotypes.PubKey{
		createCosmosPubKeyFromString("030E2D06C41677C9DA5FDD41A61281AFCF17AD8D58072BCFDC1D392C1CE1C40AD9"),
		createCosmosPubKeyFromString("033A8A30F9B821CEF322508CC4BCE77FFAFAB447ACF0EC8A2C0C9567D34B928751"),
		createCosmosPubKeyFromString("038A85353B44EB2CC630A4AC3F6FA27AF2813CD63895FBE30007A1C9CA221C3144"),
	}

	sort.Slice(pubKeys, func(i, j int) bool {
		return bytes.Compare(pubKeys[i].Address(), pubKeys[j].Address()) < 0
	})

	multiSigPubKey := kmultisig.NewLegacyAminoPubKey(1, pubKeys)
	t.Log("multiSigAddr", cosmoscommon.FromAddress(chain, multiSigPubKey.Address().Bytes()).String())
}

func createCosmosPubKeyFromString(pubKeyString string) *secp256k1.PubKey {
	pubKeyBytes, err := hex.DecodeString(pubKeyString)
	if err != nil {
		return nil
	}

	// Cosmos SDK에서는 공개 키를 바이트로 직접 사용
	return &secp256k1.PubKey{Key: pubKeyBytes}
}

func Test_HexToBech32(t *testing.T) {
	hex := "0xfa0734e376011ea8b9286fc9625ad5a97d750fe6"
	bech32ToHexDecode, err := bech32.ConvertAndEncode("link", common.HexToAddress("ETH", hex).Bytes()) //hrp :link
	assert.NoError(t, err)

	t.Log("address bytes", common.BytesToAddress("ETH", common.HexToAddress("ETH", hex).Bytes()))
	t.Log("address", bech32ToHexDecode)
}
