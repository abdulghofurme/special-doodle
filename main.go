package main

import (
	"log"

	"github.com/abdulghofurme/special-doodle/database"
	"github.com/abdulghofurme/special-doodle/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
