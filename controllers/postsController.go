package controllers

import (
	"blogpostApi/initializers"
	"blogpostApi/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data from request body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with retireved posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// respond with retireved posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id from URL
	id := c.Param("id")

	// get data from request body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	// find post being updated
	var post models.Post
	initializers.DB.First(&post, id)

	// update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// respond with update
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get id from URL
	id := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)
}