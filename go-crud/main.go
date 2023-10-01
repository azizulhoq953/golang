package main

import (
	"Go_3-Month/golang/go-crud/controllers"
	"Go_3-Month/golang/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.Run() // listen and serve on 0.0.0.0:8080

}
