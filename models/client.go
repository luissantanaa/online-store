package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
