package main

import (
	"log"

	"github.com/adrienBdx/chore-todos/gofiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	// - ToDos:
	//app.Use(cors.New())
	// database
	// redis

	log.Fatal(app.Listen(":3000"))
}
