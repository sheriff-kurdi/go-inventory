package stock_domain_entities

import (
	"gorm.io/gorm"
	"kurdi-go/domain/domain_entities"
)

type StockItem struct {
	domain_entities.Entity
	CostPrice      float32 `json:"cost_price"`
	SellingPrice   float32 `json:"selling_price"`
	Discount       float32 `json:"discount"`
	TotalStock     int     `json:"total_stock"`
	AvailableStock int     `json:"available_stock"`
	ReservedStock  int     `json:"reserved_stock"`
	IsDiscounted   bool    `json:"is_discounted"`
}

func (stockItem *StockItem) AfterFind(tx *gorm.DB) (err error) {
	if stockItem.IsDiscounted {
		stockItem.SellingPrice -= stockItem.Discount
	}
	return
}
