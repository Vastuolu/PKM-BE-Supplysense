package service

import (
	"supplysense/database"
	"supplysense/internal/User/model"

)

func GetAllUsers()([]model.User, error){
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil{
		return nil, err
	}
	return users, nil
}

func Register(user *model.User) error {
	return database.DB.Create(user).Error
}



