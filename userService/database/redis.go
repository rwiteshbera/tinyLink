package database

import (
	"context"
	"userService/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateRedisClient(config *config.Config, databaseNumber int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_DB_ADDRESS,
		Password: config.REDIS_DB_PASSWORD,
		DB:       databaseNumber,
	})

	return redisClient
}
