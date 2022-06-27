package main

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", controller.Ping)
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)

	panic(router.Run(":8080"))
}
