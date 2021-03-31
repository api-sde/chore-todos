package handlers

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "Bob"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("SECRET")) //config.Config("SECRET")))
	if err != nil {
		//return c.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
