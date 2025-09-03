package common

import (
	"encoding/base64"
	"testing"
)

func TestGenerateRandomToken(t *testing.T) {
	n := 32
	token, err := GenerateRandomToken(n)
	if err != nil {
		t.Fatalf("GenerateRandomToken returned an error: %v", err)
	}

	if token == "" {
		t.Fatal("GenerateRandomToken returned an empty token")
	}

	decoded, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		t.Fatalf("Failed to decode token: %v", err)
	}

	if len(decoded) != n {
		t.Fatalf("Expected token length %d, got %d", n, len(decoded))
	}
}
