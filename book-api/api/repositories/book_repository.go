package repositories

import (
	"bookapi/infrastructure"
	"bookapi/models"
)

type BookRepository struct {
	db infrastructure.Database
}

func NewBookRepository(
	db infrastructure.Database,
) BookRepository {
	return BookRepository{
		db: db,
	}
}

func (b BookRepository) FindAll() (models.Books, int64, error) {
	var totalRows int64 = 0
	var books []models.Book
	queryBuilder := b.db.DB.Model(&models.Book{})

	err := queryBuilder.
		Where(books).Find(&books).
		Count(&totalRows).Error

	return books, totalRows, err
}
