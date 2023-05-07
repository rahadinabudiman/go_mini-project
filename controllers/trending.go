package controllers

import (
	"encoding/json"
	"go_mini-project/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetTrendingController(c echo.Context) error {
	url := "https://api.themoviedb.org/3/trending/movie/week?api_key=0bf8630ff9d3ff478b4f4bb3b8029338"

	// membuat request ke API TMDB
	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to fetch trending movies"})
	}
	defer resp.Body.Close()

	// membaca response body dan mengambil data
	var data models.TrendingMoviesResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to parse trending movies"})
	}

	// mengembalikan response ke client
	return c.JSON(http.StatusOK, data.Results)
}
