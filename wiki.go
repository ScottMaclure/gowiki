package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
)

// Represents a page in the wiki.
type Page struct {
	Title string
	// io library works with bytes, so that's why we use a byte slice.
	Body []byte
}

// Persistence function.
// "p" is the receiver of this function. Interesting.
// Will return nil if save works.
func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load data from filesystem.
func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		// If we attempt to view a page that doesn't exist, lets create one.
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit"):]
	p, err := loadPage(title)
	tmpl := "edit"
	if err != nil {
		p = &Page{Title: title}
		tmpl = "create"
	}
	renderTemplate(w, tmpl, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}

