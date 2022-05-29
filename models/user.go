package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName"`
	Email string `json:"email" validate:"required,email"` 
	Phone string `json:"phone"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

func (b *User) TableName() string {
	return "users"
}