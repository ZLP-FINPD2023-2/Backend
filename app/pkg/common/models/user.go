package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Age      uint8
	Gender   bool
	Email    string
	Password string
}
