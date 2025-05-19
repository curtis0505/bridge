package neo_keystore

import (
	"crypto/rand"
	"fmt"
	"github.com/curtis0505/bridge/libs/crypto/prompt"
	"math/big"
	"regexp"

	"github.com/klaytn/klaytn/console"
)

const (
	strictMinLength       = 13
	strictMinSpecialChars = 1
	strictMinBigLetters   = 1
	strictMinDigits       = 1
)

func randPassphrase(n int) (string, error) {
	var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+{}|[]")
	b := make([]rune, n)
	max := new(big.Int).SetUint64(uint64(len(runes)))
	for i := range b {
		if n, err := rand.Int(rand.Reader, max); err != nil {
			return "", nil
		} else {
			b[i] = runes[n.Uint64()]
		}

	}
	return string(b), nil
}

func strictPassphrase(prompt string, reservedPrompt *prompt.ReservedPromptInput) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
		fmt.Println(fmt.Sprintf(" ⓥ %v or more characters", strictMinLength))
		fmt.Println(fmt.Sprintf(" ⓥ %v or more special characters", strictMinSpecialChars))
		fmt.Println(fmt.Sprintf(" ⓥ %v or more big letters", strictMinBigLetters))
		fmt.Println(fmt.Sprintf(" ⓥ %v or more digits", strictMinDigits))
	}
	var (
		passphrase, confirm string
		err                 error
	)
	if reservedPrompt == nil {
		passphrase, err = console.Stdin.PromptPassword("Passphrase: ")
	} else {
		passphrase, err = reservedPrompt.PromptPassword("Passphrase: ")
	}
	if err != nil {
		return "", fmt.Errorf("Failed to read passphrase: %v", err)
	} else if err = verifyStrictPassphrase(passphrase); err != nil {
		return "", err
	}
	if reservedPrompt == nil {
		confirm, err = console.Stdin.PromptPassword("Repeat Passphrase: ")
	} else {
		confirm, err = reservedPrompt.PromptPassword("Repeat Passphrase: ")
	}
	if err != nil {
		return "", fmt.Errorf("Failed to read passphrase confirmation: %v", err)
	}
	if passphrase != confirm {
		return "", fmt.Errorf("passphrase mismatched")
	}
	return passphrase, nil
}

func verifyStrictPassphrase(passphrase string) error {
	if !regexp.MustCompile("[[:alpha:]]").MatchString(passphrase) {
		return fmt.Errorf("require only ascii")
	} else if !regexp.MustCompile(fmt.Sprintf(".{%d,}", strictMinLength)).MatchString(passphrase) {
		return fmt.Errorf("require %v or more characters", strictMinLength)
	} else if !regexp.MustCompile(fmt.Sprintf("[^A-Za-z0-9]{%d,}", strictMinSpecialChars)).MatchString(passphrase) {
		return fmt.Errorf("require %v or more special characters", strictMinSpecialChars)
	} else if !regexp.MustCompile(fmt.Sprintf("[A-Z]{%d,}", strictMinBigLetters)).MatchString(passphrase) {
		return fmt.Errorf("require %v or more big letters", strictMinBigLetters)
	} else if !regexp.MustCompile(fmt.Sprintf("[0-9]{%d,}", strictMinDigits)).MatchString(passphrase) {
		return fmt.Errorf("require %v or more digits", strictMinDigits)
	} else {
		return nil
	}
}
