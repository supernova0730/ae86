package storage

import "gorm.io/gorm"

type OrderItemStorage struct {
	db *gorm.DB
}

func NewOrderItemStorage(db *gorm.DB) *OrderItemStorage {
	return &OrderItemStorage{db: db}
}
