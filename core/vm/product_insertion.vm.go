package vm

import "github.com/sheriff-kurdi/inventory/core/models/products"

type ProductSavingVM struct {
	Id      uint                        `gorm:"primary"`
	ProductDetails []ProductDetailsInsertionVM `json:"product_details"`
	products.ProductQuantity
	products.ProductPrice
}

type ProductDetailsInsertionVM struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
}
