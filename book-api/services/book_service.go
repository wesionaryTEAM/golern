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
