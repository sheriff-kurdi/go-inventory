package products

import (
	"github.com/sheriff-kurdi/inventory/core/contracts/repositories"
	"github.com/sheriff-kurdi/inventory/core/vm"

	"gorm.io/gorm"
)

type MockedProductsRepository struct {
	products []vm.ProductVM
}

func NewMockedProductsRepository(products []vm.ProductVM) MockedProductsRepository {
	repository := MockedProductsRepository{
		products: products,
	}
	return repository
}

func (repository MockedProductsRepository) SelectAll(connection *gorm.DB) []vm.ProductVM {
	return repository.products
}

func (repository MockedProductsRepository) DeleteById(connection *gorm.DB, productId int) (err error) {

	return
}

func (repository MockedProductsRepository) SelectByCriteria(connection *gorm.DB, searchCriteria repositories.ProductsSearcheCriteria) []vm.ProductVM {
	var productsVM []vm.ProductVM

	return productsVM
}

func (repository MockedProductsRepository) SelectAllByDetails(connection *gorm.DB, languageCode string) []vm.ProductVM {
	productsVM := []vm.ProductVM{}

	return productsVM
}

func (repository MockedProductsRepository) SelectAllById(connection *gorm.DB, id int) (vm.ProductVM, error) {
	var productVM vm.ProductVM
	for _, product := range repository.products{
		if int(product.Id) == id {
			return product, nil
		}
	}

	return productVM, nil
}

func (repository MockedProductsRepository) Save(connection *gorm.DB, productVM vm.ProductSavingVM) (productId int, err error) {

	return
}
