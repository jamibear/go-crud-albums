package config

import (
	"github.com/jamibear/go-crud-albums/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=jami dbname=album port=5432 sslmode=disable TimeZone=Asia/Manila"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Album{})
	DB = db
}
