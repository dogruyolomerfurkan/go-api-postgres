package models

import "gorm.io/gorm"

type Post struct {
	//Gorm model has default values for entities
	gorm.Model
	Body  string
	Title string
}
