package storage

import "gorm.io/gorm"

type CategoryStorage struct {
	db *gorm.DB
}

func NewCategoryStorage(db *gorm.DB) *CategoryStorage {
	return &CategoryStorage{db: db}
}
