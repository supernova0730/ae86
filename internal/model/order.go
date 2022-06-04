package model

import "ae86/internal/enums"

type Order struct {
	Model
	Address            string              `gorm:"column:address;not null"`
	State              enums.OrderState    `gorm:"column:state;not null"`
	PaymentMethod      enums.PaymentMethod `gorm:"column:payment_method;not null"`
	CancellationReason string              `gorm:"column:cancellation_reason"`
	CustomerID         int64               `gorm:"column:customer_id"`
	StoreID            int64               `gorm:"column:store_id"`

	Customer *Customer
	Store    *Store
}

func (Order) TableName() string {
	return "orders"
}
