package repositories

import (
	"kurdi-go/core/vm"

	"gorm.io/gorm"
)

type IProductsRepository interface {
	SelectAll() []vm.ProductVM
	SelectByCriteria(searchCriteria ProductsSearcheCriteria) []vm.ProductVM
	Save(connection *gorm.DB, productVM vm.ProductSavingVM) (productId int, err error)
	SelectAllByDetails(languageCode string) []vm.ProductVM
	DeleteById(connection *gorm.DB, productId int) (err error)

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
