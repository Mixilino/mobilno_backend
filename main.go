package main

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware)
	{
		protected.GET("/tasks", controller.GetAllTasks)
		protected.POST("/task", controller.CreateTask)
		protected.DELETE("/task/:task_id", controller.DeleteTask)
	}

	panic(router.Run(":8080"))
}
