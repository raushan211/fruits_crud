package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func RedisConnect() {

	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := Client.Ping().Result()

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", ping)
	}

}
