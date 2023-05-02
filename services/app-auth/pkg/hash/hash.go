package hash

import (
	log "github.com/sirupsen/logrus"
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
		log.Error("error comparing hashes ", err)
		return false
	}
	return true
}
