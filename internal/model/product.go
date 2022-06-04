package model

type Product struct {
	Model
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
	Image       string `gorm:"column:image"`
	IsActive    bool   `gorm:"column:is_active"`
	CategoryID  int64  `gorm:"column:category_id"`

	Category *Category
}

func (Product) TableName() string {
	return "product"
}
