package handlers

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
)

/*
TODOs:
- Verify is user exist before creation
- Generic Redis operation interfac
*/

// user/:email
func GetUser(ctx *fiber.Ctx) error {

	email := ctx.Params("email")

	userId, err := persistence.Redis.HGet(context.Background(), "Users", email).Result()

	if err == redis.Nil || userId == "" {
		return ctx.Status(404).JSON(fiber.Map{"message": "User not found for " + email})
	}

	userData, err := persistence.Redis.HGet(context.Background(), "Users", userId).Result()

	userResult := new(models.User)
	json.Unmarshal([]byte(userData), userResult)

	return ctx.JSON(userResult)
}

func GetUsers(ctx *fiber.Ctx) error {

	allUsers, err := persistence.Redis.HGetAll(context.Background(), "Users").Result()

	return ctx.JSON(fiber.Map{
		"count": len(allUsers) / 2,
		"error": err,
		"data":  allUsers,
	})
}

func CreateUser(ctx *fiber.Ctx) error {

	newUser := new(models.User)

	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Couldn't create user", "data": err})
	}

	newUser.UserId = uuid.New().String()

	newUserJson, err := json.Marshal(newUser)

	if err != nil {
		return ctx.SendStatus(500)
	}

	persistence.Redis.HSet(context.Background(), "Users",
		newUser.Email, newUser.UserId,
		newUser.UserId, newUserJson,
	)

	return ctx.JSON(newUser)
}
