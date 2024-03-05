package middleware

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
			"error":   err.Error(),
		})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	id, err := strconv.Atoi(fmt.Sprint(claims["sub"]))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error processing sub claim",
		})
	}

	client := models.Client{ID: uint(id)}
	db.DB.Db.Find(&client)

	// Extract user information from the token and store it in the context
	c.Locals("user", client)
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	client := c.Locals("user")
	details, ok := client.(models.Client)

	if details.Role != "admin" || !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Invalid Permissions",
			"role":    details,
		})
	}

	return c.Next()
}
