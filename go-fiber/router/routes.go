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
	auth.Post("/logout", handlers.Logout)

	/// All API endpoints ///
	// For unprotected endpoints: api := app.Group("/api")
	api := app.Group("/api",

		func(ctx *fiber.Ctx) error {
			middleware.Protected()

			// will throw if invalid token
			_, err := middleware.GetUserClaims(ctx)
			if err != nil {
				return err
			}

			ctx.Next()
			return nil
	})

	// User
	user := api.Group("/user")
	user.Get("/:email", handlers.GetUser)
	user.Get("/", handlers.GetUsers)
	user.Post("/", handlers.CreateUser)

	// To Dos
	todo := api.Group("/todo")
	todo.Get("/all", handlers.GetAllToDos)
	todo.Get("/user", handlers.GetToDoByUser)
	todo.Get("/:todoId", handlers.GetToDoById)
	todo.Post("/", handlers.CreateToDo)

}
