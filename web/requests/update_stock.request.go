package requests

// import (
// 	stockDomainEntities "github.com/sheriff-kurdi/inventory/domain/domain_entities/stock"
// 	"time"
// )

// type UpdateStockRequest struct {
// 	UpdateStockItemRequest        UpdateStockItemRequest          `json:"stock_item"`
// 	UpdateStockItemDetailsRequest []UpdateStockItemDetailsRequest `json:"stock_item_details"`
// }

// type UpdateStockItemRequest struct {
// 	Id             int     `json:"id"`
// 	CostPrice      float32 `json:"cost_price"`
// 	SellingPrice   float32 `json:"selling_price"`
// 	Discount       float32 `json:"discount"`
// 	TotalStock     int     `json:"total_stock"`
// 	AvailableStock int     `json:"available_stock"`
// 	ReservedStock  int     `json:"reserved_stock"`
// 	IsDiscounted   bool    `json:"is_discounted"`
// }
// type UpdateStockItemDetailsRequest struct {
// 	Name         string `json:"name"`
// 	Description  string `json:"description"`
// 	LanguageCode string `json:"language_code"`
// }

// func (updateStockRequest *UpdateStockRequest) ToStockItemEntity() (stockItem stockDomainEntities.StockItem) {
// 	stockItem.Id = uint(updateStockRequest.UpdateStockItemRequest.Id)
// 	stockItem.CostPrice = updateStockRequest.UpdateStockItemRequest.CostPrice
// 	stockItem.SellingPrice = updateStockRequest.UpdateStockItemRequest.SellingPrice
// 	stockItem.TotalStock = updateStockRequest.UpdateStockItemRequest.TotalStock
// 	stockItem.AvailableStock = updateStockRequest.UpdateStockItemRequest.AvailableStock
// 	stockItem.ReservedStock = updateStockRequest.UpdateStockItemRequest.ReservedStock
// 	stockItem.Discount = updateStockRequest.UpdateStockItemRequest.Discount
// 	stockItem.IsDiscounted = updateStockRequest.UpdateStockItemRequest.IsDiscounted
// 	return
// }

// func (updateStockRequest *UpdateStockRequest) ToStockItemDetailsEntities(stockItemId int) (stockItemDetails []stockDomainEntities.StockItemDetails) {
// 	for _, stockItemDetailsRequest := range updateStockRequest.UpdateStockItemDetailsRequest {
// 		stockItemDetailsEntity := stockDomainEntities.StockItemDetails{}
// 		stockItemDetailsEntity.LanguageCode = stockItemDetailsRequest.LanguageCode
// 		stockItemDetailsEntity.Name = stockItemDetailsRequest.Name
// 		stockItemDetailsEntity.StockItemId = stockItemId
// 		stockItemDetailsEntity.Description = stockItemDetailsRequest.Description
// 		stockItemDetailsEntity.UpdatedAt = time.Now()
// 		stockItemDetails = append(stockItemDetails, stockItemDetailsEntity)
// 	}
// 	return
// }
