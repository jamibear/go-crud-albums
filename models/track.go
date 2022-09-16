package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model

	ID       int    `json:"id" gorm:"primary_key"`
	Album_id int    `json:"album_id"`
	Title    string `json:"title"`
}
