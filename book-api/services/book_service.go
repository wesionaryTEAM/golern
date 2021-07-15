package services

import (
	"bookapi/api/repositories"
	"bookapi/models"
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

func (b BookService) GetAllBooks() (models.Books, int64, error) {
	return b.repository.FindAll()
}
