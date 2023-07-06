package controllers

import (
	"fmt"
	"go-gin/config"
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAlbums(c *gin.Context, DB *mongo.Client) {

	collectionName := config.GetCollection(DB, "albums")

	collectionName.Find(c, bson.M{})

	albums := []models.Album{}

	results, err := collectionName.Find(c, bson.M{})
	if err != nil {
		return
	}
	defer results.Close(c)
	for results.Next(c) {
		var singleAlbum models.Album
		if err = results.Decode(&singleAlbum); err != nil {
			return
		}

		albums = append(albums, singleAlbum)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func SaveAlbums(c *gin.Context, DB *mongo.Client) {
	var album models.Album

	error := c.BindJSON(&album)
	if error != nil {
		fmt.Println("Error", error)
		return
	}

	collectionName := config.GetCollection(DB, "albums")
	collectionName.InsertOne(c, album)
	c.Writer.WriteHeader(http.StatusInternalServerError)
	c.IndentedJSON(http.StatusOK, album)
}

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Java", Artist: "Rajasekhar", Price: 56.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
