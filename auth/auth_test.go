package auth

import (
	"testing"
)

func TestGenSalt(t *testing.T) {
	salt, err := GenSalt()
	if err != nil {
		t.Fatalf("GenSalt() returned an error: %v", err)
	}
	if len(salt) != 16 {
		t.Errorf("GenSalt() returned a salt of length %d, want 16", len(salt))
	}
}

func TestHashAndEqual(t *testing.T) {
	password := "password"
	salt, err := GenSalt()
	if err != nil {
		t.Fatalf("GenSalt() returned an error: %v", err)
	}

	hashedPassword := Hash(password, salt)

	if !Equal(password, salt, hashedPassword) {
		t.Errorf("Equal() returned false for the correct password")
	}

	if Equal("wrongpassword", salt, hashedPassword) {
		t.Errorf("Equal() returned true for an incorrect password")
	}
}
