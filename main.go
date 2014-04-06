package main

import (
	"log"
	"net/http"
	"./lib/wiki"
)

// Load http server, bind handlers.
func main() {

	port := ":8080"

	// Application handlers
	http.HandleFunc("/", wiki.DefaultHandler)
	// Wrap these handlers with common functionality.
	http.HandleFunc("/view/", wiki.MakeHandler(wiki.ViewHandler))
	http.HandleFunc("/edit/", wiki.MakeHandler(wiki.EditHandler))
	http.HandleFunc("/save/", wiki.MakeHandler(wiki.SaveHandler))

	log.Println("Server running, access via http://localhost" + port + "\n")

	http.ListenAndServe(port, nil)

}