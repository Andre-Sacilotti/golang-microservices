package utils

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"fmt"
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
	ciphertext, _ := b64.StdEncoding.DecodeString(EncryptedText)
	nonce := make([]byte, 12)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("111111111111111")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("1112222222222")
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}
