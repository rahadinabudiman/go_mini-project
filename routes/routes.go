package routes

import (
	"go_mini-project/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	user := e.Group("/users")
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserByIdController)
	user.POST("", controllers.CreateUserController)
	user.PUT("/:id", controllers.UpdateUserByIdController)
	user.DELETE("/:id", controllers.DeleteUserByIdController)

	// Review Routes
	review := e.Group("/reviews")
	review.POST("", controllers.CreateReviewController)
	review.GET("", controllers.GetReviewsController)
	review.GET("/:id", controllers.GetReviewByIdController)
	review.DELETE("/:id", controllers.DeleteReviewByIdController)
	review.PUT("/:id", controllers.UpdateReviewByIdController)

	return e
}