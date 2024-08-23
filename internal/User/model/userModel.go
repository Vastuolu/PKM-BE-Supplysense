package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;size:255"`
	Username string `gorm:"not null;size:100" validate:"required,max=100"` 
	Email string `gorm:"unique;not null;size:100" validate:"required,email"`
	Password *string `validate:"min=8"`
	Provider string `gorm:"not null; size:100"`
	Firstname string `gorm:"not null; size:256"`
	Lastname string `gorm:"not null; size:256"`
	AvatarUrl *string `validate:url`
	gorm.Model
}