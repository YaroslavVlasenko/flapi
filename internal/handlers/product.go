package handlers

import (
	"app-backend/internal/database"
	"app-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts query all products
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.DB.Preload("Translations.Locale").Find(&products)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// GetProduct query product by ID
func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")

	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

// CreateProduct create a new product
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}
	database.DB.Create(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

// DeleteProduct delete product
func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")

	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	database.DB.Delete(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
