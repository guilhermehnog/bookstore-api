package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehnog/bookstore-api/database"
	"github.com/guilhermehnog/bookstore-api/models"
)

func GetBooksHandler(c *gin.Context) {

	var books []models.Book
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM books`
	rows, err := conn.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.ID, &book.Author, &book.Title, &book.Quantity)
		if err != nil {
			continue
		}
		books = append(books, book)
	}
	c.IndentedJSON(http.StatusOK, books)
}
