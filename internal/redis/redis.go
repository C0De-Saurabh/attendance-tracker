package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var RedisClient *redis.Client
var ctx = context.Background()

func Connect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // default DB
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Unable to connect to Redis: %v", err)
	}
}
