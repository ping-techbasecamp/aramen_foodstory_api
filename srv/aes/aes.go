package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/forgoer/openssl"
)

func Encrypt(key, plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	bPlainText := []byte(plainText)
	cipherText := make([]byte, aes.BlockSize+len(bPlainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(bPlainText))

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(key, cryptoText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}

func DecryptECB(key, encryptedText string) (string, error) {
	src, _ := base64.StdEncoding.DecodeString(encryptedText)
	newkey := []byte(key)
	dst, _ := openssl.AesECBDecrypt(src, newkey, openssl.PKCS7_PADDING)
	plainText, err := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(dst))
	return string(plainText), err
}

func EncryptECB(key, plainText string) (string, error) {
	src := []byte(plainText)
	newkey := []byte(key)
	dst, err := openssl.AesECBEncrypt(src, newkey, openssl.PKCS7_PADDING)
	return base64.StdEncoding.EncodeToString(dst), err
}
