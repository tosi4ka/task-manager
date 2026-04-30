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

// @Summary  Creating a task
// @Tags     tasks
// @Accept   json
// @Produce  json
// @Param    body body     task.CreateTaskRequest  true "Task creation payload"
// @Success  201  {object} task.Task
// @Failure  400  {object} ErrorType
// @Router   /task/createTask [post]
// @Security BearerAuth
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

// @Summary  Update a task
// @Tags     tasks
// @Accept   json
// @Produce  json
// @Param    id   path     string                 true "ID of the task to update"
// @Param    body body     task.UpdateTaskRequest true "Fields to update (all optional)"
// @Success  200  {object} task.Task
// @Failure  400  {object} ErrorType
// @Router /task/updateTask [patch]
// @Security BearerAuth
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

// @Summary Show all users tasks
// @Tags tasks
// @Accept       json
// @Produce json
// @Param id path string true "ID of the user whose tasks to retrieve"
// @Success 200 {array} task.Task
// @Failure 400 {object} task.AppError
// @Router /task/tasksList/{id} [get]
// @Security     BearerAuth
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

// @Summary Delete task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID of the task to delete"
// @Success 204
// @Failure 400 {object} task.AppError
// @Router /task/deleteTask/{id} [delete]
// @Security     BearerAuth
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	ctx := c.Request.Context()
	err = h.service.DeleteTask(ctx, id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(204)
}

// @Summary Find task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID of the task to retrieve"
// @Success 200 {object} task.Task
// @Failure 400 {object} task.AppError
// @Router /task/getById/{id} [get]
// @Security     BearerAuth
func (h *TaskHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	ctx := c.Request.Context()
	task, err := h.service.GetByID(ctx, id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, task)
}
