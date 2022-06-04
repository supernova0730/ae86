package model

type OrderItem struct {
	ID        int64 `gorm:"column:id"`
	Amount    int   `gorm:"column:amount;not null;default:1;check:amount > 0"`
	OrderID   int64 `gorm:"column:order_id"`
	ProductID int64 `gorm:"column:product_id"`

	Order   *Order
	Product *Product
}

func (OrderItem) TableName() string {
	return "order_item"
}
