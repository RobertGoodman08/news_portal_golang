package main 


import (
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	"tidy/initializers"
	"tidy/controllers"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func main() {


	r := gin.Default()


	// Categories
	r.GET("/category", controllers.CategoriesIndex)
	r.GET("/show/:id", controllers.CategoriesShow)
	r.POST("/post", controllers.CategoriesCreate)
	r.PUT("/update/:id", controllers.CategoriesUpdate)
	r.DELETE("/delete/:id", controllers.CategoriesDelete)


	// News
	r.GET("/", controllers.NewsPage)
	r.GET("/show_news/:id", controllers.NewsDetail)
	r.POST("/post_news", controllers.NewsCreate)
	r.PUT("/update_news/:id", controllers.NewsUpdate)
	r.DELETE("/delete_news/:id", controllers.NewsDelete)

	

	r.Run() 
}