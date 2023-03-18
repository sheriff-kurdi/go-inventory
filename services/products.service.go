package services

import (
	"kurdi-go/core/entities/products"
	"kurdi-go/core/services"
	"kurdi-go/infrastructure/repositories/postgres"
)

type StockService struct {
	productsService services.ProductsService
}

func NewStockService() StockService {
	return StockService{
		productsService: services.NewProductsService(postgres.NewProductsRepository()),
	}
}

func (service StockService) ListAll(languageCode string) []products.Product {
	stockList := service.productsService.ListAll()
	return stockList
}
