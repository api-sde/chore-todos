package handlers

import (
	"fmt"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetHello(ctx *fiber.Ctx) error {

	helloUser := models.User{
		Ip:   ctx.IP(),
		Name: "Hello Bob",
	}

	return ctx.JSON(helloUser)
}

func SetHelloRedis(ctx *fiber.Ctx) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx.Context(), "hello", "hello from redis container", 0).Err()
	if err != nil {
		panic(err)
	}

	return ctx.JSON("Hello key set!")
}

func GetHelloRedis(ctx *fiber.Ctx) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx.Context(), "hello").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	return ctx.JSON(val)
}
