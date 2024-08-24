package auth

import (
	"net/http"
	"supplysense/database"
	"supplysense/helper"
	"supplysense/internal/User/model"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)


func LoginProvider(c echo.Context) error {
	reqWContext := c.Request()
	reqWContext = gothic.GetContextWithProvider(reqWContext, c.Param("provider"))
	urlAuth, err := gothic.GetAuthURL(c.Response().Writer,reqWContext)
	if err != nil {
		resMap := helper.JsonResponse(500, nil, 0,err)
		return c.JSON(500, resMap)
	}
	resMap := helper.JsonResponse(200,  helper.InterfaceMaker("urlAuth", urlAuth), 1,nil)
	return c.JSON(http.StatusOK,resMap)
}

func LoginProviderCallback(c echo.Context) error {
	//complete user login and take information
    reqWContext := gothic.GetContextWithProvider(c.Request(), c.Param("provider"))
	user, err := gothic.CompleteUserAuth(c.Response().Writer, reqWContext)
	if err != nil {
		resMap := helper.JsonResponse(500, nil, 0, err)
		return c.JSON(500, resMap)
	}

	if err := userFindorCreate(&user); err != nil{
		resMap := helper.JsonResponse(500, nil,0,err)
		return c.JSON(500, resMap)
	}

	//make jwt token
	claims := jwtPayloadInterface{
			user.UserID,
			user.NickName,
			user.Email,
			user.AvatarURL,
			user.Provider,
	}

	signedToken,err := makeJwtToken(&claims)
	if err != nil {
		resMap := helper.JsonResponse(500, nil, 0, err)
		return c.JSON(500, resMap)
	}

	resMap := helper.JsonResponse(200, helper.InterfaceMaker("token", signedToken),1, nil)
	return c.JSON(200, resMap)
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