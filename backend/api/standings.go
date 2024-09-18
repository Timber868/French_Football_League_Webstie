package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a named struct for individual team standings
type TeamStanding struct {
	Team   string `json:"team"`
	Points int    `json:"points"`
}

// Update StandingsResponse to use the named struct
type StandingsResponse struct {
	Standings []TeamStanding `json:"standings"`
}

// GetStandings is the handler function to fetch and return Ligue 1 standings
func GetStandings(c *gin.Context) {
	fmt.Println("Get Standings")
	// Define the external API URL (e.g., Football-Data.org)
	apiKey := "920fc34aa8d24922a5ba0ec8f916d8e6" // Replace with your actual API key
	url := "https://native-stats.org/competition/FL1/"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Add the API key to the request headers
	req.Header.Add("X-Auth-Token", apiKey)
	fmt.Println("\033[33mReq:", req, "\033[0m")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	// Decode the API response into our custom struct
	var standings StandingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&standings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	standings2 := StandingsResponse{
		Standings: []TeamStanding{
			{Team: "Paris Saint-Germain", Points: 1},
			{Team: "Olympique Lyonnais", Points: 2},
			{Team: "Olympique de Marseille", Points: 3},
		},
	}
	// Return the standings as JSON
	c.JSON(http.StatusOK, standings2)
}
