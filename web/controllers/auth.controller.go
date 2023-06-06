package controllers

import (
	"github.com/sheriff-kurdi/inventory/core/services"
	"github.com/sheriff-kurdi/inventory/core/vm"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	productsService services.AuthService
}

func NewAuthController() *AuthController {
	controller := AuthController{
		productsService: services.NewAuthService(),
	}
	return &controller
}

// @Description Login and get token.
// @Summary Login and get token
// @Tags Auth
// @Accept json
// @Produce json
// @Param lognVM body  vm.LoginVM true "Login View Model"
// @Success 200 {array} vm.LoginVM
// @Router /api/v1/auth/login [post]
func (controller AuthController) Login(ctx *fiber.Ctx) error {
	
	lognVM := vm.LoginVM{}

	// Check, if received JSON data is valid.
	if err := ctx.BodyParser(&lognVM); err != nil {
		// Return status 400 and error message.
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response := controller.productsService.Login(&lognVM)
	return ctx.Status(response.GetStatus()).JSON(response.GetData())
}