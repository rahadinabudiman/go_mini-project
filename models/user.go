package models

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type User struct {
	gorm.Model
	Name     string   `json:"name" form:"name"`
	Username string   `json:"username" form:"username"`
	Password string   `json:"password" form:"password"`
	Token    string   `gorm:"-" json:"token" form:"token"`
	Review   []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// For Token Only
type UerToken struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Token    string `gorm:"-" json:"token" form:"token"`
}

type CustomValidator struct {
	Validators *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validators.Struct(i)

	if err != nil {
		var sb strings.Builder
		sb.WriteString("Validation error:\n")

		for _, err := range err.(validator.ValidationErrors) {
			sb.WriteString(fmt.Sprintf("- %s\n", err))
		}

		return echo.NewHTTPError(http.StatusBadRequest, sb.String())
	}

	return nil
}
