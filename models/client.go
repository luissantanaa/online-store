package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AuxClient struct {
	ID       uint
	Username string
}

type Client struct {
	gorm.Model
	ID        uint   `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Username  string `json:"username" gorm:"not null; default:null; unique;"`
	Password  string `json:"password" gorm:"not null; default:null"`
	Role      string `json:"role" gorm:"default:client"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Client) String() string {
	return fmt.Sprintf(`{ID:%v, Username:%v}`, c.ID, c.Username)
}

func (c Client) OmitInfo() AuxClient {
	return AuxClient{ID: c.ID, Username: c.Username}
}
