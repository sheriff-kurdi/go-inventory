package products

import (
	"kurdi-go/core/entities"

	"gorm.io/gorm"
)

type Product struct {
	Id           uint    `gorm:"primary"`
	CostPrice    float32 `json:"cost_price"`
	SellingPrice float32 `json:"selling_price"`
	Discount     float32 `json:"discount"`
	IsDiscounted bool    `json:"is_discounted"`
	entities.TimeStamps
}

func (product *Product) AfterFind(tx *gorm.DB) (err error) {
	if product.IsDiscounted {
		product.SellingPrice -= product.Discount
	}
	return
}
