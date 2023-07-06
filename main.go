package main

import (
	"go-gin/config"
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Client

func main() {
	router := gin.Default()

	DB = config.ConnectDB()

	router.GET("/albums", func(c *gin.Context) {
		controllers.FindAlbums(c, DB)
	})
	router.POST("/albums", func(c *gin.Context) {
		controllers.SaveAlbums(c, DB)
	})

	router.Run("localhost:4000")
}
