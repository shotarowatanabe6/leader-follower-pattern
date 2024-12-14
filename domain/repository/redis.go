package domain

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type IRedisRepository interface {
	Get(ctx *gin.Context, key string) (string, error)
	Set(ctx *gin.Context, key, value string) error
}

type Redis struct {
	client *redis.Client
}

func NewRedisRepository() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return Redis{client}
}

func (r Redis) Get(ctx *gin.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("key not found: %v", err)
	}
	return value, nil
}

func (r Redis) Set(ctx *gin.Context, key, value string) error {
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set key: %v", err)
	}
	return nil
}
