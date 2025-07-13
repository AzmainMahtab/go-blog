// Package controller
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AzmainMahtab/go-blog/models"
	"github.com/AzmainMahtab/go-blog/pkg/handler"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		handlers.Error(c, err)
		return
	}

	if err := models.Create(&post); err != nil {
		handlers.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result, err := models.GetPosts()
	if err != nil {
		handlers.Error(c, err)
		return
	}
	posts = *result

	handlers.Success(c, posts, nil, "Post created successfully")
}
