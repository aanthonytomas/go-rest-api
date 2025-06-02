package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RDB *redis.Client

func InitRedis() *redis.Client {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password by default
		DB:       0,
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}
	log.Println("Connected to Redis!")
	return RDB
}
