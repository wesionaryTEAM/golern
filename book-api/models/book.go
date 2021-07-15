package models

import "github.com/gin-gonic/gin"

type Book struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Author         string `json:"author"`
	Published_Year int    `json:"published_year"`
	ISBN           string `json:"isbn"`
}

func (b Book) ToMap() gin.H {
	return gin.H{
		"title":          b.Title,
		"description":    b.Description,
		"author":         b.Author,
		"published_year": b.Published_Year,
		"isbn":           b.ISBN,
	}
}

type Books []Book

// ToMap results
func (b Books) ToMap() []gin.H {
	results := []gin.H{}
	for _, club := range b {
		results = append(results, club.ToMap())
	}
	return results
}
