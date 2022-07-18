package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"desc"`
	UserID      int        `json:"created_by"`
	Products    []*Product `json:"products" gorm:"many2many:product_categories;"`
}
