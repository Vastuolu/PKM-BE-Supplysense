package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"supplysense/config"
	"supplysense/database"
	"supplysense/internal/User/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)


type jwtTokenInterface struct{
	ID string
	Username string
	Email string
	Avatar string
	Provider string
	jwt.RegisteredClaims
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
	//complete user login and take information
    reqWContext := gothic.GetContextWithProvider(c.Request(), c.Param("provider"))
	user, err := gothic.CompleteUserAuth(c.Response().Writer, reqWContext)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := userFindorCreate(&user); err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	//make jwt token
	claims := jwtTokenInterface{
			user.UserID,
			user.NickName,
			user.Email,
			user.AvatarURL,
			user.Provider,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    config.GetEnv("WEB_URL"),
				Subject:   user.UserID,
			},

	}

	signedToken,err := makeJwtToken(&claims)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, signedToken)
}


func makeJwtToken(claims *jwtTokenInterface) (string, error){
	signSecret := []byte(config.GetEnv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signSecret)
	if err != nil {
		log.Printf("error: %v", err)
		return "", err
	}
	return signedToken, nil
}

func userFindorCreate(userGoth *goth.User) error {
	avatarURL := userGoth.AvatarURL
	user := &model.User{
		ID: userGoth.UserID,
		Username: userGoth.NickName,
		Email: userGoth.Email,
		Provider: userGoth.Provider,
		Firstname: userGoth.FirstName,
		Lastname: userGoth.LastName,
		AvatarUrl: &avatarURL,
	}
	result := database.DB.Where(&model.User{ID: userGoth.UserID, Provider: userGoth.Provider}).FirstOrCreate(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}