package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	CategoryId   int       `json:"category_id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
}
