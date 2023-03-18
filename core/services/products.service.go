package services

import (
	repository "kurdi-go/core/contracts/repositories"
	"kurdi-go/core/entities/products"
)

type ProductsService struct {
	repository repository.IProductsRepository
}

func NewProductsService(productRepository repository.IProductsRepository) ProductsService {
	service := ProductsService{
		repository: productRepository,
	}
	return service
}

func (service ProductsService) ListAll() ([]products.Product) {
	return service.repository.SelectAll()

}
