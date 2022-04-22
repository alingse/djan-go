package app

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UserModel) TableName() string {
	return "users"
}
