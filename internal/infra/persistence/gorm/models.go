package gorm

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type OrderModel struct {
	ID           string      `gorm:"primaryKey;size:14"`
	Status       string      `gorm:"not null"`
	CustomerName string      `gorm:"not null"`
	Phone        string      `gorm:"not null"`
	Address      string      `gorm:"not null"`
	Items        []OrderItemModel `gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time
}

type OrderItemModel struct {
	ID           string `gorm:"primaryKey;size:14"`
	OrderID      string `gorm:"index;size:14;not null"`
	Size         string `gorm:"not null"`
	Pizza        string `gorm:"not null"`
	Instructions string
}

func (o *OrderModel) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItemModel) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}

func (OrderModel) TableName() string { return "orders" }
func (OrderItemModel) TableName() string { return "order_items" }

type UserModel struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (UserModel) TableName() string { return "users" }
