package server

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	r.GET("health", healthHandler)
}
