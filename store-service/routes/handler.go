package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"
)

func GetItems(c *fiber.Ctx) error {
	items := []models.Item{}
	db.DB.Db.Find(&items)
	return c.Status(200).JSON(items)
}

func AddItem(c *fiber.Ctx) error {
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	db.DB.Db.Create(&item)
	return c.Status(200).JSON(item)
}
