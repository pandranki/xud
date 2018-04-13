package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

type Order struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	PairID    string
	PeerID    string
	Quantity  string
	Price     string
}

func (o *Order) GetOrders(db *gorm.DB) []Order {
	orders := []Order{}
	db.Find(&orders)
	return orders
}

func (o *Order) AddOrder(db *gorm.DB, order Order) (string, error) {
	db.Create(order)
	return "", nil
}
