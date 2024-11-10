package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate a hashed password with a cost of 12
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
