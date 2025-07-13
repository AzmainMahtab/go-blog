package main

import (
	"github.com/AzmainMahtab/go-blog/db"
	"github.com/AzmainMahtab/go-blog/models"
)

func init() {
	// Initialize the database connection
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.Post{}) // Automatically migrate the database schema
}
