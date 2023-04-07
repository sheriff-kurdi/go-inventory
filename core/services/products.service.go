package services

import (
	repository "kurdi-go/core/contracts/repositories"
	"kurdi-go/core/vm"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/infrastructure/repositories/postgres"
	"kurdi-go/web/utils"

	"gorm.io/gorm"
)

type ProductsService struct {
	repository repository.IProductsRepository
	connection *gorm.DB
}

func NewProductsService() ProductsService {
	service := ProductsService{
		repository: postgres.NewProductsRepository(postgresDatabse.Connect()),
		connection: postgresDatabse.Connect(),
	}
	return service
}

func (service ProductsService) ListAll(languageCode string) []vm.ProductVM {
	return service.repository.SelectAllByDetails(service.connection, languageCode)
}

func (service ProductsService) FindById(id int, languageCode string) *vm.ProductVM {
	products := service.repository.SelectByCriteria(service.connection, repository.ProductsSearcheCriteria{
		Id:           &id,
		LanguageCode: &languageCode,
	})
	if len(products) == 0 {
		return nil
	}
	return &products[0]
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
