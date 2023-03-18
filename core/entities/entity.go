package entities

import (
	"time"
)

type Entity struct {
	Id        uint      `gorm:"primary"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time
}
