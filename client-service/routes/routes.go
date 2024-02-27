package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("boas!")
	})

	route.Get("/clients", GetClients)

	route.Delete("/delete", DeleteClient)

	route.Post("/signup", AddClient)

	//route.Get("/login", Login)
}
