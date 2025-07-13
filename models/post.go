// Package models
package models

import (
	"gorm.io/gorm"

	"github.com/AzmainMahtab/go-blog/db"
)

type Post struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null; unique"`
	Body   string `json:"body" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
}

func (p *Post) TableName() string {
	return "posts"
}

func Create(p *Post) error {
	result := db.DB.Create(&p)
	return result.Error
}

func GetPosts() (*[]Post, error) {
	var posts []Post
	result := db.DB.Find(&posts)
	return &posts, result.Error
}
