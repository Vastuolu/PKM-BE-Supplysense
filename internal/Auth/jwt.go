package auth

import (
	"supplysense/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtPayloadInterface struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatarUrl"`
	Provider string `json:"provider"`
}

type jwtClaimsInterface struct{
	jwtPayloadInterface
	jwt.RegisteredClaims
}

func makeJwtToken(jwtInterface *jwtPayloadInterface) (string, error){
	claims := jwtClaimsInterface{
		*jwtInterface,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.GetEnv("WEB_URL"),
			Subject:   jwtInterface.ID,
		},
	}
	signSecret := []byte(config.GetEnv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}