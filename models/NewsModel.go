package models


import "gorm.io/gorm"


type Categories struct {
  gorm.Model
  Title string
}



type News struct {
	gorm.Model
	Title string
	Content string
	Photo string
	IsPublished *bool
	CategoryRefer int
 	Category  Categories `gorm:"foreignKey:CategoryRefer"`
	Views int
}




