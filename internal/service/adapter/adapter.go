package adapter

import "ae86/internal/model"

type StorageContainer interface {
	Manager() ManagerStorage
	Store() StoreStorage
	Category() CategoryStorage
	Product() ProductStorage
	Customer() CustomerStorage
	Order() OrderStorage
	OrderItem() OrderItemStorage
}

type ManagerStorage interface {
	GetByID(id uint) (result model.Manager, err error)
	GetByUsername(username string) (result model.Manager, err error)
	Create(manager model.Manager) (id uint, err error)
	Update(id uint, manager model.Manager) (err error)
	Delete(id uint) (err error)
}

type StoreStorage interface {
	GetByID(id uint) (result model.Store, err error)
	Create(store model.Store) (id uint, err error)
	Update(id uint, store model.Store) (err error)
	Delete(id uint) (err error)
}

type CategoryStorage interface {
	GetByID(id uint) (result model.Category, err error)
	GetAllByStoreID(storeID uint) (result []model.Category, err error)
	Create(category model.Category) (id uint, err error)
	Update(id uint, category model.Category) (err error)
	Delete(id uint) (err error)
}

type ProductStorage interface {
	GetByID(id uint) (result model.Product, err error)
	GetAllBy(filter ProductFilter) (result []model.Product, err error)
	Create(product model.Product) (id uint, err error)
	Update(id uint, product model.Product) (err error)
	Delete(id uint) (err error)
}

type CustomerStorage interface {
	GetByID(id uint) (result model.Customer, err error)
	GetByExternalID(externalID uint) (result model.Customer, err error)
	Create(customer model.Customer) (id uint, err error)
}

type OrderStorage interface{}

type OrderItemStorage interface{}
