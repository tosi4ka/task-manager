package main

import (
	"task-manager/internal/auth"
	"task-manager/internal/config"
	"task-manager/internal/db"
	"task-manager/internal/server"
	"task-manager/internal/task"

	_ "task-manager/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// @title           Task Manager API
// @version         1.0
// @description     REST API

// @host            localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name            Authorization
func main() {
	cfg := config.Load()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	database, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	redis, err := db.ConnectRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer redis.Close()

	err = db.RunMigrations(database, cfg.MigrationsPath)
	if err != nil {
		panic(err)
	}

	repo := auth.NewUserRepository(database)
	service := auth.NewUserService(repo, cfg.JWTSecret)
	authHandler := server.NewAuthHandler(service)

	taskRepo := task.NewTaskRepository(database)
	taskService := task.NewTaskService(taskRepo)
	taskHandler := task.NewTaskHandler(taskService)

	server.SetupRouter(r, authHandler, cfg.JWTSecret, taskHandler, redis)

	r.Run(":" + cfg.Port)
}
