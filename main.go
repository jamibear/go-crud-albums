package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamibear/go-crud-albums/config"
	"github.com/jamibear/go-crud-albums/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}
