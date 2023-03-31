package repositories

import (
	"kurdi-go/core/vm"
)

type IProductsRepository interface {
	SelectAll() []vm.Product
	SelectByCriteria(searchCriteria ProductsSearcheCriteria) []vm.Product
	SelectAllByDetails(languageCode string) []vm.Product
	//SelectById(query string) products.Product
	// SelectAllByCondition(query string) []products.Product
	// SelectByCondition(query string) products.Product
	// Insert(product products.Product) products.Product
	// Update(product products.Product) products.Product
	// Delete(id int) int
}

type ProductsSearcheCriteria struct {
	Id            *int
	CostPriceFrom *float32
	CostPriceTo   *float32
	IsDiscounted  *bool
}
