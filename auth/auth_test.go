package auth_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/auth"
)

func TestGenSalt(t *testing.T) {
	salt, err := auth.GenSalt()
	if err != nil {
		t.Fatalf("GenSalt() returned an error: %v", err)
	}

	if len(salt) != 16 {
		t.Errorf("GenSalt() returned a salt of length %d, want 16", len(salt))
	}
}

func TestHashAndEqual(t *testing.T) {
	password := "password"

	salt, err := auth.GenSalt()
	if err != nil {
		t.Fatalf("GenSalt() returned an error: %v", err)
	}

	hashedPassword := auth.Hash(password, salt)

	if !auth.Equal(password, salt, hashedPassword) {
		t.Errorf("Equal() returned false for the correct password")
	}

	if auth.Equal("wrongpassword", salt, hashedPassword) {
		t.Errorf("Equal() returned true for an incorrect password")
	}

	// Test with empty password
	emptyPassword := ""

	hashedEmptyPassword := auth.Hash(emptyPassword, salt)
	if !auth.Equal(emptyPassword, salt, hashedEmptyPassword) {
		t.Errorf("Equal() returned false for an empty password")
	}

	if auth.Equal("somepassword", salt, hashedEmptyPassword) {
		t.Errorf(
			"Equal() returned true for an incorrect password when comparing with empty password hash",
		)
	}

	// Test with password containing special characters
	specialCharPassword := "!@#$%^&*()"

	hashedSpecialCharPassword := auth.Hash(specialCharPassword, salt)
	if !auth.Equal(specialCharPassword, salt, hashedSpecialCharPassword) {
		t.Errorf("Equal() returned false for a password with special characters")
	}

	if auth.Equal("somepassword", salt, hashedSpecialCharPassword) {
		t.Errorf(
			"Equal() returned true for an incorrect password when comparing with special char password hash",
		)
	}

	// Test with a different valid hashed password
	password2 := "anotherpassword"

	salt2, err := auth.GenSalt()
	if err != nil {
		t.Fatalf("GenSalt() returned an error: %v", err)
	}

	hashedPassword2 := auth.Hash(password2, salt2)
	if auth.Equal(password, salt2, hashedPassword2) {
		t.Errorf("Equal() returned true for a different password with different salt")
	}
}
