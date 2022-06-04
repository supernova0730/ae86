package controllers

import "ae86/internal/transport/adapter"

type ProductController struct {
	service adapter.ServiceContainer
}

func NewProductController(service adapter.ServiceContainer) *ProductController {
	return &ProductController{service: service}
}
