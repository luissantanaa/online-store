package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luissantanaa/online-store/models"
	"github.com/luissantanaa/online-store/pkg/db"
)

func PlaceOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	item_exists := models.Item{}
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if order.ClientID == 0 || order.Items == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Fields",
		})
	}

	for _, item := range order.Items {
		if item.ItemID == 0 || item.Quantity <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid Item or quantity",
			})
		}
		db.DB.Db.Table("items").Find(&item_exists, "id=?", item.ItemID)

		if item_exists.ID != item.ItemID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Item does not exist",
			})
		}

		if item_exists.Quantity < item.Quantity {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("Sorry, not enough stock for item %v", item_exists.Name),
			})
		}
	}

	db.DB.Db.Create(&order)
	return c.Status(200).JSON(order)
}

func CancelOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if order.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No order ID passed",
		})
	}

	res := db.DB.Db.Find(&order, order.ID)

	if res == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	db.DB.Db.Unscoped().Delete(&order)
	return c.Status(200).JSON(order)
}

func SendOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if order.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No order ID passed",
		})
	}

	res := db.DB.Db.Find(&order, order.ID)
	if res == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	order.Sent = true
	db.DB.Db.Save(&order)
	return c.Status(200).JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	db.DB.Db.Find(&orders)
	return c.Status(200).JSON(orders)
}
