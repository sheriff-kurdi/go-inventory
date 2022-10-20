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

func (repository StockRepository) FindById(stockItemId int, languageCode string) (stock responses.StockResponse, err error) {
	if languageCode == "" {
		languageCode = os.Getenv("DEFAULT_LANGUAGE")
	}
	ListAllByLanguageQuery := `SELECT * FROM stock_items 
								JOIN stock_item_details ON stock_item_details.stock_item_id	 = stock_items.id AND stock_item_details.language_code = ? WHERE stock_items.id = ?`

	err = infrastructure_database.PostgresDB.Raw(ListAllByLanguageQuery, languageCode, stockItemId).Scan(&stock).Error
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

func (repository StockRepository) UpdateDetails(connection *gorm.DB, stockItemDetailsEntities []stockDomainEntities.StockItemDetails) (err error) {

	createStockItemDetailsJson, err := json.Marshal(stockItemDetailsEntities)
	if err != nil {
		return
	}
	createStockQuery := `UPDATE stock_item_details stockItemDetail SET 
                              "name" = itemDetail."name",
                              "description" = itemDetail."description",
                              "updated_at" = itemDetail."UpdatedAt"
						 FROM (SELECT * FROM json_to_recordset(?) AS X ("name" varchar, "description" varchar, language_code varchar, "stock_item_id" int, "UpdatedAt" timestamp)) AS itemDetail
						 WHERE stockItemDetail.stock_item_id = itemDetail.stock_item_id AND stockItemDetail.language_code = itemDetail.language_code;`

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

func (repository StockRepository) DeleteStockDetails(connection *gorm.DB, stockItemId int) (err error) {
	deleteStockDetailsQuery := `DELETE FROM stock_item_details WHERE stock_item_id = ? ;`

	result := connection.Exec(deleteStockDetailsQuery, stockItemId)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
