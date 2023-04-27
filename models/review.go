package models

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id"`
	MovieID int    `json:"movie_id" form:"movie_id"`
	Title   string `json:"title" form:"title"`
	Ulasan  string `json:"ulasan" form:"review_user"`
	Rating  int    `json:"rating" form:"rating"`
}