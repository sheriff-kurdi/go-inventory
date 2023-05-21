package services

import (
	repository "github.com/sheriff-kurdi/inventory/core/contracts/repositories"
	"github.com/sheriff-kurdi/inventory/core/vm"
	"github.com/sheriff-kurdi/inventory/web/utils"

	"gorm.io/gorm"
)

type ProductsService struct {
	repository repository.IProductsRepository
	connection *gorm.DB
}

func NewProductsService(repository repository.IProductsRepository, connection *gorm.DB) ProductsService {
	service := ProductsService{
		repository: repository,
		connection: connection ,
	}
	return service
}

func (service ProductsService) ListAll(languageCode string) []vm.ProductVM {
	
	return service.repository.SelectAllByDetails(service.connection, languageCode)
}

func (service ProductsService) FindById(id int, languageCode string) (*vm.ProductVM, error) {
	product, err := service.repository.SelectAllById(service.connection, id)

	return &product, err
}

func (service ProductsService) DeleteById(productId int) (err error) {
	err = service.repository.DeleteById(service.connection, productId)
	return 
}



func (service ProductsService) Save(productVM vm.ProductSavingVM) (productId int, err error) {
	transasction := service.connection.Begin()
	productId, err = service.repository.Save(transasction, productVM)
	if err != nil {
		transasction.Rollback()
		utils.Logger().Info(err.Error())
		return
	}
	transasction.Commit()
	return
}

func (service ProductsService) GetById(number int) (productVM vm.ProductVM, err error ) {
	productVM, err = service.repository.GetById(number)
	if err != nil {
		utils.Logger().Info(err.Error())
		return
	}
	return
}

func (service ProductsService) GetByIdV2(name *string, number int) (productVM vm.ProductVM, err error ) {
	productVM, err = service.repository.GetByIdV2(name, number)
	if err != nil {
		utils.Logger().Info(err.Error())
		return
	}
	return
}
