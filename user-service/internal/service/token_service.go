package service

import (
	"fmt"
	"time"

	"user-service/internal/models"

	"github.com/golang-jwt/jwt"
)

type TokenService interface {
	GenerateTokens(accessPayload jwt.MapClaims, refreshPayload jwt.MapClaims) (string, string, error)
	CreateAccessPayload(user *models.User) jwt.MapClaims
	CreateRefreshPayload(user *models.User) jwt.MapClaims
	ValidateAccessToken(tokenStr string) bool
	ParseRefreshToken(tokenStr string) *models.User
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

func (s *tokenService) ValidateAccessToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtAccessSecret, nil
	})

	return err == nil && token.Valid
}

func (s *tokenService) ParseRefreshToken(tokenStr string) *models.User {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtRefreshSecret, nil
	})
	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	if claims["user_id"] == nil || claims["role"] == nil {
		return nil
	}

	return &models.User{
		ID:   claims["user_id"].(string),
		Role: claims["role"].(string),
	}
}
