package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehnog/bookstore-api/database"
	"github.com/guilhermehnog/bookstore-api/models"
)

func DeleteBookHandler(c *gin.Context) {
	// Validate JSON input with book struct
	var deleteBook models.Book
	if err := c.BindJSON(&deleteBook); err != nil {
		return
	}

	// Opens database connection
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	//Execute delete logic
	var rows int64
	if res, err := conn.Exec(`DELETE FROM books WHERE id=$1`, deleteBook.ID); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		rows, err = res.RowsAffected()
	}
	c.IndentedJSON(http.StatusOK, gin.H{"Book Deleted": deleteBook, "Rows Affected": rows})
}
