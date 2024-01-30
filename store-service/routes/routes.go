package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("boas!")
	})

	app.Get("/items", GetItems)

	app.Post("/item", AddItem)

	app.Get("/populate", PopulateItems)

	app.Delete("/delete", DeleteItem)
}
