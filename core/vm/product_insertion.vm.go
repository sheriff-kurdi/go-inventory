package vm

import "kurdi-go/core/models/products"

type ProductInsertionVM struct {
	Details []ProductDetailsInsertionVM `json:"details"`
	products.ProductQuantity
}

type ProductDetailsInsertionVM struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
}
