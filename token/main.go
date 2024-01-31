package token

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomToken() (string, error) {
	tokenBytes := make([]byte, 32)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(tokenBytes), nil
}

func DecodeRandomToken(token string) ([]byte, error) {
	decodedBytes, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	return decodedBytes, nil
}
