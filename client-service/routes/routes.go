package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/client-service/middleware"
)

func SetupRoutes(app *fiber.App) {
	// admin := fiber.New()
	// admin.Use(middleware.AdminMiddleware)
	// app.Mount("/api/v1", admin)

	route := app.Group("/api/v1")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("boas!")
	})

	route.Get("/clients", middleware.AuthMiddleware, middleware.AdminMiddleware, GetClients)

	route.Delete("/delete", middleware.AuthMiddleware, middleware.AdminMiddleware, DeleteClient)

	route.Post("/signup", AddClient)

	route.Post("/login", Login)
}
