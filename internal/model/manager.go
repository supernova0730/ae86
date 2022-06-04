package model

type Manager struct {
	Model
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Phone     string `gorm:"column:phone"`
}

func (Manager) TableName() string {
	return "manager"
}
