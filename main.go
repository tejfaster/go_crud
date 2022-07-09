package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Lager blu", Artist: "John rime", Price: 96.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for index, a := range albums {
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.1.2"})
	// router.GET("/albums", getAlbums)
	// router.POST("/albums",createAlbum)
	// router.GET("/albums/:id",getAlumsByID)
	// router.DELETE("/albums/:id",deleteAlbumsByID)
	// router.Run("localhost:8081")
}
