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
	Connection *gorm.DB
}

func NewProductsService() AuthService {
	service := AuthService{
		repository: postgres.NewProductsRepository(postgresDatabse.Connect()),
		Connection: postgresDatabse.Connect(),
	}
	return service
}

func (service AuthService) ListAll(languageCode string) []vm.ProductVM {
	return service.repository.SelectAllByDetails(languageCode)

}

func (service AuthService) FindById(id int, languageCode string) *vm.ProductVM {
	products := service.repository.SelectByCriteria(repository.ProductsSearcheCriteria{
		Id:           &id,
		LanguageCode: &languageCode,
	})
	if len(products) == 0 {
		return nil
	}
	return &products[0]

}

func (service AuthService) Insert(productVM vm.ProductSavingVM) (productId int, err error) {
	transasction := service.Connection.Begin()
	productId, err = service.repository.Upsert(transasction, productVM)
	if err != nil {
		transasction.Rollback()
		utils.Logger().Info(err.Error())
		return
	}
	transasction.Commit()
	return
}
