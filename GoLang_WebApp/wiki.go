package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

const (
	viewURLSourcePath = "/views/"
	saveURLSourcePath = "/save/"
	editURLSourcePath = "/edit/"
)

// Page represents a web page
type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func (p *Page) save() error {
	filename := p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadpage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(viewURLSourcePath):]
	p, err := loadpage(title)
	if err != nil {
		http.Redirect(w, r, editURLSourcePath+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(editURLSourcePath):]
	p, err := loadpage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(editURLSourcePath):]
	body := r.FormValue("body")

	p := &Page{}
	p.Title = string(title)
	p.Body = []byte(body)
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, viewURLSourcePath, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc(viewURLSourcePath, viewHandler)
	http.HandleFunc(editURLSourcePath, editHandler)
	http.HandleFunc(saveURLSourcePath, saveHandler)
	http.ListenAndServe(":8080", nil)
}
