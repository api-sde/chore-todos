package main

import (
	"log"

	"github.com/adrienBdx/chore-todos/gofiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	router.SetupRoutes(app)

	// - ToDos:
	//app.Use(cors.New())

	//store := redis.New()
	//redisPing(store)

	log.Fatal(app.Listen(":3000"))
}

// func redisPing(store *redis.Storage) {

// 	store.Set("pong", []byte("ping"), time.Duration(time.Microsecond.Hours()))

// }
