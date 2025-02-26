package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Structure type for an album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// To get all the albums Data
func GetAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// To get an specific album by It's ID
func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.JSON(http.StatusOK, album)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Album not founded"})
}


// To create an New Album
func CreateAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "All fields are required to create an album"})
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// To Update An Album By It's specific ID
func UpdateAlbum(c *gin.Context) {
	albumId := c.Param("id")
	var updatedAlbum Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Alll fields of album are required"})
	}
	for i, album := range albums {
		if album.ID == albumId {
			albums[i] = updatedAlbum
			c.JSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// Delete an album
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Album deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
