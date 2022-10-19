package infrastructure_repositories

import (
	"kurdi-go/domain/domain_entities/stock"
	"kurdi-go/http/responses"
	"kurdi-go/infrastructure/infrastructure_database"
)

type StockRepository struct {
}

func NewStockRepository() *StockRepository {
	repository := StockRepository{}
	return &repository
}
func (repository StockRepository) ListAllDiscounted() (discountedBooks []responses.StockResponse, err error) {
	err = infrastructure_database.PostgresDB.Model(&stock_domain_entities.StockItem{}).Where("is_discounted = ?", true).Find(&discountedBooks).Error
	return
}

func (repository StockRepository) ListAll(languageCode string) (stockList []responses.StockResponse, err error) {
	ListAllByLanguageQuery := `SELECT * FROM stock_items 
								JOIN stock_item_details ON stock_item_details.id = stock_items.id AND stock_item_details.language_code = ?`

	err = infrastructure_database.PostgresDB.Raw(ListAllByLanguageQuery, languageCode).Scan(&stockList).Error
	return
}

func (repository StockRepository) ListAllDiscountedModel() (discountedBooks []stock_domain_entities.StockItem, err error) {
	err = infrastructure_database.PostgresDB.Model(&stock_domain_entities.StockItem{}).Where("is_discounted = ?", true).Find(&discountedBooks).Error
	return
}
