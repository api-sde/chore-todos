package handlers

import (
	"context"
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

var ctback = context.Background()

func SetHelloRedis(ctx *fiber.Ctx) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctback, "hello", "hello from redis container", 0).Err()
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

	val, err := rdb.Get(ctback, "hello").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)

	return ctx.JSON(val)
	//return ctx.JSON("connected" + rdb.String() + " ")
}

func ClearRedis(ctx *fiber.Ctx) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.FlushDB(ctback)

	return ctx.SendStatus(200)
}
