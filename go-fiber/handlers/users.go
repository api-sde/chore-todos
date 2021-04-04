package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/store"
)

// user/:email
func GetUser(ctx *fiber.Ctx) error {

	email := ctx.Params("email")

	userId, err := persistence.GetHashValue(store.Users, email)

	if err != nil {
		return ctx.Status(404).JSON(err)
	}

	userJson, err := persistence.GetHashValue(store.Users, userId)
	userResult := models.ToModel(new(models.User), userJson)

	return ctx.JSON(userResult)
}

func GetUsers(ctx *fiber.Ctx) error {

	usersMap, err := persistence.GetAllHash(store.Users)

	if err != nil {
		return ctx.Status(404).JSON(err)
	}

	modelList := models.ToCollectionModel(models.User{}, usersMap)

	return ctx.JSON(fiber.Map{
		"count": len(usersMap) / 2,
		"error": err,
		"data":  modelList,
	})
}

func CreateUser(ctx *fiber.Ctx) error {

	newUser := new(models.User)

	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "Couldn't parse user", "error": err})
	}

	if persistence.ExistInHash(store.Users, newUser.Email) {
		return ctx.Status(400).JSON(fiber.Map{"message": "A user already exist with this email."})
	}

	newUser.UserId = uuid.New().String()

	newUserJson, err := json.Marshal(newUser)

	if err != nil {
		return ctx.SendStatus(500)
	}

	// Insert 2 pairs: email | userId & userId | user(json)
	persistence.InsertInHash(store.Users,
		newUser.Email, newUser.UserId,
		newUser.UserId, newUserJson,
	)

	return ctx.JSON(newUser)
}
