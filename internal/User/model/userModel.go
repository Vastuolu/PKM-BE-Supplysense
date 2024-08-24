package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;size:255"`
	Username string `json:"username" gorm:"not null;size:100" validate:"required,max=100"` 
	Email string `json:"email" gorm:"unique;not null;size:100" validate:"required,email"`
	Password *string `validate:"required,min=8"`
	Provider string `json:"provider" gorm:"not null; size:100"`
	Firstname string `json:"firstname" gorm:"not null; size:256" validate:"required"`
	Lastname string `json:"lastname" gorm:"not null; size:256" validate:"required"`
	AvatarUrl *string `json:"avatarUrl" validate:"omitempty,url"`
	gorm.Model
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.NewString()
	if user.Provider == ""{
		user.Provider = "standard"
	}
    return
}