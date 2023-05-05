package db

import (
	"stockify/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB - setup database for sharing to all api
func SetupDB() {
	dsn := "root:admin1234@tcp(127.0.0.1:3306)/stockify?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.User{}) // pass by reference
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})

	db = database
}
