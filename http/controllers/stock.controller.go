package controllers

import (
	"github.com/gofiber/fiber/v2"
	"kurdi-go/helpers/resources"
	"kurdi-go/http/requests"
	"kurdi-go/services"
	"strconv"
)

type StockController struct {
	service *services.StockService
}

func NewStockController() *StockController {
	controller := StockController{
		service: services.NewStockService(),
	}
	return &controller
}

// GetAll GET /books
// GetAll all books

func (controller StockController) GetAll(ctx *fiber.Ctx) error {
	//get all
	languageCode := ctx.GetReqHeaders()["language"]
	response := controller.service.ListAll(languageCode)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// FindById GET /books/:id
// Find a book
func (controller StockController) FindById(ctx *fiber.Ctx) error {
	bookId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//find by id
	languageCode := ctx.GetReqHeaders()["language"]
	response := controller.service.FindById(bookId, languageCode)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Create POST /books
// Create new book
func (controller StockController) Create(ctx *fiber.Ctx) error {
	var request requests.CreateStockRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//create
	response := controller.service.Create(request)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Update PATCH /books/:id
// Update a book
func (controller StockController) Update(ctx *fiber.Ctx) error {
	//check param
	stockItemId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	// Validate request
	var request requests.UpdateStockRequest
	if err := ctx.BodyParser(&request); err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//update
	if request.UpdateStockItemRequest.Id != stockItemId {
		response := resources.GetError400Resource("stock item id should be equal request id")
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	response := controller.service.Update(request, stockItemId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}

// Delete DELETE /books/:id
// Delete a book
func (controller StockController) Delete(ctx *fiber.Ctx) error {
	//check param
	bookId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		response := resources.GetError500Resource(err.Error())
		return ctx.Status(response.GetStatus()).JSON(response.GetData())
	}
	//delete
	response := controller.service.Delete(bookId)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())

}
