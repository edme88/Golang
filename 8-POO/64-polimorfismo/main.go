package main

import (
	"fmt"
	"library-manager/book"
)

func main() {
	myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

	myBook.PrintInfo()
	myBook.SetTitle("Moby Dick (Edición especial)")
	fmt.Println(myBook.GetTitle())
	myBook.PrintInfo()

	myTextBook := book.NewTextBook("Aprendiendo a leer", "Maria Gonzales", 261, "Santillana", "Inicial")
	myTextBook.PrintInfo()

	fmt.Println("=====================")
	book.Print(myBook)
	book.Print(myTextBook)
}