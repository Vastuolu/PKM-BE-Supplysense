package handler

import (
	"net/http"
	"supplysense/internal/User/model"
	"supplysense/internal/User/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func GetAllUsers(c echo.Context) error {
	users, err := service.GetAllUsers()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Validation failed",
            "error":   err.Error(),
        })
    }
    return c.JSON(http.StatusOK, users)
}

func Register(c echo.Context) error {
	var user model.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Invalid input",
            "error":   err.Error(),
        })
    }

	// if err := validate.Struct(&user); err != nil {
    //     return c.JSON(http.StatusBadRequest, map[string]string{
    //         "message": "Validation failed",
    //         "error":   err.Error(),
    //     })
    // }

    err := service.Register(&user)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "message": "Validation failed",
            "error":   err.Error(),
        })
    }
	return c.JSON(http.StatusCreated, user)
}