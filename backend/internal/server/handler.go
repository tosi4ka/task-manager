package server

import (
	"task-manager/internal/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	service *auth.UserService
}

func NewAuthHandler(service *auth.UserService) *AuthHandler{
	return &AuthHandler{service: service}
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func (h *AuthHandler) Register (c *gin.Context){
	var newUser auth.RegisterRequest
	
	if err := c.ShouldBindJSON(&newUser); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.Register(ctx, newUser)
	if err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

		c.JSON(201, response)
}

func (h *AuthHandler) Login (c *gin.Context){
	var user auth.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.Login(ctx, user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}