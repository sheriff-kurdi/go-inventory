package products

import (
	"github.com/sheriff-kurdi/inventory/core/models"
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	Id             uint                  `gorm:"primary"`
	ProductDetails []ProductDetailsModel `gorm:"foreignKey:ProductId" json:"product_details"`
	ProductQuantity
	ProductPrice
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

func (m *ProductModel) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *ProductModel) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	return nil
}

func (m *ProductModel) TableName() string {
	return "products"
}
