package routes

import (
	"bookapi/api/controllers"
	"bookapi/infrastructure"
)

type BooksRoute struct {
	controller controllers.BooksController
	router     infrastructure.Router
}

func NewBookRoute(
	controller controllers.BooksController,
	router infrastructure.Router,
) BooksRoute {
	return BooksRoute{
		controller: controller,
		router:     router,
	}
}

func (b BooksRoute) Setup() {
	books := b.router.Gin.Group("/books")
	{
		books.GET("/", b.controller.HandleGetAllBooks())
		books.POST("/", b.controller.CreateBook())
	}
}
