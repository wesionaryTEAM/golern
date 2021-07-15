package controllers

import (
	"bookapi/api/responses"
	"bookapi/infrastructure"
	"bookapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	logger  infrastructure.Logger
	service services.BookService
}

func NewBooksController(
	logger infrastructure.Logger,
	service services.BookService,
) BooksController {
	return BooksController{
		logger:  logger,
		service: service,
	}
}

func (bc BooksController) HandleGetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {

		books, count, err := bc.service.GetAllBooks()

		if err != nil {
			bc.logger.Zap.Error("Failed to load Clubs", err.Error())
			responses.ErrorJSON(c, http.StatusBadRequest, "Failed to Load Books")
			return
		}
		responses.JSON(c, http.StatusOK, gin.H{
			"count": count,
			"data":  books.ToMap(),
		})
	}
}
