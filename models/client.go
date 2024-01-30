package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
