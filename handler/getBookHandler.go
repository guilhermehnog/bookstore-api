package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guilhermehnog/bookstore-api/database"
	"github.com/guilhermehnog/bookstore-api/models"
)

func GetBookHandler(c *gin.Context) {
	// Validate JSON input with book struct
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Bad Request": "Please Inform a valid type:INT id number"})
	}

	// Opens database connection
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Execute select logic
	var book models.Book
	sql := `SELECT * FROM books where id=$1`
	row := conn.QueryRow(sql, id)
	if err = row.Scan(&book.ID, &book.Author, &book.Title, &book.Quantity); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Bad Request": "Book ID not found", "Book ID": id})
	} else {
		c.IndentedJSON(http.StatusOK, book)
	}

}
