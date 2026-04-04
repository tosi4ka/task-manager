package main

import (
	"task-manager/internal/config"
	"task-manager/internal/db"
	"task-manager/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	db, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server.SetupRouter(r)

	r.Run(":" + cfg.Port)
}
