package book

import "fmt"

type Book struct {
	Title string
	Author string
	Pages int
}

//simular un constructor
// Un constructor es un método llamado por el tiempo de ejecución 
// cuando se crea una instancia de la clase o de la estructura . 
// Una clase o estructura puede tener varios constructores que toman argumentos diferentes. 
// Los constructores permiten asegurarse de que las instancias del tipo son válidas cuando se crean.
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title: title,
		Author: author,
		Pages: pages,
	}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s,\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}