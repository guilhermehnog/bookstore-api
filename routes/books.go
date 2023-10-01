package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermehnog/bookstore-api/handler"
)

func ConfigureBooksRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/books", handler.GetBooksHandler)

	v1.GET("/book/:id", handler.GetBookHandler)

	v1.POST("/book", handler.PostBookHandler)

	v1.PATCH("/book", handler.PatchBookHandler)

	v1.DELETE("/book", handler.DeleteBookHandler)
}
