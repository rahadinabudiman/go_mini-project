package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name" form:"name" validate:"required"`
	Username string   `json:"username" form:"username" validate:"required"`
	Password string   `json:"password" form:"password" validate:"required"`
	Review   []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// For Response User
type UserResponse struct {
	Name     string   `json:"name" form:"name" validate:"required"`
	Username string   `json:"username" form:"username" validate:"required"`
	Review   []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// For Response User Create
type UserCreateResponse struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
}

// For Token Only
type UerToken struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Token    string `gorm:"-" json:"token" form:"token"`
}
