package services

import (
	repository "kurdi-go/core/contracts/repositories"
	"kurdi-go/core/vm"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/infrastructure/repositories/postgres"

	"gorm.io/gorm"
)

type ProductsService struct {
	repository repository.IProductsRepository
	Connection *gorm.DB
}

func NewProductsService() ProductsService {
	service := ProductsService{
		repository: postgres.NewProductsRepository(postgresDatabse.Connect()),
	}
	return service
}

func (service ProductsService) ListAll(languageCode string) []vm.Product {
	return service.repository.SelectAllByDetails(languageCode)

}

func (service ProductsService) FindById(id int) *vm.Product {
	products := service.repository.SelectByCriteria(repository.ProductsSearcheCriteria{
		Id: &id,
	})
	if len(products) == 0 {
		return nil
	}
	return &products[0]

}
