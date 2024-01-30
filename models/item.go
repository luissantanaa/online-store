package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID        uint
	Name      string
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
