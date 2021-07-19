package controllers

import (
	"bookapi/api/responses"
	"bookapi/infrastructure"
	"bookapi/models"
	"bookapi/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/form"
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

func DecodePostForm(c *gin.Context) (*models.Book, error) {
	var book models.Book
	decoder := form.NewDecoder()

	err := decoder.Decode(&book, c.Request.PostForm)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (bc BooksController) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		book := &models.Book{}

		// if err := c.ShouldBindJSON(&book); err != nil {
		// 	bc.logger.Zap.Error("Book Create Error: ", err)
		// 	responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
		// 	return
		// }
		// err := c.Request.ParseMultipartForm(200000)
		// if err != nil {
		// 	return
		// }

		file, header, err := c.Request.FormFile("cover_image")
		filename := header.Filename
		defer file.Close()

		book, err = DecodePostForm(c)
		book.Cover_Image = filename

		books, err := bc.service.CreateBook(*book)
		if err != nil {
			bc.logger.Zap.Error("Error while creating Book: ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println(books, "Books from Controller")

		responses.JSON(c, http.StatusOK, gin.H{
			"msg":  "Book Created successfully.",
			"book": books.ToMap(),
		})
	}
}

func (bc BooksController) HandleGetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {

		books, count, err := bc.service.GetAllBooks()

		if err != nil {
			bc.logger.Zap.Error("Failed to load Books", err.Error())
			responses.ErrorJSON(c, http.StatusBadRequest, "Failed to Load Books")
			return
		}

		responses.JSON(c, http.StatusOK, gin.H{
			"count": count,
			"data":  books.ToMap(),
		})
	}
}

func (bc BooksController) GetOneBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("id")
		binaryId, err := models.StringToBinary16(bookId)
		if err != nil {
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
		}

		book, err := bc.service.GetOneBook(binaryId)
		if err != nil {
			bc.logger.Zap.Error("Failed to get Book", err.Error())
			responses.ErrorJSON(c, http.StatusInternalServerError, "Failed to get Book")
			return
		}

		responses.JSON(c, http.StatusOK, gin.H{
			"data": book.ToMap(),
		})
	}
}

func (bc BooksController) UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId := c.Param("id")
		binaryId, err := models.StringToBinary16(bookId)
		if err != nil {
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
		}

		book, err := bc.service.GetOneBook(binaryId)
		if err != nil {
			bc.logger.Zap.Error("Failed to get Book", err.Error())
			responses.ErrorJSON(c, http.StatusInternalServerError, "Failed to get Book")
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			responses.ErrorJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		err = bc.service.UpdateBook(binaryId, &book)
		if err != nil {
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		responses.SuccessJSON(c, http.StatusOK, "Book Successfully Updated")
	}
}
