package resources

import (
	"github.com/gofiber/fiber/v2"
)

type Error500 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func ServerError(message string) IResource {
	errors := make(map[string]interface{})
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    errors,
	}
	// logger, _ := zap.NewProduction()
	// defer logger.Sync()
	// logger.Error("Internal Server Error", zap.Any(message, errors))
	resource := Error500{Status: 500, Message: "", Data: dataJson}
	return &resource
}

func (error500 Error500) GetStatus() int {
	return error500.Status
}

func (error500 Error500) GetData() *fiber.Map {
	return error500.Data
}
