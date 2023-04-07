package services

import (
	repository "kurdi-go/core/contracts/repositories"
	"kurdi-go/core/vm"
	postgresDatabse "kurdi-go/infrastructure/database/postgres"
	"kurdi-go/infrastructure/repositories/postgres"
	"kurdi-go/web/resources"
	"kurdi-go/web/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	repository repository.IProductsRepository
	Connection *gorm.DB
}

func NewAuthService() AuthService {
	service := AuthService{
		repository: postgres.NewProductsRepository(postgresDatabse.Connect()),
	}
	return service
}

func (service AuthService) Login(loginVM *vm.LoginVM) resources.IResource {

	// Generate a new Access token.
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 and token generation error.
		return resources.BadRequest(err.Error())
	}
	return resources.Ok(token, "")
}


