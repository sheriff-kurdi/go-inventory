package repositories

import (
	"kurdi-go/core/vm"

	"gorm.io/gorm"
)

type IProductsRepository interface {
	SelectAll() []vm.ProductVM
	SelectByCriteria(searchCriteria ProductsSearcheCriteria) []vm.ProductVM
	Insert(connection *gorm.DB, productVM vm.ProductInsertionVM) (productId int, err error)
	SelectAllByDetails(languageCode string) []vm.ProductVM
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
