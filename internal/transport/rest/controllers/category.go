package controllers

import "ae86/internal/transport/adapter"

type CategoryController struct {
	service adapter.ServiceContainer
}

func NewCategoryController(service adapter.ServiceContainer) *CategoryController {
	return &CategoryController{service: service}
}
