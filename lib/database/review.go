package database

import (
	"go_mini-project/config"
	"go_mini-project/models"
)

// Create Review by User Id
func CreateReview(review models.Review) (models.Review, error) {
	err := config.DB.Create(&review).Error

	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

// Get All Review
func GetReview() (reviews []models.Review, err error) {
	err = config.DB.Find(&reviews).Error

	if err != nil {
		return []models.Review{}, err
	}

	//make to return blog data
	return
}

// Get Review by id
func GetReviewById(id any) (models.Review, error) {
	var review models.Review

	err := config.DB.Where("id = ?", id).First(&review).Error

	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}

// Delete Review By Id
func DeleteReview(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Review{}).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Update Review By Id
func UpdateReview(review models.Review, id any) (models.Review, error) {
	err := config.DB.Table("reviews").Where("id = ?", id).Updates(&review).Error

	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}
