package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PatchBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PATCH Book"})
}
