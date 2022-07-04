package util

import (
	"backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ExtractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(tokenString string) (*jwt.Token, *RestError) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	// Proverava da li je token validan i dalje
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, &RestError{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		}
	}
	return token, nil

}
func TokenValid(tokenString string) (*jwt.Token, *RestError) {
	if tokenString == "" {
		return nil, NewRestErrUnauthorized()
	}
	token, err := verifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractUserFromToken(token *jwt.Token) (*model.User, *RestError) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, NewRestErrBadRequest("No userId")
		}
		username, ok := claims["username"].(string)
		if !ok {
			return nil, NewRestErrBadRequest("No username")
		}
		return &model.User{
			Model: gorm.Model{
				ID: uint(userId),
			},
			Username: username,
		}, nil
	}
	return nil, NewRestErrInternalServerError()
}

func CreateToken(user model.User) (string, *RestError) {
	secret := os.Getenv("JWT_SECRET")
	// AccessToken
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", NewRestErrInternalServerError()
	}
	return accessTokenString, nil
}
