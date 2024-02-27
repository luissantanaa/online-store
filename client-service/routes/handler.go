package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

// func login() {

// }

func GetClients(c *fiber.Ctx) error {
	clients := []models.Client{}
	db.DB.Db.Find(&clients)
	return c.Status(200).JSON(clients)
}

func AddClient(c *fiber.Ctx) error {
	client := new(models.Client)
	client_exists := models.Client{Username: client.Username}
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if client.Username == "" || client.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Fields",
		})
	}

	if db.DB.Db.Find(&client_exists) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	hashed, hash_error := bcrypt.GenerateFromPassword([]byte(client.Password), 8)
	if hash_error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error hashing password",
		})
	}

	client.Password = string(hashed)

	db.DB.Db.Create(&client)
	return c.Status(200).JSON(client.String())
}

func DeleteClient(c *fiber.Ctx) error {
	client := new(models.Client)
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	res := db.DB.Db.Find(&client, client.ID)

	if res == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Client not found",
		})
	}

	db.DB.Db.Delete(&client)

	return c.Status(200).JSON(client)
}
