package services

import (
	"kurdi-go/domain/domain_entities/stock"
	resources2 "kurdi-go/helpers/resources"
	"kurdi-go/http/requests"
	"kurdi-go/http/responses"
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

func (service StockService) FindById(bookId int) (response resources2.IResource) {
	var book responses.StockResponse
	if err := infrastructure_database.PostgresDB.Model(stock_domain_entities.StockItem{}).Where("id = ?", bookId).First(&book).Scan(&book).Error; err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(book, "")
}

func (service StockService) Create(request requests.CreateStockRequest) (response resources2.IResource) {
	var books []responses.StockResponse
	err := infrastructure_database.PostgresDB.Model(&stock_domain_entities.StockItem{}).Scan(&books).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	stockModel := stock_domain_entities.StockItem{}
	//stockModel.Name = request.Author

	err = infrastructure_database.PostgresDB.Create(&stockModel).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(stockModel, "created")
}

func (service StockService) Update(request requests.CreateStockRequest, bookId int) (response resources2.IResource) {
	var bookModel stock_domain_entities.StockItem
	//bookModel.Name = request.Title
	bookModel.Id = uint(bookId)
	err := infrastructure_database.PostgresDB.Save(&bookModel).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(bookModel, "updated")
}

func (service StockService) Delete(bookId int) (response resources2.IResource) {
	err := infrastructure_database.PostgresDB.Delete(&stock_domain_entities.StockItem{}, bookId).Error
	if err != nil {
		return resources2.GetError500Resource(err.Error())
	}
	return resources2.GetSuccess200Resource(bookId, "Deleted")

}
