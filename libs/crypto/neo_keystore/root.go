package neo_keystore

import (
	"bytes"
	"crypto/ecdsa"
	crand "crypto/rand"
	"fmt"
	"github.com/curtis0505/bridge/libs/crypto/neo_aes"
	"github.com/curtis0505/bridge/libs/crypto/prompt"
	"io/ioutil"
	"math/big"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/codahale/sss"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/klaytn/klaytn/accounts"
	"github.com/klaytn/klaytn/accounts/keystore"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/crypto"
)

type NeopinKeystore struct {
	KeystoreDir    string
	ks             *keystore.KeyStore
	keepLock       bool
	keyhashDB      *leveldb.DB
	ReservedPrompt *prompt.ReservedPromptInput
}

const (
	KeyHashDB = "keyhash"
)

func NewNeopinKeystore(keystoreDir string, keepLock bool, useKeyHash bool, reservedPromptInputs ...string) (*NeopinKeystore, error) {
	reservedPrompt := prompt.NewReservedPromptInput(reservedPromptInputs...)

	var err error
	if keystoreDir == "" {
	enter_keystore:
		if keystoreDir, err = reservedPrompt.PromptInput(fmt.Sprintf("Enter the path where the keystore will be stored. <keystore>: ")); err != nil {
			return nil, err
		} else if keystoreDir == "" {
			goto enter_keystore
		}
	}

	if keystoreDir[0] == byte('~') {
		keystoreDir = path.Join(HomeDir(), keystoreDir[1:])
	}

	var keyhashDB *leveldb.DB
	if useKeyHash == true {
		if db, err := leveldb.OpenFile(path.Join(keystoreDir, KeyHashDB), nil); err != nil {
			return nil, err
		} else {
			keyhashDB = db
		}
	}

	r := &NeopinKeystore{
		KeystoreDir:    keystoreDir,
		ks:             keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP),
		keepLock:       keepLock,
		keyhashDB:      keyhashDB,
		ReservedPrompt: reservedPrompt,
	}
	return r, nil
}

func (p *NeopinKeystore) Close() {
	if p.keyhashDB != nil {
		p.keyhashDB.Close()
	}
}

func (p *NeopinKeystore) MakePassphrase() (string, error) {
enter_split:
	if num, err := p.ReservedPrompt.PromptInput(fmt.Sprintf("How many share of password? if less than 2, a normal password is generated. <splitN> : ")); err != nil {
		return "", err
	} else if num == "" {
		goto enter_split
	} else if splitN, err := strconv.Atoi(num); err != nil {
		return "", err
	} else if splitN >= 10 || splitN < 0 {
		fmt.Println("splitN is a positive integer less than 10.")
		goto enter_split
	} else {
		var (
			passphrase string
			err        error
		)
		switch {
		case splitN > 1:
			passphrase, err = p.sharedPassphrase(splitN)
		default:
			passphrase, err = p.normalPassphrase()
		}

		if err != nil {
			return "", err
		}

		return passphrase, nil
	}
}

func (p *NeopinKeystore) GenerateKey() (accounts.Account, error) {
	if passphrase, err := p.MakePassphrase(); err != nil {
		return accounts.Account{}, err
	} else {
		return p.newAccount(passphrase)
	}
}

func (p *NeopinKeystore) UpdatePassphrase(address common.Address, name string, threshold int) error {
	if account, err := p.Find(address); err != nil {
		return err
	} else if oldPassphrase, _, err := p.getPassphrase(address, name, threshold); err != nil {
		return err
	} else if err := p.ks.Unlock(account, oldPassphrase); err != nil {
		return err
	} else if err := p.ks.Lock(address); err != nil {
		return err
	} else {
		fmt.Println("From now on, we will create a new passphrase.")
		if newPassphrase, err := p.MakePassphrase(); err != nil {
			return err
		} else if err := p.ks.Update(account, oldPassphrase, newPassphrase); err != nil {
			return err
		}
	}
	return nil
}

func (p *NeopinKeystore) GenerateKeyEncrypted() error {
	if passphrase, err := p.MakePassphrase(); err != nil {
		return err
	} else {
		//하나의 passphrase로 키를 몇개를 만들것인가?
		keyNum, err := func() (int, error) {
			fmt.Println(fmt.Sprintf("How many privateKeys will you create with the passphrase?"))
			if inputKeyNum, err := p.ReservedPrompt.PromptInput("num: "); err != nil {
				return 0, err
			} else if keyNum, err := strconv.Atoi(inputKeyNum); err != nil {
				return 0, err
			} else {
				return keyNum, nil
			}
		}()
		if err != nil {
			return err
		}

		//ecdsa key파일을 만들것인가?
		saveECDSA, err := func() (bool, error) {
			fmt.Println(fmt.Sprintf("create ecdsa hex file?"))
			if yn, err := p.ReservedPrompt.PromptInput("(y/N): "); err != nil {
				return false, err
			} else if yn == "y" || yn == "Y" {
				return true, nil
			} else {
				return false, nil
			}
		}()
		if err != nil {
			return err
		}

		for i := 0; i < keyNum; i++ {
			if privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), crand.Reader); err != nil { //privateKey를 생성.
				return err
			} else {
				//keyJson을 생성.
				key := &keystore.KeyV4{
					Id:          uuid.New(),
					Address:     crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
					PrivateKeys: [][]*ecdsa.PrivateKey{{privateKeyECDSA}},
				}
				if keyjson, err := keystore.EncryptKey(key, passphrase, keystore.StandardScryptN, keystore.StandardScryptP); err != nil {
					return err
				} else if encrypted, err := neo_aes.Encrypt([]byte(passphrase), keyjson); err != nil { //passphrase로 암호화.
					return err
				} else {
					newAddress := key.GetAddress()

					//address와 확장자 enc를 붙여서 base64 형태로 파일에 저장한다.
					file := path.Join(p.KeystoreDir, fmt.Sprintf("%s.enc", common.ToHex(newAddress[:])))

					if err := ioutil.WriteFile(file, []byte(encrypted), 0700); err != nil {
						return err
					} else if bytesRead, err := ioutil.ReadFile(file); err != nil { //파일에 저장이 잘되었는지 확인.
						return err
					} else if bytes.Equal([]byte(encrypted), bytesRead) == false {
						return fmt.Errorf("mismatch encrypted key json with bytes from file")
					}
					fmt.Println("new account: ", common.ToHex(newAddress[:]))

					if saveECDSA == true {
						if err := crypto.SaveECDSA(path.Join(p.KeystoreDir, fmt.Sprintf("%s.key", common.ToHex(newAddress[:]))), key.GetPrivateKey()); err != nil {
							return err
						} else {
							fmt.Println("saved ecdsa")
						}
					}

					//초기화..
					b := privateKeyECDSA.D.Bits()
					for i := range b {
						b[i] = 0
					}
				}
			}
		}
	}
	return nil
}

func (p *NeopinKeystore) Keystore() *keystore.KeyStore {
	return p.ks
}

func (p *NeopinKeystore) deriveKeyFromFile(fileName string) ([]byte, error) {
	if bytesData, err := ioutil.ReadFile(path.Join(fileName, ShareJsonScheme, ShareJsonFile)); err != nil {
		return []byte{}, err
	} else {
		return crypto.Keccak256Hash(bytesData).Bytes(), nil
	}
}

func (p *NeopinKeystore) deriveValueFromShare(share []byte) []byte {
	return crypto.Keccak256(share)[:2]
}

func (p *NeopinKeystore) getPassphrase(address common.Address, name string, threshold int) (string, *leveldb.Batch, error) {
	addressHex := hexutil.Encode(address[:])
	batch := (*leveldb.Batch)(nil)

	if threshold < 2 {
		fmt.Println(fmt.Sprintf("unlocking %s account %s", name, addressHex))
		if password, err := p.ReservedPrompt.PromptPassword("password: "); err != nil {
			return "", nil, err
		} else {
			return password, nil, nil
		}
	} else {
		subset := make(map[byte][]byte)
		fmt.Println(fmt.Sprintf("unlocking %s account %s", name, addressHex))
		for i := 0; i < threshold; i++ {
		read_share_key:
			if fileName, err := p.ReservedPrompt.PromptInput(fmt.Sprintf("Enter the path to read the share file (%d/%d): ", i+1, threshold)); err != nil {
				return "", nil, err
			} else {
				//trim.
				fileName = strings.TrimRight(fileName, " ")
				//파일에서 share를 읽어온다.
				shareJson := new(ShareJson)
				if err := shareJson.ReadFile(fileName, p.ReservedPrompt); err != nil {
					color.Red("read file error, %v", err)
					goto read_share_key
				}

				//ldb에 저장둔 key hash가 있다면 미리 체크한다.
				if p.keyhashDB != nil {
					if key, err := p.deriveKeyFromFile(fileName); err != nil {
						color.Red("read key hash db error %v", err)
					} else if value, err := p.keyhashDB.Get(key, nil); err == nil {
						if bytes.Equal(value, p.deriveValueFromShare(shareJson.Share)) == false {
							color.Red("not the same as keyhash.")
							goto read_share_key
						}
					} else if err.Error() == "leveldb: not found" { //키값이 없으면 batch에 넣는다. 만약 unlock에 성공을 하면 batch를 실행하도록 한다.
						if batch == nil {
							batch = new((leveldb.Batch))
						}
						batch.Put(key, p.deriveValueFromShare(shareJson.Share))
					} else {
						color.Red("read key hash db error %v", err)
					}
				}
				subset[shareJson.Index] = shareJson.Share
			}
		}
		fmt.Println("All shares are received up to the thresholds.")
		return string(sss.Combine(subset)), batch, nil
	}
}

func (p *NeopinKeystore) Unlock(account accounts.Account, name string, threshold int) error {
	if passphrase, batch, err := p.getPassphrase(account.Address, name, threshold); err != nil {
		return err
	} else if err := p.ks.Unlock(account, passphrase); err != nil {
		return err
	} else if p.ks.IsUnlocked(account.Address) == false {
		return fmt.Errorf("Failed to unlock %s", hexutil.Encode(account.Address[:]))
	} else {
		//unlock에 성공했다면 keyhash의 batch를 실행한다.
		if batch != nil && p.keyhashDB != nil {
			if err := p.keyhashDB.Write(batch, nil); err != nil {
				color.Yellow("write key hash error %v", err)
			}
		}
		fmt.Println(hexutil.Encode(account.Address[:]), "unlocked")
		return nil
	}
}

func (p *NeopinKeystore) Lock(address common.Address) error {
	if err := p.ks.Lock(address); err != nil {
		return err
	} else {
		fmt.Println(hexutil.Encode(address[:]), "locked")
		return nil
	}
}

func (p *NeopinKeystore) Find(address common.Address) (accounts.Account, error) {
	return p.ks.Find(accounts.Account{
		Address: address,
	})
}

func (p *NeopinKeystore) GetPublicKey(address common.Address, name string, threshold int, keepLock bool) (ecdsa.PublicKey, accounts.Account, error) {
	if a, err := p.Find(address); err != nil {
		return ecdsa.PublicKey{}, accounts.Account{}, err
	} else if passphrase, _, err := p.getPassphrase(address, name, threshold); err != nil {
		return ecdsa.PublicKey{}, accounts.Account{}, err
	} else if keyjson, err := p.ks.Export(a, passphrase, passphrase); err != nil {
		return ecdsa.PublicKey{}, accounts.Account{}, err
	} else if key, err := keystore.DecryptKey(keyjson, passphrase); err != nil {
		return ecdsa.PublicKey{}, accounts.Account{}, err
	} else {
		defer key.ResetPrivateKey()

		if keepLock == false {
			if err := p.ks.Unlock(a, passphrase); err != nil {
				return ecdsa.PublicKey{}, accounts.Account{}, err
			}
		}

		return key.GetPrivateKey().PublicKey, a, nil
	}
}

func (p *NeopinKeystore) SignTx(a accounts.Account, tx *types.Transaction, chainID *big.Int, name string, threshold int) (*types.Transaction, error) {
	if p.ks.IsUnlocked(a.Address) == false {
		if err := p.Unlock(a, name, threshold); err != nil {
			return nil, err
		}
	}
	if p.keepLock == true {
		defer p.ks.Lock(a.Address)
	}

	return p.ks.SignTx(a, tx, chainID)
}

func (p *NeopinKeystore) SignTxAsFeePayer(a accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if p.ks.IsUnlocked(a.Address) == false {
		if err := p.Unlock(a, "payer", 0); err != nil {
			return nil, err
		}
	}
	return p.ks.SignTxAsFeePayer(a, tx, chainID)
}

func (p *NeopinKeystore) SignHash(a accounts.Account, hash []byte, name string, threshold int) ([]byte, error) {
	if p.ks.IsUnlocked(a.Address) == false {
		if err := p.Unlock(a, name, threshold); err != nil {
			return nil, err
		}
	}
	if p.keepLock == true {
		defer p.ks.Lock(a.Address)
	}
	if sig, err := p.ks.SignHash(a, hash); err != nil {
		return nil, err
	} else {
		if sig[64] == 0 || sig[64] == 1 {
			sig[64] += 27
		}
		return sig, nil
	}
}

func (p *NeopinKeystore) normalPassphrase() (string, error) {
	if strict, err := p.ReservedPrompt.PromptInput(fmt.Sprintf("Choose whether it is a plain password(p) or a strict password(default) : ")); err != nil {
		return "", err
	} else {
		var (
			passphrase string
			err        error
		)

		switch strict {
		case "":
			passphrase, err = strictPassphrase("Enter the password of new account. : ", p.ReservedPrompt)
		default:
			passphrase, err = p.ReservedPrompt.PromptPassword("Enter the password of new account. : ")
		}
		if err != nil {
			return "", err
		}
		return passphrase, nil
	}
}

func (p *NeopinKeystore) sharedPassphrase(splitN int) (string, error) {
	fmt.Println("splitN:", splitN)

	//threshold를 정한다.
enter_threshold:
	if num, err := p.ReservedPrompt.PromptInput(fmt.Sprintf("Enter the threshold, must be less than or equal to %v. <threshold>: ", splitN)); err != nil {
		return "", nil
	} else if num == "" {
		goto enter_threshold
	} else if threshold, err := strconv.Atoi(num); err != nil {
		return "", nil
	} else if threshold <= 1 || threshold > splitN {
		color.Yellow("threshold must be greater than 1 and less than splitN.")
		goto enter_threshold
	} else {
		fmt.Println("threshold:", threshold)

		var (
			passphrase string
			shares     map[byte][]byte
			err        error
		)

		if passphrase, err = randPassphrase(16); err != nil { //passphrase를 생성.
			return "", err
		} else if shares, err = sss.Split(byte(splitN), byte(threshold), []byte(passphrase)); err != nil { //split한다.
			return "", err
		}
		fmt.Println("passphrase splited to", len(shares))

		//usb의 경로를 모두 저장한다.
		usbs := make(map[string]int)
		usbIndex := make(map[byte]string)
		for i := 1; i <= splitN; i++ {
		enter_path:
			if usbPath, err := p.ReservedPrompt.PromptInput(fmt.Sprintf("Enter the path to store the share (%d/%d) : ", i, splitN)); err != nil {
				return "", err
			} else if usbPath == "" {
				goto enter_path
			} else {
				usb := strings.TrimSpace(usbPath)
				if _, ok := usbs[usb]; ok == true {
					return "", fmt.Errorf("found equal usb directory")
				} else {
					usbs[usb] = i
					usbIndex[byte(i)] = usb

					if err := NewShareJson(byte(i), shares[byte(i)]).WriteFile(usb, p.ReservedPrompt); err != nil {
						return "", err
					}
					fmt.Println(i, "share of passphrase transfered to", usb)

					//share의 2 byte는 이미 체크할수있도록 ldb에 저장을 한다.
					if p.keyhashDB != nil {
						if key, err := p.deriveKeyFromFile(usb); err != nil {
							color.Yellow("read file error %v", err)
						} else {
							p.keyhashDB.Put(key, p.deriveValueFromShare(shares[byte(i)]), nil)
						}
					}
				}
			}
		}
		return passphrase, nil
	}
}

func (p *NeopinKeystore) newAccount(passphrase string) (accounts.Account, error) {
	//key 생성
	if account, err := p.ks.NewAccount(passphrase); err != nil {
		return accounts.Account{}, err
	} else {
		fmt.Println("account generated")
		fmt.Println("address:", hexutil.Encode(account.Address[:]))
		fmt.Println("keystore:", account.URL.Path)
		if p.ks.IsUnlocked(account.Address) == true {
			fmt.Println("keystore is unlocked")
		}
		return account, nil
	}
}

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func (p *NeopinKeystore) TimedUnlock(a accounts.Account, passphrase string, timeout time.Duration) error {
	if p.Keystore().IsUnlocked(a.Address) == false {
		return p.Keystore().TimedUnlock(a, passphrase, timeout)
	}

	return nil
}
