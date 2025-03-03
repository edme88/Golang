package book

import "fmt"

//Todos los métodos deben implementar esta función
type Printable interface {
	PrintInfo()
}

func Print(p Printable){
	p.PrintInfo()
}

type Book struct {
	title string
	author string
	pages int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title: title,
		author: author,
		pages: pages,
	}
}

//Getters
func (b *Book) GetTitle() string{
	return b.title
}

func (b *Book) GetAuthor() string{
	return b.author
}

func (b *Book) GetPages() int{
	return b.pages
}

//Setters
func (b *Book) SetTitle(title string){
	b.title = title
}

func (b *Book) SetAuthor(author string){
	b.author = author
}

func (b *Book) SetPages(pages int){
	b.pages = pages
}

//MÉTODOS
func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s,\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	nivel	  string
}


func NewTextBook(title, author string, pages int, editorial, level string) *TextBook{
	return &TextBook{
		Book: Book{title, author, pages},
		editorial: editorial,
		nivel: level,
	}
}

func (b *TextBook) PrintInfo() {
	b.Book.PrintInfo()
	fmt.Printf("Editorial: %s\nNivel: %s\n", b.editorial, b.nivel)
}

//Polimorfimos: capacidad de los elementos de diferentes clases de responder el mismo mensaje o método
//Una interfaz es un conjunto de métodos que debe satisfacer una interfaz
