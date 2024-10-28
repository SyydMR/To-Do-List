package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret")

type Claims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(id int64) (string, error) {
	expirationTime := time.Now().Add(3 * time.Hour)
	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyJWT(tokenString string) (int64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, fmt.Errorf("signature invalid")
		}
		return 0, fmt.Errorf("could not parse token: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.ID, nil
}
