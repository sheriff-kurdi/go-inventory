package products

import (
	"kurdi-go/core/entities"
)

type ProductDetails struct {
	entities.Entity
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
	StockItemId  int    `json:"stock_item_id"`
}
