package routes

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/http/controllers"
	"kurdi-go/http/middlewares"
)

func StockRoutes(router *fiber.App) {
	bookController := controllers.NewStockController()

	stockApp := router.Group("/stock")
	stockApp.Use(middlewares.AuthenticationMiddleware)

	stockApp.Get("/", bookController.GetAll)
	stockApp.Get("/:id", bookController.FindById)
	stockApp.Post("/", bookController.Create)
	stockApp.Patch("/:id", bookController.Update)
	stockApp.Delete("/:id", bookController.Delete)

}
