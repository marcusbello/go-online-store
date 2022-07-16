package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}
