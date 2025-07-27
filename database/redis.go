package database

import (
	"log"
	"sync"
	"relay/config"

	"github.com/redis/go-redis/v9"
)

var redisInstance *redis.Client
var redisOnce sync.Once

func GetRedisInstance() *redis.Client {
	redisOnce.Do(func() {
		redisURL := config.ENV.RedisURL

		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("Unable to parse Redis connection string: %v\n", err)
		}

		redisInstance = redis.NewClient(opt)
		log.Println("Redis client initialized successfully.")
	})

	return redisInstance
}
