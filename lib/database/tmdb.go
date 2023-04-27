package database

import (
	"encoding/json"
	"fmt"
	"go_mini-project/models"
	"net/http"
)

func getMovieID(movieID string, apiKey string) (int, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result models.TMBDMovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.ID, nil
}
