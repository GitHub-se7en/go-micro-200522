package cache

import (
	"fmt"
	reidsgo "github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis"
	"testing"
)

func TestRedis(t *testing.T) {
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
func TestRedis2(t *testing.T) {
	c, err := reidsgo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
}
