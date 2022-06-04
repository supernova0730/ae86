package model

type Product struct {
	Model
	Title       string `gorm:"column:title;not null"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price;not null;check:price > 0"`
	Image       string `gorm:"column:image;not null"`
	IsActive    bool   `gorm:"column:is_active;not null;default:true"`
	CategoryID  uint   `gorm:"column:category_id"`

	Category *Category
}

func (Product) TableName() string {
	return "product"
}
