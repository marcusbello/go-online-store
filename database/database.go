package database

import (
	"go-online-store/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbMigrate := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{})
	if dbMigrate != nil {
		panic(dbMigrate)
	}
	DB = db
}
