package controllers

import (
	"go_mini-project/lib/database"
	"go_mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
	movieID, err := database.GetMovieID(review.Title)
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
