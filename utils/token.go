package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID uint) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*jwt.Token, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	return jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}
