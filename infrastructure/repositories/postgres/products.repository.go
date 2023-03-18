package postgres

import (
	"kurdi-go/core/entities/products"
	"kurdi-go/infrastructure/database"
)

type ProductsRepository struct {
}

func NewProductsRepository() ProductsRepository {
	repository := ProductsRepository{}
	return repository
}

func (repository ProductsRepository) SelectAll() []products.Product {
	var products []products.Product
	query := `SELECT * FROM products 
								JOIN products_details ON products_details.product_id = products.id `

	database.PostgresDB.Raw(query).Scan(&products)
	return products
}
