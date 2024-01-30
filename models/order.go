package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Client    Client `gorm:"foreignKey:ID"`
	Items     []Item
	CreatedAt time.Time
	UpdatedAt time.Time
	Sent      bool
}
