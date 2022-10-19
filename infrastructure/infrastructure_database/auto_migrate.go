package infrastructure_database

import (
	"kurdi-go/domain/domain_entities"
	"kurdi-go/domain/domain_entities/stock"
)

func PostgresAutoMigrate() {
	err := PostgresDB.AutoMigrate(&stock_domain_entities.StockItem{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&stock_domain_entities.StockItemDetails{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&domain_entities.Language{})
	if err != nil {
		return
	}
}
