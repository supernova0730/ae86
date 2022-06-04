package controllers

import "ae86/internal/transport/adapter"

type CustomerController struct {
	service adapter.ServiceContainer
}

func NewCustomerController(service adapter.ServiceContainer) *CustomerController {
	return &CustomerController{service: service}
}
