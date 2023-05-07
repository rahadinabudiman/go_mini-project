package routes

import (
	"go_mini-project/controllers"
	m "go_mini-project/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Login dan Register users
	e.POST("/users/register", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)
	e.GET("/trending", controllers.GetTrendingController)

	// User Routes with JWT
	user := e.Group("/users")
	user.GET("", controllers.GetUsersController, m.CheckLogin)
	user.GET("/:id", controllers.GetUserByIdController)
	user.PUT("/:id", controllers.UpdateUserByIdController)
	user.DELETE("/:id", controllers.DeleteUserByIdController)
	user.GET("/profile/:id", controllers.GetUserDetailController)

	// Review Routes
	review := e.Group("/reviews")
	review.POST("", controllers.CreateReviewController)
	review.GET("", controllers.GetReviewsController)
	review.GET("/:id", controllers.GetReviewByIdController)
	review.DELETE("/:id", controllers.DeleteReviewByIdController)
	review.PUT("/:id", controllers.UpdateReviewByIdController)
	review.GET("/:title", controllers.GetReviewByTitle)

	return e
}
