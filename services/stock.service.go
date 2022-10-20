package services

import (
	"kurdi-go/domain/domain_entities/stock"
	resources2 "kurdi-go/helpers/resources"
	"kurdi-go/http/requests"
	"kurdi-go/infrastructure/infrastructure_database"
	"kurdi-go/infrastructure/infrastructure_repositories"
)

type StockService struct {
	repository *infrastructure_repositories.StockRepository
}

func NewStockService() *StockService {
	service := StockService{
		repository: infrastructure_repositories.NewStockRepository(),
	}
	return &service
}

func (service StockService) ListAll(languageCode string) (response resources2.IResource) {
	stockList, err := service.repository.ListAll(languageCode)
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stockList, "")
}

func (service StockService) FindById(stockId int, languageCode string) (response resources2.IResource) {
	stock, err := service.repository.FindById(stockId, languageCode)
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stock, "")
}

func (service StockService) Create(request requests.CreateStockRequest) (response resources2.IResource) {
	connection := infrastructure_database.PostgresDB
	transaction := connection.Begin()

	// create stock item
	stockItemEntity := request.ToStockItemEntity()
	err := transaction.Model(&stock_domain_entities.StockItem{}).Save(&stockItemEntity).Error
	if err != nil {
		transaction.Rollback()
		return resources2.GetError500Resource(err.Error())
	}

	// create stock details
	stockItemDetailsEntities := request.ToStockItemDetailsEntities(int(stockItemEntity.Id))
	err = service.repository.CreateDetails(transaction, stockItemDetailsEntities)
	if err != nil {
		transaction.Rollback()
		return resources2.GetError500Resource(err.Error())
	}
	transaction.Commit()
	return resources2.GetSuccess201Resource(request, "created")
}

func (service StockService) Update(request requests.UpdateStockRequest, stockItemId int) (response resources2.IResource) {
	connection := infrastructure_database.PostgresDB
	transaction := connection.Begin()

	stockItemEntity := request.ToStockItemEntity()
	stockItemEntity.Id = uint(stockItemId)

	err := transaction.Save(&stockItemEntity).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	// update details
	stockItemDetailsEntities := request.ToStockItemDetailsEntities(int(stockItemEntity.Id))
	err = service.repository.UpdateDetails(transaction, stockItemDetailsEntities)
	if err != nil {
		transaction.Rollback()
		return resources2.GetError500Resource(err.Error())
	}
	transaction.Commit()
	return resources2.GetSuccess200Resource(request, "updated")
}

func (service StockService) Delete(stockItemId int) (response resources2.IResource) {
	connection := infrastructure_database.PostgresDB
	transaction := connection.Begin()
	err := transaction.Delete(&stock_domain_entities.StockItem{}, stockItemId).Error
	if err != nil {
		transaction.Rollback()
		return resources2.GetError500Resource(err.Error())
	}
	err = service.repository.DeleteStockDetails(transaction, stockItemId)
	if err != nil {
		transaction.Rollback()
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stockItemId, "Deleted")

}
