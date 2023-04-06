package sqlserver

import (
	"kurdi-go/core/contracts/repositories"
	"kurdi-go/core/models/products"

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

func (repository ProductsRepository) SelectAll() []products.ProductModel {
	var products []products.ProductModel
	query := `SELECT * FROM products ;`

	repository.Connection.Raw(query).Scan(&products)
	return products
}

func (repository ProductsRepository) SelectByCriteria(searchCriteria repositories.ProductsSearcheCriteria) []products.ProductModel {
	var products []products.ProductModel
	query := `SELECT * FROM products `
	params := make([]interface{}, 0)

	if searchCriteria.Id != nil  || searchCriteria.CostPriceFrom != nil || searchCriteria.CostPriceTo != nil || searchCriteria.IsDiscounted != nil{
		query += "WHERE "
	}

	if searchCriteria.Id != nil {
		params = append(params, &searchCriteria.Id) 
		query += "products.id = ?"
	}

	if searchCriteria.CostPriceFrom != nil {
		params = append(params, &searchCriteria.CostPriceFrom) 
		query += "products.cost_price >= ?"
	}

	if searchCriteria.CostPriceTo != nil {
		params = append(params, &searchCriteria.CostPriceFrom) 
		query += "products.cost_price <= ?"
	}

	if searchCriteria.IsDiscounted != nil {
		params = append(params, &searchCriteria.IsDiscounted) 
		query += "products.is_discounted = ?"
	}

	repository.Connection.Raw(query, params...).Scan(&products)

	return products
}
