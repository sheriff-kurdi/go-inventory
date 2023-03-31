package routes

import (
	"kurdi-go/web/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductsRoutes(router *fiber.App) {
	productsController := controllers.NewProductsController()

	productsRoutes := router.Group("/api/v1/products")
	productsRoutes.Get("/", productsController.GetAll)
	productsRoutes.Get("/:id", productsController.FindById)

}
