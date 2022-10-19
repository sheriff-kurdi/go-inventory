package infrastructure_repositories

import (
	"kurdi-go/http/responses"
	"kurdi-go/infrastructure/infrastructure_database"
	"os"
)

type StockRepository struct {
}

func NewStockRepository() *StockRepository {
	repository := StockRepository{}
	return &repository
}

func (repository StockRepository) ListAll(languageCode string) (stockList []responses.StockResponse, err error) {
	if languageCode == "" {
		languageCode = os.Getenv("DEFAULT_LANGUAGE")
	}
	ListAllByLanguageQuery := `SELECT * FROM stock_items 
								JOIN stock_item_details ON stock_item_details.stock_item_id	 = stock_items.id AND stock_item_details.language_code = ?`

	err = infrastructure_database.PostgresDB.Raw(ListAllByLanguageQuery, languageCode).Scan(&stockList).Error
	return
}

func (repository StockRepository) FindById(stockItemId int, languageCode string) (stockList []responses.StockResponse, err error) {
	if languageCode == "" {
		languageCode = os.Getenv("DEFAULT_LANGUAGE")
	}
	ListAllByLanguageQuery := `SELECT * FROM stock_items 
								JOIN stock_item_details ON stock_item_details.stock_item_id	 = stock_items.id AND stock_item_details.language_code = ? WHERE stock_items.id = ?`

	err = infrastructure_database.PostgresDB.Raw(ListAllByLanguageQuery, languageCode, stockItemId).Scan(&stockList).Error
	return
}
