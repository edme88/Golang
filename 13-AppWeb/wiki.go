package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error{
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola! %s", r.URL.Path[1:])
}

func main(){
	// p1 := &Page{Title: "TestPage", Body: []byte("Esta es una página de muestra")}
	// p1.save()

	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil)) //detiene el servidor si hay error
}