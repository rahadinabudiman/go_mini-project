package controllers

import (
	"go_mini-project/lib/database"
	"go_mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Create Review
func CreateReviewController(c echo.Context) error {
	review := models.Review{}
	c.Bind(&review)

	if err := c.Validate(&review); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	id, ok := c.Get("userId").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "userId tidak ada",
		})
	}

	review.UserID = uint(id)

	// Fetch movie ID from TMDB API
	movieID, err := database.GetMovieID(review.Title)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	review.MovieID = movieID

	judulBaru, err := database.GetMovieIDToTitle(review.MovieID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	review.Title = judulBaru

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

	if len(reviews) == 0 {
		return c.JSON(http.StatusOK, models.Response{
			Message: "No reviews found",
			Data:    nil,
		})
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

func GetReviewByTitle(c echo.Context) error {
	title := c.Param("title")

	reviews, err := database.GetReviewByTitle(title)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	name, err := database.GetUserByIdReview(reviews[0].UserID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var reviewResponses []models.ReviewResponse
	for _, review := range reviews {
		reviewResponses = append(reviewResponses, models.ReviewResponse{
			Name:   name.Name,
			Title:  review.Title,
			Rating: review.Rating,
			Ulasan: review.Ulasan,
		})
	}

	return c.JSON(http.StatusOK, reviewResponses)
}

// Delete Review by Id
func DeleteReviewByIdController(c echo.Context) error {
	ReviewId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Mengambil ID dari Review
	reviewID, err := database.GetReviewById(strconv.Itoa(int(ReviewId)))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, ok := c.Get("userId").(int)
	if !ok || id != int(reviewID.UserID) {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "ID berbeda dengan user yang login",
		})
	}

	// Mengambil ID dari Review
	reviewIDBanget, err := database.GetReviewById(strconv.Itoa(int(ReviewId)))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validasi jika mengedit review punya orang
	if int(reviewIDBanget.UserID) != id {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Tidak Bisa Menghapus Review Orang Lain",
		})
	}

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

	if err := c.Validate(&review); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
		})
	}

	// Mengambil ID dari Review
	reviewIDBanget, err := database.GetReviewById(ReviewId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	TitleReviewDatabase := reviewIDBanget.Title
	userIDDatabase := reviewIDBanget.UserID

	id, ok := c.Get("userId").(int)
	if !ok || id != int(userIDDatabase) {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "ID berbeda dengan user yang login",
		})
	}

	// Mengambil UserID dari Review
	user, err := database.GetReviewByUserId(reviewIDBanget.UserID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validasi jika mengedit review punya orang
	if int(reviewIDBanget.UserID) != id {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Tidak Bisa Mengedit Review Orang Lain",
		})
	}

	review.MovieID = user.MovieID
	review.Title = TitleReviewDatabase
	review.UserID = uint(id)

	// Jika Judul Film Berbeda Pada Saat diedit, maka tidak bisa disimpan
	if TitleReviewDatabase != review.Title {
		return c.JSON(http.StatusBadRequest, models.Response{
			Message: "Judul Film Tidak Boleh Diubah",
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
