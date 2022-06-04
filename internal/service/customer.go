package service

import "ae86/internal/service/adapter"

type CustomerService struct {
	storage adapter.StorageContainer
}

func NewCustomerService(storage adapter.StorageContainer) *CustomerService {
	return &CustomerService{storage: storage}
}
