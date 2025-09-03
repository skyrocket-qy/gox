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

	// Test with empty password
	emptyPassword := ""
	hashedEmptyPassword := Hash(emptyPassword, salt)
	if !Equal(emptyPassword, salt, hashedEmptyPassword) {
		t.Errorf("Equal() returned false for an empty password")
	}
	if Equal("somepassword", salt, hashedEmptyPassword) {
		t.Errorf("Equal() returned true for an incorrect password when comparing with empty password hash")
	}

	// Test with password containing special characters
	specialCharPassword := "!@#$%^&*()"
	hashedSpecialCharPassword := Hash(specialCharPassword, salt)
	if !Equal(specialCharPassword, salt, hashedSpecialCharPassword) {
		t.Errorf("Equal() returned false for a password with special characters")
	}
	if Equal("somepassword", salt, hashedSpecialCharPassword) {
		t.Errorf("Equal() returned true for an incorrect password when comparing with special char password hash")
	}

	// Test with a different valid hashed password
	password2 := "anotherpassword"
	salt2, err := GenSalt()
	if err != nil {
		t.Fatalf("GenSalt() returned an error: %v", err)
	}
	hashedPassword2 := Hash(password2, salt2)
	if Equal(password, salt2, hashedPassword2) {
		t.Errorf("Equal() returned true for a different password with different salt")
	}
}
