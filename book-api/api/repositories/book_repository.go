package repositories

import (
	"bookapi/infrastructure"
	"bookapi/models"
	"fmt"
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

func (b BookRepository) Create(book models.Book) error {
	fmt.Println(book, "Book in Repo")
	return b.db.DB.Create(&book).Error
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

func (b BookRepository) FindOne(bookId models.BINARY16) (book models.Book, err error) {
	queryBuilder := b.db.DB.Model(&models.Book{})

	return book, queryBuilder.
		Where("id = ?", bookId).
		First(&book).Error
}

func (b BookRepository) Update(bookId models.BINARY16, book *models.Book) error {
	queryBuilder := b.db.DB.Model(&models.Book{})
	return queryBuilder.
		Where("id = ?", bookId).
		Updates(map[string]interface{}{
			"title":          book.Title,
			"description":    book.Description,
			"author":         book.Author,
			"published_year": book.Published_Year,
			"isbn":           book.ISBN,
			"language":       book.Language,
			"country":        book.Country,
			"publisher":      book.Publisher,
			"cover_image":    book.Cover_Image,
		}).Error
}

func (b BookRepository) Delete(bookId models.BINARY16) error {
	return b.db.DB.Where("id = ?", bookId).Delete(&models.Book{}).Error
}
