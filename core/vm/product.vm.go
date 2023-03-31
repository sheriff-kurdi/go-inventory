package vm

import (
	"kurdi-go/core/entities"
	"kurdi-go/core/entities/products"
)

type Product struct {
	Id           uint              `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	LanguageCode string            `json:"language_code"`
	products.ProductQuantity
	entities.TimeStamps
}
