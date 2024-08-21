package auth

import (
	"log"
	"supplysense/config"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)
func init(){
		config.LoadEnv()
		sessionStore := sessions.NewCookieStore([]byte(config.GetEnv("SECRET_KEY")))
		sessionStore.Options = &sessions.Options{
			Path:     "/",               // Path untuk cookie
			MaxAge:   3600 * 30,             // Durasi cookie dalam detik
			HttpOnly: true,             // Hanya bisa diakses melalui HTTP
			Secure:   false,            // Gunakan HTTPS
		}
		goth.UseProviders(
			google.New(config.GetEnv("GOOGLE_CLIENT"), config.GetEnv("GOOGLE_SECRET"), config.GetEnv("WEB_URL")+config.GetEnv("WEB_PORT")+"/api/login/google/callback", "email", "profile"),
		)
		gothic.Store = sessionStore
		
	}

func Login(c echo.Context) error {
	reqWContext := c.Request()
	reqWContext = gothic.GetContextWithProvider(reqWContext, c.Param("provider"))
	urlAuth, err := gothic.GetAuthURL(c.Response().Writer,reqWContext)
	if err != nil {
		log.Printf("error: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusTemporaryRedirect, urlAuth)
}

func LoginCallback(c echo.Context) error {
    reqWContext := gothic.GetContextWithProvider(c.Request(), c.Param("provider"))
	user, err := gothic.CompleteUserAuth(c.Response().Writer, reqWContext)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}