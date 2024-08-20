package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;size:100" validate:"required,max=100"` 
	Email string `gorm:"unique;not null;size:100" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=8"`
}