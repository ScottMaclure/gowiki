package wiki

import (
	"net/http"
	"html/template"
	"regexp"
	"errors"
	"../page"
)

// Load templates ONCE from filesystem, rather than on every request.
var templates = template.Must(template.ParseFiles(
	"templates/create.html", 
	"templates/edit.html", 
	"templates/view.html",
))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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
func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {

	err := templates.ExecuteTemplate(w, tmpl + ".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Example of abstracting shared code using a function literal.
func MakeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// Extra page title from request
		// Call the provided handler with that title.
		title, err := getTitle(w, r)
		
		if err != nil {
			return
		}

		fn(w, r, title)

	}
}

// Handle a redirect from base url to viewing a default wiki page.
// i.e. if user accesses http://localhost:8080/, redirect to http://localhost:8080/view/index
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/index", http.StatusFound)
}

// View a wiki page in html!
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := page.LoadPage(title)

	if err != nil {
		// If we attempt to view a page that doesn't exist, lets create one.
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)

}


// Either edit an existing wiki page, else create a new one if it doesn't exist.
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := page.LoadPage(title)

	tmpl := "edit"

	if err != nil {
		p = &page.Page{Title: title}
		tmpl = "create"
	}

	renderTemplate(w, tmpl, p)

}

// Save a wiki page to filesystem.
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")

	p := &page.Page{Title: title, Body: []byte(body)}

	err := p.Save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/" + title, http.StatusFound)

}



