package services

import (
	"bookapi/api/repositories"
	"bookapi/models"

	"github.com/google/uuid"
)

type BookService struct {
	repository repositories.BookRepository
}

func NewBookService(
	repository repositories.BookRepository,
) BookService {
	return BookService{
		repository: repository,
	}
}

func (b BookService) CreateBook(
	book models.Book,
) (models.Book, error) {
	id := uuid.New()
	book.ID = models.BINARY16(id)
	err := b.repository.Create(book)

	return book, err
}

func (b BookService) GetAllBooks() (models.Books, int64, error) {
	return b.repository.FindAll()
}

func (b BookService) GetOneBook(bookId models.BINARY16) (models.Book, error) {
	return b.repository.FindOne(bookId)
}

func (b BookService) UpdateBook(bookId models.BINARY16, book *models.Book) error {
	return b.repository.Update(bookId, book)
}

func (b BookService) DeleteBook(bookId models.BINARY16) error {
	return b.repository.Delete(bookId)
}
