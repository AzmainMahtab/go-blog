package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AzmainMahtab/go-blog/db"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	db.ConnectDB() // Initialize the database connection
	server := gin.Default()

	server.GET("/hello", hello)

	server.Run(":8080")
}
