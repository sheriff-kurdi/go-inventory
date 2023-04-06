package postgres

import (
	"kurdi-go/core/contracts/repositories"
	"kurdi-go/core/models/products"
	"kurdi-go/core/vm"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repository ProductsRepository) SelectAll() []vm.ProductVM {
	var products []vm.ProductVM
	query := `SELECT * FROM products ;`

	repository.Connection.Raw(query).Scan(&products)
	return products
}

func (repository ProductsRepository) DeleteById(connection *gorm.DB, productId int) (err error) {
	query := `
		delete from products where id = ? cascade;
	`
	err = connection.Exec(query, productId).Error
	return
}

func (repository ProductsRepository) SelectByCriteria(searchCriteria repositories.ProductsSearcheCriteria) []vm.ProductVM {
	var products []vm.ProductVM
	query := `SELECT * FROM products `
	params := make([]interface{}, 0)
	if searchCriteria.LanguageCode != nil && len(*searchCriteria.LanguageCode) != 0 {
		params = append(params, &searchCriteria.LanguageCode)
		query += "join product_details on product_details.language_code = ? and product_details.product_id = products.id "
	}

	if searchCriteria.Id != nil || searchCriteria.CostPriceFrom != nil || searchCriteria.CostPriceTo != nil || searchCriteria.IsDiscounted != nil {
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

func (repository ProductsRepository) SelectAllByDetails(languageCode string) []vm.ProductVM {
	var productsList []vm.ProductVM
	if languageCode == "" {
		languageCode = os.Getenv("DEFAULT_LANGUAGE")
	}
	query := `SELECT * FROM products
		join product_details on product_details.language_code = ? and product_details.product_id = products.id;`
	repository.Connection.Raw(query, languageCode).Scan(&productsList)

	return productsList
}

func (repository ProductsRepository) Save(connection *gorm.DB, productVM vm.ProductSavingVM) (productId int, err error) {
	productModel := products.ProductModel{
		Id:              productVM.Id,
		ProductQuantity: productVM.ProductQuantity,
		ProductPrice:    productVM.ProductPrice,
	}

	if err = connection.Save(&productModel).Error; err != nil {
		return
	}
	productId = int(productModel.Id)

	productDetails := []products.ProductDetailsModel{}
	for _, productDetailsVM := range productVM.Details {
		productDetails = append(productDetails, products.ProductDetailsModel{
			Name:         productDetailsVM.Name,
			Description:  productDetailsVM.Description,
			LanguageCode: productDetailsVM.LanguageCode,
			ProductId:    productId,
		})

	}

	if err = connection.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "product_id"}, {Name: "name"}, {Name: "language_code"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "description", "updated_at"}),
	}).Save(&productDetails).Error; err != nil {
		return
	}
	return
}
