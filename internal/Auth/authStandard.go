package auth

import (
	"errors"
	"fmt"
	"supplysense/database"
	"supplysense/helper"
	"supplysense/internal/User/model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginStandard(c echo.Context) error {
	return nil
}

func RegisterStandard(c echo.Context) error {
	var user model.User
	var temp model.User

	if err := c.Bind(&user); err != nil{
		resMap := helper.JsonResponse(500, nil, 0, err)
		return c.JSON(500, resMap)
	}

	result := database.DB.Where(&model.User{Email: user.Email, Provider: "standard"}).First(&temp)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		resMap := helper.JsonResponse(422, nil, 0, "Email already Registered")
		return c.JSON(422, resMap)
	}
		
	if err := c.Validate(&user); err != nil {
		resMap := helper.JsonResponse(400, nil, 0, helper.MapValidationErr(err))
		return c.JSON(400, resMap)
	}

	rawHashedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil{
		resMap := helper.JsonResponse(500, nil, 0, err.Error())
		return c.JSON(500, resMap)		
	}
	hashedPassword := string(rawHashedPassword)
	user.Password = &hashedPassword
	fmt.Println(user)
	createdData := database.DB.Create(&user)
	if createdData.Error != nil {
		resMap := helper.JsonResponse(500, nil, 0, createdData.Error)
		return c.JSON(500, resMap)
	}
	resMap := helper.JsonResponse(201, nil, 0, nil)
	return c.JSON(201, resMap)
}
