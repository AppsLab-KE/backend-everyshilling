package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// GenerateHash Create a secure password hash from a plaintext password
func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CompareHash Compares a password to a hash to see of they match
func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
