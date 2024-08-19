package route

import (
	userHandler "supplysense/api/handlers"

	"github.com/labstack/echo"
)

func RegisterRoute(echo *echo.Echo) {
	Api := echo.Group("api/")
	Api.GET("/user", userHandler.GetAllUser) // Debugging
	Api.GET("/user") // Debugging
	Api.GET("/user")
	Api.GET("/user")
}