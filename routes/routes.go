package routes

import (
	"log"
	auth "supplysense/internal/Auth"
	userHandler "supplysense/internal/User/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(echo *echo.Echo) {
    log.Println("register runned")
	Api := echo.Group("/api")
	
	//Standard Auth Route
	Api.POST("/login", auth.LoginStandard)
	Api.POST("/register", auth.RegisterStandard)
	
	//Provider Auth Route
	Api.GET("/login/:provider", auth.LoginProvider)
	Api.GET("/login/:provider/callback", auth.LoginProviderCallback)
	
	Api.POST("/user", userHandler.Register)   // Debugging
	Api.GET("/user", userHandler.GetAllUsers) // Debugging
}
