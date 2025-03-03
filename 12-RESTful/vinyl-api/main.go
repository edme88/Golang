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

func postAlbums(ctx *gin.Context){
	var newAlbum album
	if err := ctx.BindJSON(&newAlbum); err != nil{
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func main(){
	fmt.Println("Bienvenido a la API de vinyl-api!")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}