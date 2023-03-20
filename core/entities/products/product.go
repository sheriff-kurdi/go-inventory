package products

import (
	"kurdi-go/core/entities"
)

type Product struct {
	Id             uint             `gorm:"primary"`
	ProductDetails []ProductDetails `gorm:"foreignKey:ProducId" json:"product_details"`
	Quantity       ProductQuantity  `json:"quantity"`
	entities.TimeStamps
}

//-------stock quantities section-----
func (product *Product)AddStock(quantity int){
	product.Quantity.TotalStock += quantity
	product.Quantity.AvailableStock += quantity
}
func (product *Product)ReserveStock(quantity int){
	product.Quantity.ReservedStock += quantity
	product.Quantity.AvailableStock -= quantity
}
func (product *Product)CancelReservation(quantity int){
	product.Quantity.ReservedStock -= quantity
	product.Quantity.AvailableStock += quantity
}
func (product *Product)Selling(quantity int){
	product.Quantity.TotalStock -= quantity
	product.Quantity.AvailableStock -= quantity
}
//-------end stock quantities section-----
