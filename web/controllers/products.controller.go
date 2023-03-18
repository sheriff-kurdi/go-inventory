package controllers

import (
	"kurdi-go/core/services"
	"kurdi-go/infrastructure/repositories/postgres"
	"kurdi-go/web/resources"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct {
	productsService services.ProductsService
}

func NewProductsController() *ProductsController {
	controller := ProductsController{
		productsService: services.NewProductsService(postgres.NewProductsRepository()),
	}
	return &controller
}

// GetAll GET /books
// GetAll all books

func (controller ProductsController) GetAll(ctx *fiber.Ctx) error {
	//get all
	products := controller.productsService.ListAll()
	response := resources.Ok(products, "")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}
