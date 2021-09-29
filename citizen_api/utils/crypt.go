package utils

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"encoding/hex"
	"os"
)

func Encrypt(PlainText string) string {
	key := []byte(os.Getenv("SECRET"))
	plaintext := []byte(PlainText)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, 12)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return string(b64.StdEncoding.EncodeToString(ciphertext))
}

func Decrypt(EncryptedText string) string {
	key := []byte(os.Getenv("SECRET"))
	ciphertext, _ := hex.DecodeString(EncryptedText)
	nonce := make([]byte, 12)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	sDec, _ := b64.StdEncoding.DecodeString(string(plaintext))
	return string(sDec)
}
