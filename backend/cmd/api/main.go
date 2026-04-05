package main

import (
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

	server.SetupRouter(r)

	r.Run(":" + cfg.Port)
}
