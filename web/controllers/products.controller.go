package controllers

import (
	postgresDatabse "github.com/sheriff-kurdi/inventory/infrastructure/database/postgres"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sheriff-kurdi/inventory/core/services"
	"github.com/sheriff-kurdi/inventory/core/vm"
	"github.com/sheriff-kurdi/inventory/infrastructure/repositories/postgres"
	"github.com/sheriff-kurdi/inventory/web/resources"
	"github.com/sheriff-kurdi/inventory/web/utils"
)

type ProductsController struct {
	productsService services.ProductsService
}

func NewProductsController() *ProductsController {
	controller := ProductsController{
		productsService: services.NewProductsService(postgres.NewProductsRepository(postgresDatabse.Connect()), postgresDatabse.Connect()),
	}
	return &controller
}

// @Description Get all products.
// @Summary get all products
// @Tags Products
// @Accept json
// @Produce json
// @Param language_code query string false "Language Code"
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
// @Param language_code query string false "Language Code"
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
	product, err := controller.productsService.FindById(productId, languageCode)
	var response resources.IResource

	if err != nil {
		response = resources.ServerError(err.Error())
	} else if product == nil {
		response = resources.NotFound("")
	} else {
		response = resources.Ok(product, "")

	}
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// @Description Delete product.
// @Summary Delete product.
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Product Id"
// @Success 200 {object} vm.ProductVM
// @Security ApiKeyAuth
// @Router /api/v1/products/{id} [delete]
func (controller ProductsController) DeleteById(ctx *fiber.Ctx) error {

	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.ServerError(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//delete by id
	err = controller.productsService.DeleteById(productId)
	if err != nil {
		response := resources.ServerError(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	response := resources.Ok(productId, "Deleted Successfully")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// @Description Save Product.
// @Summary Save Product
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param productVM body vm.ProductSavingVM true "Product VM"
// @Success 201 {object} vm.ProductSavingVM
// @Security ApiKeyAuth
// @Router /api/v1/products/ [post]
func (controller ProductsController) Save(ctx *fiber.Ctx) error {

	productVM := vm.ProductSavingVM{}
	//binding
	if err := ctx.BodyParser(&productVM); err != nil {
		utils.Logger().Info(err.Error())
		response := resources.BadRequest("GENERAL.THERE_IS_AN_ERROR")
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	//find by id
	productId, err := controller.productsService.Save(productVM)
	var response resources.IResource

	if err != nil {
		response := resources.ServerError("GENERAL.THERE_IS_AN_ERROR")
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}

	response = resources.Ok(productId, "Saved Successfully")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}
