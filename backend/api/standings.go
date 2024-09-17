package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define the structure of the standings response we expect from the API
type StandingsResponse struct {
	Standings []struct {
		Team   string `json:"team"`
		Points int    `json:"points"`
	} `json:"standings"`
}

// GetStandings is the handler function to fetch and return Ligue 1 standings
func GetStandings(c *gin.Context) {
	// Define the external API URL (e.g., Football-Data.org)
	apiKey := "YOUR_API_KEY" // Replace with your actual API key
	url := "https://api.football-data.org/v4/competitions/FL1/standings"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Add the API key to the request headers
	req.Header.Add("X-Auth-Token", apiKey)

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

	// Return the standings as JSON
	c.JSON(http.StatusOK, standings)
}
