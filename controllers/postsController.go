package controllers

import (
	"backend/initializers"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func PostCreat(c *gin.Context) {

	// get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it
	c.JSON(200, gin.H{
		"post": "post",
	})
}
