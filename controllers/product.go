package controllers

import (
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/database"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetProductsAll(c *fiber.Ctx) error {
	db := database.DBConn
	product := []models.Product{}
	if err := db.Find(&product).Error; err != nil {
		// error handling...
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	product := []models.Product{}
	if err := db.Find(&product, id).Error; err != nil {
		// error handling...
		return c.Status(400).JSON(err.Error())
	}

	if len(product) == 0 {
		return c.Status(400).SendString("ไม่พบข้อมูล")
	}

	return c.Status(200).JSON(product)

}

func CreateProduct(c *fiber.Ctx) error {

	db := database.DBConn

	product := models.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := db.Create(&product).Error; err != nil {
		// error handling...
		return c.Status(400).JSON(err.Error())
	}

	// db.Create(&product)

	return c.Status(201).JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBConn
	product := models.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).SendString("No Product Found with ID")
	}

	if err := db.Model(&product).Where("id = ?", id).Updates(&product).Error; err != nil {
		// error handling...
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(&product)

}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	product := models.Product{}

	if err := db.First(&product, id).Error; err != nil {
		// error handling...
		return err
	} else {
		if product.P_Name == "" {
			return c.Status(500).SendString("No Product Found with ID")

		}
		// Soft  permanently
		if err := db.Delete(&product, id).Error; err != nil {
			// error handling...
			return c.Status(400).JSON(err.Error())
		}
		// Delete permanently
		// db.Unscoped().Delete(&product, id)
	}

	return c.Status(200).SendString("Product Successfully deleted")

}
