package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Aquí validarías el JWT (simplificado)
		if token[1] != "valid-token" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}