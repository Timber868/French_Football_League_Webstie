package main

import (
	"ligue1-backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Set up the route for fetching Ligue 1 standings
	router.GET("/api/standings", api.GetStandings)

	// Start the server on port 8080
	router.Run(":8080")
}
