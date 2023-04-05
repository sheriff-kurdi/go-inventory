package routes

import (
	"kurdi-go/web/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	authController := controllers.NewAuthController()

	productsRoutes := app.Group("/api/v1/auth")
	productsRoutes.Post("/login", authController.Login)


}
