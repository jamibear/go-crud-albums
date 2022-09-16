package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album data
type album struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Artist string   `json:"artist"`
	Tracks []string `json:"tracks"`
}

// create album records
var albums = []album{
	{ID: "1", Title: "The Album", Artist: "BLACKPINK", Tracks: []string{"Track 1", "Track 2", "Track 3", "Track 4", "Track 5"}},
	{ID: "2", Title: "Between 1&2", Artist: "Twice", Tracks: []string{"Track 1", "Track 2", "Track 3", "Track 4", "Track 5"}},
	{ID: "3", Title: "The Second Step: Chapter One", Artist: "TREASURE", Tracks: []string{"Track 1", "Track 2", "Track 3", "Track 4", "Track 5"}},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds new set of albums to the JSON file
func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind new album data to json
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// add the new album data to album slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// get a specific album by id
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// look for the specific album by looping through the slice
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}
