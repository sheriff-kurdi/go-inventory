package repositories

import "kurdi-go/core/entities/products"

type IProductsRepository interface {
	SelectAll() []products.Product
	SelectByCriteria(searchCriteria ProductsSearcheCriteria) []products.Product
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
