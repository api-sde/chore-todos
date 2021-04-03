package persistence

/// See also https://github.com/uptrace/go-treemux-realworld-example-app/blob/3a0c6d30a2f931118c555dd1c54aa3b744df1c43/rwe/redis.go
// := vs =

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func ConnectRedis() {

	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Redis.Ping(context.Background()).Result()
	fmt.Println("Ping? " + pong)

	if !(err == nil || pong == "PONG") {
		log.Fatal("Cannot connect to Redis.")
		panic(err)
	}

	fmt.Println("Connected to Redis instance:" + Redis.String())
}
