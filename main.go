package main

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}},
	))
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)

	protected := router.Group("/task")
	protected.Use(middleware.AuthMiddleware)
	{
		protected.GET("/all", controller.GetAllTasks)
		protected.GET("/:task_id", controller.GetTask)
		protected.POST("/", controller.CreateTask)
		protected.DELETE("/:task_id", controller.DeleteTask)
		protected.PUT("/:task_id", controller.ModifyTask)
	}
	panic(router.Run(":8080"))
}
