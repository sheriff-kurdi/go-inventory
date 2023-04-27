package routes

import (
	"github.com/sheriff-kurdi/inventory/web/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductsRoutes(app *fiber.App) {
	productsController := controllers.NewProductsController()

	productsRoutes := app.Group("/api/v1/products")
	// productsRoutes.Use(
	// 	middlewares.JWTProtected(),
	// 	middlewares.AuthenticationMiddleware(),
	// )

	productsRoutes.Post("/", productsController.Save)
	productsRoutes.Get("/", productsController.GetAll)
	productsRoutes.Get("/:id", productsController.FindById)
	productsRoutes.Delete("/:id", productsController.DeleteById)


}
