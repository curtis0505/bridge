package neo_aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"strings"
)

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func Encrypt(key []byte, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := Pad(text)
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], msg)
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
	return finalMsg, nil
}

func Decrypt(key []byte, text string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return []byte{}, err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return []byte{}, errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := Unpad(msg)
	if err != nil {
		return []byte{}, err
	}

	return unpadMsg, nil
}

func CBCEncrypt(key []byte, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := Pad(text)
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], msg)

	finalMsg := hex.EncodeToString(iv) + ":" + hex.EncodeToString(ciphertext[aes.BlockSize:])

	return finalMsg, nil

}

func CBCDecrypt(key []byte, text string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	iv, _ := hex.DecodeString(strings.Split(text, ":")[0])
	msg, _ := hex.DecodeString(strings.Split(text, ":")[1])

	if (len(msg) % aes.BlockSize) != 0 {
		return []byte{}, errors.New("blocksize must be multipe of decoded message length")
	}

	cbc := cipher.NewCBCDecrypter(block, iv)

	cbc.CryptBlocks(msg, msg)

	unpadMsg, err := Unpad(msg)

	if err != nil {
		return []byte{}, err
	}

	return unpadMsg, nil
}
