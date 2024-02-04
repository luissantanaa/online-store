package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("boas!")
	})

	route.Get("/items", GetItems)

	route.Post("/item", AddItem)

	route.Get("/populate", PopulateItems)

	route.Delete("/delete", DeleteItem)
}
