package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	OrderID  int
	ItemID   uint // ID of the item
	Quantity int  // Quantity of the item
}

type Order struct {
	gorm.Model
	ID        uint        `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	ClientID  uint        `json:"client"`
	Items     []OrderItem `json:"items" gorm:"not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Sent      bool `json:"sent" gorm:"default:false"`
}

func (o Order) String() string {
	return fmt.Sprintf(`{"ID":%v,"Client":"%v", "Items":%v, "Sent":%v}`, o.ID, o.ClientID, o.Items, o.Sent)
}
