package routes

import (
	"go_mini-project/constants"
	"go_mini-project/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Login dan Register users
	e.POST("/users/register", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.GET("/trending", controllers.GetTrendingController)

	// User Routes with JWT
	user := e.Group("/users")
	user.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserByIdController)
	user.PUT("/:id", controllers.UpdateUserByIdController)
	user.DELETE("/:id", controllers.DeleteUserByIdController)
	user.GET("/profile/:id", controllers.GetUserDetailController)

	// Review Routes
	review := e.Group("/reviews")
	review.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	review.POST("", controllers.CreateReviewController)
	review.GET("", controllers.GetReviewsController)
	review.GET("/:id", controllers.GetReviewByIdController)
	review.DELETE("/:id", controllers.DeleteReviewByIdController)
	review.PUT("/:id", controllers.UpdateReviewByIdController)
	review.GET("/:title", controllers.GetReviewByTitle)

	return e
}
