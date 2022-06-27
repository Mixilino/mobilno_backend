package main

import (
	"backend/controller"
	"backend/middleware"
	"backend/model"
	"backend/util"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", controller.Ping)
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware)
	{
		protected.GET("/task", func(c *gin.Context) {
			u, ok := c.Get("user")
			if !ok {
				c.JSON(util.NewRestErrInternalServerError().StatusCode, util.NewRestErrInternalServerError())
			}
			user := u.(*model.User)
			c.JSON(200, gin.H{"id": user.ID})
		})
	}

	panic(router.Run(":8080"))
}
