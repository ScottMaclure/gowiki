package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"log"
	"regexp"
	"errors"
)

// Represents a page in the wiki.
type Page struct {
	Title string
	// io library works with bytes, so that's why we use a byte slice.
	Body []byte
}

// Load templates ONCE from filesystem, rather than on every request.
var templates = template.Must(template.ParseFiles(
	"templates/create.html", 
	"templates/edit.html", 
	"templates/view.html",
))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

// Ensure we get a valid page title from the URL.
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	
	m := validPath.FindStringSubmatch(r.URL.Path)
	
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid page title.")
	}

	return m[2], nil // The title is the second subexpression

}

// Helper function, given a Page, render using a certain template.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl + ".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Handle a redirect from base url to viewing a default wiki page.
// i.e. if user accesses http://localhost:8080/, redirect to http://localhost:8080/view/index
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/index", http.StatusFound)
}

// View a wiki page in html!
func viewHandler(w http.ResponseWriter, r *http.Request) {

	title, err := getTitle(w, r)

	if err != nil {
		return
	}

	p, err := loadPage(title)

	if err != nil {
		// If we attempt to view a page that doesn't exist, lets create one.
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)

}


// Either edit an existing wiki page, else create a new one if it doesn't exist.
func editHandler(w http.ResponseWriter, r *http.Request) {

	title, err := getTitle(w, r)

	if err != nil {
		return
	}

	p, err := loadPage(title)

	tmpl := "edit"

	if err != nil {
		p = &Page{Title: title}
		tmpl = "create"
	}

	renderTemplate(w, tmpl, p)

}

// Save a wiki page to filesystem.
func saveHandler(w http.ResponseWriter, r *http.Request) {

	title, err := getTitle(w, r)

	if err != nil {
		return
	}

	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	err = p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)

}

// Load http server, bind handlers.
func main() {

	port := ":8080"

	// Application handlers
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	log.Println("Server running, access via http://localhost" + port + "\n")

	http.ListenAndServe(port, nil)

}

