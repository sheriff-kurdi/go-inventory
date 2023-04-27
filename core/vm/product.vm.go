package vm

import (
	"github.com/sheriff-kurdi/inventory/core/models"
	"github.com/sheriff-kurdi/inventory/core/models/products"
)

type ProductVM struct {
	Id           uint   `json:"id"`
	products.ProductQuantity
	products.ProductPrice
	ProductDetails []ProductDetailsVM `json:"product_details"`
	models.TimeStamps
}

type ProductDetailsVM struct {
	Name         string               `json:"name"`
    Description  string               `json:"description"`
    LanguageCode string               `gorm:"primaryKey" json:"language_code"`
    models.TimeStamps
}
