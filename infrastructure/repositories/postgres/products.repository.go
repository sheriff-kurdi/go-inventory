package postgres

import (
	"kurdi-go/core/entities/products"
	"kurdi-go/infrastructure/database/postgres"
)

type ProductsRepository struct {
}

func NewProductsRepository() ProductsRepository {
	repository := ProductsRepository{}
	return repository
}

func (repository ProductsRepository) SelectAll() []products.Product {
	var products []products.Product
	query := `SELECT * FROM products ;`

	postgres.Connect().Raw(query).Scan(&products)
	return products
}
