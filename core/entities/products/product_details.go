package products

import (
	"kurdi-go/core/entities"
)

type ProductDetails struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	LanguageCode string            `gorm:"primaryKey" json:"language_code"`
	Language     entities.Language `gorm:"foreignKey:language_code"`
	ProducId     int               `gorm:"primaryKey" json:"product_id"`
	entities.TimeStamps
}
