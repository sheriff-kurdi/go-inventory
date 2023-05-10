package products

import (
	"testing"

	"github.com/sheriff-kurdi/inventory/core/models/products"
	"github.com/sheriff-kurdi/inventory/core/services"
	"github.com/sheriff-kurdi/inventory/core/vm"

	"github.com/stretchr/testify/assert"
)

//adding stock
func TestGetProductById(t *testing.T) {
	//arrange
	var productsList []vm.ProductVM

	productsList = append(productsList, vm.ProductVM{
		Id: 1,
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
		ProductPrice: products.ProductPrice{
			CostPrice: 10,
			SellingPrice: 20,
			Discount: 5,
			IsDiscounted: true,
		},

	})

	productsList = append(productsList, vm.ProductVM{
		Id: 2,
		ProductQuantity: products.ProductQuantity{
			TotalStock:     200,
			AvailableStock: 200,
			ReservedStock:  0,
		},
		ProductPrice: products.ProductPrice{
			CostPrice: 20,
			SellingPrice: 30,
			Discount: 10,
			IsDiscounted: false,
		},

	})

	productService :=  services.NewProductsService(MockedProductsRepository{
		products: productsList,
	}, nil)

	//action
	product, _ := productService.FindById(1, "")

	//assert
	expectedProductId := 1
	expectedProductSellingPrice := 20
	
	actualProductId := product.Id
	actualProductSellingPrice := product.SellingPrice

	assert.Equal(t, expectedProductId, int(actualProductId))
	assert.Equal(t, expectedProductSellingPrice, int(actualProductSellingPrice))

}


