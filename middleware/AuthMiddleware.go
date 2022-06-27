package middleware

import (
	"backend/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := util.ExtractToken(c)
	token, err := util.TokenValid(tokenString)
	if err != nil {
		c.JSON(err.StatusCode, err)
		c.Abort()
		return
	}
	user, err := util.ExtractUserFromToken(token)
	if err != nil {
		c.JSON(err.StatusCode, err)
		c.Abort()
		return
	}
	c.Set("user", user)
	c.Next()
}
