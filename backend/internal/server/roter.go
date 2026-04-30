package server

import (
	"task-manager/internal/task"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine, auth *AuthHandler, jwtSecret string, task *task.TaskHandler) {
	r.GET("health", healthHandler)
	r.POST("/auth/register", auth.Register)
	r.POST("/auth/login", auth.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := r.Group("/")
	protected.Use(AuthMiddleware(jwtSecret))
	{
		protected.POST("/task/createTask", task.CreateTask)
		protected.PATCH("/task/updateTask/:id", task.UpdateTask)
		protected.GET("/task/tasksList/:id", task.ListTasks)
		protected.GET("/task/getById/:id", task.GetByID)
		protected.DELETE("/task/deleteTask/:id", task.DeleteTask)
	}
}
