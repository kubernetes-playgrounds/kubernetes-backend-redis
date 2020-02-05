package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
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
