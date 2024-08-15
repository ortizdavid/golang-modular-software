package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ortizdavid/golang-modular-software/common/config"
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
	expiration := config.JwtExpiration()
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(expiration)).Unix(), // Token expires in 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
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
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}
	return signedToken, nil
}

// ParseToken parses and validates the JWT token
func (s *JwtService) ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse JWT token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	return token, *claims, nil
}
