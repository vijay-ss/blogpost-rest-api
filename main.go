package main

import (
	"blogpostApi/controllers"
	"blogpostApi/iniitalizers"

	"github.com/gin-gonic/gin"
)

  func init() {
	iniitalizers.LoadEnvVariables()
	iniitalizers.ConnectToDB()
  }
  
  func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }