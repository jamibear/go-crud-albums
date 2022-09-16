package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jamibear/go-crud-albums/config"
	"github.com/jamibear/go-crud-albums/models"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	albums := []models.Album{}
	config.DB.Find(&albums)
	c.JSON(200, &albums)
}

// postAlbums adds new set of albums to the JSON file
func PostAlbums(c *gin.Context) {
	c.String(200, "post albums")

	var album models.Album
	c.BindJSON(&album)
	config.DB.Create(&album)
	c.JSON(200, &album)
}

// get a specific album by id
func GetAlbumById(c *gin.Context) {
	c.String(200, "get album by id")

	var album models.Album
	config.DB.Find(&album, c.Param("id"))
	c.JSON(200, &album)
}

// delete a specific album by id
func DeleteAlbumById(c *gin.Context) {
	c.String(200, "Delete album by id")

	var album models.Album
	config.DB.Where("id = ?", c.Param("id")).Delete(&album)
	c.JSON(200, &album)
}

// edit an album
func UpdateAlbum(c *gin.Context) {
	c.String(200, "Update album")

	var album models.Album
	config.DB.Where("id = ?", c.Param("id")).First(&album)
	c.BindJSON(&album)
	config.DB.Save(&album)
	c.JSON(200, &album)
}
