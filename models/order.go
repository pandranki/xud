package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	PairID string
	PeerID string
	Quantity string
	Price string
}

func (o *Order) GetOrders(db *gorm.DB) []Order {
	orders := []Order{}
	db.Find(&orders)
	return orders
}

func (o *Order) AddOrder(db *gorm.DB,order Order) (string ,error){
	db.Create(order)
	return "",nil
}