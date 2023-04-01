package routes

import (
	"kurdi-go/web/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductsRoutes(app *fiber.App) {
	productsController := controllers.NewProductsController()

	productsRoutes := app.Group("/api/v1/products")
	productsRoutes.Get("/", productsController.GetAll)
	productsRoutes.Get("/:id", productsController.FindById)

}
