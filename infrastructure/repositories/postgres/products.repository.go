package postgres

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

func (repository ProductsRepository) SelectAllByDetails(languageCode string) []products.Product {
	var products []products.Product
	query := `SELECT * FROM products 
		join product_details on product_details.language_code = ? and product_details.product_id = products.id;`

	repository.Connection.Raw(query, languageCode).Scan(&products)
	return products
}