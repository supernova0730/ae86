package controllers

import "ae86/internal/transport/adapter"

type OrderController struct {
	service adapter.ServiceContainer
}

func NewOrderController(service adapter.ServiceContainer) *OrderController {
	return &OrderController{service: service}
}
