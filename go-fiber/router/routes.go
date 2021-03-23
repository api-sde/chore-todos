package router

import (
	"github.com/adrienBdx/chore-todos/gofiber/handlers"
	"github.com/adrienBdx/chore-todos/gofiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Group("*", logger.New())
	app.Get("/hello", handlers.GetHello)
	app.Get("/hello-protected", middleware.Protected(), handlers.GetHello)

	// Auth
	auth := app.Group("/auth")
	auth.Get("/login", handlers.Login)

}
