package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Client    Client `json:"client" gorm:"foreignKey:ID"`
	Items     []Item `json:"items" gorm:"not null; default:null; unique;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Sent      bool `json:"sent" gorm:"default:false"`
}

func (o Order) String() string {
	return fmt.Sprintf(`{"ID":"%v","Client":"%v", "Items":%v, "Sent":%v}`, o.ID, o.Client, o.Items, o.Sent)
}
