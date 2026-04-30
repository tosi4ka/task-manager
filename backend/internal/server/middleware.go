package server

import (
	"strings"
	"task-manager/internal/auth"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
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

func RateLimitMiddleware() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  60,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	return mgin.NewMiddleware(instance)
}
