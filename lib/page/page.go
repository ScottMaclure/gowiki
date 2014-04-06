package page

import (
	"io/ioutil"
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
func (p *Page) Save() error {

	filename := "data/" + p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)

}

// Load data from filesystem.
func LoadPage(title string) (*Page, error) {

	filename := "data/" + title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil

}
