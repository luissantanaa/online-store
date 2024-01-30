package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID        uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
