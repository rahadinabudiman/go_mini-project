package routes

import (
	"go_mini-project/controllers"
	m "go_mini-project/middlewares"
	"go_mini-project/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	cv := &util.CustomValidator{Validators: validator.New()}
	e.Validator = cv
	// Login dan Register users
	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUserController)
	e.GET("/trending", controllers.GetTrendingController)

	// User Routes with JWT
	user := e.Group("/users")
	user.GET("", controllers.GetUsersController, m.CheckLogin)
	user.GET("/:id", controllers.GetUserByIdController, m.CheckLogin)
	user.PUT("/:id", controllers.UpdateUserByIdController, m.CheckLogin)
	user.DELETE("/:id", controllers.DeleteUserByIdController, m.CheckLogin)

	// Review Routes
	review := e.Group("/reviews")
	review.POST("", controllers.CreateReviewController, m.CheckLogin)
	review.GET("", controllers.GetReviewsController, m.CheckLogin)
	review.GET("/:id", controllers.GetReviewByIdController, m.CheckLogin)
	review.DELETE("/:id", controllers.DeleteReviewByIdController, m.CheckLogin)
	review.PUT("/:id", controllers.UpdateReviewByIdController, m.CheckLogin, m.JWTValidator)
	review.GET("/title/:title", controllers.GetReviewByTitle, m.CheckLogin)

	return e
}
