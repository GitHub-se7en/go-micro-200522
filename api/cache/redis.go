package cache

import (
	"github.com/go-redis/redis"
	"log"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		log.Fatalf("redis conn error %v",err)
	}
}