package vm

import (
	"kurdi-go/core/models"
	"kurdi-go/core/models/products"
)

type ProductVM struct {
	Id           uint              `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	LanguageCode string            `json:"language_code"`
	products.ProductQuantity
	models.TimeStamps
}
