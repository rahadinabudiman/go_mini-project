package controllers

import (
	"go_mini-project/lib/database"
	"go_mini-project/middlewares"
	"go_mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Get all user controller
func GetUsersController(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all user",
		Data:    users,
	})
}

// Get user by id controller
func GetUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	user, err := database.GetUserById(UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get user by id",
		Data:    user,
	})
}

// Create user controller
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create user",
		Data:    user,
	})
}

// Update user by id controller
func UpdateUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	user := models.User{}
	c.Bind(&user)

	user, err := database.UpdateUser(user, UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success update user by id",
		Data:    user,
	})
}

// Delete user by id controller-
func DeleteUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	_, err := database.DeleteUser(UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete user by id",
	})
}

// Login User With JWT
func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, e := database.LoginUser(user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Response{
			Message: "failed login user",
			Data:    e.Error(),
		})
	}

	token, err := middlewares.CreateToken(user.ID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Response{
			Message: "failed login user",
			Data:    err.Error(),
		})
	}
	middlewares.CreateCookie(c, token)

	respontoken := models.UerToken{
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success login user",
		Data:    respontoken,
	})
}

// Detail Users Check by Id
func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Response{
			Message: "failed get user detail",
			Data:    err.Error(),
		})
	}

	users, err := database.GetDetailUsers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.Response{
			Message: "failed get user detail",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get user detail",
		Data:    users,
	})
}
