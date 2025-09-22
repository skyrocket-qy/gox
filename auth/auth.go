package auth

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func Hash(password string, salt []byte) (hashedPassword []byte) {
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

func Equal(password string, salt, hashedPassword []byte) bool {
	return bytes.Equal(Hash(password, salt), hashedPassword)
}

func GenSalt() (salt []byte, err error) {
	salt = make([]byte, 16)
	_, err = rand.Read(salt)

	return salt, err
}
