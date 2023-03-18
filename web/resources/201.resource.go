package resources

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Success201 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func GetSuccess201Resource(data interface{}, message string) IResource {

	dataJson := &fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	}

	resource := Success201{Status: http.StatusCreated, Message: "", Data: dataJson}
	return &resource
}

func (success201 Success201) GetStatus() int {
	return success201.Status
}

func (success201 Success201) GetData() *fiber.Map {
	return success201.Data
}
