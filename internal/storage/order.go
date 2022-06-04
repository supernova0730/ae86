package storage

import "gorm.io/gorm"

type OrderStorage struct {
	db *gorm.DB
}

func NewOrderStorage(db *gorm.DB) *OrderStorage {
	return &OrderStorage{db: db}
}
