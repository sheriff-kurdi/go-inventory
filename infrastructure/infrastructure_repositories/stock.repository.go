package infrastructure_repositories

import (
	"kurdi-go/http/responses"
	"kurdi-go/infrastructure/infrastructure_database"
)

type StockRepository struct {
}

func NewStockRepository() *StockRepository {
	repository := StockRepository{}
	return &repository
}

func (repository StockRepository) ListAll(languageCode string) (stockList []responses.StockResponse, err error) {
	if languageCode == "" {
		languageCode = "ar"
	}
	ListAllByLanguageQuery := `SELECT * FROM stock_items 
								JOIN stock_item_details ON stock_item_details.id = stock_items.id AND stock_item_details.language_code = ?`

	err = infrastructure_database.PostgresDB.Raw(ListAllByLanguageQuery, languageCode).Scan(&stockList).Error
	return
}
