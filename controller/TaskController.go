package controller

import (
	"backend/model"
	"backend/service"
	"backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var user *model.User
	user, err := extractUser(c)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	var task model.Task
	task.UserID = user.ID
	if err := c.ShouldBindJSON(&task); err != nil {
		fmt.Println(err.Error())
		restErr := util.NewRestErrBadRequest("Invalid body data for task.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	if err := service.CreateTask(&task); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"task": "created"})
}

func GetAllTasks(c *gin.Context) {
	var user *model.User
	user, err := extractUser(c)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	tasks, err := service.GetAllTasks(user.ID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": model.ConvertTaskArrayToDto(tasks)})
}

func extractUser(c *gin.Context) (*model.User, *util.RestError) {
	u, ok := c.Get("user")
	if !ok {
		return nil, util.NewRestErrInternalServerError()
	}
	user := u.(*model.User)
	return user, nil
}
