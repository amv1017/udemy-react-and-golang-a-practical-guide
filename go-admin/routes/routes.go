package routes

import (
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/controllers"
	"github.com/amv1017/udemy-building-modern-web-apps-with-golang/go-admin/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login",    controllers.Login)
	app.Use(middlewares.IsAuthenticated)
	app.Get( "/api/user",     controllers.User)
	app.Post("/api/logout",   controllers.Logout)
}

