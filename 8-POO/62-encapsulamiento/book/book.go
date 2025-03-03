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

//minuscula => privado
//mayuscula => puede ser accedido fuera del paquete


//Se usan getters y setters para cambiar y obtener los valores porque son atributos Privados 
//Entonces no hay forma de accederlos desde fuera del modulo
//Evita escritura o lectura accidental
//en Go no es obligatorio usar getters y setters, 
//pero pueden ser útiles cuando necesitas control, validación, transformación o encapsulación.

//Recomendación: Un método Set puede validar los datos antes de asignarlos.
func (b *Book) SetTitle(title string){
	b.title = title
}

//Recomendación: Un método Get puede hacer cálculos o formateos antes de devolver el valor.}
func (b *Book) GetTitle() string{
	return b.title
}

func (b *Book) SetAuthor(author string){
	b.author = author
}

func (b *Book) GetAuthor() string{
	return b.author
}

func (b *Book) SetPages(pages int){
	b.pages = pages
}

func (b *Book) GetPages() int{
	return b.pages
}


func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s,\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}