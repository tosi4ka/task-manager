package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrl          string
	JWTSecret      string
	MigrationsPath string
	RedisURL       string
}

func Load() *Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbUrl := os.Getenv("DB_URL")

	jwtSecret := os.Getenv("JWT_SECRET")

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		migrationsPath = "migrations"
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}

	return &Config{
		Port:           port,
		DBUrl:          dbUrl,
		JWTSecret:      jwtSecret,
		MigrationsPath: migrationsPath,
		RedisURL:       redisURL,
	}
}
