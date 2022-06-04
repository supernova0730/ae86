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

type CategoryStorage interface{}

type ProductStorage interface{}

type CustomerStorage interface{}

type OrderStorage interface{}

type OrderItemStorage interface{}
