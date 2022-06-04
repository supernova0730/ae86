package model

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	IsDeleted bool `gorm:"column:is_deleted;not null;default:false"`
}
