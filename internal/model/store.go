package model

import "time"

type Store struct {
	Model
	Title            string        `gorm:"column:title"`
	Info             string        `gorm:"column:info"`
	Address          string        `gorm:"column:address"`
	Image            string        `gorm:"column:image"`
	AvgDeliveryTime  time.Duration `gorm:"column:avg_delivery_time"`
	WorkingHourBegin time.Time     `gorm:"column:working_hour_begin"`
	WorkingHourEnd   time.Time     `gorm:"column:working_hour_end"`
	MinOrderPrice    int           `gorm:"column:min_order_price"`
	DeliveryPrice    int           `gorm:"column:delivery_price"`
	ContactPhone     string        `gorm:"column:contact_phone"`
	ManagerID        int64         `gorm:"column:manager_id"`

	Manager *Manager
}

func (Store) TableName() string {
	return "store"
}
