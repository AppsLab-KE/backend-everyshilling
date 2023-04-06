package tokens

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId string, secretKey string, expiryMinutes int) (string, error) {
	claims := jwt.MapClaims{}
	claims["uuid"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expiryMinutes)).UnixNano()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(jwtToken string, secretKey string) (userId string, err error) {
	parsedToken, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected sign algorithm")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", errors.New("invalid parsedToken")
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid parsedToken")
	}

	mapClaims := parsedToken.Claims.(jwt.MapClaims)
	userId, ok := mapClaims["uuid"].(string)

	if !ok {
		return "", errors.New("invalid parsedToken")
	}

	return userId, nil
}
