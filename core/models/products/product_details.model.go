package products

import (
	"kurdi-go/core/models"
	"time"

	"gorm.io/gorm"
)

type ProductDetailsModel struct {
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	LanguageCode string               `gorm:"primaryKey" json:"language_code"`
	Language     models.LanguageModel `gorm:"foreignKey:language_code"`
	ProductId     int                  `gorm:"primaryKey" json:"product_id"`
	models.TimeStamps
}

func (m *ProductDetailsModel) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *ProductDetailsModel) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	return nil
}

func (m *ProductDetailsModel) TableName() string {
	return "product_details"
}
