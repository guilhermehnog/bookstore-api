package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET Books"})
}
