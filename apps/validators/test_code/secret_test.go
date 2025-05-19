package test_code

// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getSecret() (string, error) {

	cfg, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}
	//
	password := cfg.Account.KeystoreInfo.GetPassword()
	return password, nil
}

func TestSecret(t *testing.T) {
	secret, err := getSecret()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("secret", secret)
}

func Test_Verification_Bech32ToHex(t *testing.T) {
	vault := "link1sqctf7u55yhu4s00d4f33nj7k2a09tjx4x0ev74u83zvpqkhvfqsvqqnpz"
	minter := "link1un8cmeumu2fw77xzc36uz88n67zgzf8zlyt66hkp9nayua7dkj8qtwyr4a"

	_, vaultDecode, err := bech32.DecodeAndConvert(vault) //hrp :link
	assert.NoError(t, err)

	_, minterDecode, err := bech32.DecodeAndConvert(minter) //hrp :link
	assert.NoError(t, err)
	t.Log("vault", hexutil.Encode(vaultDecode))
	t.Log("vault", hexutil.Encode(minterDecode))

	vmHashByte := crypto.Keccak256(
		//common.HexToAddress("FNSA", fmt.Sprintf("%s", vaultDecode)).Bytes(),
		//common.HexToAddress("FNSA", fmt.Sprintf("%s", minterDecode)).Bytes(),
		//common.HexToAddress(hexutil.Encode(vaultDecode)).Bytes(),
		//common.HexToAddress(hexutil.Encode(minterDecode)).Bytes(),
		vaultDecode, minterDecode,
	)

	fmt.Println(hexutil.Encode(vmHashByte))

	t.Log("address", cosmos.FromAddress("FNSA", vmHashByte))
}
