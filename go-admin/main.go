package main

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/database"
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/routes"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))

}

