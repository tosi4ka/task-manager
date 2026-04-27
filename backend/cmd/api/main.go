package main

import (
	"task-manager/internal/auth"
	"task-manager/internal/config"
	"task-manager/internal/db"
	"task-manager/internal/server"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.Load()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	database, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	err = db.RunMigrations(database, cfg.MigrationsPath)
	if err != nil {
		panic(err)
	}

	repo := auth.NewUserRepository(database)
	service := auth.NewUserService(repo, cfg.JWTSecret)
	authHandler := server.NewAuthHandler(service)

	server.SetupRouter(r, authHandler, cfg.JWTSecret)

	r.Run(":" + cfg.Port)
}
