package controllers

import (
	"encoding/json"
	"fmt"
	"go_mini-project/lib/database"
	"go_mini-project/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Movie ID TMDB
func getMovieID(title string) (int, error) {
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
	body, err := ioutil.ReadAll(resp.Body)
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

// Create Review
func CreateReviewController(c echo.Context) error {
	review := models.Review{}
	c.Bind(&review)

	if review.UserID == 0 || review.Title == "" || review.Ulasan == "" || review.Rating == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "All fields are required",
			Data:    nil,
		})
	}

	// Fetch movie ID from TMDB API
	movieID, err := getMovieID(review.Title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	review.MovieID = movieID

	// Save review to database
	review, err = database.CreateReview(review)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success add review",
		Data:    review,
	})
}

// Get All Review
func GetReviewsController(c echo.Context) error {
	reviews, err := database.GetReview()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all review",
		Data:    reviews,
	})
}

// Get Review by Id
func GetReviewByIdController(c echo.Context) error {
	ReviewId := c.Param("id")

	review, err := database.GetReviewById(ReviewId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get review by id",
		Data:    review,
	})
}

// Delete Review by Id
func DeleteReviewByIdController(c echo.Context) error {
	ReviewId := c.Param("id")

	review, err := database.DeleteReview(ReviewId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete review by id",
		Data:    review,
	})
}

// Update Review by Id
func UpdateReviewByIdController(c echo.Context) error {
	ReviewId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Invalid ID",
			Data:    nil,
		})
	}

	review := models.Review{}
	c.Bind(&review)

	if review.UserID == 0 || review.Title == "" || review.Ulasan == "" || review.Rating == 0 {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "All fields are required",
			Data:    nil,
		})
	}

	if review.UserID != uint(ReviewId) {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "UserID tidak boleh dirubah",
			Data:    nil,
		})
	}

	review, err = database.UpdateReview(review, ReviewId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update review by id",
		Data:    review,
	})
}
