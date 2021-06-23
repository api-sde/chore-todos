package handlers

import (
	"encoding/json"

	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/services"
	"github.com/adrienBdx/chore-todos/gofiber/store"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// user/:email
func GetUser(ctx *fiber.Ctx) error {

	email := ctx.Params("email")

	if len(email) == 0 {
		return ctx.Status(400).JSON(fiber.Map{"message": "Parameter error, invalid email."})
	}

	if !services.IsUserExisting(email) {
		return ctx.Status(404).JSON(fiber.Map{"message": "User not found."})
	}

	userResult, err := services.GetUserByEmail(email)

	if err != nil {
		return ctx.Status(500).JSON(err)
	}

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

	if services.IsUserExisting(newUser.Email) {
		return ctx.Status(400).JSON(fiber.Map{"message": "A user already exist with this email."})
	}

	newUser.UserId = uuid.New().String()
	newUser.Ip = ctx.IP()
	hashedPassword, errEncrypt := services.EncryptPassword(newUser.Password)
	newUser.Password = hashedPassword

	newUserJson, err := json.Marshal(newUser)

	if err != nil || errEncrypt != nil {
		return ctx.SendStatus(500)
	}

	// Insert 2 pairs:
	// - email | userId
	// - userId | user(json)
	persistence.InsertInHash(store.Users,
		newUser.Email, newUser.UserId,
		newUser.UserId, newUserJson,
	)

	// Clean up secrets before returning the response
	newUser.Password = ""

	return ctx.JSON(newUser)
}
