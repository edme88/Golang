package book

import "fmt"

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

// ! Composicion permite crear nuevas estructuras que poseen campos y métodos de otras estructuras
//Nueva estructura para libros escolares
type TextBook struct {
	Book
	editorial string
	nivel	  string
}


//Método que contiene la estructura anterior
func NewTextBook(title, author string, pages int, editorial, level string) *TextBook{
	return &TextBook{
		//Se puede usar Book o NewBook porque estamos en el mismo paquete
		Book: Book{title, author, pages},
		editorial: editorial,
		nivel: level,
	}
}

func (b *TextBook) PrintInfo() {
	b.Book.PrintInfo()
	fmt.Printf("Editorial: %s\nNivel: %s\n", b.editorial, b.nivel)
}
