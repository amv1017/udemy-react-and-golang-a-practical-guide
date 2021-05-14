package controllers

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/database"
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page","1"))
	
	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}


