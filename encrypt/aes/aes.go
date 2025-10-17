package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

var ErrInvalidBlockSize = errors.New("invalid block size")

func EncryptCBC(plainEncryptText, secretKey, iv []byte) ([]byte, error) {
	plainEncryptText = pkcs7Padding(plainEncryptText, aes.BlockSize)

	if len(plainEncryptText)%aes.BlockSize != 0 {
		return nil, ErrInvalidBlockSize
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(plainEncryptText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainEncryptText)

	return cipherText, nil
}

func EncryptCBCWithBase64(plainEncryptText, secretKey, iv []byte) (string, error) {
	cipherText, err := EncryptCBC(plainEncryptText, secretKey, iv)
	if err != nil {
		return "", err
	}

	result := base64.StdEncoding.EncodeToString(cipherText)

	return result, nil
}

func DecryptCBC(plainDecryptText, secretKey, iv []byte) ([]byte, error) {
	if len(plainDecryptText)%aes.BlockSize != 0 {
		return nil, ErrInvalidBlockSize
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(plainDecryptText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(out, plainDecryptText)

	out = pkcs7Unpacking(out)

	return out, nil
}

func DecryptWithBase64(payload string, secretKey, iv []byte) ([]byte, error) {
	plainDecryptText, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	return DecryptCBC(plainDecryptText, secretKey, iv)
}
