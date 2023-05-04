package db

import (
	"stockify/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB - setup database for sharing to all api
func SetupDB() {
	database, err := gorm.Open(sqlite.Open("stock.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.User{}) // pass by reference
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})

	db = database
}
