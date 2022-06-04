package adapter

type ServiceContainer interface {
	Manager() ManagerService
	Store() StoreService
	Category() CategoryService
	Product() ProductService
	Customer() CustomerService
	Order() OrderService
	OrderItem() OrderItemService
}

type ManagerService interface{}

type StoreService interface{}

type CategoryService interface{}

type ProductService interface{}

type CustomerService interface{}

type OrderService interface{}

type OrderItemService interface{}
