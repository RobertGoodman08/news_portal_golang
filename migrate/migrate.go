package main

import (
	"tidy/initializers"
	"tidy/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}



func main(){
	initializers.DB.AutoMigrate(&models.Categories{})
	initializers.DB.AutoMigrate(&models.News{})
}