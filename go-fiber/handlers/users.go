package handlers

import (
	"context"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
TODOs:
- Verify is user exist before creation
- Redis instance
- Generic Redis operation interfac
*/

// user/:email
func GetUser(ctx *fiber.Ctx) error {

	email := ctx.Params("email")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	userId, err := rdb.Get(context.Background(), email).Result()

	if err == redis.Nil || userId == "" {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "User Id found for " + email, "data": userId})
}

func CreateUser(ctx *fiber.Ctx) error {

	newUser := new(models.User)

	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser.UserId = uuid.New().String()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rdb.Set(context.Background(), newUser.Email, newUser.UserId, 0)

	rdb.HSet(context.Background(), "User-"+newUser.UserId,
		"UserId", newUser.UserId,
		"Email", newUser.Email,
		"Name", newUser.Name,
	)

	return ctx.JSON(newUser)
}
