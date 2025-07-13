// Package db used for db settings	and connection
package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	log.Println("Connected to the database successfully")
}
