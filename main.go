package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	server := gin.Default()

	server.GET("/hello", hello)

	server.Run(":8080")
}
