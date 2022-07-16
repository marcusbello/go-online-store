package models

import (
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	Name   string `json:"name" gorm:"default:member"`
	UserID uint
}

type User struct {
	gorm.Model
	UserName string   `json:"user_name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Token    string   `json:"token"`
	UserRole UserRole `json:"user_role"`
	Profile  Profile  `json:"profile"`
}

type Profile struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Zipcode   int    `json:"zipcode"`
}
