package stock_domain_entities

import "kurdi-go/domain/domain_entities"

type StockItemDetails struct {
	domain_entities.Entity
	Name         string `json:"name"`
	Description  string `json:"description"`
	LanguageCode string `json:"language_code"`
	StockItemId  string `json:"stock_item_id"`
}
