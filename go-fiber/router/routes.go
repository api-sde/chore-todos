package router

import (
	"github.com/adrienBdx/chore-todos/gofiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Group("*", logger.New())
	app.Get("/hello", handlers.GetHello)
}
