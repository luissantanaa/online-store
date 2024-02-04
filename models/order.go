package models

import (
	"fmt"
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

func (o Order) String() string {
	return fmt.Sprintf(`{"ID":"%v","Client":"%v", "Items":%v, "Sent":%v}`, o.ID, o.Client, o.Items, o.Sent)
}
