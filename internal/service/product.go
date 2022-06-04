package service

import "ae86/internal/service/adapter"

type ProductService struct {
	storage adapter.StorageContainer
}

func NewProductService(storage adapter.StorageContainer) *ProductService {
	return &ProductService{storage: storage}
}
