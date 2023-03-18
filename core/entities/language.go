package entities

type Language struct {
	LanguageCode string `gorm:"unique" json:"language_code"`
	Name         string `json:"name"`
	TimeStamps
}
