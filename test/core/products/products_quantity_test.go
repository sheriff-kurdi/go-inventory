package products

import (
	"kurdi-go/core/models/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

//adding stock
func TestTotallStockAfterAddingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.AddStock(10)

	//assert
	expected := 110
	actual := product.ProductQuantity.TotalStock
	assert.Equal(t, expected, actual)
}
func TestAvailableStockAfterAddingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.AddStock(10)

	//assert
	expected := 110
	actual := product.ProductQuantity.AvailableStock
	assert.Equal(t, expected, actual)
}

//selling stock
func TestTotallStockAftersellingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.Selling(10)

	//assert
	expected := 90
	actual := product.ProductQuantity.TotalStock
	assert.Equal(t, expected, actual)
}
func TestAvailableStockAftersellingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.Selling(10)

	//assert
	expected := 90
	actual := product.ProductQuantity.AvailableStock
	assert.Equal(t, expected, actual)
}

//reserve
func TestAvailableStockAfterReservingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.ReserveStock(10)

	//assert
	expected := 90
	actual := product.ProductQuantity.AvailableStock
	assert.Equal(t, expected, actual)
}
func TestTotallStockAfterReservingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.ReserveStock(10)

	//assert
	expected := 100
	actual := product.ProductQuantity.TotalStock
	assert.Equal(t, expected, actual)
}
func TestReservedStockAfterReservingStock(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 100,
			ReservedStock:  0,
		},
	}
	//action
	product.ReserveStock(10)

	//assert
	expected := 10
	actual := product.ProductQuantity.ReservedStock
	assert.Equal(t, expected, actual)
}

//cancel reservation
func TestAvailableStockAfterCancelReservation(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 90,
			ReservedStock:  10,
		},
	}
	//action
	product.CancelReservation(10)

	//assert
	expected := 100
	actual := product.ProductQuantity.AvailableStock
	assert.Equal(t, expected, actual)
}
func TestTotallStockAfterCancelReservation(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 90,
			ReservedStock:  10,
		},
	}
	//action
	product.CancelReservation(10)

	//assert
	expected := 100
	actual := product.ProductQuantity.TotalStock
	assert.Equal(t, expected, actual)
}
func TestReservedStockAfterCancelReservation(t *testing.T) {
	//prepare
	product := products.ProductModel{
		ProductQuantity: products.ProductQuantity{
			TotalStock:     100,
			AvailableStock: 90,
			ReservedStock:  10,
		},
	}
	//action
	product.CancelReservation(10)

	//assert
	expected := 0
	actual := product.ProductQuantity.ReservedStock
	assert.Equal(t, expected, actual)
}
