package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func Encrypt(text []byte, key []byte) (string, error) {
	b, err := EncryptHex(text, key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
func EncryptHex(text []byte, key []byte) ([]byte, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, text)

	return encrypted, nil
}
func EncryptSting(text string, key []byte) (string, error) {
	return Encrypt([]byte(text), key)
}

func Decrypt(encrypted string, key []byte) ([]byte, error) {
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	return DecryptHex(src, key)
}

func DecryptHex(encryptedBytes []byte, key []byte) ([]byte, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(encryptedBytes))
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, encryptedBytes)
	return decrypted, nil
}

func DecryptSting(encrypted string, key []byte) (string, error) {
	b, e := Decrypt(encrypted, key)
	return string(b), e
}
