package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model

	ID     uint   `json:"ID" gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
