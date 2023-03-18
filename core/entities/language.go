package entities

type Language struct {
	Entity
	LanguageCode string `gorm:"unique" json:"language_code"`
	Name         string `json:"name"`
}
