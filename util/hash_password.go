package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword hashes a password using SHA256 and returns its hexadecimal representation
func HashPassword(password string) string {
	// Generate a random salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}

	// Combine the password and the salt
	saltedPassword := append([]byte(password), salt...)

	// Hash the salted password using SHA256
	hash := sha256.Sum256(saltedPassword)

	// Return the hexadecimal representation of the hash
	return hex.EncodeToString(hash[:])
}

// CheckPassword compares a hashed password with a plain-text password
func CheckPassword(hashedPassword, password string) bool {
	// Convert the hexadecimal representation of the hashed password back to bytes
	bytes, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	// Generate the hash of the plain-text password
	hash := sha256.Sum256(append([]byte(password), bytes[32:]...))

	// Compare the hashed password with the generated hash
	return hashedPassword == hex.EncodeToString(hash[:])
}
