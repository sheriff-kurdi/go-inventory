package vm

import (
	"github.com/sheriff-kurdi/inventory/core/models"
	"github.com/sheriff-kurdi/inventory/core/models/products"
)

type ProductVM struct {
	Id           uint              `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	LanguageCode string            `json:"language_code"`
	products.ProductQuantity
	models.TimeStamps
}
