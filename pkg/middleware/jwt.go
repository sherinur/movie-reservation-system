package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var ErrEmptyHeader = errors.New("authorization header is empty")
var ErrNoBearer = errors.New("authorization header does not contain 'Bearer'")
var ErrInvalidToken = errors.New("invalid token")

var secret []byte

func SetSecret(jwtSecret []byte) {
	secret = jwtSecret
}

func GetSecret() []byte {
	return secret
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		tokenString, err := getToken(header)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		jwtToken, err := jwt.Parse(tokenString, keyFunc)

		if err != nil || !jwtToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	return []byte(secret), nil
}

func getToken(header string) (string, error) {
	if err := validateHeader(header); err != nil {
		return "", err
	}

	return strings.TrimPrefix(header, "Bearer "), nil
}

func validateHeader(header string) error {
	if header == "" {
		return ErrEmptyHeader
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return ErrNoBearer
	}

	return nil
}
