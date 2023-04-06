package products

import "kurdi-go/core/models"

type Product struct {
	Id             uint             `gorm:"primary"`
	ProductDetails []ProductDetails `gorm:"foreignKey:ProducId" json:"product_details"`
	ProductQuantity
	models.TimeStamps
}

//-------stock quantities section-----
func (product *Product) AddStock(quantity int) {
	product.TotalStock += quantity
	product.AvailableStock += quantity
}
func (product *Product) ReserveStock(quantity int) {
	product.ReservedStock += quantity
	product.AvailableStock -= quantity
}
func (product *Product) CancelReservation(quantity int) {
	product.ReservedStock -= quantity
	product.AvailableStock += quantity
}
func (product *Product) Selling(quantity int) {
	product.TotalStock -= quantity
	product.AvailableStock -= quantity
}

//-------end stock quantities section-----
