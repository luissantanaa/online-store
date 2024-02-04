package models

import (
	"fmt"
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

func (c Client) String() string {
	return fmt.Sprintf(`{"ID":"%v", "Username":"%v"}`, c.ID, c.Username)
}
