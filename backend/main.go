package main

import (
	"ligue1-backend/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Set up the route for fetching Ligue 1 standings
	router.GET("/api/standings", func(c *gin.Context) {
		standings2 := api.StandingsResponse{
			Standings: []api.TeamStanding{
				{Team: "Paris Saint-Germain", Points: 1},
				{Team: "Olympique Lyonnais", Points: 2},
				{Team: "Olympique de Marseille", Points: 3},
			},
		}

		c.JSON(http.StatusOK, standings2)
	})

	// Start the server on port 8080
	router.Run(":8080")
}
