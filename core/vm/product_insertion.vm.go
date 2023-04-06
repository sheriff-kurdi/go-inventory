package vm

import "kurdi-go/core/models/products"

type ProductSavingVM struct {
	Id      uint                        `gorm:"primary"`
	Details []ProductDetailsInsertionVM `json:"details"`
	products.ProductQuantity
	products.ProductPrice
}

type ProductDetailsInsertionVM struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
}
