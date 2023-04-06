package products

import "kurdi-go/core/models"

type ProductModel struct {
	Id             uint             `gorm:"primary"`
	ProductDetails []ProductDetailsModel `gorm:"foreignKey:ProducId" json:"product_details"`
	ProductQuantity
	models.TimeStamps
}

//-------stock quantities section-----
func (product *ProductModel) AddStock(quantity int) {
	product.TotalStock += quantity
	product.AvailableStock += quantity
}
func (product *ProductModel) ReserveStock(quantity int) {
	product.ReservedStock += quantity
	product.AvailableStock -= quantity
}
func (product *ProductModel) CancelReservation(quantity int) {
	product.ReservedStock -= quantity
	product.AvailableStock += quantity
}
func (product *ProductModel) Selling(quantity int) {
	product.TotalStock -= quantity
	product.AvailableStock -= quantity
}

//-------end stock quantities section-----
