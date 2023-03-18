package resources

import "github.com/gofiber/fiber/v2"

type NoContent204 struct {
	Status  int
	Message string
	Data    *fiber.Map
}

func NoContent(message string) IResource {
	dataJson := &fiber.Map{
		"success": true,
		"message": message,
	}
	resource := NoContent204{Status: 204, Message: "", Data: dataJson}
	return &resource
}

func (NoContent204 NoContent204) GetStatus() int {
	return NoContent204.Status
}

func (NoContent204 NoContent204) GetData() *fiber.Map {
	return NoContent204.Data
}
