package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Name     string
	Username string `gorm:"unique"`
	Age      json.Number
	Role     bool      `gorm:"default:0"`
	Articles []Article `gorm:"foreignKey:user_id"'`
}
