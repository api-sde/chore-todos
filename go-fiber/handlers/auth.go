package handlers

import (
	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/adrienBdx/chore-todos/gofiber/services"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {

	userToLogin := new(models.User)

	if err := ctx.BodyParser(userToLogin); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": "Couldn't parse user", "error": err})
	}

	if !(len(userToLogin.Email) > 0) ||
		!(len(userToLogin.Name) > 0) ||
		!(services.IsUserExisting(userToLogin.Email)) {

		return ctx.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	userModel, err := services.GetUserByEmail(userToLogin.Email)

	if err != nil || !services.ValidatePassword(userToLogin.Password, userModel.Password) {
		return ctx.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userToLogin.Name
	claims["email"] = userToLogin.Email
	claims["userid"] = userModel.UserId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["validation_trace"] = uuid.New() // For logout, to implement

	newToken, err := token.SignedString([]byte("SECRET")) //config.Config("SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "Success login", "data": newToken})
}

func Logout(ctx *fiber.Ctx) error {
	return nil
}

func Authorize(ctx *fiber.Ctx) error {

	authBearer := ctx.Get(fiber.HeaderAuthorization)
	jwtToken := strings.Fields(authBearer)[1]
	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userid"].(string)
	email := claims["email"].(string)
	username := claims["username"].(string)
	validation_trace := claims["validation_trace"].(string)

	ctx.Locals("CurrentUserId", userId)
	ctx.Locals("CurrentUserEmail", email)
	ctx.Locals("CurrentUserName", username)
	ctx.Locals("ValidationTrace", validation_trace)

	return nil
}
