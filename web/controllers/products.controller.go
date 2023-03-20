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

// GetAll GET /products
// GetAll all products
// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  []products.Product
// @Failure      400  {object}  resources.IResource
// @Failure      404  {object}  resources.IResource
// @Failure      500  {object}  resources.IResource
// @Router       /products/{id} [get]
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
