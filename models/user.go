package models

import (
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	ID   uint   `json:"ID"`
	Name string `json:"name"`
}

type User struct {
	gorm.Model
	UserName  string      `json:"user_name"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone"`
	Address   UserAddress `json:"address" gorm:"embedded"`
	Password  string      `json:"password"`
	Token     string      `json:"token"`
	UserRole  UserRole    `json:"user_role" gorm:"foreignKey:ID;"`
}

type UserAddress struct {
	gorm.Model
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zipcode int    `json:"zipcode"`
}
