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
	Api.GET("/user", userHandler.GetAllUsers) // Debugging
	Api.POST("/user", userHandler.Register)   // Debugging
	Api.GET("/login/:provider", auth.Login)
	Api.GET("/login/:provider/callback", auth.LoginCallback)
}
