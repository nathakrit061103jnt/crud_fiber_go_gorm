package controllers

import (
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/database"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetProductsAll(c *fiber.Ctx) error {
	db := database.DBConn
	var product []models.Product
	if err := db.Find(&product).Error; err != nil {
		// error handling...
		return err
	}
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var product []models.Product
	if err := db.Find(&product, id).Error; err != nil {
		// error handling...
		return err
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {

	db := database.DBConn

	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return err
	}

	if err := db.Create(&product).Error; err != nil {
		// error handling...
		return err
	}

	// db.Create(&product)

	return c.JSON(product)

}

func UpdateProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DBConn
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString("No Product Found with ID")
	}

	if err := db.Model(product).Where("id = ?", id).Updates(product).Error; err != nil {
		// error handling...
		return err
	}

	return c.Status(200).SendString("Update Product Successfully")

}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var product models.Product

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
			return err
		}
		// Delete permanently
		// db.Unscoped().Delete(&product, id)
	}

	return c.SendString("Product Successfully deleted")

}
