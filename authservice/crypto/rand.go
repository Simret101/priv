package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateSecureSecret(length int) string {
	// Create a byte slice with the specified length
	secret := make([]byte, length)
	// Read random data from crypto/rand
	_, err := rand.Read(secret)
	if err != nil {
		fmt.Println("Error generating secret:", err)
		return ""
	}
	// Encode the secret in base64 for easy use in your .env file
	return base64.StdEncoding.EncodeToString(secret)
}

func main() {
	// Generate a 32-byte secret
	secret := generateSecureSecret(32)
	fmt.Println("Generated Refresh Secret:", secret)
}
