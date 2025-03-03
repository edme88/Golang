package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year int `json:"year"`
	Price string `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Youthanasia", Artist: "Medadeth", Year: 1994, Price: "u$d 13.65"},
	{ID: "2", Title: "Somewhere in Time", Artist: "Iron Maiden", Year: 1986, Price: "u$d 40"},
	{ID: "3", Title: "The Number of the Beast", Artist: "Iron Maiden", Year: 1982, Price: "u$d 25.19"},
}

func getAlbums(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK,albums)
}

func getAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK,a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func postAlbums(ctx *gin.Context){
	var newAlbum album
	if err := ctx.BindJSON(&newAlbum); err != nil{
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")
	var modifyAlbum album
	if err := ctx.BindJSON(&modifyAlbum); err != nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Datos inválidos"})
		return
	}
	
	for i, a := range albums {
		if a.ID == id {
			albums[i] = modifyAlbum
			albums[i].ID = id
			ctx.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func deleteAlbumsByID(ctx *gin.Context){
	id := ctx.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1])
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album eliminado"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func main(){
	fmt.Println("Bienvenido a la API de vinyl-api!")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbumsByID) //Quedo de Tarea
	router.DELETE("/albums/:id", deleteAlbumsByID) //Quedo de Tarea
	router.Run("localhost:8080")
}