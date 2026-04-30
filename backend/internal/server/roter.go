package server

import (
	"task-manager/internal/task"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, auth *AuthHandler, jwtSecret string, task *task.TaskHandler) {
	r.GET("health", healthHandler)
	r.POST("/auth/register", auth.Register)
	r.POST("/auth/login", auth.Login)

	protected := r.Group("/")
	protected.Use(AuthMiddleware(jwtSecret))
	{
		protected.POST("/task/createTask", task.CreateTask)
		protected.PATCH("/task/updateTask", task.UpdateTask)
		protected.GET("/task/tasksList/:id", task.ListTasks)
		protected.DELETE("/task/deleteTask/:id", task.DeleteTask)
	}
}
