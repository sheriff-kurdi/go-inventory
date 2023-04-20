package services

import (
	repository "github.com/sheriff-kurdi/inventory/core/contracts/repositories"
	"github.com/sheriff-kurdi/inventory/core/vm"
	postgresDatabse "github.com/sheriff-kurdi/inventory/infrastructure/database/postgres"
	"github.com/sheriff-kurdi/inventory/infrastructure/repositories/postgres"
	"github.com/sheriff-kurdi/inventory/web/resources"
	"github.com/sheriff-kurdi/inventory/web/utils"

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


