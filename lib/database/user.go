package database

import (
	"go_mini-project/config"
	"go_mini-project/middlewares"
	"go_mini-project/models"
)

// Get all user
func GetUser() (users []models.User, err error) {
	err = config.DB.Preload("Review").Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	//make to return blog data

	return
}

// Create user
func CreateUser(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Get user by id
func GetUserById(id any) (models.User, error) {
	var user models.User

	err := config.DB.Preload("Review").Where("id = ?", id).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Update user by id
func UpdateUser(user models.User, id any) (models.User, error) {
	err := config.DB.Table("users").Where("id = ?", id).Updates(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Delete user by id
func DeleteUser(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete user", nil
}

// Login User With JWT
// func LoginUser(user models.User) (models.User, error) {

// 	err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error

// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	return user, nil
// }

// Get Detail User JWT
func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

// Login User JWT
func LoginUser(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
