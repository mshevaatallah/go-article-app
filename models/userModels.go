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
	Articles []Article `gorm:"foreignKey:user_id"'`
}
