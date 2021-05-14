package controllers

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/database"
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/models"
	"math"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page","1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64

	var products []models.Product

	database.DB.Offset(offset).Limit(limit).Find(&products) // loading foreign key

	database.DB.Model(&models.Product{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": products, 
		"meta": fiber.Map{
			"total": total,
			"page": page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)
	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product {
		Id: uint(id),
	}
	database.DB.Find(&product)
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product {
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Model(&product).Updates(product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product {
		Id: uint(id),
	}
	database.DB.Delete(&product)
	return nil
}

