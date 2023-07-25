package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Representa os dados de um album
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

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Chama BindJSON para vincluar o JSON recebido à variável newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Adiciona o album ao JSON de albums
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// Realiza uma iteração na lista de albums procurando
	// pelo ID fornecido na URL
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album não encontrado"})
}
