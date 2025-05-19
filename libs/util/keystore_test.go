package util

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/klaytn/klaytn/accounts/keystore"
	"github.com/klaytn/klaytn/crypto"
)

var passphrase = ""

func TestKeystoreGenerate(t *testing.T) {

	passphrase = RandPassword(20)
	println("passphrase:", passphrase)

	if privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader); err != nil { //privateKey를 생성.
		t.Fail()
	} else {
		//keyJson을 생성.
		key := &keystore.KeyV4{
			//Id:          uuid.NewRandom(),
			Address:     crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
			PrivateKeys: [][]*ecdsa.PrivateKey{{privateKeyECDSA}},
		}
		if keyjson, err := keystore.EncryptKey(key, passphrase, keystore.StandardScryptN, keystore.StandardScryptP); err != nil {
			t.Fail()
		} else {
			println(string(keyjson))
		}
	}
}

func TestKeystoreDecrypt(t *testing.T) {
	const passphrase = "11]QeuD!%E(k|fuPXAkf" // test only
	const keystorePath = "./test_keystore2.json"

	addressParsed, _, err := readAccount(keystorePath)
	if err == nil {
		println("addressParsed:", addressParsed.Hex())
	} else {
		println("err:", err)
		t.Fail()
	}

	if _, err := decryptKey(keystorePath, passphrase); err != nil {
		t.Fail()
	} else {
		println("privateKey decrypted:")
		//println(pk.PublicKey.)
	}

}

func TestKeystoreString(t *testing.T) {
	//NOT WORKING for web3signer const ks = "{\"crypto\":{\"kdf\":{\"function\":\"scrypt\",\"params\":{\"dklen\":32,\"n\":262144,\"r\":8,\"p\":1,\"salt\":\"2a7ca18634ac28b93f29d92ce422009ba183edff4c5b33f170169031b756fe13\"},\"message\":\"\"},\"checksum\":{\"function\":\"sha256\",\"params\":{},\"message\":\"e3fe72bb3bf66beb34866efb4f7a97a58fd0751993d764400302e86746d66d27\"},\"cipher\":{\"function\":\"aes-128-ctr\",\"params\":{\"iv\":\"84ff32315e3e2ce63a5b096888b51a6c\"},\"message\":\"d3c95b5881ebac5fb626660aeaef823d914e645b5718d1c2a7d477a751965d0d\"}},\"description\":\"\",\"pubkey\":\"a478b6ece186ebc82bbecb6307a575f5e1381ef685e90d91bebc45950c159ca5a38ac4edc85f2d34f56098ab4659263d\",\"path\":\"m/12381/3600/0/0/0\",\"uuid\":\"7fcd5a5d-21b3-4292-b387-c455ac8c8ae9\",\"version\":4}"
	const ks = "{\\\"crypto\\\":{\\\"kdf\\\":{\\\"function\\\":\\\"scrypt\\\",\\\"params\\\":{\\\"dklen\\\":32,\\\"n\\\":262144,\\\"r\\\":8,\\\"p\\\":1,\\\"salt\\\":\\\"2a7ca18634ac28b93f29d92ce422009ba183edff4c5b33f170169031b756fe13\\\"},\\\"message\\\":\\\"\\\"},\\\"checksum\\\":{\\\"function\\\":\\\"sha256\\\",\\\"params\\\":{},\\\"message\\\":\\\"e3fe72bb3bf66beb34866efb4f7a97a58fd0751993d764400302e86746d66d27\\\"},\\\"cipher\\\":{\\\"function\\\":\\\"aes-128-ctr\\\",\\\"params\\\":{\\\"iv\\\":\\\"84ff32315e3e2ce63a5b096888b51a6c\\\"},\\\"message\\\":\\\"d3c95b5881ebac5fb626660aeaef823d914e645b5718d1c2a7d477a751965d0d\\\"}},\\\"description\\\":\\\"\\\",\\\"pubkey\\\":\\\"a478b6ece186ebc82bbecb6307a575f5e1381ef685e90d91bebc45950c159ca5a38ac4edc85f2d34f56098ab4659263d\\\",\\\"path\\\":\\\"m/12381/3600/0/0/0\\\",\\\"uuid\\\":\\\"7fcd5a5d-21b3-4292-b387-c455ac8c8ae9\\\",\\\"version\\\":4}"
	url := "http://localhost:9000/eth/v1/keystores"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"keystores":["%s"],"passwords":["%s"]}`, ks, "spdhvls123!"))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func web3signerImport(ks string) {

	//url := "http://localhost:9000/eth/v1/keystores"
	url := "https://web3signer.dq.neopin.nwz.cloud/eth/v1/keystores"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"keystores":["%s"],"passwords":["%s"]}`, ks, "spdhvls123!"))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func TestWeb3SignerImportKeystore(t *testing.T) {

	var ks string
	directory := "/Users/ntesla/Dev/ethereum/20240401_keys_holesky/holesky0_dq/validator_keys"

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(directory, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				continue
			}
			ks = strings.ReplaceAll(string(content), "\"", "\\\"")

			fmt.Println(ks)
		}
		if i > 1 {
			//	break
		}
		web3signerImport(ks)
	}

	//NOT WORKING for web3signer  ks = "{\"crypto\":{\"kdf\":{\"function\":\"scrypt\",\"params\":{\"dklen\":32,\"n\":262144,\"r\":8,\"p\":1,\"salt\":\"2a7ca18634ac28b93f29d92ce422009ba183edff4c5b33f170169031b756fe13\"},\"message\":\"\"},\"checksum\":{\"function\":\"sha256\",\"params\":{},\"message\":\"e3fe72bb3bf66beb34866efb4f7a97a58fd0751993d764400302e86746d66d27\"},\"cipher\":{\"function\":\"aes-128-ctr\",\"params\":{\"iv\":\"84ff32315e3e2ce63a5b096888b51a6c\"},\"message\":\"d3c95b5881ebac5fb626660aeaef823d914e645b5718d1c2a7d477a751965d0d\"}},\"description\":\"\",\"pubkey\":\"a478b6ece186ebc82bbecb6307a575f5e1381ef685e90d91bebc45950c159ca5a38ac4edc85f2d34f56098ab4659263d\",\"path\":\"m/12381/3600/0/0/0\",\"uuid\":\"7fcd5a5d-21b3-4292-b387-c455ac8c8ae9\",\"version\":4}"
	//WORKING for web3signer      ks = "{\\\"crypto\\\":{\\\"kdf\\\":{\\\"function\\\":\\\"scrypt\\\",\\\"params\\\":{\\\"dklen\\\":32,\\\"n\\\":262144,\\\"r\\\":8,\\\"p\\\":1,\\\"salt\\\":\\\"2a7ca18634ac28b93f29d92ce422009ba183edff4c5b33f170169031b756fe13\\\"},\\\"message\\\":\\\"\\\"},\\\"checksum\\\":{\\\"function\\\":\\\"sha256\\\",\\\"params\\\":{},\\\"message\\\":\\\"e3fe72bb3bf66beb34866efb4f7a97a58fd0751993d764400302e86746d66d27\\\"},\\\"cipher\\\":{\\\"function\\\":\\\"aes-128-ctr\\\",\\\"params\\\":{\\\"iv\\\":\\\"84ff32315e3e2ce63a5b096888b51a6c\\\"},\\\"message\\\":\\\"d3c95b5881ebac5fb626660aeaef823d914e645b5718d1c2a7d477a751965d0d\\\"}},\\\"description\\\":\\\"\\\",\\\"pubkey\\\":\\\"a478b6ece186ebc82bbecb6307a575f5e1381ef685e90d91bebc45950c159ca5a38ac4edc85f2d34f56098ab4659263d\\\",\\\"path\\\":\\\"m/12381/3600/0/0/0\\\",\\\"uuid\\\":\\\"7fcd5a5d-21b3-4292-b387-c455ac8c8ae9\\\",\\\"version\\\":4}"
	//fmt.Println(ks)
	//return

}
