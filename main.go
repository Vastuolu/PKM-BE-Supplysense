package main

import (
	"fmt"
	"supplysense/database"
	"supplysense/routes"
    "github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
)
var validate *validator.Validate

func main(){
	e := echo.New()
	e.Debug = true
	validate = validator.New(validator.WithRequiredStructEnabled())
	routes.RegisterRoute(e)
	database.DBConnect()
	e.Logger.Error(e.Start(":3000"))
	fmt.Println("test2")
}