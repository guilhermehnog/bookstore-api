package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureBooksRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GET Books"})
	})

	v1.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GET Book"})
	})

	v1.POST("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "POST Books"})
	})

	v1.PATCH("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PATCH Books"})
	})

	v1.DELETE("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "DELETE Books"})
	})
}
