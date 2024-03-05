package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/tools/middleware"
)

func SetupRoutes(app *fiber.App) {
	// admin := fiber.New()
	// admin.Use(middleware.AdminMiddleware)
	// app.Mount("/api/v1", admin)

	route := app.Group("/api/v1")

	route.Post("/order", middleware.AuthMiddleware, PlaceOrder)

	route.Delete("/cancel", middleware.AuthMiddleware, CancelOrder)

	route.Put("/send", middleware.AuthMiddleware, middleware.AdminMiddleware, SendOrder)

	route.Get("/orders", middleware.AuthMiddleware, GetOrders)
}
