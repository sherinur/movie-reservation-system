package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS

type CorsConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

var corsConfig *CorsConfig

func SetCorsConfig(config *CorsConfig) {
	corsConfig = config
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		for _, allowedOrigin := range corsConfig.AllowedOrigins {
			if origin == allowedOrigin {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.Header("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowedMethods, ", "))
		c.Header("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowedHeaders, ", "))

		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusOK)
			c.Abort()
			return
		}

		c.Next()
	}
}
