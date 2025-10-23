package neo_keystore

import (
	"crypto/aes"
	"encoding/json"
	"fmt"
	"github.com/curtis0505/bridge/libs/crypto/prompt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/kaiachain/kaia/console"
	"github.com/kaiachain/kaia/crypto"
)

const (
	ShareJsonScheme = "passphrase"
	ShareJsonFile   = "share.sss"
)

type ShareJson struct {
	Index byte
	Share []byte
}

func NewShareJson(index byte, share []byte) *ShareJson {
	return &ShareJson{
		Index: index,
		Share: share,
	}
}

func (p *ShareJson) WriteFile(dir string, reservedPrompt *prompt.ReservedPromptInput) error {
	if err := p.encrypt(reservedPrompt); err != nil {
		fmt.Println("share key json encrypt error, ", err)
		return err
	}

	dir = strings.TrimSpace(dir)
	endPath := path.Join(dir, ShareJsonScheme)
	if err := os.RemoveAll(endPath); err != nil {
		return err
	} else if err := os.MkdirAll(endPath, 0700); err != nil {
		return err
	} else if bytes, err := json.Marshal(p); err != nil {
		return err
	} else if err := ioutil.WriteFile(path.Join(endPath, ShareJsonFile), bytes, 0700); err != nil {
		return err
	}
	return nil
}

func (p *ShareJson) ReadFile(dir string, reservedPrompt *prompt.ReservedPromptInput) error {
	dir = strings.TrimSpace(dir)
	endPath := path.Join(dir, ShareJsonScheme)
	if bytes, err := ioutil.ReadFile(path.Join(endPath, ShareJsonFile)); err != nil {
		return err
	} else if err := json.Unmarshal(bytes, p); err != nil {
		return err
	} else if err := p.decrypt(reservedPrompt); err != nil {
		return err
	}
	return nil
}

func (p *ShareJson) encrypt(reservedPrompt *prompt.ReservedPromptInput) error {
	if passphrase, err := strictPassphrase("Enter the password of new share. : ", reservedPrompt); err != nil {
		return err
	} else {
		key := crypto.Keccak256([]byte(passphrase))[:len(p.Share)]

		if block, err := aes.NewCipher([]byte(key)); err != nil {
			return err
		} else {
			ciphertext := make([]byte, block.BlockSize())
			block.Encrypt(ciphertext, p.Share)
			p.Share = ciphertext
		}
	}
	return nil
}

func (p *ShareJson) decrypt(reservedPrompt *prompt.ReservedPromptInput) error {
	fmt.Println("unlocking share")
	var (
		passphrase string
		err        error
	)
	if reservedPrompt == nil {
		passphrase, err = console.Stdin.PromptPassword("password: ")
	} else {
		passphrase, err = reservedPrompt.PromptPassword("password: ")
	}
	if err != nil {
		return err
	}

	key := crypto.Keccak256([]byte(passphrase))[:len(p.Share)]

	if block, err := aes.NewCipher(key); err != nil {
		return err
	} else {
		plaintext := make([]byte, block.BlockSize())
		block.Decrypt(plaintext, p.Share)
		p.Share = plaintext
	}
	return nil
}
