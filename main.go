package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "localhost:6379",
		Password: "root", // no password set
		DB:       0,      // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>


}