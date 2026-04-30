package task

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	service *TaskService
}

func NewTaskHandler(service *TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func handleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Status, appErr)
		return
	}
	c.JSON(500, ErrInternal)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var newTask CreateTaskRequest

	if err := c.ShouldBindJSON(&newTask); err != nil {
		handleError(c, err)
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.CreateTask(ctx, newTask)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, response)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var upgTask UpdateTaskRequest

	if err := c.ShouldBindJSON(&upgTask); err != nil {
		handleError(c, err)
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.UpdateTask(ctx, id, upgTask)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, response)
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	ctx := c.Request.Context()
	response, err := h.service.ListTasks(ctx, id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, response)
}
