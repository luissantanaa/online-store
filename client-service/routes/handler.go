package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	client := new(models.Client)
	client_exists := models.Client{Username: client.Username}
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db.DB.Db.Find(&client_exists)

	if client_exists.Username != client.Username {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username does not exist",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(client_exists.Password), []byte(client.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid username or Password", "pass": client_exists.Password})
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	dur, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Error parsing jwt duration"})
	}

	claims["sub"] = client_exists.ID
	claims["exp"] = now.Add(dur).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

func GetClients(c *fiber.Ctx) error {
	clients := []models.Client{}
	db.DB.Db.Find(&clients)
	return c.Status(200).JSON(omitInfoClients(clients))
}

func AddClient(c *fiber.Ctx) error {
	client := new(models.Client)
	client_exists := models.Client{Username: client.Username}
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if client.Username == "" || client.Password == "" || client.Role != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Fields",
		})
	}

	db.DB.Db.Find(&client_exists)

	if client_exists.Username == client.Username {
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
	return c.Status(200).JSON(client.OmitInfo())
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

	return c.Status(200).JSON(client.OmitInfo())
}

func omitInfoClients(clients []models.Client) []models.AuxClient {
	var auxClients []models.AuxClient
	for _, client := range clients {
		auxClients = append(auxClients, client.OmitInfo())
	}

	return auxClients
}
