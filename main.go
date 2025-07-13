package main

import (
	"github.com/AzmainMahtab/go-blog/db"
	"github.com/AzmainMahtab/go-blog/routers"
)

func init() {
	db.ConnectDB()
}

// func hello(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello, World!",
// 	})
// }

func main() {
	r := routers.Routes()
	r.Run(":8080") // Start the server on port 8080
}
