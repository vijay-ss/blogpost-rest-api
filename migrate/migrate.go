package main

import (
	"blogpostApi/iniitalizers"
	"blogpostApi/models"
	"fmt"
)

func init() {
	iniitalizers.LoadEnvVariables()
	iniitalizers.ConnectToDB()
}

func main() {
	iniitalizers.DB.AutoMigrate(&models.Post{})
	fmt.Println("Successfully performed migration.")
}