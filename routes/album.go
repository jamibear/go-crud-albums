package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jamibear/go-crud-albums/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.PostAlbums)
	router.GET("/albums/:id", controller.GetAlbumById)
	router.PUT("/albums/:id", controller.UpdateAlbum)
	router.DELETE("/albums/:id", controller.DeleteAlbumById)
}
