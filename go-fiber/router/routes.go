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

	app.Get("/redis-set", handlers.SetHelloRedis)
	app.Get("/redis-get", handlers.GetHelloRedis)
	app.Get("/redis-clear", handlers.ClearRedis)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", handlers.Login)

	// User
	user := app.Group("/user")
	user.Get("/:email", handlers.GetUser)
	user.Get("/", handlers.GetUsers)
	user.Post("/", handlers.CreateUser)

	// ToDo
	todo := app.Group("/todo")

	todo.Get("/all", handlers.GetAllToDos)
	todo.Get("/user", func(ctx *fiber.Ctx) error {
		middleware.Protected()
		middleware.GetUserClaims(ctx)
		ctx.Next()
		return nil
	}, handlers.GetToDoByUser)

	todo.Get("/:todoId", handlers.GetToDoById)

	todo.Post("/", func(ctx *fiber.Ctx) error {
		middleware.Protected()
		middleware.GetUserClaims(ctx)
		ctx.Next()
		return nil
	}, handlers.CreateToDo)

}
