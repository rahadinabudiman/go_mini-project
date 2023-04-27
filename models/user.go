package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string   `json:"name" form:"name"`
	Username string   `json:"username" form:"username"`
	Password string   `json:"password" form:"password"`
	Token    string   `gorm:"-" json:"token" form:"token"`
	Review   []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
