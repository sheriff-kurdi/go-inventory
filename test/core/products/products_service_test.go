package products

import (
	"fmt"
	"testing"

	"github.com/sheriff-kurdi/inventory/core/models/products"
	"github.com/sheriff-kurdi/inventory/core/services"
	"github.com/sheriff-kurdi/inventory/core/vm"
	"github.com/sheriff-kurdi/inventory/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductById(t *testing.T) {
	//arrange
	mockedRepository := &mocks.IProductsRepository{}
	productService := services.NewProductsService(mockedRepository, nil)

	expectedProductVM := vm.ProductVM{
		Id: 1,
		ProductQuantity: products.ProductQuantity{
			TotalStock:     200,
			AvailableStock: 200,
			ReservedStock:  0,
		},
		ProductPrice: products.ProductPrice{
			CostPrice:    20,
			SellingPrice: 30,
			Discount:     10,
			IsDiscounted: false,
		},
	}

	mockedRepository.On("SelectAllById", mock.Anything, 1).Return(expectedProductVM, nil)

	//action
	actualProductVM, _ := productService.FindById(1, "")

	//assert



	assert.Equal(t, expectedProductVM.Id, actualProductVM.Id)
	assert.Equal(t, expectedProductVM.SellingPrice, actualProductVM.SellingPrice)

}


func TestGetById(t *testing.T) {
	//arrange
	expectedProductVM := vm.ProductVM{
		Id: 1,
		ProductQuantity: products.ProductQuantity{
			TotalStock:     200,
			AvailableStock: 200,
			ReservedStock:  0,
		},
		ProductPrice: products.ProductPrice{
			CostPrice:    20,
			SellingPrice: 30,
			Discount:     10,
			IsDiscounted: false,
		},
	}
	mockedRepository := &mocks.IProductsRepository{}
	productService := services.NewProductsService(mockedRepository, nil)

	mockedRepository.On("GetById", 1).Return(expectedProductVM, nil)

	//action
	actualProductVM, _ := productService.GetById(1)

	//assert


	assert.Equal(t, expectedProductVM.Id, actualProductVM.Id)

}

func TestGetByIdV2(t *testing.T) {
	//arrange
	expectedProductVM := vm.ProductVM{
		Id: 1,
		ProductQuantity: products.ProductQuantity{
			TotalStock:     200,
			AvailableStock: 200,
			ReservedStock:  0,
		},
		ProductPrice: products.ProductPrice{
			CostPrice:    20,
			SellingPrice: 30,
			Discount:     10,
			IsDiscounted: false,
		},
	}
	mockedRepository := &mocks.IProductsRepository{}
	productService := services.NewProductsService(mockedRepository, nil)

	string := "string"
	fmt.Print(string)
	mockedRepository.On("GetByIdV2", &string, 1).Return(expectedProductVM, nil)
	//action
	actualProductVM, _ := productService.GetByIdV2(&string, 1)

	//assert


	assert.Equal(t, expectedProductVM.Id, actualProductVM.Id)

}