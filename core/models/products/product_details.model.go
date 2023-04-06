package products

import "kurdi-go/core/models"

type ProductDetailsModel struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	LanguageCode string            `gorm:"primaryKey" json:"language_code"`
	Language     models.LanguageModel `gorm:"foreignKey:language_code"`
	ProducId     int               `gorm:"primaryKey" json:"product_id"`
	models.TimeStamps
}
