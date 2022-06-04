package storage

import "gorm.io/gorm"

type ProductStorage struct {
	db *gorm.DB
}

func NewProductStorage(db *gorm.DB) *ProductStorage {
	return &ProductStorage{db: db}
}
