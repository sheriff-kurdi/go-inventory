package database

import (
	"github.com/sheriff-kurdi/inventory/core/models"
	"github.com/sheriff-kurdi/inventory/core/models/products"

	"gorm.io/gorm"
)

func AutoMigrate(connection *gorm.DB) {

	err := connection.AutoMigrate(&products.ProductModel{})
	if err != nil {
		return
	}
	err = connection.AutoMigrate(&products.ProductDetailsModel{})
	if err != nil {
		return
	}
	err = connection.AutoMigrate(&models.LanguageModel{})
	if err != nil {
		return
	}
}
