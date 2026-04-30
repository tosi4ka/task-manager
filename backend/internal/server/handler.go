package server

import (
	"errors"
	"task-manager/internal/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *auth.UserService
}

func NewAuthHandler(service *auth.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

func handleError(c *gin.Context, err error) {
	var appErr *auth.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Status, appErr)
		return
	}
	c.JSON(500, auth.ErrInternal)
}

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

// @Summary User registration
// @Tags auth
// @Accept json
// @Produce json
// @Param  body body auth.RegisterRequest true "User Data"
// @Success 201 {object} auth.AuthResponse
// @Failure 400 {object} auth.AppError
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var newUser auth.RegisterRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		handleError(c, err)
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.Register(ctx, newUser)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, response)
}

// @Summary Login
// @Tags auth
// @Accept json
// @Produce json
// @Param body body auth.LoginRequest true "Email and password"
// @Success 200 {object} auth.AuthResponse
// @Failure 400 {object} auth.AppError
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var user auth.LoginRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err)
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.Login(ctx, user)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, response)
}
