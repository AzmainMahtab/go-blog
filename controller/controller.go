// Package controller
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AzmainMahtab/go-blog/models"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.Create(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result, err := models.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		return
	}
	posts = *result

	c.JSON(http.StatusOK, gin.H{"message": "Posts retrieved successfully", "data": posts})
}
