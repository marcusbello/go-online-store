package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"desc"`
	Price       string      `json:"price"`
	Categories  []*Category `json:"categories" gorm:"many2many:product_categories;"`
	UserID      int         `json:"created_by" gorm:"default:2"`
}
