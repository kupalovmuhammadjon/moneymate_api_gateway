package redis

import (
	"api_gateway/configs"
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectDB(cfg *configs.Config) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
