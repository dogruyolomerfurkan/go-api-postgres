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

//https://youtu.be/lf_kiH_NPvM?t=1352
