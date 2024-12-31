package core

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("your-secret-key")

// JWTService is responsible for JWT token operations
type JWTService struct {}

// NewJWTService creates a new JWT service
func NewJWTService() *JWTService {
	return &JWTService{}
}

// GenerateToken generates a JWT token for a user
func (s *JWTService) GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token
func (s *JWTService) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}
	return false, nil
}
