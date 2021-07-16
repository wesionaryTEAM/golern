package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Base
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	Author         string `json:"author" form:"author"`
	Published_Year int    `json:"published_year" form:"published_year"`
	ISBN           string `json:"isbn" form:"isbn"`
	Language       string `json:"language" form:"language"`
	Country        string `json:"country" form:"country"`
	Publisher      string `json:"publisher" form:"publisher"`
	Cover_Image    string `json:"cover_image"`
}

func (b Book) ToMap() gin.H {
	return gin.H{
		"id":             b.ID,
		"title":          b.Title,
		"description":    b.Description,
		"author":         b.Author,
		"published_year": b.Published_Year,
		"isbn":           b.ISBN,
		"language":       b.Language,
		"country":        b.Country,
		"publisher":      b.Publisher,
		"cover_image":    b.Cover_Image,
	}
}

type Books []Book

// ToMap results
func (b Books) ToMap() []gin.H {
	results := []gin.H{}
	for _, book := range b {
		results = append(results, book.ToMap())
	}
	return results
}

func (b *Book) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	b.ID = BINARY16(id)
	return err
}

func (b Book) TableName() string {
	return "books"
}
