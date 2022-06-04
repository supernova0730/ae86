package controllers

import "ae86/internal/transport/adapter"

type ManagerController struct {
	service adapter.ServiceContainer
}

func NewManagerController(service adapter.ServiceContainer) *ManagerController {
	return &ManagerController{service: service}
}
