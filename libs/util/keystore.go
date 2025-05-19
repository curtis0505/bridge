package util

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/types"
	etheraccounts "github.com/ethereum/go-ethereum/accounts"
	etherkeystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/accounts/keystore"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/console"
)

const (
	//Parameters for password
	minLength       = 13
	minSpecialChars = 1
	minBigLetters   = 1
	minDigits       = 1
)

func MakePassPhrase(prompt string) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
		fmt.Println(fmt.Sprintf(" ⓥ %v or more characters", minLength))
		fmt.Println(fmt.Sprintf(" ⓥ %v or more special characters", minSpecialChars))
		fmt.Println(fmt.Sprintf(" ⓥ %v or more big letters", minBigLetters))
		fmt.Println(fmt.Sprintf(" ⓥ %v or more digits", minDigits))
	}
	password, err := console.Stdin.PromptPassword("Password: ")
	if err != nil {
		return "", fmt.Errorf("Failed to read password: %v", err)
	} else if err = verifyPassword(password); err != nil {
		return "", err
	}
	confirm, err := console.Stdin.PromptPassword("Repeat password: ")
	if err != nil {
		return "", fmt.Errorf("Failed to read password confirmation: %v", err)
	}
	if password != confirm {
		return "", fmt.Errorf("password do not match")
	}
	return password, nil
}

func verifyPassword(password string) error {
	if !regexp.MustCompile("[[:alpha:]]").MatchString(password) {
		return fmt.Errorf("require only ascii")
	} else if !regexp.MustCompile(fmt.Sprintf(".{%d,}", minLength)).MatchString(password) {
		return fmt.Errorf("require %v or more characters", minLength)
	} else if !regexp.MustCompile(fmt.Sprintf("[^A-Za-z0-9]{%d,}", minSpecialChars)).MatchString(password) {
		return fmt.Errorf("require %v or more special characters", minSpecialChars)
	} else if !regexp.MustCompile(fmt.Sprintf("[A-Z]{%d,}", minBigLetters)).MatchString(password) {
		return fmt.Errorf("require %v or more big letters", minBigLetters)
	} else if !regexp.MustCompile(fmt.Sprintf("[0-9]{%d,}", minDigits)).MatchString(password) {
		return fmt.Errorf("require %v or more digits", minDigits)
	} else {
		return nil
	}
}

func RandPassword(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}|[]")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func KeyFileName(keyAddr common.Address) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), hex.EncodeToString(keyAddr[:]))
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}

func FindKeystore(account common.Address, keystore string) (string, error) {
	files, err := ioutil.ReadDir(keystore)
	if err != nil { //keystore 파일들이 있는 directory안에서 모든 파일을 가져온다.
		return "", err
	} else if len(files) == 0 {
		return "", errors.New("There is no keystore file")
	}

	for _, v := range files {
		if nonKeyFile(v) == true {
			continue
		}
		if addressParsed, file, err := readAccount(path.Join(keystore, v.Name())); err != nil {
			continue
		} else if addressParsed == account {
			return file, nil
		}
	}
	return "", errors.New("failed to find keystore by given address")
}

func UnlockKeystore(account accounts.Account, ks *keystore.KeyStore, header string, trials int) error {
	addressHex := hexutil.Encode(account.Address[:])
	for trials := 0; trials < 3; trials++ {
		prompt := fmt.Sprintf("unlocking %s account %s | attempt %d/%d", header, addressHex, trials+1, 3)
		if pw, err := GetPassPhrase(prompt); err != nil {
			fmt.Println(err)
		} else {
			if err := ks.Unlock(account, pw); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(addressHex, "unlocked")
				return nil
			}
		}
	}
	return fmt.Errorf("Failed to unlock account %s", addressHex)
}

func UnlockKey(address common.Address, keystore string, header string, trials int) (*ecdsa.PrivateKey, error) {
	file, err := FindKeystore(address, keystore)
	if err != nil {
		return nil, err
	}
	for trials := 0; trials < 3; trials++ {
		prompt := fmt.Sprintf("unlocking %s account %s | attempt %d/%d", header, hexutil.Encode(address[:]), trials+1, 3)
		if pw, err := GetPassPhrase(prompt); err != nil {
			fmt.Println(err)
		} else {
			if pk, err := decryptKey(file, pw); err != nil {
				fmt.Println(err)
			} else {
				return pk, nil
			}
		}
	}
	return nil, fmt.Errorf("Failed to unlock account %s", hexutil.Encode(address[:]))
}

func GetPrivateKeyFromKeystore(account *common.Address, keystore string, password string) (*ecdsa.PrivateKey, error) {
	if files, err := ioutil.ReadDir(keystore); err != nil { //keystore 파일들이 있는 directory안에서 모든 파일을 가져온다.
		return nil, err
	} else {
		if len(files) == 0 {
			return nil, errors.New("There is no keystore file")
		}
		for _, v := range files {
			if nonKeyFile(v) == true {
				continue
			}

			if account != nil {
				addressParsed, _, err := readAccount(path.Join(keystore, v.Name()))
				if err == nil {
					if addressParsed != *account {
						continue
					}
				} else {
					continue
				}
			}

			if privateKey, err := decryptKey(filepath.Join(keystore, v.Name()), password); err != nil {
				return nil, err
			} else {
				return privateKey, nil
			}
		}
	}
	return nil, errors.New("fail to unlock key")
}

func GetPassPhrase(prompt string) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
	}

	if password, err := console.Stdin.PromptPassword("password: "); err != nil {
		return "", err
	} else {
		return password, nil
	}
}

func nonKeyFile(fi os.FileInfo) bool {
	// Skip editor backups and UNIX-style hidden files.
	if strings.HasSuffix(fi.Name(), "~") || strings.HasPrefix(fi.Name(), ".") {
		return true
	}
	// Skip misc special files, directories (yes, symlinks too).
	if fi.IsDir() || fi.Mode()&os.ModeType != 0 {
		return true
	}
	return false
}

func decryptKey(jsonFile string, password string) (*ecdsa.PrivateKey, error) {
	if json, err := ioutil.ReadFile(jsonFile); err != nil {
		log.Error("signer error", "func", "decryptKey", "msg", "ioutil.ReadFile", "jsonFile", jsonFile)
		return nil, err
	} else {
		if key, err := keystore.DecryptKey(json, password); err != nil {
			return nil, err
		} else {
			return key.GetPrivateKey(), nil
		}
	}
}

func readAccount(path string) (common.Address, string, error) {
	var (
		buf = new(bufio.Reader)
		key struct {
			Address string `json:"address"`
		}
	)

	fd, err := os.Open(path)
	if err != nil {
		return common.Address{}, "", err
	}
	defer fd.Close()
	buf.Reset(fd)
	// Parse the address.
	key.Address = ""
	err = json.NewDecoder(buf).Decode(&key)
	addr := common.HexToAddress(key.Address)
	switch {
	case err != nil:
		return common.Address{}, "", err
	case (addr == common.Address{}):
		return common.Address{}, "", fmt.Errorf("Failed to decode keystore key, missing or zero address")
	default:
		return addr, fd.Name(), nil
	}
}

func SigHash(messageHash []byte, key *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(messageHash, key)
	if err != nil {
		return []byte{}, err
	}
	sig[64] += 27
	return sig, nil
}

func GetAccountFromKeystore(keyStore *etherkeystore.KeyStore, accountConfig types.AccountConfig) (*etheraccounts.Account, error) {
	ksAccounts := keyStore.Accounts()
	if len(ksAccounts) == 0 {
		return nil, fmt.Errorf("no accounts found in the keystore")
	}
	var account etheraccounts.Account
	for _, currentAccount := range ksAccounts {
		if strings.ToLower(currentAccount.Address.String()) == strings.ToLower(accountConfig.Address) {
			account = currentAccount
			continue
		}
	}
	zeroAddress := common.Address{}
	if account.Address.String() == zeroAddress.String() {
		return nil, fmt.Errorf("keystore: not found account")
	}
	return &account, nil
}
