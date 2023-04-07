package products

import ()

type ProductPrice struct {
	CostPrice    float32 `json:"cost_price"`
	SellingPrice float32 `json:"selling_price"`
	Discount     float32 `json:"discount"`
	IsDiscounted bool    `json:"is_discounted"`
}
