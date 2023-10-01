package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehnog/bookstore-api/database"
	"github.com/guilhermehnog/bookstore-api/models"
)

func PostBookHandler(c *gin.Context) {
	// Validate JSON input with book struct
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Opens database connection
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	//Execute insert logic
	sql := `INSERT INTO books (author, title, quantity) VALUES ($1, $2, $3) RETURNING id`
	if err = conn.QueryRow(sql, newBook.Author, newBook.Title, newBook.Quantity).Scan(&newBook.ID); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, newBook)

}
