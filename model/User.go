package model

import "time"

// https://gorm.io/docs/models.html
// struct tags

type User struct {
	ID        uint      `gorm:"primary_key"`
	Username  string    `gorm:"unique" from:"username" binding:"required"`
	Password  string    `form:"password" binding:"required"`
	Level     string    `gorm:"default:normal"`
	CreatedAt time.Time `gorm:"autoCreatedTime"`
}
