package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	viewURLSourcePath    = "/views/"
	saveURLSourcePath    = "/save/"
	editURLSourcePath    = "/edit/"
	validURLRegexPattern = "^/(edit|save|view)/([a-zA-Z0-9]+)$"
)

// Page represents a web page
type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile(validURLRegexPattern)

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

/********************
Handlers
*********************/

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadpage(title)
	if err != nil {
		http.Redirect(w, r, editURLSourcePath+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadpage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

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
	http.HandleFunc(viewURLSourcePath, makeHandler(viewHandler))
	http.HandleFunc(editURLSourcePath, makeHandler(editHandler))
	http.HandleFunc(saveURLSourcePath, makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}
