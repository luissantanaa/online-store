package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        uint   `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Username  string `json:"username" gorm:"not null; default:null; unique;"`
	Password  string `json:"password" gorm:"not null; default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Client) String() string {
	return fmt.Sprintf(`{ID:%v, Username:%v}`, c.ID, c.Username)
}
