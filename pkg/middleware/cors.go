package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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

		if corsConfig != nil {
			for _, allowedOrigin := range corsConfig.AllowedOrigins {
				if origin == allowedOrigin {
					c.Header("Access-Control-Allow-Origin", origin)
					c.Header("Access-Control-Allow-Credentials", "true")
					break
				}
			}

			c.Header("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowedMethods, ", "))
			c.Header("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowedHeaders, ", "))
			c.Header("Access-Control-Expose-Headers", "Authorization")

			// Preflight-запрос (OPTIONS)
			if c.Request.Method == http.MethodOptions {
				c.Header("Access-Control-Max-Age", "86400")
				c.Status(http.StatusNoContent)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
