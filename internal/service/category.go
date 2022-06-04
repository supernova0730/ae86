package service

import "ae86/internal/service/adapter"

type CategoryService struct {
	storage adapter.StorageContainer
}

func NewCategoryService(storage adapter.StorageContainer) *CategoryService {
	return &CategoryService{storage: storage}
}
