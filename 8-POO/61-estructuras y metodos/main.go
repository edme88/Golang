// go mod init library-manager
package main

import "library-manager/book"

func main() {

	//En lugar de esto podemos usar un constructor
	// var myBook = book.Book{
	// 	Title: "Moby Dick",
	// 	Author: "Herman Melville",
	// 	Pages: 300,
	// }
	myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

	myBook.PrintInfo()
}