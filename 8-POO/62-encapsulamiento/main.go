// go mod init library-manager
package main

import (
	"fmt"
	"library-manager/book"
)

func main() {
	//El constructor,
	//book := book.Book{"Hamlet","Shakespeare",200}
	myBook := book.NewBook("Moby Dick", "Herman Melville", 300)

	myBook.PrintInfo()
	myBook.SetTitle("Moby Dick (Edición especial)")
	fmt.Println(myBook.GetTitle())
	myBook.PrintInfo()
}