package controller

import (
	"backend/model"
	"backend/service"
	"backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := util.NewRestErrBadRequest("Invalid body data. No username and/or password.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	token, restErr := service.SignIn(user)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"jwt": token})

}

func SignUp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := util.NewRestErrBadRequest("Invalid body data. No username and/or password.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	str, err := service.SignUp(user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": str})

}
