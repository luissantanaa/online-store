package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/client-service/routes"
	"github.com/luissantanaa/online-store/pkg/db"
)

func main() {
	db.ConnectDb()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8082")

}
