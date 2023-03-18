package database

import (
	"kurdi-go/core/entities"
	"kurdi-go/core/entities/products"
)

func PostgresAutoMigrate() {
	err := PostgresDB.AutoMigrate(&products.Product{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&products.ProductDetails{})
	if err != nil {
		return
	}
	err = PostgresDB.AutoMigrate(&entities.Language{})
	if err != nil {
		return
	}
}
