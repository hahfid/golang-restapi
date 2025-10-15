package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang-restapi/pkg/jwt"
)

type AuthMiddleware struct {
	tokens *jwt.TokenManager
}

func NewAuthMiddleware(tokens *jwt.TokenManager) *AuthMiddleware {
	return &AuthMiddleware{tokens: tokens}
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		parts := strings.SplitN(authHeader, "Bearer ", 2)
		if len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		if _, err := m.tokens.ValidateToken(strings.TrimSpace(parts[1])); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}
}
