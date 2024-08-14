package services

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{
		secretKey: secretKey,
	}
}

// GenerateJwtToken creates a new JWT token for the given user ID
func (s *JwtService) GenerateJwtToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Printf("Error signing JWT token: %v", err)
		return "", err
	}
	return signedToken, nil
}

// GenerateRefreshToken creates a new refresh token for the given user ID
func (s *JwtService) GenerateRefreshToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), // Refresh token expires in 30 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		log.Printf("Error signing refresh token: %v", err)
		return "", err
	}
	return signedToken, nil
}

// ParseToken parses and validates the JWT token
func (s *JwtService) ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return nil, nil, err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, *claims, nil
}

