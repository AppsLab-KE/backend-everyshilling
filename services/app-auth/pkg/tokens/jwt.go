package tokens

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const (
	publicKeyPath  = "/etc/auth-service/public.pem"
	privateKeyPath = "/etc/auth-service/private.pem"
)

func GenerateToken(userId string, expiryMinutes, refreshExpiryDays int) (string, string, error) {
	// open private key
	file, err := os.Open(privateKeyPath)
	if err != nil {
		return "", "", err
	}

	defer file.Close()

	// read public key
	privateKey, err := io.ReadAll(file)
	if err != nil {
		return "", "", err
	}

	// parse pub
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", "", err
	}

	claims := jwt.MapClaims{}
	claims["uuid"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expiryMinutes)).UnixNano()
	claims["type"] = "access"

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		return "", "", err
	}

	claims["exp"] = time.Now().Add(time.Hour * 24 * time.Duration(refreshExpiryDays)).UnixNano()
	claims["type"] = "refresh"

	refreshTokenString, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}

func VerifyToken(jwtToken string, isTypeRefreshToken bool) (userId string, err error) {

	// open public key
	file, err := os.Open(publicKeyPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// read public key
	publicKey, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// parse public key
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return "", err
	}

	// Parse the signed JWT and verify it with the RSA public key
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexepcted signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	if err != nil {
		log.Errorf("failed to parse token: %v", err)
		return "", errors.New("invalid token")
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	mapClaims := parsedToken.Claims.(jwt.MapClaims)
	userId, ok := mapClaims["uuid"].(string)

	if !ok {
		return "", errors.New("invalid token")
	}

	if isTypeRefreshToken {
		tokenType, ok := mapClaims["type"].(string)
		if !ok {
			return "", errors.New("invalid token type")
		}

		if tokenType != "refresh" {
			return "", errors.New("invalid token type")
		}
	} else {
		tokenType, ok := mapClaims["type"].(string)
		if !ok {
			return "", errors.New("invalid token type")
		}

		if tokenType != "access" {
			return "", errors.New("invalid token type")
		}
	}

	expiryTime, ok := mapClaims["exp"].(float64)
	if !ok {
		return "", errors.New("invalid/expired token")
	}

	if int64(expiryTime) <= time.Now().UnixNano() {
		return "", errors.New("expired token")
	}

	return userId, nil
}
