package middlewares

import (
	"go_mini-project/constants"
	"go_mini-project/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

var CheckLogin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SECRET_JWT),
	TokenLookup:   "cookie:JWTCookie",
	ErrorHandler: func(err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	},
},
)

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}
	return 0
}

func JWTValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("JWTCookie")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET_JWT), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		id, ok := claims["userId"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, models.Response{
				Message: "Unauthorized",
			})
		}
		c.Set("userId", int(id))
		return next(c)
	}
}
