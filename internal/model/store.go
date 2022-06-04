package model

import "time"

type Store struct {
	Model
	Title            string        `gorm:"column:title;not null"`
	Info             string        `gorm:"column:info"`
	Address          string        `gorm:"column:address;not null"`
	Image            string        `gorm:"column:image;not null"`
	AvgDeliveryTime  time.Duration `gorm:"column:avg_delivery_time;not null"`
	WorkingHourBegin time.Time     `gorm:"column:working_hour_begin;not null"`
	WorkingHourEnd   time.Time     `gorm:"column:working_hour_end;not null"`
	MinOrderPrice    int           `gorm:"column:min_order_price;check:min_order_price > 0;not null"`
	DeliveryPrice    int           `gorm:"column:delivery_price;check:delivery_price > 0;not null"`
	ContactPhone     string        `gorm:"column:contact_phone;not null"`
	ManagerID        int64         `gorm:"column:manager_id"`

	Manager *Manager
}

func (Store) TableName() string {
	return "store"
}
