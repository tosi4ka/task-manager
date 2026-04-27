package server

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, auth *AuthHandler, jwtSecret string) {
	r.GET("health", healthHandler)
	r.POST("/auth/register", auth.Register)
	r.POST("/auth/login", auth.Login)

	protected := r.Group("/")
	protected.Use(AuthMiddleware(jwtSecret))
	{
	}
}
