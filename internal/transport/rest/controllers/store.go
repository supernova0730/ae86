package controllers

import "ae86/internal/transport/adapter"

type StoreController struct {
	service adapter.ServiceContainer
}

func NewStoreController(service adapter.ServiceContainer) *StoreController {
	return &StoreController{service: service}
}
