package main

import (
	"log"

	"github.com/adrienBdx/chore-todos/gofiber/persistence"
	"github.com/adrienBdx/chore-todos/gofiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	persistence.ConnectRedis()

	router.SetupRoutes(app)

	// - ToDos:
	//app.Use(cors.New())

	log.Fatal(app.Listen(":3000"))
}
