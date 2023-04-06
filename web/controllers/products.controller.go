package controllers

import (
	"kurdi-go/core/services"
	"kurdi-go/core/vm"
	"kurdi-go/web/resources"
	"kurdi-go/web/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductsController struct {
	productsService services.AuthService
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
// @Success 200 {array} vm.ProductVM
// @Security ApiKeyAuth
// @Router /api/v1/products [get]
func (controller ProductsController) GetAll(ctx *fiber.Ctx) error {
	languageCode := ctx.Query("language_code")
	products := controller.productsService.ListAll(languageCode)
	response := resources.Ok(products, "")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// @Description Get product.
// @Summary get product.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product Id"
// @Success 200 {object} vm.ProductVM
// @Security ApiKeyAuth
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

// @Description Create Product.
// @Summary Create Product
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param productVM body vm.ProductSavingVM true "Product VM"
// @Success 200 {object} vm.ProductSavingVM
// @Security ApiKeyAuth
// @Router /api/v1/products/ [post]
func (controller ProductsController) Insert(ctx *fiber.Ctx) error {

	productVM := vm.ProductSavingVM{}
	//binding
	if err := ctx.BodyParser(&productVM); err != nil {
		utils.Logger().Info(err.Error())
		response := resources.BadRequest("GENERAL.THERE_IS_AN_ERROR")
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	//find by id
	productId, err := controller.productsService.Insert(productVM)
	var response resources.IResource

	if err != nil {
		response := resources.ServerError("GENERAL.THERE_IS_AN_ERROR")
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	response = resources.Ok(productId, "Created Successfully")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}
