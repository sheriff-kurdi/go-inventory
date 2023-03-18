package repositories

import "kurdi-go/core/entities/products"

type IProductsRepository interface {
	SelectAll() []products.Product
	// SelectAllByCondition(query string) []products.Product
	// SelectByCondition(query string) products.Product
	// Insert(product products.Product) products.Product
	// Update(product products.Product) products.Product
	// Delete(id int) int
}
