package main

import (
	"github.com/labstack/echo"
	"supplysense/database"
)

func main(){
	e := echo.New()

	database.DBConnect()
	e.Logger.Error(e.Start(":3000"))
}