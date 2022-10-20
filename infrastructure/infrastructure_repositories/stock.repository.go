package infrastructure_repositories

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	stockDomainEntities "kurdi-go/domain/domain_entities/stock"
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

func (repository StockRepository) CreateDetails(connection *gorm.DB, stockItemDetailsEntities []stockDomainEntities.StockItemDetails) (err error) {

	createStockItemDetailsJson, err := json.Marshal(stockItemDetailsEntities)
	if err != nil {
		return
	}
	createStockQuery := `INSERT INTO  stock_item_details ("name", "description", "language_code", "stock_item_id", "created_at", "updated_at", "deleted_at") 
						 SELECT * FROM json_to_recordset(?) AS X ("name" varchar, "description" varchar, language_code varchar, "stock_item_id" int, "CreatedAt" timestamp, "UpdatedAt" timestamp, "DeletedAt" timestamp);`

	result := connection.Exec(createStockQuery, string(createStockItemDetailsJson))
	if result.Error != nil {
		err = result.Error
		return
	}
	if result.RowsAffected < int64(len(stockItemDetailsEntities)) {
		return errors.New("all items details not created")
	}
	return
}
