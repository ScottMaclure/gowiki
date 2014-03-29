package main

import (
	"fmt"
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
// TODO Write a function to convert Title to a file-friendly name. "Hello, World" -> hello-world.
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

func main() {

	p1 := &Page{Title: "TestPage", Body: []byte("This is a test page!")}
	p1.save();

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

}

