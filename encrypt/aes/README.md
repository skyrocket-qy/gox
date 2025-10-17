# encrypt/aes

The `encrypt/aes` package provides robust AES-256 encryption and decryption capabilities using the Cipher Block Chaining (CBC) mode. It includes essential utilities for PKCS7 padding and unpadding, ensuring data integrity for block cipher operations. Additionally, it offers convenient functions for handling Base64 encoding and decoding of encrypted data.

## Features

*   **AES-256 CBC Encryption:** Encrypts plaintext using the AES algorithm in CBC mode with a 256-bit key.
*   **AES-256 CBC Decryption:** Decrypts ciphertext that was encrypted using AES-256 CBC.
*   **PKCS7 Padding:** Implements PKCS7 padding to ensure plaintext length is a multiple of the AES block size.
*   **PKCS7 Unpadding:** Removes PKCS7 padding after decryption.
*   **Base64 Integration:** Provides helper functions to encrypt and decrypt data directly with Base64 encoding/decoding, simplifying common workflows.

## Usage Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/skyrocket-qy/ciri/encrypt/aes"
)

func main() {
	// Your 32-byte (256-bit) secret key
	secretKey := []byte("averysecretkeythatis32byteslong")
	// Your 16-byte (128-bit) initialization vector
	iv := []byte("thisis16bytesiv!")

	plainText := []byte("Hello, secure world! This is a secret message.")
	fmt.Printf("Original Plaintext: %s\n", plainText)

	// --- Encryption ---
	// Encrypt using CBC mode, then Base64 encode the result
	encryptedBase64, err := aes.EncryptCBCWithBase64(plainText, secretKey, iv)
	if err != nil {
		log.Fatalf("Error encrypting: %v", err)
	}
	fmt.Printf("Encrypted (Base64): %s\n", encryptedBase64)

	// --- Decryption ---
	// Decrypt the Base64 encoded ciphertext
	decryptedText, err := aes.DecryptWithBase64(encryptedBase64, secretKey, iv)
	if err != nil {
		log.Fatalf("Error decrypting: %v", err)
	}
	fmt.Printf("Decrypted Plaintext: %s\n", decryptedText)

	// Verify if decrypted text matches original
	if string(plainText) == string(decryptedText) {
		fmt.Println("Verification successful: Decrypted text matches original.")
	} else {
		fmt.Println("Verification failed: Decrypted text does NOT match original.")
	}

	// --- Direct CBC Encryption/Decryption (without Base64 helper) ---
	fmt.Println("\n--- Direct CBC Encryption/Decryption ---")
	cipherText, err := aes.EncryptCBC(plainText, secretKey, iv)
	if err != nil {
		log.Fatalf("Error direct encrypting: %v", err)
	}
	fmt.Printf("Ciphertext (raw bytes): %x\n", cipherText)

	decryptedRawText, err := aes.DecryptCBC(cipherText, secretKey, iv)
	if err != nil {
		log.Fatalf("Error direct decrypting: %v", err)
	}
	fmt.Printf("Decrypted Raw Plaintext: %s\n", decryptedRawText)
}
```