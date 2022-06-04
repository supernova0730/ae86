package model

type Order struct {
	Model
	Address            string `gorm:"column:address"`
	State              string `gorm:"column:state"`
	PaymentMethod      string `gorm:"column:payment_method"`
	CancellationReason string `gorm:"column:cancellation_reason"`
	CustomerID         int64  `gorm:"column:customer_id"`
	StoreID            int64  `gorm:"column:store_id"`

	Customer *Customer
	Store    *Store
}

func (Order) TableName() string {
	return "orders"
}
