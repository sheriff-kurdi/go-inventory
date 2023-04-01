package controllers

import (
	"kurdi-go/core/services"
	"kurdi-go/web/resources"
	"kurdi-go/web/utils"
	"strconv"
	"time"

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
	//get all
	//s := products.Product{}
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		// Return status 500 and JWT parse error.
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	languageCode := ctx.Query("language_code")
	products := controller.productsService.ListAll(languageCode)
	response := resources.Ok(products, "")
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
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
