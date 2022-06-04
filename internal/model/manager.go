package model

type Manager struct {
	Model
	Username  string `gorm:"column:username;unique;not null"`
	Password  string `gorm:"column:password;not null"`
	FirstName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name;not null"`
	Phone     string `gorm:"column:phone;not null"`
}

func (Manager) TableName() string {
	return "manager"
}
