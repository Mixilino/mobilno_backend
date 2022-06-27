package util

import (
	"backend/model"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func CreateToken(user model.User) (string, *RestError) {
	secret := os.Getenv("JWT_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", NewRestErrInternalServerError()
	}
	return token, nil
}
