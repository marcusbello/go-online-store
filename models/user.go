package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserID       int       `json:"user_id" gorm:"primary_key"`
	Username     string    `json:"username"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Token        string    `json:"token"`
	UserRole     string    `json:"user_role"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
}
