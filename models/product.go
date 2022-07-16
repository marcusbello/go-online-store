package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       string   `json:"price"`
	Stock       uint     `json:"stock"`
	Category    Category `json:"category" gorm:"foreignKey:ID;"`
	CreatedBy   string   `json:"created_by"`
}
