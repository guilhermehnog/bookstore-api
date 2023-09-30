package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GET Books"})
	})

	router.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GET Book"})
	})

	router.POST("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "POST Books"})
	})

	router.PATCH("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PATCH Books"})
	})

	router.DELETE("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "DELETE Books"})
	})

	router.Run("localhost:8080")
}
