package postgres

import (
	"os"

	"github.com/sheriff-kurdi/inventory/core/contracts/repositories"
	"github.com/sheriff-kurdi/inventory/core/models/products"
	"github.com/sheriff-kurdi/inventory/core/vm"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductsRepository struct {
}

func NewProductsRepository(connection *gorm.DB) ProductsRepository {
	repository := ProductsRepository{}
	return repository
}

func (repository ProductsRepository) SelectAll(connection *gorm.DB) []vm.ProductVM {
	var products []vm.ProductVM
	query := `SELECT * FROM products ;`

	connection.Raw(query).Scan(&products)
	return products
}

func (repository ProductsRepository) DeleteById(connection *gorm.DB, productId int) (err error) {
	query := `
		delete from products where id = ? cascade;
	`
	err = connection.Exec(query, productId).Error
	return
}

func (repository ProductsRepository) SelectByCriteria(connection *gorm.DB, searchCriteria repositories.ProductsSearcheCriteria) []vm.ProductVM {
	var productsVM []vm.ProductVM

	query := `SELECT * FROM products `
	params := make([]interface{}, 0)

	if searchCriteria.LanguageCode != nil && len(*searchCriteria.LanguageCode) != 0 {
		params = append(params, &searchCriteria.LanguageCode)
		query += "join product_details on product_details.language_code = ? and product_details.product_id = products.id "
	} else {
		defaultLanguage := os.Getenv("DEFAULT_LANGUAGE")
		params = append(params, &defaultLanguage)
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

	connection.Raw(query, params...).Scan(&productsVM)

	return productsVM
}

func (repository ProductsRepository) SelectAllByDetails(connection *gorm.DB, languageCode string) []vm.ProductVM {
	productsVM := []vm.ProductVM{}
	var productsModel []products.ProductModel
	if languageCode == "" || len(languageCode) == 0 {
		languageCode = os.Getenv("DEFAULT_LANGUAGE")
	}


	connection.Preload("ProductDetails", "language_code = ?", languageCode).Find(&productsModel)

	for _, productModel := range productsModel {
		productVM := vm.ProductVM{}

		productVM.Id = productModel.Id
		productVM.TotalStock = productModel.TotalStock
		productVM.AvailableStock = productModel.AvailableStock
		productVM.ReservedStock = productModel.ReservedStock

		productVM.CostPrice = productModel.CostPrice
		productVM.SellingPrice = productModel.SellingPrice
		productVM.Discount = productModel.Discount
		productVM.IsDiscounted = productModel.IsDiscounted

		productVM.ProductDetails = []vm.ProductDetailsVM{}

		for _, detail := range productModel.ProductDetails {
			productVM.ProductDetails = append(productVM.ProductDetails, vm.ProductDetailsVM{
				Name:         detail.Name,
				Description:  detail.Description,
				LanguageCode: detail.LanguageCode,
				TimeStamps:   detail.TimeStamps,
			})
		}

		productsVM = append(productsVM, productVM)
	}

	return productsVM
}

func (repository ProductsRepository) SelectAllById(connection *gorm.DB, id int) vm.ProductVM {
	var productVM vm.ProductVM
	var productModel products.ProductModel
	connection.Preload("ProductDetails").Where("id = ?", id).First(&productModel)
	productVM.Id = productModel.Id
	productVM.TotalStock = productModel.TotalStock
	productVM.AvailableStock = productModel.AvailableStock
	productVM.ReservedStock = productModel.ReservedStock

	productVM.CostPrice = productModel.CostPrice
	productVM.SellingPrice = productModel.SellingPrice
	productVM.Discount = productModel.Discount
	productVM.IsDiscounted = productModel.IsDiscounted
	productVM.ProductDetails = []vm.ProductDetailsVM{}
	for _, detail := range productModel.ProductDetails {
		productVM.ProductDetails = append(productVM.ProductDetails, vm.ProductDetailsVM{
			Name:         detail.Name,
			Description:  detail.Description,
			LanguageCode: detail.LanguageCode,
			TimeStamps:   detail.TimeStamps,
		})
	}

	return productVM
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
	for _, productDetailsVM := range productVM.ProductDetails {
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
