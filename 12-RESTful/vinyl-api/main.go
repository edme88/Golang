package main

import (
	"fmt"

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
}

func main(){
	fmt.Println("Bienvenido a la API de vinyl-api!")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Mundo",
		})
	})
	router.Run("localhost:8080")
}