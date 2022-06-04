package model

type OrderItem struct {
	ID        int64 `gorm:"column:id"`
	Amount    int   `gorm:"column:amount"`
	OrderID   int64 `gorm:"column:order_id"`
	ProductID int64 `gorm:"column:product_id"`

	Order   *Order
	Product *Product
}

func (OrderItem) TableName() string {
	return "order_item"
}
