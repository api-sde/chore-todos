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

	//store := redis.New()
	//redisPing(store)

	log.Fatal(app.Listen(":3000"))
}

// func redisPing(store *redis.Storage) {

// 	store.Set("pong", []byte("ping"), time.Duration(time.Microsecond.Hours()))

// }
