package routes

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login",    controllers.Login)
	app.Get( "/api/user",     controllers.User)
}

