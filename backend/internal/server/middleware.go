package server

import (
	"strings"
	"task-manager/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if bearerToken == "" {
			c.JSON(401, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		tokens := strings.TrimPrefix(bearerToken, "Bearer")

		if tokens == "" {
			c.JSON(401, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		userId, err := auth.ValidateToken(tokens, jwtSecret)
		if err != nil {
			c.JSON(401, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
