package responses

type StockResponse struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	CostPrice      float32 `json:"costPrice"`
	SellingPrice   float32 `json:"sellingPrice"`
	Discount       float32 `json:"discount"`
	TotalStock     int     `json:"totalStock"`
	AvailableStock int     `json:"availableStock"`
	ReservedStock  int     `json:"reservedStock"`
	IsDiscounted   bool    `json:"discounted"`
}
