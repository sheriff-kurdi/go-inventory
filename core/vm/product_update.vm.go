package vm

import "github.com/sheriff-kurdi/inventory/core/models/products"

type ProductUpdateVM struct {
	Details []ProductDetailsUpdateVM `json:"details"`
	products.ProductQuantity
}

type ProductDetailsUpdateVM struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
}
