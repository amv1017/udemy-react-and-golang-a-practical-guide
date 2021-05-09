package controllers

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
		LastName: "Doe",
	}

	return c.JSON(user)
}

