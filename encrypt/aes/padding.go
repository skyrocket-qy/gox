package aes

import "bytes"

func pkcs7Padding(ciphertext []byte, blockSize int) ([]byte, error) {
	padding := blockSize - len(ciphertext)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, paddingText...), nil
}

func pkcs7Unpacking(origData []byte) []byte {
	length := len(origData)
	unpacking := int(origData[length-1])

	return origData[:(length - unpacking)]
}
