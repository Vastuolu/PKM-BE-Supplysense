package routes

import (
	userHandler "supplysense/internal/User/handler"
	"log"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(echo *echo.Echo) {
    log.Println("register runned")
	Api := echo.Group("/api")
	Api.GET("/user", userHandler.GetAllUsers) // Debugging
	Api.POST("/user", userHandler.Register)   // Debugging
	// Api.GET("/user")
	// Api.GET("/user")
}
