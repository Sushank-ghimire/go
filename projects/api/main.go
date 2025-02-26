package main

import (
	"fmt"
	"gin-apis/controllers" // Importing the controllers package
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	PORT       = 3000
	backendUrl = "http://localhost:3000"
)

func main() {
	fmt.Println("-------Basic Golang Web Apis-----------")
	router := gin.Default()
	fmt.Println("Server is listining on : ", backendUrl)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Album store is open explore your albums"})
	})
	router.GET("/api/albums", controllers.GetAllAlbums)
	router.POST("/api/albums", controllers.CreateAlbum)
	router.DELETE("/api/albums/:id", controllers.DeleteAlbum)
	router.PUT("/api/albums/:id", controllers.UpdateAlbum)
	router.PUT("/api/album/get/:id", controllers.GetAlbumById)

	router.Run(":3000")
}
