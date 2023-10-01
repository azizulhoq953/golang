package controllers

import (
	"Go_3-Month/golang/go-crud/initializers"
	"Go_3-Month/golang/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get Data Of Req Body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Body, Body: body.Body}
	// post := models.Post{Title: "Hello", Body: "Post body"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get The Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with Them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get Id Of url
	id := c.Param("id")

	//Get The posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//Get The Data of the url

	id := c.Param("id")

	//Get The Data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find The Post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get The Id of the url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
