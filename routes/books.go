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

	v1.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "POST Book"})
	})

	v1.PATCH("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "PATCH Book"})
	})

	v1.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "DELETE Book"})
	})
}
