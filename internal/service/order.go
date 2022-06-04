package service

import "ae86/internal/service/adapter"

type OrderService struct {
	storage adapter.StorageContainer
}

func NewOrderService(storage adapter.StorageContainer) *OrderService {
	return &OrderService{storage: storage}
}
