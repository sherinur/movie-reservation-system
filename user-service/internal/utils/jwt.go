package utils

import (
	"time"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt"
)

var expHours = 1440

func GenerateJWT(user *models.User, jwtSecretKey []byte) (string, error) {
	payload := jwt.MapClaims{
		"sub":    user.Email,
		"status": "client",
		"exp":    time.Now().Add(time.Hour * time.Duration(expHours)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
