package handlers

import (
	"crypto/rand"
	"encoding/json"
	"math"
	"strconv"
	"strings"

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
	newUser.Password, newUser.UserKey, _ = BlurPassword(newUser.Email, newUser.Password)

	newUserJson, err := json.Marshal(newUser)

	if err != nil {
		return ctx.SendStatus(500)
	}

	// Insert 2 pairs: email | userId & userId | user(json)
	persistence.InsertInHash(store.Users,
		newUser.Email, newUser.UserId,
		newUser.UserId, newUserJson,
	)

	// Clean up secrets before returning the response
	newUser.Password = ""
	newUser.UserKey = ""

	return ctx.JSON(newUser)
}

/// Heuristic unsecure blur
func BlurPassword(email string, password string) (blurred string, key string, err error) {

	// User specific value, can be common
	size := len(email) + len(password)

	// Unique user key
	userRandomKey, err := rand.Prime(rand.Reader, size)

	if err != nil {
		return "", "", err
	}

	seed := ((float64(size) / math.Pi) * float64(userRandomKey.Int64()))
	seedString := strconv.FormatFloat(seed, 'G', 128, 64)

	var blurredPassword strings.Builder
	var switchToAscii bool = false
	var switchRune bool = false

	for position, char := range password {

		var result rune

		charInt := int(char)

		var index int
		if switchRune {
			index = charInt + position
			switchRune = !switchRune
		} else {
			index = charInt - position
		}

		for len(seedString) <= index {
			index = index / 2
		}

		result = rune(seedString[index])

		if switchToAscii {
			blurredPassword.WriteString(string(result))
		} else {
			blurredPassword.WriteRune(result)
		}

		switchToAscii = !switchToAscii
	}

	return blurredPassword.String(), string(userRandomKey.Int64()), nil
}

/* To Do: bcrypt:
https://hackernoon.com/how-to-store-passwords-example-in-go-62712b1d2212
https://gowebexamples.com/password-hashing/
*/
