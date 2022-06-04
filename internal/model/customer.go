package model

type Customer struct {
	Model
	ExternalID uint   `gorm:"column:external_id;not null"`
	Username   string `gorm:"column:username"`
	Phone      string `gorm:"column:phone"`
	FirstName  string `gorm:"column:first_name"`
	LastName   string `gorm:"column:last_name"`
}

func (Customer) TableName() string {
	return "customer"
}
