package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/pkg/db"
	"github.com/luissantanaa/online-store/store-service/routes"
)

func main() {
	db.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8081")

}
