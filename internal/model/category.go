package model

type Category struct {
	Model
	Title   string `gorm:"column:title;not null"`
	StoreID int64  `gorm:"column:store_id"`

	Store *Store
}

func (Category) TableName() string {
	return "category"
}
