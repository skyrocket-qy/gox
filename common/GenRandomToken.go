package common

import (
	"crypto/rand"
	"encoding/base64"
)

// Generate URL-Safe random string
func GenerateRandomToken(n int) (string, error) {
	tokenBytes := make([]byte, n)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(tokenBytes), nil
}
