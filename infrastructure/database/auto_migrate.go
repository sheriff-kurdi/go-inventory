package database

import (
	"kurdi-go/core/entities"
	"kurdi-go/core/entities/products"

	"gorm.io/gorm"
)

func AutoMigrate(connection *gorm.DB) {

	err := connection.AutoMigrate(&products.Product{})
	if err != nil {
		return
	}
	err = connection.AutoMigrate(&products.ProductDetails{})
	if err != nil {
		return
	}
	err = connection.AutoMigrate(&entities.Language{})
	if err != nil {
		return
	}
}
