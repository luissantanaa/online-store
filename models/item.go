package models

import (
	"fmt"
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

func (i Item) String() string {
	//return fmt.Sprintf(`{"ID":"%v", "Name":"%v", "Quantity":%v}`, i.Name, i.Quantity)
	return fmt.Sprintf(`{"Name":"%v"}`, i.Name)
}
