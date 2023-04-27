package services

import (
	repository "github.com/sheriff-kurdi/inventory/core/contracts/repositories"
	"github.com/sheriff-kurdi/inventory/core/vm"
	postgresDatabse "github.com/sheriff-kurdi/inventory/infrastructure/database/postgres"
	"github.com/sheriff-kurdi/inventory/infrastructure/repositories/postgres"
	"github.com/sheriff-kurdi/inventory/web/utils"

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
	product := service.repository.SelectAllById(service.connection, id)

	return &product
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
