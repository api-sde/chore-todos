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
	claims["validation_trace"] = uuid.New() // For logout

	newToken, err := token.SignedString([]byte("SECRET")) //config.Config("SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "Success login", "data": newToken})
}

func Logout(ctx *fiber.Ctx) error {

	token, _ := GetJWT(ctx)
	var validationTrace string
	validationTrace = token.Claims.(jwt.MapClaims)["validation_trace"].(string)

	if !services.InvalidateToken(validationTrace) {
		return ctx.Status(500).JSON(fiber.Map{"message": "Error while trying to log out."})
	}

	return ctx.Status(200).JSON(fiber.Map{"message": "Successfully logged out."})
}

func Authorize(ctx *fiber.Ctx) error {

	token, err := GetJWT(ctx)

	if err != nil {
		return err
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userid"].(string)
	email := claims["email"].(string)
	username := claims["username"].(string)
	validation_trace := claims["validation_trace"].(string)

	if !services.VerifyTokenValidity(validation_trace) {
		err := fiber.NewError(401)
		err.Message = "Invalidated JWT, please try to login again. IP: " + ctx.IP()
		return err
	}

	ctx.Locals("CurrentUserId", userId)
	ctx.Locals("CurrentUserEmail", email)
	ctx.Locals("CurrentUserName", username)
	ctx.Locals("ValidationTrace", validation_trace)

	return nil
}

func GetJWT(ctx *fiber.Ctx) (*jwt.Token, error) {
	authBearer := ctx.Get(fiber.HeaderAuthorization)
	jwtToken := strings.Fields(authBearer)

	if len(jwtToken) < 1 {
		return nil, fiber.NewError(401)
	}

	token, _ := jwt.Parse(jwtToken[1], func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	return token, nil
}
