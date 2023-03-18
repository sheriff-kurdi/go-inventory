package products

import (
	"kurdi-go/core/entities"
	"gorm.io/gorm"
)

type Product struct {
	entities.Entity
	CostPrice      float32 `json:"cost_price"`
	SellingPrice   float32 `json:"selling_price"`
	Discount       float32 `json:"discount"`
	TotalStock     int     `json:"total_stock"`
	AvailableStock int     `json:"available_stock"`
	ReservedStock  int     `json:"reserved_stock"`
	IsDiscounted   bool    `json:"is_discounted"`
}

func (product *Product) AfterFind(tx *gorm.DB) (err error) {
	if product.IsDiscounted {
		product.SellingPrice -= product.Discount
	}
	return
}
