package routes

import (
	"github.com/sheriff-kurdi/inventory/web/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	authController := controllers.NewAuthController()

	productsRoutes := app.Group("/api/v1/auth")
	productsRoutes.Post("/login", authController.Login)


}
