package routes

import "github.com/gin-gonic/gin"

func InitializeRoutes() {
	// Initialize Router
	router := gin.Default()
	// Configure Routes
	ConfigureBooksRoutes(router)
	// Run the server
	router.Run("localhost:8080")
}
