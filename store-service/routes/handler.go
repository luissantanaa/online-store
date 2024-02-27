package routes

import (
	"log"

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
	item_exists := models.Item{Name: item.Name}
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if item.Name == "" || item.Quantity == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Fields",
		})
	}

	if db.DB.Db.Find(&item_exists) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Item already exists",
		})
	}

	db.DB.Db.Create(&item)
	return c.Status(200).JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	res := db.DB.Db.Find(&item, item.ID)

	if res == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	db.DB.Db.Delete(&item)

	return c.Status(200).JSON(item)
}

func PopulateItems(c *fiber.Ctx) error {
	items := []models.Item{
		{Name: "Bananas", Quantity: 50},
		{Name: "Carrots", Quantity: 50},
		{Name: "Peppers", Quantity: 50},
		{Name: "Oranges", Quantity: 50},
		{Name: "Eggs", Quantity: 50},
		{Name: "Lettuces", Quantity: 50},
		{Name: "Garlic", Quantity: 50},
	}

	result := db.DB.Db.Create(items)
	log.Printf("Result: %v", result)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.Status(200).JSON(items)
}
