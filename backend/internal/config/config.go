package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBUrl     string
	JWTSecret string
}

func Load() *Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbUrl := os.Getenv("DB_URL")

	jwtSecret := os.Getenv("JWT_SECRET")

	return &Config{
		Port:      port,
		DBUrl:     dbUrl,
		JWTSecret: jwtSecret,
	}
}
