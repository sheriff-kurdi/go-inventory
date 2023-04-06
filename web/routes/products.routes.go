package routes

import (
	"kurdi-go/web/controllers"
	"kurdi-go/web/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProductsRoutes(app *fiber.App) {
	productsController := controllers.NewProductsController()

	productsRoutes := app.Group("/api/v1/products")
	productsRoutes.Use(
		middlewares.JWTProtected(),
		middlewares.AuthenticationMiddleware(),
	)
	
	productsRoutes.Post("/", productsController.Insert)
	productsRoutes.Get("/", productsController.GetAll)
	productsRoutes.Get("/:id", productsController.FindById)

}
