package repositories

import (
	"github.com/sheriff-kurdi/inventory/core/vm"

	"gorm.io/gorm"
)

type IProductsRepository interface {
	SelectAll(connection *gorm.DB) []vm.ProductVM
	SelectByCriteria(connection *gorm.DB, searchCriteria ProductsSearcheCriteria) []vm.ProductVM
	Save(connection *gorm.DB, productVM vm.ProductSavingVM) (productId int, err error)
	SelectAllByDetails(connection *gorm.DB, languageCode string) []vm.ProductVM
	DeleteById(connection *gorm.DB, productId int) (err error)
	SelectAllById(connection *gorm.DB, id int) (vm.ProductVM, error)
	//SelectById(query string) products.Product
	// SelectAllByCondition(query string) []products.Product
	// SelectByCondition(query string) products.Product

	// Update(product products.Product) products.Product
	// Delete(id int) int
}

type ProductsSearcheCriteria struct {
	Id            *int
	CostPriceFrom *float32
	CostPriceTo   *float32
	IsDiscounted  *bool
	LanguageCode  *string
}
