package service

import (
	"fmt"
	"time"

	"user-service/internal/models"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	GenerateTokens(payload jwt.MapClaims) (string, string, error)
	CreatePayload(user *models.User) jwt.MapClaims
}

type tokenService struct {
	jwtAccessSecret  []byte
	jwtRefreshSecret []byte
	jwtExpiration    int
}

func NewTokenService(accessSecret string, refreshSecret string, expiration int) TokenService {
	return &tokenService{
		jwtAccessSecret:  []byte(accessSecret),
		jwtRefreshSecret: []byte(refreshSecret),
		jwtExpiration:    expiration,
	}
}

func (s *tokenService) GenerateTokens(payload jwt.MapClaims) (string, string, error) {
	fmt.Println("payload:", payload)
	fmt.Println("accessSecret:", s.jwtAccessSecret)
	fmt.Println("refreshSecret:", s.jwtRefreshSecret)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	accessTokenStr, err := accessToken.SignedString(s.jwtAccessSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	refreshTokenStr, err := refreshToken.SignedString(s.jwtRefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func (s *tokenService) CreatePayload(user *models.User) jwt.MapClaims {
	return jwt.MapClaims{
		"role":    user.Role,
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Second * time.Duration(s.jwtExpiration)).Unix(),
	}
}
