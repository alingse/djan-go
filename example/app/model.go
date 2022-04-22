package app

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name  string
	Email string
}

func (UserModel) TableName() string {
	return "users"
}
