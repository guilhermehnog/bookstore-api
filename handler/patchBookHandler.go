package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehnog/bookstore-api/database"
	"github.com/guilhermehnog/bookstore-api/models"
)

func PatchBookHandler(c *gin.Context) {
	// Validate JSON input with book struct
	var updatedBook models.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	// Opens database connection
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Execute upate logic
	var rows int64
	sql := `UPDATE books SET author=$2, title=$3, quantity=$4 WHERE id=$1`

	res, err := conn.Exec(sql, updatedBook.ID, updatedBook.Author, updatedBook.Title, updatedBook.Quantity)

	if err != nil {
		return
	}
	rows, _ = res.RowsAffected()
	c.IndentedJSON(http.StatusOK, gin.H{"Book Updated": updatedBook, "Rows Affected": rows})

}
