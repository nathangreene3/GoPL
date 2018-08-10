package main

import (
	"fmt"
	"log"
	"net/http"
)

// main():nil
// Initiates a server on localhost:8000. Prints the URL to stdout.
func main() {
	// WebServer 1
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler(http.ResponseWriter, *http.Request):nil
// Prints the URL to stdout.
func handler(w http.ResponseWriter, r *http.Request) {
	// WebServer 1
	fmt.Fprintf(w, "URL.Path: %q\n", r.URL.Path)
}
