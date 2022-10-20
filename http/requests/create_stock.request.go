package requests

import (
	stockDomainEntities "kurdi-go/domain/domain_entities/stock"
	"time"
)

type CreateStockRequest struct {
	CreateStockItemRequest        CreateStockItemRequest          `json:"stock_item"`
	CreateStockItemDetailsRequest []CreateStockItemDetailsRequest `json:"stock_item_details"`
}

type CreateStockItemRequest struct {
	CostPrice      float32 `json:"cost_price"`
	SellingPrice   float32 `json:"selling_price"`
	Discount       float32 `json:"discount"`
	TotalStock     int     `json:"total_stock"`
	AvailableStock int     `json:"available_stock"`
	ReservedStock  int     `json:"reserved_stock"`
	IsDiscounted   bool    `json:"is_discounted"`
}
type CreateStockItemDetailsRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
}

func (createStockRequest *CreateStockRequest) ToStockItemEntity() (stockItem stockDomainEntities.StockItem) {
	stockItem.CostPrice = createStockRequest.CreateStockItemRequest.CostPrice
	stockItem.SellingPrice = createStockRequest.CreateStockItemRequest.SellingPrice
	stockItem.TotalStock = createStockRequest.CreateStockItemRequest.TotalStock
	stockItem.AvailableStock = createStockRequest.CreateStockItemRequest.AvailableStock
	stockItem.ReservedStock = createStockRequest.CreateStockItemRequest.ReservedStock
	stockItem.Discount = createStockRequest.CreateStockItemRequest.Discount
	stockItem.IsDiscounted = createStockRequest.CreateStockItemRequest.IsDiscounted
	return
}

func (createStockRequest *CreateStockRequest) ToStockItemDetailsEntities(stockItemId int) (stockItemDetails []stockDomainEntities.StockItemDetails) {
	for _, stockItemDetailsRequest := range createStockRequest.CreateStockItemDetailsRequest {
		stockItemDetailsEntity := stockDomainEntities.StockItemDetails{}
		stockItemDetailsEntity.LanguageCode = stockItemDetailsRequest.LanguageCode
		stockItemDetailsEntity.Name = stockItemDetailsRequest.Name
		stockItemDetailsEntity.StockItemId = stockItemId
		stockItemDetailsEntity.Description = stockItemDetailsRequest.Description
		stockItemDetailsEntity.CreatedAt = time.Now()
		stockItemDetailsEntity.UpdatedAt = time.Time{}
		stockItemDetailsEntity.DeletedAt = time.Time{}
		stockItemDetails = append(stockItemDetails, stockItemDetailsEntity)
	}
	return
}
