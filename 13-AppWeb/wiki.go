package main

import (
	"html/template"
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

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	//t, _ := template.ParseFiles("view.html")
	//*fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)*////
	//t.Execute(w, p)
	renderTemplates(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, _ := loadPage(title)
	renderTemplates(w, "edit", p)
}

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page){
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func main(){
	// p1 := &Page{Title: "TestPage", Body: []byte("Esta es una página de muestra")}
	// p1.save()

	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) //detiene el servidor si hay error
}