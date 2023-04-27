package controllers 


import (
	// "time"
	"github.com/gin-gonic/gin"
	"tidy/models"
	"tidy/initializers"
)



func CategoriesCreate(c *gin.Context){ // обрабатывает запрос на создание новой категории


	var news struct { // создаем пустую структуру news, которая будет заполнена данными запроса
		Title string
	}


	c.Bind(&news) // привязываем данные запроса к структуре news

	category := models.Categories{Title: news.Title} // создаем новую переменную category типа models.Categories с заданным названием

	result := initializers.DB.Create(&category)  // создаем новую запись в базе данных для категории


	// проверяем наличие ошибок при создании записи
	if result.Error != nil {
		c.Status(400)
		return 
	}


	c.JSON(200, gin.H{
		"category": category,
	})
}



func CategoriesIndex(c *gin.Context) { // запрос на получение списка всех категорий
	
	
	var news []models.Categories // создаем пустой слайс news, который будет заполнен данными из базы данных
	initializers.DB.Find(&news) // находим все записи в таблице категорий и заполняем ими слайс news

	c.JSON(200, gin.H{
		"news": news,
	})
}


func CategoriesShow(c *gin.Context){

	id := c.Param("id")

	var new models.Categories
	result := initializers.DB.First(&new, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Страница не найдена",
		})
		return
	}
	

	c.JSON(200, gin.H{
		"new": new,
	})
}



 

func CategoriesUpdate(c *gin.Context){ //  запрос на обновление данных категории
	// извлекаем id из параметров запроса
	id := c.Param("id")

	// создаем пустую структуру news, которая будет заполнена данными запроса
	var news struct {
		Title string
	}


	c.Bind(&news) // привязываем данные запроса к структуре news


	// создаем новую переменную new типа models.Categories и находим категорию с заданным id
	var new models.Categories
	initializers.DB.First(&new, id)


	// обновляем данные в модели new, указывая только те поля, которые были изменены
	initializers.DB.Model(&new).Updates(models.Categories{
		Title: news.Title,
	})


	
	c.JSON(200, gin.H{
		"news": news,
	})

}


func CategoriesDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Categories{}, id)

	c.Status(200)
}