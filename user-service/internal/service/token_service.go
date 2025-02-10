package service

import (
	"time"

	"user-service/internal/models"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	GenerateTokens(accessPayload jwt.MapClaims, refreshPayload jwt.MapClaims) (string, string, error)
	CreateAccessPayload(user *models.User) jwt.MapClaims
	CreateRefreshPayload(user *models.User) jwt.MapClaims
}

type tokenService struct {
	jwtAccessSecret      []byte
	jwtRefreshSecret     []byte
	jwtAccessExpiration  int
	jwtRefreshExpiration int
}

func NewTokenService(accessSecret string, refreshSecret string, accessExpiration int, refreshExpiration int) TokenService {
	return &tokenService{
		jwtAccessSecret:      []byte(accessSecret),
		jwtRefreshSecret:     []byte(refreshSecret),
		jwtAccessExpiration:  accessExpiration,
		jwtRefreshExpiration: refreshExpiration,
	}
}

func (s *tokenService) GenerateTokens(accessPayload jwt.MapClaims, refreshPayload jwt.MapClaims) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessPayload)
	accessTokenStr, err := accessToken.SignedString(s.jwtAccessSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshPayload)
	refreshTokenStr, err := refreshToken.SignedString(s.jwtRefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func (s *tokenService) CreateAccessPayload(user *models.User) jwt.MapClaims {
	return jwt.MapClaims{
		"role":    user.Role,
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Second * time.Duration(s.jwtAccessExpiration)).Unix(),
	}
}

func (s *tokenService) CreateRefreshPayload(user *models.User) jwt.MapClaims {
	return jwt.MapClaims{
		"role":    user.Role,
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Second * time.Duration(s.jwtRefreshExpiration)).Unix(),
	}
}
