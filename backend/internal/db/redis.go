package db

import (
	"task-manager/internal/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg *config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return client, nil
}
