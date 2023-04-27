package controllers 


import (
	"github.com/gin-gonic/gin"
	"tidy/models"
	"tidy/initializers"
)


func NewsCreate(c *gin.Context) { // Создание новостей
    var news struct {
        Title         string
        Content       string
        Photo         string
        IsPublished   *bool
        CategoryRefer int
        Views         int
    }

    c.Bind(&news)

    news_s := models.News{
        Title:         news.Title,
        Content:       news.Content,
        Photo:         news.Photo,
        IsPublished:   news.IsPublished,
        CategoryRefer: news.CategoryRefer,
        Views:         news.Views,
    }

    result := initializers.DB.Create(&news_s)

    if result.Error != nil {
		c.Status(400)
		return 
	}

    c.JSON(200, gin.H{
        "news_s": news_s,
    })
}


func NewsPage(c *gin.Context) { // Страница с новостями
	
	
	var news []models.News 
	initializers.DB.Find(&news) 

	c.JSON(200, gin.H{
		"news": news,
	})
}


func NewsDetail(c *gin.Context){ // подробно о новостей

	id := c.Param("id")

	var new models.News
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



 

func NewsUpdate(c *gin.Context){ // Редактировать новости
	// извлекаем id из параметров запроса
	id := c.Param("id")


	var news struct {
        Title         string
        Content       string
        Photo         string
        IsPublished   *bool
        CategoryRefer int
        Views         int
    }


	c.Bind(&news) 



	var new models.News
	initializers.DB.First(&new, id)


	// обновляем данные в модели new, указывая только те поля, которые были изменены
	initializers.DB.Model(&new).Updates(models.News{
		Title:         news.Title,
        Content:       news.Content,
        Photo:         news.Photo,
        IsPublished:   news.IsPublished,
        CategoryRefer: news.CategoryRefer,
        Views:         news.Views,
	})


	
	c.JSON(200, gin.H{
		"news": news,
	})

}


func NewsDelete(c *gin.Context) { // Удаление новостей
	id := c.Param("id")

	initializers.DB.Delete(&models.News{}, id)

	c.Status(200)
}