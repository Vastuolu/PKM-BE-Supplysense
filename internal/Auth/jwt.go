package auth

import (
	"log"
	"supplysense/config"

	"github.com/golang-jwt/jwt/v5"
)

type jwtTokenInterface struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	Provider string `json:"provider"`
	jwt.RegisteredClaims
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