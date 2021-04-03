package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
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

	err := persistence.Redis.Set(ctback, "Hello", "Hello from redis container", 0).Err()
	if err != nil {
		panic(err)
	}

	return ctx.JSON("Hello key set!")
}

func GetHelloRedis(ctx *fiber.Ctx) error {

	val, err := persistence.Redis.Get(ctback, "Hello").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello|" + val)

	return ctx.JSON(val)
}

func ClearRedis(ctx *fiber.Ctx) error {

	if strings.HasPrefix(ctx.Hostname(), "localhost") {
		persistence.Redis.FlushDB(ctback)

		return ctx.JSON("Redis successfully flushed")
	}

	return ctx.JSON("Invalid request from " + ctx.IP())
}
