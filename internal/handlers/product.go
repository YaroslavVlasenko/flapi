package handlers

import (
	"fmt"
	"github.com/YaroslavVlasenko/flapi/internal/database"
	"github.com/YaroslavVlasenko/flapi/internal/models"
	"github.com/YaroslavVlasenko/flapi/internal/responses"
	"github.com/YaroslavVlasenko/flapi/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"strconv"
)

// GetAllProducts query all products
func GetAllProducts(c *fiber.Ctx) error {
	var products *[]models.Product
	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit

	var count int64
	database.DB.Preload("Translations.Locale").Find(&products).Count(&count)
	//paginationInfo := utils.GeneratePaginationInfo(count, pagination)

	database.DB.Offset(offset).Limit(pagination.Limit).Order(pagination.Sort).Preload("Translations.Locale").Find(&products)

	return c.JSON(responses.Success("success", "All products", products))
}

// GetProduct query product by ID
func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Print("Error loading param: id, from Ctx.Params.")
	}

	if err = database.DB.Preload("Translations.Locale").First(&product, id).Error; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("error", "No product found with ID", err.Error()))
	}

	return c.JSON(responses.Success("success", "Product found", product))
}

// CreateProduct create a new product
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("error", "Couldn't create product", err.Error()))
	}
	database.DB.Create(&product)
	return c.JSON(responses.Success("success", "Created product", product))
}

// UpdateProduct update product
func UpdateProduct(c *fiber.Ctx) error {
	var upi models.UpdateProductInput
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Print("Error loading param: id, from Ctx.Params.")
	}

	if err := c.BodyParser(&upi); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.Error("error", "Review your input data.", err.Error()))
	}

	var product models.Product

	database.DB.First(&product, id)
	_ = copier.Copy(&product, &upi)
	database.DB.Save(&product)

	return c.JSON(responses.Success("success", "Product successfully updated.", product))
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
