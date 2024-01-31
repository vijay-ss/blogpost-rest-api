package main

import (
	"blogpostApi/initializers"
	"blogpostApi/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	fmt.Println("Successfully performed migration.")
}