package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func Encrypt(text []byte, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, text)
	return hex.EncodeToString(encrypted), nil
}
func EncryptSting(text string, key []byte) (string, error) {
	return Encrypt([]byte(text), key)
}

func Decrypt(encrypted string, key []byte) ([]byte, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return decrypted, nil
}

func DecryptSting(encrypted string, key []byte) (string, error) {
	b, e := Decrypt(encrypted, key)
	return string(b), e
}
