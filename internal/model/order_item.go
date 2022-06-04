package model

type OrderItem struct {
	ID        uint `gorm:"column:id"`
	Amount    int  `gorm:"column:amount;not null;default:1;check:amount > 0"`
	OrderID   uint `gorm:"column:order_id"`
	ProductID uint `gorm:"column:product_id"`

	Order   *Order
	Product *Product
}

func (OrderItem) TableName() string {
	return "order_item"
}
