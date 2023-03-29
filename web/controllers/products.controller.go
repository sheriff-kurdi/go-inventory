package controllers

import (
	"kurdi-go/core/services"
	"kurdi-go/web/resources"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct {
	productsService services.ProductsService
}

func NewProductsController() *ProductsController {
	controller := ProductsController{
		productsService: services.NewProductsService(),
	}
	return &controller
}

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} products.Product
// @Router /v1/products [get]
func (controller ProductsController) GetAll(ctx *fiber.Ctx) error {
	//get all
	//s := products.Product{}
	products := controller.productsService.ListAll()
	response := resources.Ok(products, "")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// FindById GET /products/:id
// Find a product
func (controller ProductsController) FindById(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.ServerError(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//find by id
	product := controller.productsService.FindById(productId)
	var response resources.IResource

	if product == nil {
		response = resources.NotFound("")
	} else {
		response = resources.Ok(product, "")

	}
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
