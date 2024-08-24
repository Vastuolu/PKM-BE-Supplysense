package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;size:255"`
	Username string `json:"username" gorm:"not null;size:100" validate:"required_if=IsRegister true,max=100"` 
	Email string `json:"email" gorm:"unique;not null;size:100" validate:"required,email"`
	Password *string `validate:"required,min=8"`
	Provider string `json:"provider" gorm:"not null; size:100"`
	Firstname string `json:"firstname" gorm:"not null; size:256" validate:"required_if=IsRegister true"`
	Lastname string `json:"lastname" gorm:"not null; size:256" validate:"required_if=IsRegister true"`
	AvatarUrl *string `json:"avatarUrl" validate:"omitempty,url"`
	IsRegister bool `gorm:"-" json:"-" validate:"-"`
	gorm.Model
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.NewString()
	if user.Provider == ""{
		user.Provider = "standard"
	}
    return
}