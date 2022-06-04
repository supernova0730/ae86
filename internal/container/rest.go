package container

import (
	"ae86/internal/transport/adapter"
	"ae86/internal/transport/rest/controllers"
)

type RestContainer struct {
	manager  *controllers.ManagerController
	store    *controllers.StoreController
	category *controllers.CategoryController
	product  *controllers.ProductController
	customer *controllers.CustomerController
	order    *controllers.OrderController
}

func NewRestContainer(serviceContainer adapter.ServiceContainer) *RestContainer {
	return &RestContainer{
		manager:  controllers.NewManagerController(serviceContainer),
		store:    controllers.NewStoreController(serviceContainer),
		category: controllers.NewCategoryController(serviceContainer),
		product:  controllers.NewProductController(serviceContainer),
		customer: controllers.NewCustomerController(serviceContainer),
		order:    controllers.NewOrderController(serviceContainer),
	}
}

func (c *RestContainer) Manager() *controllers.ManagerController {
	return c.manager
}

func (c *RestContainer) Store() *controllers.StoreController {
	return c.store
}

func (c *RestContainer) Category() *controllers.CategoryController {
	return c.category
}

func (c *RestContainer) Product() *controllers.ProductController {
	return c.product
}

func (c *RestContainer) Customer() *controllers.CustomerController {
	return c.customer
}

func (c *RestContainer) Order() *controllers.OrderController {
	return c.order
}
