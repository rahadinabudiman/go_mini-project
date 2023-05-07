package middlewares

import (
	"go_mini-project/constants"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
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
