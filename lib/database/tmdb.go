package database

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Movie ID TMDB
func GetMovieID(title string) (int, error) {
	API_KEY := "0bf8630ff9d3ff478b4f4bb3b8029338"

	// Replace spaces with %20 or +
	title = strings.ReplaceAll(title, " ", "%20")
	// title = strings.ReplaceAll(title, " ", "+") // alternatif

	// Build URL
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s", API_KEY, title)

	// Send HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Parse JSON response
	var result struct {
		Results []struct {
			ID int `json:"id"`
		} `json:"results"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	// Return movie ID
	if len(result.Results) > 0 {
		return result.Results[0].ID, nil
	}
	return 0, fmt.Errorf("movie not found")
}
