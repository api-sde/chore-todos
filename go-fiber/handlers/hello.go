package handlers

import (
	"github.com/adrienBdx/chore-todos/gofiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetHello(ctx *fiber.Ctx) error {

	helloUser := models.User{
		Ip:   ctx.IP(),
		Name: "Hello Bob",
	}

	return ctx.JSON(helloUser)
}
