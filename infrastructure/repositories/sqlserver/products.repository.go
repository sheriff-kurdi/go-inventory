package sqlserver

import (
	"kurdi-go/core/contracts/repositories"
	"kurdi-go/core/entities/products"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	Connection *gorm.DB
}

func NewProductsRepository(connection *gorm.DB) ProductsRepository {
	repository := ProductsRepository{
		Connection: connection,
	}
	return repository
}

func (repository ProductsRepository) SelectAll() []products.Product {
	var products []products.Product
	query := `SELECT * FROM products ;`

	repository.Connection.Raw(query).Scan(&products)
	return products
}

func (repository ProductsRepository) SelectByCriteria(searchCriteria repositories.ProductsSearcheCriteria) []products.Product {
	var products []products.Product

	return products
}
