package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

var validPath = regexp.MustCompile("^/(view|edit|save)/([A-Za-z0-9]+)+$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error){
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Título de página inválido")
	}
	return m[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//t, _ := template.ParseFiles("view.html")
	//*fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)*////
	//t.Execute(w, p)
	renderTemplates(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplates(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	http.Redirect(w,r,"/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page){
	// t, err := template.ParseFiles(tmpl + ".html")
	// if err != nil {
	// 	http.Error(w,err.Error(),http.StatusInternalServerError)
	// 	return
	// }
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	//err = t.Execute(w, p)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}

func main(){
	// p1 := &Page{Title: "TestPage", Body: []byte("Esta es una página de muestra")}
	// p1.save()

	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) //detiene el servidor si hay error
}