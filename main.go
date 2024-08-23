package main

import (
	"supplysense/config"
	"supplysense/database"
	"supplysense/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	validate *validator.Validate
)

func main(){
	config.LoadEnv()
	e := echo.New()
	config.GothSetup()
	validate = validator.New(validator.WithRequiredStructEnabled())
	routes.RegisterRoute(e)
	database.DBConnect()
	e.Logger.Fatal(e.Start(config.GetEnv("WEB_PORT")))
}