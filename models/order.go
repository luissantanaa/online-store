package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        uint
	Client    Client `gorm:"foreignKey:ID"`
	Items     []Item `gorm:"foreignKey:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Sent      bool
}
