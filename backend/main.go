package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {
	env := os.Getenv("ENV")
	var endpoint string
	switch env {
	case "PRODUCTION":
		endpoint = "redis-test"
	default:
		endpoint = "0.0.0.0"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     endpoint + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for {
		pong, err := client.Ping().Result()
		fmt.Println(pong)
		if err != nil {
			fmt.Println("err: ", err)
		}
		time.Sleep(5 * time.Second)
	}

}
