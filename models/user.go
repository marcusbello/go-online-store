package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Token    string  `json:"token"`
	UserRole string  `json:"user_role" gorm:"default:member"`
	Profile  Profile `json:"profile"`
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
