package requests

type CreateStockRequest struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	CostPrice      float32 `json:"cost_price"`
	SellingPrice   float32 `json:"selling_price"`
	Discount       float32 `json:"discount"`
	TotalStock     int     `json:"total_stock"`
	AvailableStock int     `json:"available_stock"`
	ReservedStock  int     `json:"reserved_stock"`
	IsDiscounted   bool    `json:"is_discounted"`
}
