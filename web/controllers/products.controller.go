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

// @Description Get all products.
// @Summary get all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} products.Product
// @Router /api/v1/products [get]
func (controller ProductsController) GetAll(ctx *fiber.Ctx) error {
	//get all
	//s := products.Product{}
	languageCode := ctx.Query("language_code")
	products := controller.productsService.ListAll(languageCode)
	response := resources.Ok(products, "")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} products.Product
// @Router /api/v1/products/{id} [get]
func (controller ProductsController) FindById(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	languageCode := ctx.Query("language_code")

	if err != nil {
		response := resources.ServerError(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//find by id
	product := controller.productsService.FindById(productId, languageCode)
	var response resources.IResource

	if product == nil {
		response = resources.NotFound("")
	} else {
		response = resources.Ok(product, "")

	}
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
