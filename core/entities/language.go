package entities

type Language struct {
	LanguageCode string `gorm:"primaryKey" json:"language_code"`
	Name         string `json:"name"`
	TimeStamps
}
