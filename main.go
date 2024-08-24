package main

import (
	"fmt"
	"supplysense/config"
	"supplysense/database"
	"supplysense/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct{
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs.Error())
	  // Optionally, you could return the error to give each route more control over the status code
	  return errs
	}
	return nil
  }

func main(){
	config.LoadEnv()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	config.GothSetup()
	routes.RegisterRoute(e)
	database.DBConnect()
	e.Logger.Fatal(e.Start(config.GetEnv("WEB_PORT")))
}