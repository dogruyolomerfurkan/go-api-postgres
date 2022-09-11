package controllers

import (
	"apiexample/initializers"
	"apiexample/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(ctx *gin.Context) {
	//Get data off request body
	var newPost models.Post

	//Create post
	if err := ctx.BindJSON(&newPost); err != nil {
		return
	}

	// post := models.Post{Body: "Merhaba", Title: "My title"}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	//Return
	ctx.JSON(200, gin.H{"post": newPost})
}

func PostGetList(ctx *gin.Context) {

	var posts []models.Post
	//Basic Dto definition
	type GetPostDto struct {
		ID    int
		Title string
		Body  string
	}

	var postList []GetPostDto
	// initializers.DB.Select("id", "title", "body").Where("id = ?", 2).Find(&posts)
	initializers.DB.Select("id", "title", "body").Find(&posts)
	for i := 0; i < len(posts); i++ {
		newPost := GetPostDto{
			ID:    int(posts[i].ID),
			Title: posts[i].Title,
			Body:  posts[i].Body,
		}
		postList = append(postList, newPost)
	}
	ctx.JSON(200, gin.H{"posts": postList})
}

//https://youtu.be/lf_kiH_NPvM?t=1352
