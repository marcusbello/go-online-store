package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ProductId    int       `json:"product_id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Category     string    `json:"category"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	LastModified time.Time `json:"last_modified"`
}
